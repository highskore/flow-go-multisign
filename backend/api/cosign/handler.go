package cosign

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lukaracki/flow-go-multisign/backend/pkg/interfaces"
)

type Req struct {
	Payload interfaces.Signable
}

// AuthHandler is the handler for the cosign endpoint
// endpoint: POST /cosign
func AuthHandler(c *gin.Context) {
	var req Req

	if err := c.BindJSON(&req); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	response, err := ProcessSignable(c, &req.Payload)

	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
