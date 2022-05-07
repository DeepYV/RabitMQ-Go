package queue

import (
	"deepak/api/request"
	"deepak/api/util"

	"github.com/streadway/amqp"
)

type IConnect interface {
	ConnectQ()
	Producer(body request.Msg) error
	Consumer() ([]string, error)
}

type Connect struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	q    amqp.Queue
}

func (con *Connect) ConnectQ() {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	util.Error(err, "Failed to connect to RabbitMQ")
	//defer conn.Close()

	ch, err := conn.Channel()
	util.Error(err, "Failed to open a channel")
	// defer ch.Close()

	q, err := ch.QueueDeclare(
		"msg", // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	con.conn = conn
	con.ch = ch
	con.q = q

}
