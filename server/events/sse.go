package events

import (
	"github.com/alexandrevicenzi/go-sse"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	SSE *sse.Server
)

func InitSSE() *sse.Server {
	options := &sse.Options{
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
		},
	}

	options.Logger = log.New(os.Stdout, "go-sse: ", log.Ldate|log.Ltime|log.Lshortfile)

	go Heartbeat()

	return sse.NewServer(options)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	SSE.ServeHTTP(w, r)
}

//Needed for polyfill
func Heartbeat() {
	for {
		for _, channel := range SSE.Channels() {
			SSE.SendMessage(channel, sse.SimpleMessage("heartbeat"))
		}

		<-time.After(30 * time.Second)
	}
}
