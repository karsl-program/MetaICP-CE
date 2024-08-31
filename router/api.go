package router

import (
	"metaicp/responses"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", responses.IndexNode)

	router.GET("/join", responses.JoinNode)

	router.POST("/join", responses.VerifyNode)

	router.POST("/submit", responses.SubmitNode)

	router.POST("/admin", responses.AdminNode)

	router.POST("/admin/allow", responses.AdminAllowNode)

	router.POST("/admin/ban", responses.AdminBanNode)

	router.GET("/about", responses.AboutNode)

	router.GET("/login", responses.LoginNode)

	router.GET("/select/:id", responses.SelectNode)

	router.GET("/assets/:static", responses.StaticNode)

	return router
}
