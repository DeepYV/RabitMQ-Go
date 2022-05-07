package queue

import (
	"deepak/api/request"

	"github.com/streadway/amqp"
)

func (con *Connect) Producer(body request.Msg) error {
	con.ConnectQ()

	err1 := con.ch.Publish(
		"",         // exchange
		con.q.Name, // routing key
		false,      // mandatory
		false,      // immediatec
		amqp.Publishing{

			ContentType: "text/plain",
			Body:        []byte(body.Message),
		})
	defer con.conn.Close()

	if err1 != nil {

		return err1
	}

	return nil
}
