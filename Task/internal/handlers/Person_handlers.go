package Handler

import (
	"fmt"
	"html/template"
	Logic "myapp/internal/logic"
	Model "myapp/internal/model"
	"net/http"
)

func GetAllInfo(w http.ResponseWriter, r *http.Request) {
	persons, err := Logic.Read()
	if err != nil {
		fmt.Fprint(w, err)
	}
	if len(persons) == 0 {
		fmt.Fprint(w, "Записей ещё нет!")
	}
	tmpl, err := template.ParseFiles("templates/GetAllPersons.html")
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	for _, p := range persons {
		tmpl.Execute(w, p)
	}
}

func Form_handler_Create(w http.ResponseWriter, r *http.Request) {
	var newPerson Model.Person
	newPerson.Email = r.FormValue("email")
	newPerson.Phone = r.FormValue("phone")
	newPerson.FirstName = r.FormValue("firstName")
	newPerson.LastName = r.FormValue("lastName")
	err := Logic.Create(newPerson)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Form_handler_GetById(w http.ResponseWriter, r *http.Request) {

	Person_Id := r.FormValue("id")
	persons, err := Logic.ReadOne(Person_Id)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	if len(persons) == 0 {
		fmt.Fprint(w, "Записи с заданным id не существует!")
	}
	tmpl, err := template.ParseFiles("templates/GetAllPersons.html")
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	for _, p := range persons {
		tmpl.Execute(w, p)
	}
}
func Form_handler_Delete_By_Id(w http.ResponseWriter, r *http.Request) {
	Person_Id := r.FormValue("id")
	persons, err := Logic.ReadOne(Person_Id)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	if len(persons) == 0 {
		fmt.Fprint(w, "Записи с заданным id не существует!")
	}
	err = Logic.Delete(Person_Id)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Form_handler_Update_By_Id(w http.ResponseWriter, r *http.Request) {
	Person_Id := r.FormValue("id")
	persons, err := Logic.ReadOne(Person_Id)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	if len(persons) == 0 {
		fmt.Fprint(w, "Записи с заданным id не существует!")
	}

	var newPerson Model.Person
	newPerson.Email = r.FormValue("email")
	newPerson.Phone = r.FormValue("phone")
	newPerson.FirstName = r.FormValue("firstName")
	newPerson.LastName = r.FormValue("lastName")
	err = Logic.Update(newPerson, Person_Id)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func MainForm(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/mainForm.html")
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	tmpl.Execute(w, nil)
}

func AddPerson(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/CreatePerson.html")
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	tmpl.Execute(w, nil)
}

func Remove(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/DeleteById.html")
	if err != nil {
		panic(err)
	}
	tmpl.Execute(w, nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/EditPerson.html")
	if err != nil {
		panic(err)
	}
	tmpl.Execute(w, nil)
}

func Get_By_Id(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/GetByID.html")
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	tmpl.Execute(w, nil)
}
