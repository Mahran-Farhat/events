package routes

import (
Controller "../controllers"
    "github.com/gin-gonic/gin"
)


//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/events-api")
	{
		grp1.GET("hello", Controller.Hello)
		grp1.POST("event-screen", Controller.ParseScreen)
		grp1.POST("section-screen", Controller.ParseSection)
		grp1.POST("page-screen", Controller.ParsePage)
		
	}

	return r
}
