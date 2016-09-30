# Toggl Time entry
 
 Toggl time entry is an interface for [toggl](https://github.com/toggl/toggl_api_docs) time entries.
 
[![Build Status](https://travis-ci.org/dougEfresh/toggl-timeentry.svg?branch=master)](https://travis-ci.org/dougEfresh/toggl-timeentry)
[![Go Report Card](https://goreportcard.com/badge/github.com/dougEfresh/toggl-timeentry)](https://goreportcard.com/report/github.com/dougEfresh/toggl-timeentry)
[![GoDoc](https://godoc.org/github.com/dougEfresh/toggl-timeentry?status.svg)](https://godoc.org/github.com/dougEfresh/toggl-timeentry)
[![Coverage Status](https://coveralls.io/repos/github/dougEfresh/toggl-timeentry/badge.svg?branch=master)](https://coveralls.io/github/dougEfresh/toggl-timeentry?branch=master)
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/dougEfresh/toggl-timeentry/master/LICENSE)

**Example:**

```sh
go get gopkg.in/dougEfresh/gtoggl.v8 gopkg.in/dougEfresh/toggl-timeentry.v8
```

```go
import "gopkg.in/dougEfresh/gtoggl.v8"
import "gopkg.in/dougEfresh/toggl-timeentry.v8"

func main() {
    thc, err := gtoggl.NewClient("token")
    ...
    tc, err := gtimeentry.NewClient(thc)
    ...
    timeentry,err := tc.Get(1)
    if err == nil {
        panic(err)
    }
}
```  
 

