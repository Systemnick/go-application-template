package main

import (
	"context"
	"fmt"
	"net/http"
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

	go func() {
		if err := a.Run(); err != http.ErrServerClosed {
			a.Logger().Fatal().Msgf("http server run error: %s", err)
			return
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	a.Logger().Info().Msg("get signal start to shutdown")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := a.Stop(ctx); err != nil {
		a.Logger().Error().Msgf("application stop error %s", err)
		return
	}
}
