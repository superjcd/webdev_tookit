package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BadRequest(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"msg": "bad request, details: " + err.Error()})
}

func InternalServerError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{"msg": "internal server error, details: " + err.Error()})
}

func Unauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{"msg": "unauthorized"})
}

func CreateOneOk(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"msg": "one created"})
}

func CreateManyOk(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"msg": "many created"})
}

func FetchOneOk(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{"msg": "ok", "data": data})
}

func FetchManyOk(c *gin.Context, datas any) {
	c.JSON(http.StatusOK, gin.H{"msg": "ok", "datas": datas})
}
