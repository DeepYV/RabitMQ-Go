package controller

import (
	"deepak/api/request"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Producer(c *gin.Context) {

	var dataFromBody request.Msg
	if err := json.NewDecoder(c.Request.Body).Decode(&dataFromBody); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return

	}

	if err := QueueService.Producer(dataFromBody); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return

	}
	c.JSON(http.StatusAccepted, "message submitted")
	return
}
