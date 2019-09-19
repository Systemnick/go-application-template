package voice_platform

type Call struct {
	sid      string
	from     string
	to       string
	scenario *Scenario
}

type Scenario []Action

type Action struct {
	command string
	params  map[string]string
	count   int
}

type Params struct {
	Events chan string
}

type IVoicePlatform interface {
	Connect(params Params) error
	CreateCall(call Call) error
	ModifyCall(call Call) error
}

type VoicePlatform struct {
	IVoicePlatform
}
