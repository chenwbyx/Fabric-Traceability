package web

import (
	"net/http"
	"fmt"
	"github.com/chenwbyx/Fabric-Traceability/web/controller"
)


// 启动Web服务并指定路由信息
func WebStart(app controller.Application)  {

	fs:= http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// 指定路由信息(匹配请求)
	http.HandleFunc("/admin", app.LoginView)
	http.HandleFunc("/login", app.Login)
	http.HandleFunc("/loginout", app.LoginOut)

	http.HandleFunc("/", app.Index)
	http.HandleFunc("/help", app.Help)

	http.HandleFunc("/addEduInfo", app.AddEduShow)
	http.HandleFunc("/addEdu", app.AddEdu)

	http.HandleFunc("/queryPage", app.QueryPage)
	//http.HandleFunc("/query", app.FindCertByNoAndName)

	http.HandleFunc("/queryPage2", app.QueryPage2)
	http.HandleFunc("/query2", app.FindByID)


	http.HandleFunc("/modifyPage", app.ModifyShow)
	http.HandleFunc("/modify", app.Modify)

	http.HandleFunc("/upload", app.UploadFile)

	fmt.Println("启动Web服务, 监听端口号为: 8001")
	err := http.ListenAndServe("localhost:8001", nil)
	if err != nil {
		fmt.Printf("Web服务启动失败: %v", err)
	}

}



