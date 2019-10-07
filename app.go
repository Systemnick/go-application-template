package main

import (
	"context"
	"os"
	"sync"
	"time"

	"git.rnd.mtt/innovation/call-initiator/input"
	"git.rnd.mtt/innovation/call-initiator/input/amqp"
	"git.rnd.mtt/innovation/call-initiator/storage"
	"git.rnd.mtt/innovation/call-initiator/storage/tarantool"
	"git.rnd.mtt/innovation/call-initiator/voice-platform/freeswitch"
	"github.com/rs/zerolog"
	"github.com/twinj/uuid"
)

type Config struct {
	WorkerCount   int
	LogFluentd    string
	VoicePlatform string
	FS            freeswitch.Params
}

type Application struct {
	id             string
	config         *Config
	logger         *zerolog.Logger
	input          input.IInput
	storage        storage.IStorage
	voicePlatform  *freeswitch.Freeswitch
	workerChannels []chan input.InForm
	wait           *sync.WaitGroup
}

func NewApplication(c *Config) (*Application, error) {
	a := &Application{}

	a.id = uuid.NewV4().String()
	a.config = c
	a.logger = initLogger()
	a.input = initInput()
	a.storage = initStorage()
	a.voicePlatform = initVoicePlatform(c.FS)
	a.wait = &sync.WaitGroup{}

	a.logger.Info().Str("application", a.id).Msg("Application started")

	return a, nil
}

func initLogger() *zerolog.Logger {
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.DurationFieldUnit = time.Millisecond
	zerolog.DurationFieldInteger = true
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	return &logger
}

func initInput() input.IInput {
	a := &amqp.AMQP{}
	a.Url = "amqp://guest:guest@localhost:5672/"
	return a
}

func initStorage() storage.IStorage {
	t := &tarantool.Tarantool{}
	t.Dsn = ":3301"
	return t
}

func initVoicePlatform(params freeswitch.Params) *freeswitch.Freeswitch {
	vp := freeswitch.Freeswitch{}
	err := vp.Connect(params)
	if err != nil {
		// todo Print something
	}
	return &vp
}

func (a *Application) Run() error {
	for i := 0; i < a.config.WorkerCount; i++ {
		c := make(chan input.InForm)
		a.workerChannels = append(a.workerChannels, c)
		w := a.NewWorker(i, c)
		a.wait.Add(1)
		go w.Run(a.wait)
	}

	return nil
}

func (a *Application) Stop(context context.Context) error {
	a.logger.Debug().Str("application", a.id).Msg("Stopping application")
	for _, channel := range a.workerChannels {
		close(channel)
		a.workerChannels = a.workerChannels[1:]
	}
	a.wait.Wait()
	a.logger.Info().Str("application", a.id).Msg("Stopping completed")

	return nil
}
