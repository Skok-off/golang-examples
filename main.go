package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil" //для чтения файла
	"net/http"
)

//чтение файла конфигурации
func readConfig(fileName string) (string, string, string) {
	cf, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Ошибка чтения файла")
	}
	var config map[string]string
	err2 := json.Unmarshal([]byte(cf), &config)
	if err2 != nil {
		fmt.Println(err)
	}
	return config["addr"], config["login"], config["password"]
}

//Загрузка html шаблона
func pageTemplate(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "index", nil)
}

//страница авторизации
func autorizationPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/login.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "login", nil)
}

//основная функция
func main() {
	//переменные из файла конфигурации
	addr, login, password := readConfig("config.json")
	fmt.Println(login, password)
	fmt.Println("Сервер запущен, прослушивается порт", addr)
	http.HandleFunc("/", pageTemplate) //pageTemplate
	http.HandleFunc("/login", autorizationPage)
	http.HandleFunc("/goLogin", func(w http.ResponseWriter, r *http.Request) {
		if (login == r.FormValue("login")) && (password == r.FormValue("password")) {
			fmt.Fprintf(w, "Авторизация пройдена успешно")
		}
	})
	http.ListenAndServe(addr, nil)
}
