package signal

import (
	"os"
	"os/signal"
)

func Notify(c chan<- os.Signal, sig ...os.Signal) {
	signal.Notify(c, sig...)
}
