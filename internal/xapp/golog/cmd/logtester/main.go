
package main

import (
	"fmt"
	"os"
	"time"
	mdcloggo "nRIC/internal/xapp/golog"
)

func main() {
	logger, _ := mdcloggo.InitLogger("myname")
	logFileMonitor := 0
	logger.MdcAdd("foo", "bar")
	logger.MdcAdd("foo2", "bar2")
	if logger.Mdclog_format_initialize(logFileMonitor) != 0 {
		logger.Error("Failed in MDC Log Format Initialize")
	}

	start := time.Now()
	for i := 0; i < 10; i++ {
		logger.Info("Some test logs")
	}
	elapsed := time.Since(start)
	fmt.Fprintf(os.Stderr, "Elapsed %v\n", elapsed)
}
