package main

import (
	"errors"
	"log"
)

type status struct {
	Message string // Indicates what exactly is starting up
	Failure error  // because i want to know what the system spit back out
}

func main() {
	log.Printf("Let's light this candle")

	r := run()
	for _, m := range r {

		if m.Failure != nil {
			m.Message = m.Message + m.Failure.Error()
		}
		log.Printf(m.Message)
	}

}

func run() (results []status) {
	results = append(results,
		initdb(),
		initweb(),
	)

	// error correction ommited - fix later

	return results
}

func initdb() (s status) {

	s.Message = "db server initializing... "
	s.Failure = errors.New("db error")
	return s
}

func initweb() (s status) {
	s.Message = "web server initializing... "
	s.Failure = errors.New("web server error")
	return s
}

func initsms() (s status) {
	s.Message = "sms server initializing... "
	// s.Failure = errors.New("sms server error")
	return s
}
