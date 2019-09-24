package voice_platform

type Call struct {
	Sid      string
	From     string
	To       string
	Scenario *Scenario
}

type Scenario []Action

type Action struct {
	Command string
	Params  map[string]string
	Count   int
}

type Params struct {
	Events chan string
}

type IVoicePlatform interface {
	Connect(params Params) error
	CreateCall(call Call) error
	ModifyCall(call Call) error

	AddParticipant(call Call) error
}

type VoicePlatform struct {
	IVoicePlatform
}
