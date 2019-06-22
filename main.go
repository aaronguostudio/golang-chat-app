package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type H struct {
	Code int `json:"code"`
	Msg  string `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Resp (w http.ResponseWriter,code int, data interface{}, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	h := H {
		Code: code,
		Msg: msg,
		Data: data,
	}

	ret, err := json.Marshal(h)
	if err != nil {
		log.Println(err.Error())
	}

	w.Write(ret)
}

func userLogin (writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	mobile := request.PostForm.Get("mobile")
	passwd := request.PostForm.Get("passwd")

	loginOk := false
	str := ""
	code := 0
	data := make(map[string]interface{})

	if mobile == "111" && passwd == "111" {
		loginOk = true
	}

	if !loginOk {
		code = -1
		str = "Password is not valid"
		Resp(writer, code, nil, str)
	} else {
		data["id"] = 1
		data["code"] = code
		data["token"] = "token"
		Resp(writer, code, data, str)
	}
}

func RegisterView () {
	// parse template
	tpl, err := template.ParseGlob("view/**/*")
	if err != nil {
		// print and exit
		log.Fatal(err.Error())
	}

	for _, v := range tpl.Templates() {
		tplname := v.Name();

		fmt.Println("HandleFunc     " + tplname)
		http.HandleFunc(tplname, func(writer http.ResponseWriter, request *http.Request) {
			err := tpl.ExecuteTemplate(writer, tplname,nil)
			if err != nil {
				log.Fatal(err.Error())
			}
		})
	}
}

func main () {
	
	// Bind request and response function

	http.HandleFunc("/user/login", userLogin)

	http.Handle("/asset/", http.FileServer(http.Dir(".")))

	RegisterView()

	// Start a server
	http.ListenAndServe(":8080", nil)

}