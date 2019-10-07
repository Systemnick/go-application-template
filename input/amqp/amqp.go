package amqp

import (
	"git.rnd.mtt/innovation/call-initiator/input"
	"github.com/streadway/amqp"
)

type AMQP struct {
	input.IInput
	Url        string
	connection amqp.Connection
}
