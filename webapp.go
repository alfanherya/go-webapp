package webapp

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("listening....")
	err := http.ListenAndServe(GetPort(), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello. This is our first Go web App on Heroku!")

}

//get port From the environment so we can run on heroku

func GetPort() string {
	var port = os.Getenv("PORT")
	//Set Default port if there is nothing in the environment
	if port == "" {
		port = "4747"
		fmt.Println("INFO: no port environment variable detected, default to " + port)

	}
	return ":" + port
}
