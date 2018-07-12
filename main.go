package main

import (
	"fmt"
	"html/template"
	"net/http"
)

//переменные нужно будет перенести в файл конфигурации в JSON-структуре
var addr = ":3000"

//Загрузка html шаблонов
func pageTemplate(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "index", nil)
}

//основная функция
func main() {
	fmt.Println("Сервер запущен, прослушивается порт", addr)
	http.HandleFunc("/", pageTemplate)
	http.ListenAndServe(addr, nil)
}
