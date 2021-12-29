package controllers

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

type JsonResponse struct {
    Type    string `json:"type"`
    Data    string `json:"data"`
    Message string `json:"message"`
}

func printMessage(message string) {
    fmt.Println("")
    fmt.Println(message)
    fmt.Println("")
}

func Hello(c *gin.Context) {

    printMessage("Say Hello ...")


    var response = JsonResponse{Type: "success", Data: "Rest Api Track Event", Message:"Hello, we are available"}

    c.JSON(http.StatusOK, response)
}
