

package xapp

import (
	"os"
	"os/signal"
	"path/filepath"
	"sync/atomic"
	"syscall"
	"time"
)

type ReadyCB func(interface{})
type ShutdownCB func()

var (
	// XApp is an application instance
	Resource      *Router
	Metric        *Metrics
	Logger        *Log
	Config        Configurator
	Subscription  *Subscriber
	readyCb       ReadyCB
	readyCbParams interface{}
	shutdownCb    ShutdownCB
	shutdownFlag  int32
	shutdownCnt   int32
)

func IsReady() bool {
	return true
}

func SetReadyCB(cb ReadyCB, params interface{}) {
	readyCb = cb
	readyCbParams = params
}

func XappReadyCb(params interface{}) {
}

func SetShutdownCB(cb ShutdownCB) {
	shutdownCb = cb
}


func InstallSignalHandler() {
	//
	// Signal handlers to really exit program.
	// shutdownCb can hang until application has
	// made all needed gracefull shutdown actions
	// hardcoded limit for shutdown is 20 seconds
	//
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	//signal handler function
	go func() {
		for range interrupt {
			if atomic.CompareAndSwapInt32(&shutdownFlag, 0, 1) {
				// close function
				go func() {
					timeout := int(20)
					sentry := make(chan struct{})
					defer close(sentry)

					// close callback
					go func() {
						//XappShutdownCb()
						sentry <- struct{}{}
					}()
					select {
					case <-time.After(time.Duration(timeout) * time.Second):
						Logger.Info("xapp-frame shutdown callback took more than %d seconds", timeout)
					case <-sentry:
						Logger.Info("xapp-frame shutdown callback handled within %d seconds", timeout)
					}
					os.Exit(0)
				}()
			} else {
				newCnt := atomic.AddInt32(&shutdownCnt, 1)
				Logger.Info("xapp-frame shutdown already ongoing. Forced exit counter %d/%d ", newCnt, 5)
				if newCnt >= 5 {
					Logger.Info("xapp-frame shutdown forced exit")
					os.Exit(0)
				}
				continue
			}

		}
	}()
}

func init() {

	Logger = NewLogger(filepath.Base(os.Args[0])) //LoadConfig()

	InstallSignalHandler()
}
