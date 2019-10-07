package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	c := initConfig()

	a, err := NewApplication(c)
	if err != nil {
		fmt.Printf("Creating application error: %s\n", err.Error())
		return
	}

	if err := a.Run(); err != nil {
		fmt.Printf("Application run error: %s\n", err.Error())
		return
	}

	quit := make(chan os.Signal)
	signal.Notify(quit,
		// os.Signal(syscall.SIGHUP),
		os.Signal(syscall.SIGINT),
		os.Signal(syscall.SIGQUIT),
		// os.Signal(syscall.SIGILL),
		// os.Signal(syscall.SIGTRAP),
		os.Signal(syscall.SIGABRT),
		// os.Signal(syscall.SIGBUS),
		// os.Signal(syscall.SIGFPE),
		// os.Signal(syscall.SIGKILL),
		// os.Signal(syscall.Signal(0xa)),
		os.Signal(syscall.SIGSEGV),
		// os.Signal(syscall.Signal(0xc)),
		// os.Signal(syscall.SIGPIPE),
		// os.Signal(syscall.SIGALRM),
		os.Signal(syscall.SIGTERM),
	)

	reload := make(chan os.Signal)
	signal.Notify(reload,
		os.Signal(syscall.SIGHUP),
		// os.Signal(syscall.SIGINT),
		// os.Signal(syscall.SIGQUIT),
		os.Signal(syscall.SIGILL),
		os.Signal(syscall.SIGTRAP),
		// os.Signal(syscall.SIGABRT),
		os.Signal(syscall.SIGBUS),
		os.Signal(syscall.SIGFPE),
		// os.Signal(syscall.SIGKILL),
		os.Signal(syscall.Signal(0xa)),
		// os.Signal(syscall.SIGSEGV),
		os.Signal(syscall.Signal(0xc)),
		os.Signal(syscall.SIGPIPE),
		os.Signal(syscall.SIGALRM),
		// os.Signal(syscall.SIGTERM),
	)

	for {
		select {
		case s := <-quit:
			fmt.Printf("Signal '%s' was received\n", s.String())
			ctx := Quit()
			if err := a.Stop(ctx); err != nil {
				fmt.Printf("Application stop error: %s\n", err.Error())
				return
			}
		case s := <-reload:
			fmt.Printf("Signal '%s' was received\n", s.String())

		}
	}
}

func Quit() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	return ctx
}

func initConfig() *Config {
	wc, err := strconv.Atoi(os.Getenv("WORKER_COUNT"))
	if err != nil {
		fmt.Printf("Environment variable WORKER_COUNT: bad integer: %s\n", err.Error())
	}

	c := &Config{}
	c.WorkerCount = wc

	return c
}
