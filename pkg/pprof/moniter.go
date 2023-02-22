package pprof

import (
	"net/http"
	_ "net/http/pprof"

	log "github.com/sirupsen/logrus"
)

func Init() {
	go func() {
		log.Fatal(http.ListenAndServe(":6060", nil))
	}()
}

//感觉要在一个稳定的ip下面搞。