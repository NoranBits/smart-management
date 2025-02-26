// /////////////////////////////////////////////
// src: ./internal/middleware/logging.go	 //
// desc: Provides Customlogging middleware	//
// //////////////////////////////////////////
package logger

import (
	"log"
)

// Init initializes the logger based on the desired log level.
func Init(level string) {

	log.Printf("Logger initialized with level: %s", level)
}
