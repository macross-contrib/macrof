package macrof

import (
	"github.com/insionng/macross"
	"net/http/pprof"
)

func Wrap(e *macross.Macross) {
	e.Get("/debug/pprof/", fromHandlerFunc(pprof.Index).Handle)
	e.Get("/debug/pprof/heap", fromHTTPHandler(pprof.Handler("heap")).Handle)
	e.Get("/debug/pprof/goroutine", fromHTTPHandler(pprof.Handler("goroutine")).Handle)
	e.Get("/debug/pprof/block", fromHTTPHandler(pprof.Handler("block")).Handle)
	e.Get("/debug/pprof/threadcreate", fromHTTPHandler(pprof.Handler("threadcreate")).Handle)
	e.Get("/debug/pprof/cmdline", fromHandlerFunc(pprof.Cmdline).Handle)
	e.Get("/debug/pprof/profile", fromHandlerFunc(pprof.Profile).Handle)
	e.Get("/debug/pprof/symbol", fromHandlerFunc(pprof.Symbol).Handle)
}

var Wrapper = Wrap
