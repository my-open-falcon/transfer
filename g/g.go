package g

import (
	"log"
	"runtime"
)

// changelog:
// 0.0.1: init project
// 0.0.4: bugfix: set replicas before add node
// 0.0.8: change receiver, mv proc cron to proc pkg, add readme, add gitversion, add config reload, add trace tools
// 0.0.9: fix bugs of conn pool(use transfer's private conn pool, named & minimum)
const (
	VERSION      = "0.0.9"
	GAUGE        = "GAUGE"
	COUNTER      = "COUNTER"
	DERIVE       = "DERIVE"
	DEFAULT_STEP = 60
	MIN_STEP     = 30
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
