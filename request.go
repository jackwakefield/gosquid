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
	"errors"
	"net/url"
	"strconv"
	"strings"
)

type Request struct {
	ChannelID int
	URL       *url.URL
}

var ErrorNoRequestSegments = errors.New("the request contains no segments")
var ErrorNoRequestUrl = errors.New("the request contains no URL")
var ErrorInvalidRequestUrl = errors.New("the request contains an invalid URL")

func parseRequest(line string) (*Request, error) {
	var rawURL string
	channelId := 1

	line = strings.Replace(line, "\n", "", -1)
	segments := strings.Split(line, " ")

	if len(segments) == 0 || segments[0] == "" {
		return nil, ErrorNoRequestSegments
	}

	if value, err := strconv.Atoi(segments[0]); err == nil {
		if len(segments) == 1 {
			return nil, ErrorNoRequestUrl
		}

		channelId = value
		rawURL = segments[1]
	} else {
		rawURL = segments[0]
	}

	if url, err := url.Parse(rawURL); err == nil {
		return &Request{
			URL:       url,
			ChannelID: channelId,
		}, nil
	}

	return nil, ErrorInvalidRequestUrl
}
