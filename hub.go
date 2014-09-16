// Copyright 2014 Jack Wakefield
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gosquid

import (
	"bufio"
	"os"
)

// Hub provides an interface for reading requests and replying with redirects
// to Squid.
type Hub struct {
	reader   *bufio.Reader
	writer   *bufio.Writer
	Request  chan *Request
	Redirect chan *Redirect
	Error    chan error
}

func NewHub() *Hub {
	return &Hub{
		reader:   bufio.NewReader(os.Stdin),
		writer:   bufio.NewWriter(os.Stdout),
		Request:  make(chan *Request),
		Redirect: make(chan *Redirect),
		Error:    make(chan error),
	}
}

// Run starts the read and writing process.
func (hub *Hub) Run() {
	go hub.redirect()

	for {
		message, err := hub.reader.ReadString('\n')

		if err != nil {
			hub.Error <- err
			break
		}

		var request *Request

		if request, err = parseRequest(message); err != nil {
			hub.Error <- err
			continue
		}

		hub.Request <- request
	}
}

func (hub *Hub) redirect() {
	for {
		select {
		case redirect := <-hub.Redirect:
			if _, err := hub.writer.WriteString(redirect.String()); err != nil {
				hub.Error <- err
			}

			hub.writer.Flush()
		}
	}
}
