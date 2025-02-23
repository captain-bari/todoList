package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"todo/service"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		syscall.SIGTERM, // "the normal way to politely ask a program to terminate"
		syscall.SIGINT,  // Ctrl+C
	)

	service := service.NewService() // Load
	service.Run()                   // Start

	for sig := range sigChan {
		fmt.Println("Signal:", sig)
		printStack()

		err := service.Stop() // Stop
		if err != nil {
			os.Exit(0)
		}

		os.Exit(1)
	}
}

func printStack() {
	buf := make([]byte, 1<<20)
	stacklen := runtime.Stack(buf, true)
	fmt.Printf("=== Stack dump requested ===\n***** goroutine dump *****\n%s\n***** end *****\n", buf[:stacklen])
}
