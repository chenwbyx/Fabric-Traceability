package controller

import (
	"net/http"
	"encoding/json"
	"github.com/chenwbyx/Fabric-Traceability/service"
	"fmt"
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

func (app *Application) AddEduShow(w http.ResponseWriter, r *http.Request)  {
	data := &struct {
		CurrentUser User
		Msg string
		Flag bool
	}{
		CurrentUser:cuser,
		Msg:"",
		Flag:false,
	}
	ShowView(w, r, "addEdu.html", data)
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
	/*
	edu := service.Education{
		Name:r.FormValue("name"),
		Gender:r.FormValue("gender"),
		Nation:r.FormValue("nation"),
		EntityID:r.FormValue("entityID"),
		Place:r.FormValue("place"),
		BirthDay:r.FormValue("birthDay"),
		EnrollDate:r.FormValue("enrollDate"),
		GraduationDate:r.FormValue("graduationDate"),
		SchoolName:r.FormValue("schoolName"),
		Major:r.FormValue("major"),
		QuaType:r.FormValue("quaType"),
		Length:r.FormValue("length"),
		Mode:r.FormValue("mode"),
		Level:r.FormValue("level"),
		Graduation:r.FormValue("graduation"),
		CertNo:r.FormValue("certNo"),
		Photo:r.FormValue("photo"),
		Time:time.Now().Format("2006-01-02 15:04:05"),
	}

	app.Setup.SaveEdu(edu)
	transactionID, err := app.Setup.SaveEdu(edu)

	data := &struct {
		CurrentUser User
		Msg string
		Flag bool
	}{
		CurrentUser:cuser,
		Flag:true,
		Msg:"",
	}

	if err != nil {
		data.Msg = err.Error()
	}else{
		data.Msg = "信息添加成功:" + transactionID
	}

	//ShowView(w, r, "addEdu.html", data)
	r.Form.Set("certNo", edu.CertNo)
	r.Form.Set("name", edu.Name)
	app.FindCertByNoAndName(w, r)
	*/
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

func (app *Application) FindCertByNoAndName(w http.ResponseWriter, r *http.Request)  {
	certNo := r.FormValue("certNo")
	name := r.FormValue("name")
	result, err := app.Setup.FindEduByCertNoAndName(certNo, name)
	var edu = service.Education{}
	json.Unmarshal(result, &edu)

	fmt.Println("根据证书编号与姓名查询信息成功：")
	fmt.Println(edu)

	data := &struct {
		Edu service.Education
		CurrentUser User
		Msg string
		Flag bool
		History bool
	}{
		Edu:edu,
		CurrentUser:cuser,
		Msg:"",
		Flag:false,
		History:false,
	}

	if err != nil {
		data.Msg = err.Error()
		data.Flag = true
	}

	ShowView(w, r, "queryResult.html", data)
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

	/*result, err := app.Setup.FindEduInfoByEntityID(entityID)
	var edu = service.Education{}
	json.Unmarshal(result, &edu)

	data := &struct {
		Edu service.Education
		CurrentUser User
		Msg string
		Flag bool
		History bool
	}{
		Edu:edu,
		CurrentUser:cuser,
		Msg:"",
		Flag:false,
		History:true,
	}

	if err != nil {
		data.Msg = err.Error()
		data.Flag = true
	}*/

	ShowView(w, r, "queryResult.html", data)
}

func (app *Application) ModifyShow(w http.ResponseWriter, r *http.Request)  {
	// 根据证书编号与姓名查询信息
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
	/*
	certNo := r.FormValue("Primarykey")
	name := r.FormValue("name")
	result, err := app.Setup.FindEduByCertNoAndName(certNo, name)

	var edu = service.Education{}
	json.Unmarshal(result, &edu)

	data := &struct {
		Edu service.Education
		CurrentUser User
		Msg string
		Flag bool
	}{
		Edu:edu,
		CurrentUser:cuser,
		Flag:true,
		Msg:"",
	}

	if err != nil {
		data.Msg = err.Error()
		data.Flag = true
	}
*/
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
	fmt.Println(com)
	/*
	edu := service.Education{
		Name:r.FormValue("name"),
		Gender:r.FormValue("gender"),
		Nation:r.FormValue("nation"),
		EntityID:r.FormValue("entityID"),
		Place:r.FormValue("place"),
		BirthDay:r.FormValue("birthDay"),
		EnrollDate:r.FormValue("enrollDate"),
		GraduationDate:r.FormValue("graduationDate"),
		SchoolName:r.FormValue("schoolName"),
		Major:r.FormValue("major"),
		QuaType:r.FormValue("quaType"),
		Length:r.FormValue("length"),
		Mode:r.FormValue("mode"),
		Level:r.FormValue("level"),
		Graduation:r.FormValue("graduation"),
		CertNo:r.FormValue("certNo"),
		Photo:r.FormValue("photo"),
		Time:time.Now().Format("2006-01-02 15:04:05"),
	}
	*/
	//transactionID, err := app.Setup.ModifyEdu(edu)
	app.Setup.ModifyCom(com)

	/*data := &struct {
		Edu service.Education
		CurrentUser User
		Msg string
		Flag bool
	}{
		CurrentUser:cuser,
		Flag:true,
		Msg:"",
	}

	if err != nil {
		data.Msg = err.Error()
	}else{
		data.Msg = "新信息添加成功:" + transactionID
	}

	ShowView(w, r, "modify.html", data)
	*/

	r.Form.Set("entityID", com.Primarykey)
	app.FindByID(w, r)
}
