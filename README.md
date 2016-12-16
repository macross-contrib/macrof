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
	return func(self *macross.Context) error {
		return self.String("Hello, World!\n")
	}
}

func main() {
    m := macross.New()
    m.Get("/", hello())

    // automatically add routers for net/http/pprof
    // e.g. /debug/pprof, /debug/pprof/heap, etc.
    macrof.Wrapper(m)
    m.Listen(":8080")
}
```
Start this server, and now visit http://127.0.0.1:8080/debug/pprof/ and you'll see what you want.



