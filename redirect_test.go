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
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

const redirectTestUrl string = "http://google.com"

func TestRedirectUrl(t *testing.T) {
	url, _ := url.Parse(redirectTestUrl)

	redirect := &Redirect{
		URL: url,
	}

	assert.Equal(t, "OK "+redirectTestUrl+"\n", redirect.String(), "the redirect result wasn't correct")
}

func TestRedirectStatus(t *testing.T) {
	url, _ := url.Parse(redirectTestUrl)

	redirect := &Redirect{
		URL:    url,
		Status: PermanantRedirect,
	}

	assert.Equal(t, "OK 301:"+redirectTestUrl+"\n", redirect.String(), "the redirect result wasn't correct")

	redirect = &Redirect{
		URL:    url,
		Status: FoundRedirect,
	}

	assert.Equal(t, "OK 302:"+redirectTestUrl+"\n", redirect.String(), "the redirect result wasn't correct")

	redirect = &Redirect{
		URL:    url,
		Status: TemporaryRedirect,
	}

	assert.Equal(t, "OK 307:"+redirectTestUrl+"\n", redirect.String(), "the redirect result wasn't correct")
}

func TestRedirectResult(t *testing.T) {
	url, _ := url.Parse(redirectTestUrl)

	redirect := &Redirect{
		URL:    url,
		Result: ErrorResult,
	}

	assert.Equal(t, "ERR "+redirectTestUrl+"\n", redirect.String(), "the redirect result wasn't correct")

	redirect = &Redirect{
		URL:    url,
		Result: BHResult,
	}

	assert.Equal(t, "BH "+redirectTestUrl+"\n", redirect.String(), "the redirect result wasn't correct")
}

func TestRedirectChannelID(t *testing.T) {
	url, _ := url.Parse(redirectTestUrl)

	redirect := &Redirect{
		ChannelID: 0,
		URL:       url,
	}

	assert.Equal(t, "OK "+redirectTestUrl+"\n", redirect.String(), "the redirect result wasn't correct")

	redirect = &Redirect{
		ChannelID: 1,
		URL:       url,
	}

	assert.Equal(t, "1 OK "+redirectTestUrl+"\n", redirect.String(), "the redirect result wasn't correct")
}
