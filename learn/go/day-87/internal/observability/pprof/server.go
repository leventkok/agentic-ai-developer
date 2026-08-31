package pprof

import (
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"
)

func Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	return mux
}

func Start(addr string) (*http.Server, error) {
	if addr == "" {
		return nil, nil
	}
	srv := &http.Server{Addr: addr, Handler: Handler()}
	go func() {
		log.Printf("pprof debug server on http://localhost%s/debug/pprof/", addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("pprof server: %v", err)
		}
	}()
	return srv, nil
}

func Shutdown(srv *http.Server) error {
	if srv == nil {
		return nil
	}
	return srv.Close()
}

func Addr(port string) string {
	if port == "" {
		return ""
	}
	return fmt.Sprintf(":%s", port)
}
