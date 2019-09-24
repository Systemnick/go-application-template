package main

import (
	"context"
	"os"
	"sync"
	"time"

	"git.rnd.mtt/innovation/call-initiator/input"
	"git.rnd.mtt/innovation/call-initiator/storage"
	"git.rnd.mtt/innovation/call-initiator/storage/tarantool"
	voicePlatform "git.rnd.mtt/innovation/call-initiator/voice-platform"
	"github.com/rs/zerolog"
	"github.com/twinj/uuid"
)

type Config struct {
	WorkerCount   int
	LogFluentd    string
	VoicePlatform string
}

type Application struct {
	id             string
	config         *Config
	logger         zerolog.Logger
	voicePlatform  voicePlatform.IVoicePlatform
	storage        storage.IStorage
	input          input.IInput
	workerChannels []chan input.InForm
	wait           *sync.WaitGroup
}

func NewApplication(c *Config) (*Application, error) {
	app := &Application{}

	app.id = uuid.NewV4().String()
	app.config = c
	app.logger = initLogger()
	app.voicePlatform = initVoicePlatform()
	app.storage = initStorage()
	app.wait = &sync.WaitGroup{}

	return app, nil
}

func initLogger() (zerolog.Logger) {
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.DurationFieldUnit = time.Millisecond
	zerolog.DurationFieldInteger = true
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	return logger
}

func initVoicePlatform() voicePlatform.IVoicePlatform {
	// switch config.VoicePlatform {
	// case "FreeSWITCH":
	// 	voicePlatform = FreeSWITCH{}
	// }
	// voicePlatform.Connect()
	return nil
}

func initStorage() storage.IStorage {
	t := &tarantool.Tarantool{}
	t.Dsn = ":3301"
	return t
}

func (a *Application) Run() error {
	for i := 0; i < a.config.WorkerCount; i++ {
		c := make(chan input.InForm)
		a.workerChannels = append(a.workerChannels, c)
		w := a.NewWorker(c)
		a.wait.Add(1)
		go w.Run(a.wait)
	}

	return nil
}

func (a *Application) Stop(context context.Context) error {
	// todo Stop all routines
	a.logger.Debug().Str("application", a.id).Msg("Stopping application")
	for _, channel := range a.workerChannels {
		close(channel)
	}
	a.wait.Wait()
	a.logger.Info().Str("application", a.id).Msg("Stopping completed")

	return nil
}
