package controllers

import (
Models "../models"
    "time"
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)


func page_infos(event Models.PageEvent) Models.Page {

var page Models.Page
page.EventType = event.EventType
page.PageName = event.PageName
page.PageStartTime = time.Unix(event.PageStartTime,0).UTC()
page.PageEndTime = time.Unix(event.PageEndTime,0).UTC()
page.TotalPageTime = time.Unix(event.TotalPageTime,0).UTC()
return page
}

func ParsePage(c *gin.Context) {
	var event Models.PageEvent
	var page Models.Page
	//var entry string
	err:= c.BindJSON(&event)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} 
	
	page = page_infos(event)
	c.JSON(http.StatusOK, page)
	write_page_event(page)
}
