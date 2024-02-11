package routes

import (
    "github.com/gin-gonic/gin"
    "structure/pkg/controllers"
	"structure/pkg/api/middleware"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    // Define routes
    organization := r.Group("/organization")
    {
		organization.Use(middleware.AuthMiddleware())
		organization.GET("", controllers.GetAllOrganizations)
        organization.POST("", controllers.CreateOrganization)
        organization.GET("/:organization_id", controllers.GetOrganizationByID)
		organization.PUT("/:organization_id", controllers.UpdateOrganization)
		organization.DELETE("/:organization_id", controllers.DeleteOrganization)
		organization.POST("/:organization_id/invite", controllers.InviteUserToOrganization)
    }

    r.POST("/signup", controllers.SignUp)
    r.POST("/signin", controllers.SignIn)
    r.POST("/refresh-token", controllers.RefreshToken)

    return r
}