package main

import (
	"fmt"
	"net/http"
	"strings"
)

func index(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	var nama []string = query["nama"]
	namatengah := query.Get("namatengah")
	fmt.Fprintf(w, strings.Join(nama, " ")+namatengah)

}

func images(w http.ResponseWriter, r *http.Request) {
	contetType := r.Header.Get("content-type")
	w.Header().Add("X-Powered-By", "trainocate")
	cookie := new(http.Cookie)
	cookie.Name = "test-cookie"
	cookie.Value = "arif"
	cookie.Path = "/"
	http.SetCookie(w, cookie)

	fmt.Fprintf(w, contetType)
}
func test(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		panic(err)
	}
	nama := r.PostForm.Get("Nama")
	if nama == "" {
		w.WriteHeader(400)
		fmt.Fprint(w, "data tidak ada")
	}
	fmt.Fprintf(w, nama)
}

func main() {
	directory := http.Dir("./resources")
	fileServer := http.FileServer(directory)

	http.Handle("/static/", http.StripPrefix("/static", fileServer))

	http.HandleFunc("/index", index)
	http.HandleFunc("/images/", images)
	http.HandleFunc("/images/test", test)
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
