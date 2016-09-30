# Toggl User
 
 Toggl User is an interface for [toggl](https://github.com/toggl/toggl_api_docs) users.
 
[![Build Status](https://travis-ci.org/dougEfresh/toggl-user.svg?branch=master)](https://travis-ci.org/dougEfresh/toggl-user)
[![Go Report Card](https://goreportcard.com/badge/github.com/dougEfresh/toggl-user)](https://goreportcard.com/report/github.com/dougEfresh/toggl-user)
[![GoDoc](https://godoc.org/github.com/dougEfresh/toggl-user?status.svg)](https://godoc.org/github.com/dougEfresh/toggl-user)
[![Coverage Status](https://coveralls.io/repos/github/dougEfresh/toggl-user/badge.svg?branch=master)](https://coveralls.io/github/dougEfresh/toggl-user?branch=master)
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/dougEfresh/toggl-user/master/LICENSE)

**Example:**

```sh
go get gopkg.in/dougEfresh/gtoggl.v8 gopkg.in/dougEfresh/toggl-user.v8
```

```go
import "gopkg.in/dougEfresh/gtoggl.v8"
import "gopkg.in/dougEfresh/toggl-user.v8"

func main() {
    thc, err := gtoggl.NewClient("token")
    ...
    tc, err := guser.NewClient(thc)
    ...
    user,err := tc.Get(false)
    if err == nil {
        panic(err)
    }
}
```  

