package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"git.rnd.mtt/innovation/call-initiator/input"
	"git.rnd.mtt/innovation/call-initiator/storage"
	"git.rnd.mtt/innovation/call-initiator/storage/tarantool"
	voicePlatform "git.rnd.mtt/innovation/call-initiator/voice-platform"
	"github.com/rs/zerolog"
)

type Config struct {
	WorkerCount   int
	LogFluentd    string
	VoicePlatform string
}

type Application struct {
	logger        zerolog.Logger
	voicePlatform voicePlatform.IVoicePlatform
	storage       storage.IStorage
	input         input.IInput
}

func (a *Application) Logger() zerolog.Logger {
	return a.logger
}

var config Config

func NewApplication() (*Application, error) {
	app := &Application{}

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
		go a.StartWorker(c)
	}

	return nil
}

func (a *Application) StartWorker(c chan input.InForm) {
	select {
	case inForm := <-c:
		existing, err := a.FindCall(inForm.CallSid)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("JSON: " + existing)
		}

	}
}

func (a *Application) FindCall(callSid string) (string, error) {
	return a.storage.FindRecord(callSid)
}

func (a *Application) Stop(context context.Context) error {
	// todo Stop all routines

	return nil
}
