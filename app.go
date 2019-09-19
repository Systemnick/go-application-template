package main

import (
	"fmt"
	"os"
	"time"

	"git.rnd.mtt/innovation/call-initiator/forms"
	"github.com/rs/zerolog"
)

type Config struct {
	WorkerCount   int
	LogFluentd    string
	VoicePlatform string
}

type Application struct {
	logger        zerolog.Logger
	voicePlatform IVoicePlatform
	storage       IStorage
	input         IInput
}

var config Config

func NewApplication() (*Application, error) {
	app := &Application{}

	app.logger = initLogger()
	app.voicePlatform = initVoicePlatform()

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

func initVoicePlatform() IVoicePlatform {
	// switch config.VoicePlatform {
	// case "FreeSWITCH":
	// 	voicePlatform = FreeSWITCH{}
	// }
	// voicePlatform.Connect()
	return nil
}

func (a *Application) Run() {
	var s []chan forms.InputForm

	for i := 0; i < config.WorkerCount; i++ {
		c := make(chan forms.InputForm)
		s = append(s, c)
		go a.StartWorker(c)
	}
}

func (a *Application) StartWorker(c chan forms.InputForm) {
	select {
	case input := <-c:
		existing, err := a.FindCall(input.CallSid)
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
