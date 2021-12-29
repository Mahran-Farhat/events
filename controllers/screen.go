package controllers

import (
Models "../models"
    "time"
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)


func screen_infos(event Models.ScreenEvent) Models.Screen {

var screen Models.Screen
screen.EventType = event.EventType
screen.ScreenName = event.ScreenName
screen.ScreenStartTime = time.Unix(event.ScreenStartTime,0).UTC()
screen.ScreenEndTime = time.Unix(event.ScreenEndTime,0).UTC()
screen.TotalScreenTime = time.Unix(event.TotalScreenTime,0).UTC()
return screen
}

func ParseScreen(c *gin.Context) {
	var event Models.ScreenEvent
	var screen Models.Screen
	//var entry string
	err:= c.BindJSON(&event)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} 
	
	screen = screen_infos(event)
	c.JSON(http.StatusOK, screen)
	write_screen_event(screen)
}
