# Toggl Project
 
 Toggl project is an interface for [toggl](https://github.com/toggl/toggl_api_docs) projects.
 
[![Build Status](https://travis-ci.org/dougEfresh/toggl-project.svg?branch=master)](https://travis-ci.org/dougEfresh/toggl-project)
[![Go Report Card](https://goreportcard.com/badge/github.com/dougEfresh/toggl-project)](https://goreportcard.com/report/github.com/dougEfresh/toggl-project)
[![GoDoc](https://godoc.org/github.com/dougEfresh/toggl-project?status.svg)](https://godoc.org/github.com/dougEfresh/toggl-project)
[![Coverage Status](https://coveralls.io/repos/github/dougEfresh/toggl-project/badge.svg?branch=master)](https://coveralls.io/github/dougEfresh/toggl-project?branch=master)
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/dougEfresh/toggl-project/master/LICENSE)

**Example:**

```sh
go get gopkg.in/dougEfresh/gtoggl.v8 gopkg.in/dougEfresh/toggl-project.v8
```

```go
import "gopkg.in/dougEfresh/gtoggl.v8"
import "gopkg.in/dougEfresh/toggl-project.v8"

func main() {
    thc, err := gtoggl.NewClient("token")
    ...
    tc, err := gproject.NewClient(thc)
    ...
    project,err := tc.Get(1)
    if err == nil {
        panic(err)
    }
}
```  

