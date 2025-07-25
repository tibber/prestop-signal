package main

import (
	"fmt"
	"os"
)

func main() {
	const fifo = "/app/prestop.pipe"
	const ack = "/app/prestop-ack.pipe"

	if err := sendAndWait(fifo, ack); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func sendAndWait(fifoPath, ackPath string) error {
	// Send signal
	signal, err := os.OpenFile(fifoPath, os.O_WRONLY, 0600)
	if err != nil {
		return fmt.Errorf("open signal pipe: %w", err)
	}
	defer signal.Close()

	if _, err := signal.Write([]byte("stop\n")); err != nil {
		return fmt.Errorf("write signal: %w", err)
	}

	// Wait for ack
	ack, err := os.OpenFile(ackPath, os.O_RDONLY, 0600)
	if err != nil {
		return fmt.Errorf("open ack pipe: %w", err)
	}
	defer ack.Close()

	buf := make([]byte, 16)
	n, err := ack.Read(buf)
	if err != nil {
		return fmt.Errorf("read ack: %w", err)
	}

	fmt.Printf("Received ack: %s\n", string(buf[:n]))
	return nil
}
