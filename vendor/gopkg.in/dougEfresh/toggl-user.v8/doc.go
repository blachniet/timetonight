// Copyright 2016 Douglas Chimento.  All rights reserved.

/*
Package guser provides access to toggl REST API


Example:
       import "gopkg.in/dougEfresh/gtoggl.v8"
       import "ggopkg.in/dougEfresh/toggl-user.v8"

       func main() {
	    thc, err := gtoggl.NewClient("token")
	    ...
	    tc, err := gtimeentry.NewClient(thc)
	    ...
	    me,err := tc.Get(false)
	    if err == nil {
	 	panic(err)
	   }
       }
*/
package guser
