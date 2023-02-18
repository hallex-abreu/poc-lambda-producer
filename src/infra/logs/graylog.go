package logs

import (
	"log"
	"time"

	"gopkg.in/Graylog2/go-gelf.v2/gelf"
)

type Graylog struct{}

func (g Graylog) Log(info string, message string) {
	// Create a new GELF message
	sending_message := &gelf.Message{
		Version:  "1.1",
		Host:     "my-hostname",
		Short:    "This is a short message",
		Full:     message,
		TimeUnix: float64(time.Now().Unix()),
	}

	// Set up the GELF UDP writer
	gelfWriter, err := gelf.NewUDPWriter("graylog-server:12201")
	if err != nil {
		log.Fatalf("Failed to create GELF UDP writer: %v", err)
	}

	// Send the message
	if err := gelfWriter.WriteMessage(sending_message); err != nil {
		log.Fatalf("Failed to send GELF message: %v", err)
	}

	log.Println("GELF message sent successfully")
}
