package serverLogs

import (
	"log"
	"os"
)

var (
	Main = log.New(os.Stdout, "server	", log.Ldate|log.Ltime)
)
