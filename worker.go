package main

import (
	"encoding/json"
	"fmt"
	"sync"

	"git.rnd.mtt/innovation/call-initiator/input"
	"git.rnd.mtt/innovation/call-initiator/storage"
	voicePlatform "git.rnd.mtt/innovation/call-initiator/voice-platform"
	"github.com/twinj/uuid"
)

type Worker struct {
	id      string
	app     *Application
	channel chan input.InForm
}

func (a *Application) NewWorker(c chan input.InForm) *Worker {
	w := &Worker{}

	w.id = uuid.NewV4().String()
	w.app = a
	w.channel = c

	a.logger.Debug().Str("application", a.id).Str("worker", w.id).
		Str("storage", fmt.Sprintf("%T", a.storage)).
		Str("voice_platform", fmt.Sprintf("%T", a.voicePlatform)).
		Msg("Worker created")

	return w
}

func (w *Worker) Run(wg *sync.WaitGroup) {
	defer wg.Done()

	a := w.app

	a.logger.Debug().Str("application", a.id).Str("worker", w.id).
		Str("storage", fmt.Sprintf("%T", a.storage)).
		Str("voice_platform", fmt.Sprintf("%T", a.voicePlatform)).
		Msg("Worker started")

	for {
		inForm, more := <-w.channel
		if !more {
			a.logger.Debug().Str("application", a.id).Str("worker", w.id).
				Str("storage", fmt.Sprintf("%T", a.storage)).
				Str("voice_platform", fmt.Sprintf("%T", a.voicePlatform)).
				Msg("No more messages, exiting")
			break
		}

		switch inForm.CallType {
		case input.NewCall:
			call := w.convertCallInputToVoicePlatform(&inForm)
			scenario := w.convertScenarioInputToVoicePlatform(inForm.Scenario)

			// todo Save new call to storage

			err := w.processNewCall(call, scenario)
			if err != nil {
				a.logger.Warn().Str("application", a.id).Str("worker", w.id).
					Str("call_sid", inForm.CallSid).
					Str("voice_platform", fmt.Sprintf("%T", a.voicePlatform)).
					Msgf("processNewCall error: %s", err)
				continue
			}

		case input.ExistingCall:
			existing, err := w.findCall(inForm.CallSid)
			if err != nil {
				a.logger.Warn().Str("application", a.id).Str("worker", w.id).
					Str("call_sid", inForm.CallSid).
					Str("voice_platform", fmt.Sprintf("%T", a.voicePlatform)).
					Msgf("findCall error: %s", err)
				continue
			}

			call := w.convertCallInputToVoicePlatform(&inForm)
			err = w.modifyCall(call)
			if err != nil {
				a.logger.Warn().Str("application", a.id).Str("worker", w.id).
					Str("call_sid", inForm.CallSid).
					Str("voice_platform", fmt.Sprintf("%T", a.voicePlatform)).
					Msgf("modifyCall error: %s", err)
				continue
			}

			a.logger.Debug().Str("application", a.id).Str("worker", w.id).
				Str("call_sid", inForm.CallSid).
				Str("voice_platform", fmt.Sprintf("%T", a.voicePlatform)).
				Msgf("JSON: %s", existing)

		case input.AddParticipant:
		}
	}
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////// //
// todo Move to new aggregate?

func (w *Worker) findCall(callSid string) (storage.Call, error) {
	a := w.app

	record, err := a.storage.FindRecord(callSid)
	if err != nil {
		a.logger.Warn().Str("application", a.id).Str("worker", w.id).
			Str("call_sid", callSid).
			Str("storage", fmt.Sprintf("%T", a.storage)).
			Msgf("storage.FindRecord error: %s", err)
	}
	call := storage.Call{}

	err = json.Unmarshal([]byte(record), &call)
	if err != nil {
		a.logger.Warn().Str("application", a.id).Str("worker", w.id).
			Str("call_sid", callSid).
			Str("storage", fmt.Sprintf("%T", a.storage)).
			RawJSON("record", []byte(record)).
			Msgf("json.Unmarshal error: %s", err)
	}

	return call, err
}

func (w *Worker) processNewCall(call *voicePlatform.Call, scenario *voicePlatform.Scenario) error {

	return nil
}

func (w *Worker) modifyCall(call *voicePlatform.Call) error {

	return nil
}

func (w *Worker) convertCallInputToVoicePlatform(inputCall *input.InForm) *voicePlatform.Call {
	v := &voicePlatform.Call{}
	// todo Fill all the rest fields
	v.Sid = inputCall.CallSid

	return v
}

func (w *Worker) convertScenarioInputToVoicePlatform(scenario *input.Scenario) *voicePlatform.Scenario {
	s := &voicePlatform.Scenario{}
	for _, command := range *scenario {
		action := voicePlatform.Action{}
		action.Command = command.Command
		action.Params = command.Params
		action.Count = command.Count
		*s = append(*s, action)
	}

	return s
}
