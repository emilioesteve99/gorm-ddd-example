package commonHttpControllers

import "github.com/gin-gonic/gin"

type InsertOneController interface {
	InsertOne(c *gin.Context)
}
