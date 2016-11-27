# macrof
A wrapper for golang web framework macross to use net/http/pprof easily.

# install
First install macrof to your GOPATH using go get:
```
go get github.com/macross-contrib/macrof
```

# Usage
```
package main

import (
    "github.com/insionng/macross"
    "github.com/macross-contrib/macrof"
)

func hello() macross.Handler {
	return func(c *macross.Context) error {
		return c.String(macross.StatusOK, "Hello, World!\n")
	}
}

func main() {
    e := macross.New()
    e.Get("/", hello())

    // automatically add routers for net/http/pprof
    // e.g. /debug/pprof, /debug/pprof/heap, etc.
    macrof.Wrapper(e)
    e.Run(":8080")
}
```
Start this server, and now visit http://127.0.0.1:8080/debug/pprof/ and you'll see what you want.



