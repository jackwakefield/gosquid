# GoSquid [![](http://img.shields.io/travis/jackwakefield/gosquid.svg?style=flat-square)](http://travis-ci.org/jackwakefield/gosquid) 

GoSquid is a package for reading and replying to Squid URL redirectors/rewriters.

## Installation

```
go get github.com/jackwakefield/gosquid
```

## Usage

[See GoDoc](http://godoc.org/github.com/jackwakefield/gosquid) for documentation.

## Example

```go
package main

import (
	"log"

	"github.com/jackwakefield/gosquid"
)

func main() {
	hub := gosquid.NewHub()
	go hub.Run()

	for {
		select {
		case request := <-hub.Request:
			var redirect *gosquid.Redirect

			if request.URL.Host == "google.com" {
				// rewrite the host to google.co.uk if google.com was requested
				redirect = &gosquid.Redirect{
					ChannelID: request.ChannelID,
					Result:    gosquid.ErrorResult,
					URL:       request.URL,
				}

				redirect.URL.Host = "google.co.uk"
			} else {
				// don't redirect the URL
				redirect = &gosquid.Redirect{
					ChannelID: request.ChannelID,
					Result:    gosquid.OkResult,
					URL:       request.URL,
				}
			}

			hub.Redirect <- redirect
		case err := <-hub.Error:
			log.Println("Error:", err)
		}
	}
}
```

## License

> Copyright 2014 Jack Wakefield
>
> Licensed under the Apache License, Version 2.0 (the "License");
> you may not use this file except in compliance with the License.
> You may obtain a copy of the License at
>
>     http://www.apache.org/licenses/LICENSE-2.0
>
> Unless required by applicable law or agreed to in writing, software
> distributed under the License is distributed on an "AS IS" BASIS,
> WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
> See the License for the specific language governing permissions and
> limitations under the License.