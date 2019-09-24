package main

import (
	"context"
	"os"
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
	id            string
	logger        zerolog.Logger
	voicePlatform voicePlatform.IVoicePlatform
	storage       storage.IStorage
	input         input.IInput
}

var config Config

func NewApplication() (*Application, error) {
	app := &Application{}

	app.id = uuid.NewV4().String()
	app.logger = initLogger()
	app.voicePlatform = initVoicePlatform()
	app.storage = initStorage()

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
	var s []chan input.InForm

	for i := 0; i < config.WorkerCount; i++ {
		c := make(chan input.InForm)
		s = append(s, c)
		w := a.NewWorker(c)
		go w.Run()
	}

	return nil
}

func (a *Application) Stop(context context.Context) error {
	// todo Stop all routines
	a.logger.Info().Str("application", a.id).Msg("Stopping application")

	return nil
}
