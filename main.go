package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dm03514/anomaly-detection-ip/blacklist"
)

func main() {
	metadata := struct {
		Name        string
		LastRefresh time.Time
	}{
		Name:        "Hello",
		LastRefresh: time.Now(),
	}

	fmt.Printf("%+v\n", metadata)

	h := blacklist.Handler{
		Blacklist: &blacklist.StubBlacklist{
			Result: blacklist.Result{
				Found:    true,
				Metadata: metadata,
			},
		},
	}
	s := &http.Server{
		Addr:           ":8080",
		Handler:        h,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Printf("starting_server: %q\n", s.Addr)
	log.Fatal(s.ListenAndServe())
}
