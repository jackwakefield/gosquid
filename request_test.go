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
	"testing"

	"github.com/stretchr/testify/assert"
)

const requestTestUrl string = "http://google.com"
const requestTestChannel int = 100

func TestRequestParse(t *testing.T) {
	request, err := parseRequest(requestTestUrl)

	assert.Nil(t, err, "parsing errored")
	assert.Equal(t, 1, request.ChannelID, "the channel ID is incorrect")
	assert.Equal(t, requestTestUrl, request.URL.String(), "the URL is incorrect")
}

func TestChannelRequestParse(t *testing.T) {
	request, err := parseRequest(fmt.Sprintf("%d %s", requestTestChannel, requestTestUrl))

	assert.Nil(t, err, "parsing errored")
	assert.Equal(t, requestTestChannel, request.ChannelID, "the channel ID is incorrect")
	assert.Equal(t, requestTestUrl, request.URL.String(), "the URL is incorrect")
}

func TestEmptyRequestParse(t *testing.T) {
	request, err := parseRequest("")

	assert.Nil(t, request, "request shouldn't be parsed")
	assert.NotNil(t, err, "the parse should have errored")
}

func TestNoURLRequestParse(t *testing.T) {
	request, err := parseRequest("1")

	assert.Nil(t, request, "request shouldn't be parsed")
	assert.NotNil(t, err, "the parse should have errored")
}
