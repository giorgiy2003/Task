package main

import (
	Repository "myapp/internal/repository"
	Handler "myapp/internal/handlers"
	"net/http"
	"log"
)

func main() {
	if err := Repository.OpenTable(); err != nil {
		log.Fatalln("Не удалось подключиться к базе данных!",err)
	}
	http.HandleFunc("/", Handler.MainForm)
	http.HandleFunc("/GetAllInfo", Handler.GetAllInfo)
	http.HandleFunc("/Add", Handler.AddPerson)
	http.HandleFunc("/form_handler_Create", Handler.Form_handler_Create)
	http.HandleFunc("/Get_By_Id", Handler.Get_By_Id)
	http.HandleFunc("/form_handler_GetById", Handler.Form_handler_GetById)
	http.HandleFunc("/Delete_By_Id", Handler.Remove)
	http.HandleFunc("/form_handler_DeleteById", Handler.Form_handler_Delete_By_Id)
	http.HandleFunc("/Update_By_Id", Handler.Edit)
	http.HandleFunc("/form_handler_Update", Handler.Form_handler_Update_By_Id)
	http.ListenAndServe(":8080", nil)
}