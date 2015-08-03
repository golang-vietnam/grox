package logs

import (
	"log"
	"os"
)

func New(prefix string) *log.Logger {
	logger := log.New(os.Stderr, prefix+": ", log.Ldate|log.Ltime|log.Lshortfile)

	return logger
}
