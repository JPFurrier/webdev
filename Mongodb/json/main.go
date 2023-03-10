package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"webdev/Mongodb/json/models"
)

func main() {
	r := httprouter.New()
	r.GET("/", index)

	//add route plus parameter
	r.GET("/user/:id", getUser)
	http.ListenAndServe("localhost:8080", r)
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	s := `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Index</title>
</head>
<body>
<a href="/user/9872309847">GO TO: http://localhost:8080/user/9872309847</a>
</body>
</html>`

	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s))
}

func getUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   "James Bond",
		Gender: "male",
		Age:    32,
		Id:     p.ByName("id"),
	}

	//Marshall into JSON
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	//Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) //200
	fmt.Fprintln(w, "%s\n", uj)
}
