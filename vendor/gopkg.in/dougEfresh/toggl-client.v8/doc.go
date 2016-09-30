// Copyright 2016 Douglas Chimento.  All rights reserved.

/*
Package gclient provides access to toggl REST API.

Example:
       import "gopkg.in/dougEfresh/gtoggl.v8"
       import "ggopkg.in/dougEfresh/toggl-client.v8"

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
*/
package gclient
