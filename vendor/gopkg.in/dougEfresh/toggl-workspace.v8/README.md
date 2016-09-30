# Toggl Workspace
 
 Toggl workspace is an interface for [toggl](https://github.com/toggl/toggl_api_docs) workspaces.
 
[![Build Status](https://travis-ci.org/dougEfresh/toggl-workspace.svg?branch=master)](https://travis-ci.org/dougEfresh/toggl-workspace)
[![Go Report Card](https://goreportcard.com/badge/github.com/dougEfresh/toggl-workspace)](https://goreportcard.com/report/github.com/dougEfresh/toggl-workspace)
[![GoDoc](https://godoc.org/github.com/dougEfresh/toggl-workspace?status.svg)](https://godoc.org/github.com/dougEfresh/toggl-workspace)
[![Coverage Status](https://coveralls.io/repos/github/dougEfresh/toggl-workspace/badge.svg?branch=master)](https://coveralls.io/github/dougEfresh/toggl-workspace?branch=master)
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/dougEfresh/toggl-workspace/master/LICENSE)

**Example:**

```sh
go get gopkg.in/dougEfresh/gtoggl.v8 gopkg.in/dougEfresh/toggl-workspace.v8
```

```go
import "gopkg.in/dougEfresh/gtoggl.v8"
import "gopkg.in/dougEfresh/toggl-workspace.v8"

func main() {
    thc, err := gtoggl.NewClient("token")
    ...
    tc, err := gworkspace.NewClient(thc)
    ...
    workspace,err := tc.Get(1)
    if err == nil {
        panic(err)
    }
}
```  

