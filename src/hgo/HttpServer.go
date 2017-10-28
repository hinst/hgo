package hgo

import (
	"context"
	"net/http"
	"strings"
	"time"
)

func StartHttpServer(address string) *http.Server {
	var server = &(http.Server{Addr: address})
	go server.ListenAndServe()
	return server
}

func StopHttpServer(server *http.Server, timeout time.Duration) error {
	var ctx, cancel = context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return server.Shutdown(ctx)
}

func WrapFixJavaScriptContentType(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, ".js") {
			w.Header().Add("Content-Type", "text/javascript")
		}
		handler.ServeHTTP(w, r)
	}
}
