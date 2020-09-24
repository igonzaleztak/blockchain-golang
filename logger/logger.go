package logger

import (
	"log"
	"os"
)

// Log: variable used to log stuff
var (
	file, _ = os.Create("./info.log")
	Log     = log.New(file, "", 0)
)
