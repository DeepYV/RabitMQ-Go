package queue

import (
	"fmt"
)

var pas []string

func (con *Connect) Consumer() ([]string, error) {
	con.ConnectQ()

	q, err := con.ch.QueueDeclare(
		"msg", // name
		false, // durable
		false, // delete when unused
		false, // exclusiveW
		false, // no-wait
		nil,   // arguments
	)

	if err != nil {

		return []string{}, err
	}

	msgs, err := con.ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {

		return []string{}, err
	}

	var msg []string

	go func() {
		for d := range msgs {
			msg = append(msg, string(d.Body))

		}

	}()

	fmt.Println(msg)
	defer con.conn.Close()
	return msg, nil
}
