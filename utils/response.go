package utils

import (
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Result interface{} `json:"result,omitempty"`
}

func (g *Gin) Response(httpCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Result: data,
	})

	return
}
