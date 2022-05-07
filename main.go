package main

import (
	"deepak/api/queue"

	"deepak/api/controller"

	"github.com/gin-gonic/gin"
)

func init() {
	controller.QueueService = &queue.Connect{}

	r := gin.Default()
	r.POST("Producer", controller.Producer)

	r.GET("Consumer", controller.Consumer)
	r.Run("localhost:8010")
}

func main() {

}

/* func (con Connect) Consumer(c *gin.Context) {

	q, err := con.ch.QueueDeclare(
		"msg", // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)

		return
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
		c.AbortWithError(http.StatusBadRequest, err)

		return
	}

	for d := range msgs {

		log.Printf("message: %s", d.Body)

	}

}
*/
// u and p -> true -> tokken -->acces t , refresh token , acces tokkenn exp time , refresh bcrypt or  rsa
