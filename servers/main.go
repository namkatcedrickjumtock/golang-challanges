package main

import "net/http"

func examplesResponse(w http.ResponseWriter, r *http.Request) {

	// check if route exist
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Example response 1 caming from server\n"))
}
func examplesResponse1(w http.ResponseWriter, r *http.Request) {

	// setting the header return type to json
	w.Header().Set("content-Type", "application/json; charset=utf-8")

	// accepting diff content-type from server

	// r.Header.Get("Accept")
	// ** use a switch case to check for :
		// 1 - case  "application/json" and set the header content type 
		// 1 - case  "application/xml" and set the header content type 
		// 1 - case  "application/html" and set the header content type 
	w.Write([]byte("response from server 2\n"))
}
func examplesResponse2(w http.ResponseWriter, r *http.Request) {

	// responding to diff request from client
	w.Write([]byte("response from server4\n"))
}
func main() {
	http.HandleFunc("/", examplesResponse)
	http.HandleFunc("/example2", examplesResponse1)
	http.HandleFunc("/example3", examplesResponse2)

	http.ListenAndServe(":8000", nil)
}
