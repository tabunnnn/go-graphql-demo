package provider

import (
	"github.com/gin-gonic/gin"

	"github.com/lucas-clemente/quic-go/http3"
)

// Start start quic server
func Start() (err error) {

	router := gin.New()

	router.Use(setHeader)

	// graphql GET handler
	router.GET("/graphql", gqlProvider.Get)
	router.POST("/graphql", gqlProvider.Post)

	return http3.ListenAndServe("localhost:8088", "./cert.pem", "./key.pem", router)
}

// setHeader set Header
func setHeader(c *gin.Context) {
	c.Header("Alternate-Protocol", "quic:8088")
	c.Next()
}
