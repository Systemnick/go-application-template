package input

type CallType int

const (
	ExistingCall CallType = iota
	NewCall
	AddParticipant
)

type InForm struct {
	CallType CallType  `json:"call_type"`
	CallSid  string    `json:"call_sid"`
	From     string    `json:"from"`
	To       string    `json:"to"`
	Channel  string    `json:"channel"`
	Timeout  int       `json:"timeout"`
	TraceId  string    `json:"trace_id"`
	Tag      string    `json:"tag"`
	Modify   string    `json:"modify"`
	Preserve string    `json:"preserve"`
	Scenario *Scenario `json:"scenario"`
}

type Scenario []Action

type Action struct {
	Command string
	Params  map[string]string
	Count   int
}
