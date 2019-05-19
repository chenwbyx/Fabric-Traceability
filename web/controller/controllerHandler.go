package controller

import (
	"net/http"
	"encoding/json"
	"github.com/chenwbyx/Fabric-Traceability/service"
	"time"
)

var cuser User

func (app *Application) LoginView(w http.ResponseWriter, r *http.Request)  {

	ShowView(w, r, "login.html", nil)
}

func (app *Application) Index(w http.ResponseWriter, r *http.Request)  {
	data := &struct {
		Flag bool
	}{
		Flag:false,
	}
	ShowView(w, r, "index.html", data)
}

func (app *Application) Help(w http.ResponseWriter, r *http.Request)  {
	data := &struct {
		CurrentUser User
	}{
		CurrentUser:cuser,
	}
	ShowView(w, r, "help.html", data)
}

func (app *Application) Login(w http.ResponseWriter, r *http.Request) {
	loginName := r.FormValue("loginName")
	password := r.FormValue("password")

	var flag bool
	for _, user := range users {
		if user.LoginName == loginName && user.Password == password {
			cuser = user
			flag = true
			break
		}
	}

	data := &struct {
		CurrentUser User
		Flag bool
	}{
		CurrentUser:cuser,
		Flag:false,
	}

	if flag {
		// 登录成功
		ShowView(w, r, "index.html", data)
	}else{
		// 登录失败
		data.Flag = true
		data.CurrentUser.LoginName = loginName
		ShowView(w, r, "login.html", data)
	}
}

func (app *Application) LoginOut(w http.ResponseWriter, r *http.Request)  {
	cuser = User{}
	ShowView(w, r, "login.html", nil)
}


func (app *Application) AddEdu(w http.ResponseWriter, r *http.Request)  {

	com := service.Commodity{
		Type:r.FormValue("docType"),
		Primarykey:r.FormValue("primarykey"),
		Name:r.FormValue("name"),
		Des:r.FormValue("des"),
		Specification:r.FormValue("specification"),
		Source:r.FormValue("source"),
		Machining:r.FormValue("machining"),
		Remarks:r.FormValue("remarks"),
		Principal:r.FormValue("principal"),
		PhoneNumber:r.FormValue("phoneNumber"),
		Photo:r.FormValue("photo"),

		ShelfLife:r.FormValue("shelfLife"),
		StorageMethod:r.FormValue("storageMethod"),
		Brand:r.FormValue("brand"),
		Vendor:r.FormValue("vendor"),
		PlaceOfProduction:r.FormValue("placeOfProduction"),
		ExecutiveStandard:r.FormValue("executiveStandard"),
		Time:time.Now().Format("2006-01-02 15:04:05"),
	}

	app.Setup.SaveCom(com)

	r.Form.Set("entityID", com.Primarykey)
	app.FindByID(w, r)
}

func (app *Application) QueryPage(w http.ResponseWriter, r *http.Request)  {
	data := &struct {
		CurrentUser User
		Msg string
		Flag bool
	}{
		CurrentUser:cuser,
		Msg:"",
		Flag:false,
	}
	ShowView(w, r, "query.html", data)
}

func (app *Application) QueryPage2(w http.ResponseWriter, r *http.Request)  {
	data := &struct {
		CurrentUser User
		Msg string
		Flag bool
	}{
		CurrentUser:cuser,
		Msg:"",
		Flag:false,
	}
	ShowView(w, r, "query2.html", data)
}

func (app *Application) FindByID(w http.ResponseWriter, r *http.Request)  {
	entityID := r.FormValue("entityID")
	result, err := app.Setup.FindComInfoByEntityID(entityID)
	var com = service.Commodity{}
	json.Unmarshal(result, &com)

	data := &struct {
		Com service.Commodity
		CurrentUser User
		Msg string
		Flag bool
		History bool
	}{
		Com:com,
		CurrentUser:cuser,
		Msg:"",
		Flag:false,
		History:true,
	}

	if err != nil {
		data.Msg = err.Error()
		data.Flag = true
	}

	ShowView(w, r, "queryResult.html", data)
}

func (app *Application) ModifyShow(w http.ResponseWriter, r *http.Request)  {
	entityID := r.FormValue("entityID")
	result, err := app.Setup.FindComInfoByEntityID(entityID)
	var com = service.Commodity{}
	json.Unmarshal(result, &com)

	data := &struct {
		Com service.Commodity
		CurrentUser User
		Msg string
		Flag bool
	}{
		Com:com,
		CurrentUser:cuser,
		Flag:true,
		Msg:"",
	}

	if err != nil {
		data.Msg = err.Error()
		data.Flag = true
	}

	ShowView(w, r, "modify.html", data)
}

func (app *Application) Modify(w http.ResponseWriter, r *http.Request) {
	com := service.Commodity{
		Type:r.FormValue("docType"),
		Primarykey:r.FormValue("primarykey"),
		Name:r.FormValue("name"),
		Des:r.FormValue("des"),
		Specification:r.FormValue("specification"),
		Source:r.FormValue("source"),
		Machining:r.FormValue("machining"),
		Remarks:r.FormValue("remarks"),
		Principal:r.FormValue("principal"),
		PhoneNumber:r.FormValue("phoneNumber"),
		Photo:r.FormValue("photo"),

		ShelfLife:r.FormValue("shelfLife"),
		StorageMethod:r.FormValue("storageMethod"),
		Brand:r.FormValue("brand"),
		Vendor:r.FormValue("vendor"),
		PlaceOfProduction:r.FormValue("placeOfProduction"),
		ExecutiveStandard:r.FormValue("executiveStandard"),
		Time:time.Now().Format("2006-01-02 15:04:05"),
	}

	app.Setup.ModifyCom(com)


	r.Form.Set("entityID", com.Primarykey)
	app.FindByID(w, r)
}
