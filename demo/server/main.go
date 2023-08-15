package server

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func login(w http.ResponseWriter, r *http.Request) {
	var fileName = "login.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("Error when parsing file", err)
		return
	}
	err = t.ExecuteTemplate(w, "login", "Please log in")
	if err != nil {
		fmt.Println("Error when executing templates")
		return
	}
}

var userDB = map[string]string{
	"Wallace": "goodPassword",
}

func loginSubmit(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if userDB[username] == password {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "You've now logged in")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Didn't find mathcing credential")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/login": //todo: input boxes & button for credentials
		login(w, r)
	case "/login-submit":
		loginSubmit(w, r)
	case "function":
		fileName := "function.html"
		funcMap := map[string]any{
			"upper": strings.ToUpper,
		}
		t, err := template.New(fileName).Funcs(funcMap).ParseFiles(fileName)
		if err != nil {
			fmt.Println(err)
			return
		}
		t.ExecuteTemplate(w, fileName, "Hello Ninjas")
	default:
		fmt.Fprintf(w, "Sup Ninjas")
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("", nil)
	// http.ListenAndServeTLS("", "cert.pem", "key.pem", nil)
	//go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host=localhost
}
