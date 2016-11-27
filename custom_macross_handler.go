package macrof

import (
	"github.com/insionng/macross"
	"net/http"
	"sync"
)

type customMacrossHandler struct {
	httpHandler http.Handler

	wrappedHandleFunc macross.Handler
	once              sync.Once
}

func (ceh *customMacrossHandler) Handle(c *macross.Context) error {
	ceh.once.Do(func() {
		ceh.wrappedHandleFunc = ceh.mustWrapHandleFunc(c)
	})
	return ceh.wrappedHandleFunc(c)
}

func (ceh *customMacrossHandler) mustWrapHandleFunc(c *macross.Context) macross.Handler {
	return NewFastHTTPMacrossAdaptor(ceh.httpHandler)
}

func fromHTTPHandler(httpHandler http.Handler) *customMacrossHandler {
	return &customMacrossHandler{httpHandler: httpHandler}
}

func fromHandlerFunc(serveHTTP func(w http.ResponseWriter, r *http.Request)) *customMacrossHandler {
	return &customMacrossHandler{httpHandler: &customHTTPHandler{serveHTTP: serveHTTP}}
}
