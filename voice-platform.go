package main

type Call struct {
	sid      string
	uuid     string
	channel  string
	fromName string
	from     string
	to       string
	timeout  int
	user     string
	scenario *Scenario
}

type Scenario []Action

// type Scenario struct {
// 	scenario []Action
// }

type Action struct {
	command string
	params  map[string]string
	count   int
}

type IVoicePlatform interface {
	Connect(params interface{})
	CreateCall(call Call) error
	ModifyCall(call Call) error
}
