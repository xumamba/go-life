package signal

import (
	"log"
	"os"
	"syscall"
	"testing"
	"time"
)

var (
	ch = make(chan os.Signal, 1)
)

func TestNotify(t *testing.T) {
	var (
		signal = syscall.Signal(0x6)
	)
	go signalHandler()
	ch <- signal
	time.Sleep(5 * time.Second)
}

func signalHandler() {
	Notify(ch, syscall.Signal(0x6))
	for {
		sig := <-ch
		log.Printf("get a signal %s, stop the process", sig.String())
		switch sig {
		case syscall.Signal(0x6):
			// release resources or do something
			log.Println("syscall.Signal(0x6)")
			time.Sleep(time.Second)
			return
		case syscall.Signal(0x7):
			log.Println("syscall.Signal(0x7)")
		default:
			log.Println(sig)
			return
		}
	}
}
