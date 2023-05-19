package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/cilium/ebpf"
)

const (
	bpfObjectName = "whatchadoin_kern"
	mapName       = "process_map"
)

type ProcessInfo struct {
	PID  uint32
	Name [16]byte
}

func main() {
	spec, err := ebpf.LoadCollectionSpec(fmt.Sprintf("bpf/%s.o", bpfObjectName))
	if err != nil {
		panic(err)
	}

	coll, err := ebpf.NewCollection(spec)
	if err != nil {
		panic(fmt.Sprintf("Failed to create new collection: %v\n", err))
	}
	defer coll.Close()

	prog := coll.Programs["whatchadoin"]
	if prog == nil {
		panic("No program named 'whatchadoin' found in collection")
	}

	// Create a signal handler to capture Ctrl+C and exit gracefully
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Starting ...")

	go func() {
		for {
			select {
			case <-signals:
				return
			}
		}
	}()

	// Wait until the user interrupts the program
	<-signals
}

func getProcessConnections(pid int) ([]string, error) {
	connections := []string{}
	return connections, nil
}
