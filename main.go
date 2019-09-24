package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	a, err := NewApplication()
	if err != nil {
		fmt.Printf("Creating application error: %s\n", err.Error())
		return
	}

	fmt.Printf("Application: %+v\n", a)

	if err := a.Run(); err != nil {
		fmt.Printf("Application run error: %s\n", err.Error())
		return
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	s := <-quit

	fmt.Printf("Signal %s was received", s.String())
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := a.Stop(ctx); err != nil {
		fmt.Printf("Application stop error: %s\n", err.Error())
		return
	}
}
