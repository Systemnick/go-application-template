package freeswitch

import (
	"fmt"
	"strings"

	voicePlatform "git.rnd.mtt/innovation/call-initiator/voice-platform"
	"github.com/0x19/goesl"
)

type Call struct {
	voicePlatform.Call
	Uuid     string
	Channel  string
	FromName string
	Timeout  int
	User     string
}

type Params struct {
	voicePlatform.Params
	Host     string
	Port     uint
	Password string
	Timeout  int
}

type Freeswitch voicePlatform.VoicePlatform

var client *goesl.Client

func (f Freeswitch) Connect(params Params) error {
	var err error

	client, err = goesl.NewClient(params.Host, params.Port, params.Password, params.Timeout)
	if err != nil {
		return fmt.Errorf("error while creating new ESL client: %s", err)
	}

	go client.Handle()

	// err = client.Send("events json ALL")
	// err = client.Send("events json CUSTOM")
	err = client.Send("events json")
	if err != nil {
		return fmt.Errorf("error while event subscribing from ESL client: %s", err)
	}

	err = client.BgApi(fmt.Sprintf("originate %s %s", "sofia/internal/1001@127.0.0.1", "&socket(192.168.1.2:8084 async full)"))
	if err != nil {
		return fmt.Errorf("error while sending BgApi command to ESL client: %s", err)
	}

	for {
		msg, err := client.ReadMessage()

		if err != nil {
			// If it contains EOF, we really dont care...
			if !strings.Contains(err.Error(), "EOF") && err.Error() != "unexpected end of JSON input" {
				// Error("Error while reading Freeswitch message: %s", err)
			}
			break
		}

		params.Events <- msg.String()
	}

	return nil
}

func (f Freeswitch) CreateCall(call Call) error {
	err := client.BgApi(fmt.Sprintf("originate %s %s", "sofia/internal/1001@127.0.0.1", "&socket(192.168.1.2:8084 async full)"))
	if err != nil {
		return fmt.Errorf("error while sending BgApi CreateCall command to ESL client: %s", err)
	}

	return nil
}

func (f Freeswitch) ModifyCall(call Call) error {
	err := client.BgApi(fmt.Sprintf("originate %s %s", "sofia/internal/1001@127.0.0.1", "&socket(192.168.1.2:8084 async full)"))
	if err != nil {
		return fmt.Errorf("error while sending BgApi CreateCall command to ESL client: %s", err)
	}

	return nil
}
