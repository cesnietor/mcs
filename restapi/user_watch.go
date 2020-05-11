// This file is part of MinIO Console Server
// Copyright (c) 2020 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package restapi

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
	mc "github.com/minio/mc/cmd"
)

type watchParams struct {
	BucketName string
	mc.WatchParams
}

// getParamsFromWsWatchPath gets bucket name, events, prefix, suffix from a websocket
// watch path if defined.
// The path should come as : `/watch/bucket1?prefix=&suffix=.jpg&events=put,get`
func getParamsFromWsWatchPath(wsPath string) watchParams {
	wParams := watchParams{}
	re := regexp.MustCompile(`(^/watch/)(.*?)(\?.*?$|$)`)
	matches := re.FindAll([]byte(wsPath), -1)
	// matches comes as e.g.
	// [["...", "/watch/" "bucket1" "?prefix=&suffix=.jpg&events=put,get"]]
	// [["/watch/bucket1" "/watch/" "bucket1" ""]]
	// [["/watch/" "/watch/" "" ""]]
	if len(matches[0]) > 2 {
		// bucket name is on the second group, third position
		wParams.BucketName = strings.TrimSpace(string(matches[0][2]))
	}
	// if matched query params (comes in third group) get fourth position
	if len(matches[0]) > 3 {
		q, err := url.ParseQuery(strings.TrimPrefix("?", string(matches[0][3])))
		if err != nil {
			log.Fatal(err)
		}
		wParams.Events = strings.Split(q.Get("events"), ",")
		wParams.Prefix = q.Get("prefix")
		wParams.Suffix = q.Get("suffix")
	}
	return wParams
}

// startWatch
func startWatch(wsc *wsS3Client, params watchParams) (mError error) {
	// a WaitGroup waits for a collection of goroutines to finish
	wg := sync.WaitGroup{}
	// a cancel context is needed to end all goroutines used
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Set number of goroutines to wait. wg.Wait()
	// waits until counter is zero (all are done)
	wg.Add(3)
	// start go routine for reading websocket heartbeat
	readErr := wsReadCheck(ctx, &wg, wsc.conn)
	// send Stream of watch events to the ws c.connection
	ch := sendWatchInfo(ctx, &wg, wsc, params)
	// If wsReadCheck returns it means that it is not possible to check
	// ws heartbeat anymore so we stop from doing Console Log, cancel context
	// for all goroutines.
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		if err := <-readErr; err != nil {
			log.Println("error on wsReadCheck:", err)
			mError = err
		}
		// cancel context for all goroutines.
		cancel()
	}(&wg)

	if err := <-ch; err != nil {
		mError = err
	}

	// if ch closes for any reason,
	// cancel context for all goroutines
	cancel()
	// wait all goroutines to finish
	wg.Wait()
	return mError
}

// sendWatchInfo sends stream of Watch Event to the ws connection
func sendWatchInfo(ctx context.Context, wg *sync.WaitGroup, wsc *wsS3Client, params watchParams) <-chan error {
	// decrements the WaitGroup counter
	// by one when the function returns
	defer wg.Done()
	ch := make(chan error)
	go func(ch chan<- error) {
		defer close(ch)
		wo, err := wsc.client.watch(params.WatchParams)
		for {
			select {
			case <-ctx.Done():
				wo.Close()
				return
			case events, ok := <-wo.Events():
				// zero value returned because the channel is closed and empty
				if !ok {
					return
				}
				for _, event := range events {
					// Serialize message to be sent
					bytes, err := json.Marshal(event)
					if err != nil {
						fmt.Println("error on json.Marshal:", err)
						ch <- err
						return
					}
					// Send Message through websocket connection
					err = wsc.conn.writeMessage(websocket.TextMessage, bytes)
					if err != nil {
						log.Println("error writeMessage:", err)
						ch <- err
						return
					}
				}
			case pErr, ok := <-wo.Errors():
				// zero value returned because the channel is closed and empty
				if !ok {
					return
				}
				if err != nil {
					log.Println("error on watch:", pErr.Cause)
					ch <- pErr.Cause
					return

				}
			}
		}
	}(ch)

	return ch
}
