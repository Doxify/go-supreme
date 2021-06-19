package gosupreme

import (
	"log"
	"net/http"
	"os"
	"time"
)

type Supreme struct {
	// Stock holds Supreme's product inventory
	Stock *Stock

	// Styles holds Supreme's product styles
	Styles *Styles

	l *log.Logger
	c *http.Client
}

// Creates a new gosupreme instance
func New() *Supreme {
	return &Supreme{
		Stock:  nil,
		Styles: nil,
		l:      log.New(os.Stdout, "gosupreme ", log.Ldate|log.Lmicroseconds),
		c: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (s *Supreme) Init() {
	s.FetchStock()
}
