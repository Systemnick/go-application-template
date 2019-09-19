package input

type InputForm struct {
	CallSid string
	From string
	To string
	Channel  string
	Timeout  int
	TraceId string
	Modify string
	Preserve string
	scenario *Scenario
}

type Scenario []Action

type Action struct {
	Command string
	Params  map[string]string
	Count   int
}
