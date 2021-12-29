package main
import (
Routes "./routes"
    "fmt"
    "log"
    "net/http"
)

// Main function
func main() {

 fmt.Println(" \n Starting Server at 8080 ...... \n \n")
    // serve the app
    r := Routes.SetupRouter()
	//running
	r.Run()
   
    log.Fatal(http.ListenAndServe(":8000", r))
}

