package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Consumer(c *gin.Context) {

	response, er := QueueService.Consumer()

	if er != nil {
		c.JSON(http.StatusInternalServerError, er)
		return

	}
	c.JSON(http.StatusAccepted, response)

}

// rabbitmq ui ands
