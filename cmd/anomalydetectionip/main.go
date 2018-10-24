package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"flag"
	"github.com/dm03514/anomaly-detection-ip/blacklist"
	_ "github.com/lib/pq"
)

func stubHandler() (blacklist.Handler, error) {
	metadata := map[string]string{
		"Name":        "Hello",
		"LastRefresh": time.Now().String(),
	}

	fmt.Printf("%+v\n", metadata)

	return blacklist.Handler{
		Blacklist: &blacklist.StubBlacklist{
			Result: blacklist.Result{
				Found: true,
				CIDR:  "X.X.X.X/X",
				Providers: []blacklist.Provider{
					{
						Name:     "test",
						Metadata: metadata,
					},
				},
			},
		},
	}, nil
}

func postgresHandler(dbConnectionString string) (blacklist.Handler, error) {
	p, err := blacklist.NewPostges(dbConnectionString)
	if err != nil {
		return blacklist.Handler{}, err
	}

	return blacklist.Handler{
		Blacklist: p,
	}, nil
}

func main() {
	handlerType := flag.String("handler-type", "stub", "")
	dbConnectionString := flag.String("db-connection-string", "", "")
	flag.Parse()
	var h blacklist.Handler
	var err error

	switch *handlerType {
	case "postgres":
		h, err = postgresHandler(*dbConnectionString)
	default:
		fmt.Printf("handlerType: %q unknown.  Defaulting to stub\n", *handlerType)
		h, err = stubHandler()
	}
	if err != nil {
		panic(err)
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
