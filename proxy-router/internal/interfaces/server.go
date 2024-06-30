package interfaces

import "github.com/gin-gonic/gin"

type Router interface {
	GET(uri string, handl ...gin.HandlerFunc) gin.IRoutes
	POST(uri string, handl ...gin.HandlerFunc) gin.IRoutes
	Use(middleware ...gin.HandlerFunc) gin.IRoutes
}
