// Copyright 2016 Douglas Chimento.  All rights reserved.

/*
Package gworkspace provides access to toggl REST API


Example:
       import "gopkg.in/dougEfresh/gtoggl.v8"
       import "ggopkg.in/dougEfresh/toggl-workspace.v8"

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
*/
package gworkspace
