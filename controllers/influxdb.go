package controllers

import (
	Models "../models"
	_"encoding/json"
	_"fmt"
	"github.com/influxdata/influxdb1-client/v2"
	"log"
	_"math/rand"
	_"time"
)

const (
	database = "events"
	username = "admin"
	password = "password"
)


func influxDBClient() client.Client {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://localhost:8086",
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	return c
}

func write_screen_event(screen Models.Screen) {
c := influxDBClient()
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  database,
		Precision: "s",
	})

	if err != nil {
		log.Fatalln("Error: ", err)
	}

            tags := map[string]string{
            				"eventType":  screen.EventType,
                            "screenName": screen.ScreenName,
            			}
			fields := map[string]interface{}{
				"screenStartTime":  screen.ScreenStartTime,
                "screenEndTime": screen.ScreenEndTime,
                "totalScreenTime": screen.TotalScreenTime,
			}

			point, err := client.NewPoint(
				"screen_track",
				tags,
				fields,
			)
			if err != nil {
				log.Fatalln("Error: ", err)
			}

			bp.AddPoint(point)


	err = c.Write(bp)
	if err != nil {
		log.Fatal(err)
	}
}

func write_page_event(page Models.Page) {
c := influxDBClient()
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  database,
		Precision: "s",
	})

	if err != nil {
		log.Fatalln("Error: ", err)
	}

            tags := map[string]string{
            				"eventType":  page.EventType,
                            "pageName": page.PageName,
            			}
			fields := map[string]interface{}{
				"pageStartTime":  page.PageStartTime,
                "pageEndTime": page.PageEndTime,
                "totalPageTime": page.TotalPageTime,
			}

			point, err := client.NewPoint(
				"page_track",
				tags,
				fields,
			)
			if err != nil {
				log.Fatalln("Error: ", err)
			}

			bp.AddPoint(point)


	err = c.Write(bp)
	if err != nil {
		log.Fatal(err)
	}
}

func write_section_event(section Models.Section) {
c := influxDBClient()
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  database,
		Precision: "s",
	})

	if err != nil {
		log.Fatalln("Error: ", err)
	}

            tags := map[string]string{
            				"eventType":  section.EventType,
                            "sectionName": section.SectionName,
            			}
			fields := map[string]interface{}{
				"sectionStartTime":  section.SectionStartTime,
                "sectionEndTime": section.SectionEndTime,
                "totalSectionTime": section.TotalSectionTime,
			}

			point, err := client.NewPoint(
				"section_track",
				tags,
				fields,
			)
			if err != nil {
				log.Fatalln("Error: ", err)
			}

			bp.AddPoint(point)


	err = c.Write(bp)
	if err != nil {
		log.Fatal(err)
	}
}

