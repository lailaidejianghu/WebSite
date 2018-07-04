package main

import (
	"html/template"
	"log"
	"net/http"
)

var (
	tpl     *template.Template
	routers = make(map[string]func(http.ResponseWriter, *http.Request))
)

type PageInfo struct {
	Title string
	Name  string
}

func init() {
	tpl = template.Must(template.ParseGlob("../templates/*html"))

	//Add request path to here
	routers["/"] = home
	routers["/signup"] = signUp
	routers["/regist"] = regist
	routers["/passwordgetback"] = passwordgetback
	routers["/aboutus"] = TemplateWrap("aboutus.html", &PageInfo{"AboutUs", "AboutUs"})
	routers["/contactwithus"] = TemplateWrap("contactwithus.html", &PageInfo{"ContactWithUs", "ContactWithUs"})
}

func main() {

	for k, v := range routers {
		http.HandleFunc(k, v)
	}

	http.Handle("/asserts/", http.StripPrefix("/asserts/", http.FileServer(http.Dir("../asserts"))))
	http.ListenAndServe(":8000", nil)
}

func signUp(w http.ResponseWriter, r *http.Request) {
	if "GET" == r.Method {
		err := tpl.ExecuteTemplate(w, "signup.html", PageInfo{Title: "SignUp", Name: "SignUp"})
		errCheck(w, err)
	}
	if "POST" == r.Method {
		r.ParseForm()
		user := r.FormValue("user")
		pwd := r.FormValue("pwd")

		userCheck(user)
		pwdCheck(pwd)
	}

}
func regist(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "regist.html", PageInfo{"Regist", "Regist"})
	errCheck(w, err)
}
func home(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "home.html", PageInfo{"Home", "Home"})
	errCheck(w, err)
}

func passwordgetback(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "passwordgetback.html", PageInfo{"PasswordGetback", "PasswordGetback"})
	errCheck(w, err)
}

func errCheck(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, "the page is not exist! :", 404)
		log.Fatal(err.Error())
	}
}

func TemplateWrap(templateName string, info *PageInfo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := tpl.ExecuteTemplate(w, templateName, *info)
		errCheck(w, err)
	}
}

//check the user is valid or not
//if user is invalid,return 0
//if user is email ,return 1
//if user is phone number ,return 2
func userCheck(usr string) int {
	return 0
}

//check the password is valid or not
//if hte password is valid ,return true,else return false
func pwdCheck(pwd string) bool {
	return false
}
