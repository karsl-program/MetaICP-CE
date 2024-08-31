package responses

import "github.com/gin-gonic/gin"

func StaticNode(c *gin.Context) {
	c.File("assets/" + c.Param("static"))
}
