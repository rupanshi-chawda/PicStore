package main


import (
	"fmt"
	"net/http"
)
// SimpleServer is ...
func SimpleServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s", r.URL.Path[1:])
}

func main() {
	// http.HandleFunc("/", SimpleServer)
	// http.ListenAndServe(":3000", nil)
	//Connect()
	
	//DownloadFile(filename)
}
