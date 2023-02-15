package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/golang/glog"
)

func log_output(remot_addr string, code int) {
	str_tmp := fmt.Sprintf("%s %s", remot_addr, strconv.Itoa(code))
	glog.V(2).Info(str_tmp)
}

func main() {
	flag.Set("v", "4")
	glog.V(2).Info("Starting http server...")
	mux := http.NewServeMux()
	mux.HandleFunc("/header", header)
	mux.HandleFunc("/variable", variable)
	mux.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", mux)
	if err != nil {
		glog.Error(err)
	}

}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	log_output(r.RemoteAddr, 200)
}

func header(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		w.Header().Set(k, strings.Join(v, ","))
	}
	w.WriteHeader(http.StatusOK)
	log_output(r.RemoteAddr, 200)
}

func variable(w http.ResponseWriter, r *http.Request) {
	env := os.Getenv("VERSION")
	w.Header().Set("VERSION", env)
	w.WriteHeader(http.StatusOK)
	log_output(r.RemoteAddr, 200)
}
