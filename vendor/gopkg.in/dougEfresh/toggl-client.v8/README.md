# Toggl Client
 
 Toggl Client interacts with [toggl](https://github.com/toggl/toggl_api_docs) client endpoint for the [Go](http://www.golang.org/) programming language.
 
[![Build Status](https://travis-ci.org/dougEfresh/toggl-client.svg?branch=master)](https://travis-ci.org/dougEfresh/toggl-client)
[![Go Report Card](https://goreportcard.com/badge/github.com/dougEfresh/toggl-client)](https://goreportcard.com/report/github.com/dougEfresh/toggl-client)
[![GoDoc](https://godoc.org/github.com/dougEfresh/toggl-client?status.svg)](https://godoc.org/github.com/dougEfresh/toggl-client)
[![Coverage Status](https://coveralls.io/repos/github/dougEfresh/toggl-client/badge.svg?branch=master)](https://coveralls.io/github/dougEfresh/toggl-client?branch=master)
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/dougEfresh/toggl-client/master/LICENSE)

**Example:**

```sh
go get gopkg.in/dougEfresh/gtoggl.v8 gopkg.in/dougEfresh/toggl-client.v8
```

```go
import "gopkg.in/dougEfresh/gtoggl.v8"
import "gopkg.in/dougEfresh/toggl-client.v8"

func main() {
	thc, err := gtoggl.NewClient("token")
	...
	tc, err := gclient.NewClient(thc)
	...
	client,err := tc.Get(1)
	if err == nil {
		panic(err)
	}
}
``` 

