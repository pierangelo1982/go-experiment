package main

import "net/http"

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Benvenuto nella mia HomePage"))
}
func main() {
	http.HandleFunc("/", homePage)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
