package interfaces

import (
	"github.com/gin-gonic/gin"
)

//Flighter contract
type Flighter interface {
	Search(c *gin.Context)
}
