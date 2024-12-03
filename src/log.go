package src

import (
	"log"
	"os"
)

var Logger = log.New(os.Stderr, "", log.LstdFlags)
