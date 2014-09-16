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
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type RedirectResult string

const (
	OkResult    RedirectResult = "OK"
	ErrorResult RedirectResult = "ERR"
	BHResult    RedirectResult = "BH"
)

type RedirectStatus int

const (
	NoRedirect        RedirectStatus = 0
	PermanantRedirect RedirectStatus = 301
	FoundRedirect     RedirectStatus = 302
	TemporaryRedirect RedirectStatus = 307
)

type Redirect struct {
	ChannelID int
	Result    RedirectResult
	Status    RedirectStatus
	URL       *url.URL
}

func (redirect *Redirect) String() string {
	values := make([]string, 0, 4)

	if redirect.ChannelID > 0 {
		values = append(values, strconv.Itoa(redirect.ChannelID))
	}

	result := redirect.Result

	if result == "" {
		result = OkResult
	}

	values = append(values, string(result))

	url := redirect.URL.String()

	if redirect.Status != NoRedirect {
		url = fmt.Sprintf("%d:%s", redirect.Status, url)
	}

	values = append(values, url)

	return strings.Join(values, " ") + "\n"
}
