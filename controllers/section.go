package controllers

import (
Models "../models"
    "time"
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)


func section_infos(event Models.SectionEvent) Models.Section {

var section Models.Section
section.EventType = event.EventType
section.SectionName = event.SectionName
section.SectionStartTime = time.Unix(event.SectionStartTime,0).UTC()
section.SectionEndTime = time.Unix(event.SectionEndTime,0).UTC()
section.TotalSectionTime = time.Unix(event.TotalSectionTime,0).UTC()
return section
}

func ParseSection(c *gin.Context) {
	var event Models.SectionEvent
	var section Models.Section
	//var entry string
	err:= c.BindJSON(&event)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} 
	
	section = section_infos(event)
	c.JSON(http.StatusOK, section)
	write_section_event(section)
}
