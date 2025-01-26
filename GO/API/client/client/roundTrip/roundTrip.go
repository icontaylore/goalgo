package client

import (
	"net/http"
	"log"
	"fmt"
)

type LoggingRoundTripper struct {
	transport http.RoundTripper
}

func (l *LoggingRoundTripper) RoundTrip(req *http.Request) (*http.Response,error) {
	log.Printf("запрос:-",req.Method,req.URL)

	resp,err := l.transport.RoundTrip(req)
	if err != nil {
		fmt.Printf("error %v", err)
		return nil,err
	}

	log.Printf("ответ-",resp.Status)
	return resp,nil
	
}