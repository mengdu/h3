package main

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"os/signal"
	"syscall"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("%s %s\n", req.Method, req.URL.String())
	buf, err := httputil.DumpRequest(req, true)
	fmt.Println(err, string(buf))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hi!\n"))
}

func main() {
	address := "./demo.sock"
	listener, err := net.Listen("unix", address)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	srv := http.Server{
		Handler: &MyHandler{},
	}

	go func() {
		if err := srv.Serve(listener); err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
