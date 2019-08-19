package controller_test

import (
	"bytes"
	"encoding/json"
	"github.com/restaurant/pkg/database/mysql"
	"github.com/restaurant/pkg/helpers"
	Server "github.com/restaurant/pkg/server"
	"net/http"
	"net/http/httptest"
	"testing"
)
type Login struct {
	Email string `json:"email"`
	Pass string `json:"pass"`
}


func TestCreateUser(t *testing.T) {
		//saLogin := Login{Email:"sourav241196@gmail.com",Pass:"zamorak"}
		type create struct {
			Name string `json:"name"`
			Pass string `json:"pass"`
			Email string `json:"email"`
		}

		db := mysql.NewMySqlDB("127.0.0.1","root","Zamorak1","3306","Restaurant_Test")
		//db.SetMaxOpenConns(5)
		defer db.Close()
		router := Server.NewRouter(db)
		s := router.Router()

	userdata := make([]create,5)
	login := make([]Login,3)
	login[0] = Login{
		Email: "sourav241196@gmail.com",
		Pass:  "zamorak",
	}
	login[1] = Login{
		Email: "user1@gmail.com",
		Pass:  "zamorak",
	}
	login[2] = Login{
		Email: "admin1@gmail.com",
		Pass:  "zamorak",
	}
	userdata[0] = create{
		Name:  "user2",
		Pass:  "zamorak",
		Email: "user3@gmail.com",
	}
	userdata[1] = create{
		Name:  "user2",
		Pass:  "",
		Email: "user2@gmail.com",
	}
	userdata[2] = create{
		Name:  "",
		Pass:  "zamorak",
		Email: "user2@gmail.com",
	}
	userdata[3] = create{
		Name:  "user2",
		Pass:  "zamorak",
		Email: "",
	}
	userdata[4] = create{
		Name:  "user2",
		Pass:  "zamorak",
		Email: "user1@gmail.com",
	}

	type testData struct {
		login Login
		c create
		expected int
	}
	data := make([]testData,15)
	data[0] = testData{login[0], userdata[0], 200}
	data[1] = testData{login[0],userdata[1],400}
	data[2] = testData{login[0],userdata[2],400}
	data[3] = testData{login[0],userdata[3],400}
	data[4] = testData{login[0],userdata[4],500}
	data[5] = testData{login[1],userdata[0],401}
	data[6] = testData{login[1],userdata[1],401}
	data[7] = testData{login[1],userdata[2],401}
	data[8] = testData{login[1],userdata[3],401}
	data[9] = testData{login[1],userdata[4],401}
	data[10] = testData{login[2],userdata[0],200}
	data[11] = testData{login[2],userdata[1],400}
	data[12] = testData{login[2],userdata[2],400}
	data[13] = testData{login[2],userdata[3],400}
	data[14] = testData{login[2],userdata[4],500}

	var request *http.Request
	for i,_ :=range data {
		println("i: ",i)
		if i < 5 {
			b, _ := json.Marshal(data[i].login)
			request,_ = http.NewRequest(http.MethodPost, "/superadmin/login",bytes.NewReader(b))
		} else if i > 4 && i < 10 {
			b, _ := json.Marshal(data[i].login)
			request,_ = http.NewRequest(http.MethodPost, "/user/login",bytes.NewReader(b))
		} else {
			b, _ := json.Marshal(data[i].login)
			request,_ = http.NewRequest(http.MethodPost, "/admin/login",bytes.NewReader(b))
		}
		request.Header.Set("Content-Type","application/json")
		response :=httptest.NewRecorder()
		s.ServeHTTP(response,request)
		token := response.Header().Get("token")
		user,_ := json.Marshal(data[i].c)
		request,_ = http.NewRequest(http.MethodPost, "/user/create",bytes.NewReader(user))
		request.Header.Set("token",token)
		request.Header.Set("Content-Type","application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response,request)
		helpers.AssertEqual(t,data[i].expected,response.Code,"")
		helpers.ResetDBafterCreate(db)
	}

}

func TestGetUser (t *testing.T) {

	saLogin := Login{Email:"sourav241196@gmail.com",Pass:"zamorak"}
	type get struct {
		Id	string	`json:"id"`
	}
	db := mysql.NewMySqlDB("127.0.0.1","root","Zamorak1","3306","Restaurant_Test")
	//db.SetMaxOpenConns(5)
	defer db.Close()
	router := Server.NewRouter(db)
	s := router.Router()
	b, _ := json.Marshal(saLogin)

	request,_ := http.NewRequest(http.MethodPost, "/superadmin/login",bytes.NewReader(b))
	request.Header.Set("Content-Type","application/json")
	response :=httptest.NewRecorder()
	s.ServeHTTP(response,request)
	token := response.Header().Get("token")
	userdata := make([]get,3)
	userdata[0] = get{Id:"1"}
	userdata[1] = get{Id:""}
	userdata[2] = get{Id:"one"}

	type testData struct {
		g get
		expected int
	}
	data := make([]testData,3)
	data[0] = testData{g:userdata[0],expected:200}
	data[1] = testData{g:userdata[1],expected:400}
	data[2] = testData{g:userdata[2],expected:500}
	for i,_ := range data {
		user,_ := json.Marshal(data[i].g)
		request,_ = http.NewRequest(http.MethodGet, "/user/get",bytes.NewReader(user))
		request.Header.Set("token",token)
		request.Header.Set("Content-Type","application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response,request)
		helpers.AssertEqual(t,data[i].expected,response.Code,"")
	}
}

func TestLoginUser(t *testing.T) {
	db := mysql.NewMySqlDB("127.0.0.1", "root", "Zamorak1", "3306", "Restaurant_Test")
	//db.SetMaxOpenConns(5)
	defer db.Close()
	router := Server.NewRouter(db)
	s := router.Router()
	login := make([]Login,5)
	login[0] = Login{
		Email: "user1@gmail.com",
		Pass:  "zamorak",
	}
	login[1] = Login{
		Email: "",
		Pass:  "gutthix",
	}
	login[2] = Login{
		Email: "user1@gmail.com",
		Pass:  "",
	}
	login[3] = Login{
		Email: "user1@gmail.com",
		Pass:  "gutthix",
	}
	login[4] = Login{
		Email: "user3@gmail.com",
		Pass:  "gutthix",
	}
	type testData struct {
		l Login
		expected int
	}
	data := make([]testData,5)
	data[0] = testData{login[0],200}
	data[1] = testData{login[1],http.StatusBadRequest}
	data[2] = testData{login[2],http.StatusBadRequest}
	data[3] = testData{login[3],http.StatusInternalServerError}
	data[4] = testData{login[4],http.StatusInternalServerError}
	for i,_ := range data{
		user,_ := json.Marshal(data[i].l)
		request,_ := http.NewRequest(http.MethodPost,"/user/login",bytes.NewReader(user))
		request.Header.Set("Content-Type","application/json")
		response := httptest.NewRecorder()
		s.ServeHTTP(response,request)
		helpers.AssertEqual(t,data[i].expected,response.Code,"")
	}
}

func TestUpdateUser (t *testing.T) {
	db := mysql.NewMySqlDB("127.0.0.1", "root", "Zamorak1", "3306", "Restaurant_Test")
	//db.SetMaxOpenConns(5)
	defer db.Close()
	router := Server.NewRouter(db)
	s := router.Router()
	type upd struct {
		Id     string `json:"id"`
		Flag   string `json:"flag"`
		Update string `json:"update"`
	}
	login := make([]Login, 2)
	login[0] = Login{
		Email: "sourav241196@gmail.com",
		Pass:  "zamorak",
	}
	login[1] = Login{
		Email: "admin1@gmail.com",
		Pass:  "zamorak",
	}
	userdata := make([]upd, 9)
	userdata[0] = upd{Id: "1", Flag: "0", Update: "user1"}
	userdata[1] = upd{Id: "1", Flag: "1", Update: "gutthix"}
	userdata[2] = upd{Id: "1", Flag: "2", Update: "user1@gmail.com"}
	userdata[3] = upd{Id: "", Flag: "0", Update: "user11"}
	userdata[4] = upd{Id: "1", Flag: "", Update: "user11"}
	userdata[5] = upd{Id: "1", Flag: "0", Update: ""}
	userdata[6] = upd{Id: "one", Flag: "0", Update: "user11"}
	userdata[7] = upd{Id: "1", Flag: "one", Update: "user11"}
	userdata[8] = upd{Id: "2", Flag: "0", Update: "user1"}
	type testData struct {
		Login    Login
		User     upd
		Expected int
	}
	data := make([]testData, 10)
	data[0] = testData{Login: login[1], User: userdata[0], Expected: 401}
	data[1] = testData{login[1], userdata[8], 200}
	data[2] = testData{Login: login[0], User: userdata[0], Expected: 200}
	data[3] = testData{Login: login[0], User: userdata[1], Expected: 200}
	data[4] = testData{Login: login[0], User: userdata[2], Expected: 200}
	data[5] = testData{Login: login[0], User: userdata[3], Expected: 400}
	data[6] = testData{Login: login[0], User: userdata[4], Expected: 400}
	data[7] = testData{Login: login[0], User: userdata[5], Expected: 400}
	data[8] = testData{Login: login[0], User: userdata[6], Expected: 500}
	data[9] = testData{Login: login[0], User: userdata[7], Expected: 500}

	var request *http.Request
	for i := range data {


		b, _ := json.Marshal(data[i].Login)
		if i <= 1 {
			request, _ = http.NewRequest(http.MethodPost, "/admin/login", bytes.NewReader(b))
		} else {
			request, _ = http.NewRequest(http.MethodPost, "/superadmin/login", bytes.NewReader(b))
		}
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		s.ServeHTTP(response, request)
		token := response.Header().Get("token")
		user,_ := json.Marshal(data[i].User)
		println(i)
		request,_ = http.NewRequest(http.MethodPut, "/user/update",bytes.NewReader(user))
		request.Header.Set("token",token)
		request.Header.Set("Content-Type","application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response,request)
		helpers.AssertEqual(t,data[i].Expected,response.Code,"")
	}
	helpers.ResetDBafterCreate(db)
}

func TestDelUser (t *testing.T) {
	type del struct {
		Id	string	`json:"id"`
	}
	login := Login{
		Email: "sourav241196@gmail.com",
		Pass:  "zamorak",
	}
	db := mysql.NewMySqlDB("127.0.0.1","root","Zamorak1","3306","Restaurant_Test")
	defer db.Close()
	router := Server.NewRouter(db)
	s := router.Router()
	b, _ := json.Marshal(login)

	request,_ := http.NewRequest(http.MethodPost, "/superadmin/login",bytes.NewReader(b))
	request.Header.Set("Content-Type","application/json")
	response :=httptest.NewRecorder()
	s.ServeHTTP(response,request)
	token := response.Header().Get("token")
	user:= make([]del,3)
	user[0] = del{Id:"1"}
	user[1] = del{Id:""}
	user[2] = del{Id:"one"}
	type testdata struct {
		D del
		Expected int
	}
	data := make([]testdata,3)
	data[0] = testdata{D:user[0],Expected:200}
	data[1] = testdata{D:user[1],Expected:400}
	data[2] = testdata{D:user[2],Expected:500}

	for i:= range user {
		println(i)
		user,_ := json.Marshal(data[i].D)
		request,_ := http.NewRequest(http.MethodDelete, "/user/del",bytes.NewReader(user))
		request.Header.Set("token",token)
		request.Header.Set("Content-Type","application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response,request)
		helpers.AssertEqual(t,data[i].Expected,response.Code,"")

	}
	helpers.ResetDBafterDelete(db,0)
}
func TestListUser(t *testing.T) {

	saLogin := Login{Email:"sourav241196@gmail.com",Pass:"zamorak"}

	db := mysql.NewMySqlDB("127.0.0.1","root","Zamorak1","3306","Restaurant_Test")
	defer db.Close()
	router := Server.NewRouter(db)
	s := router.Router()
	b, _ := json.Marshal(saLogin)

	request,_ := http.NewRequest(http.MethodPost, "/superadmin/login",bytes.NewReader(b))
	request.Header.Set("Content-Type","application/json")
	response :=httptest.NewRecorder()
	s.ServeHTTP(response,request)
	token := response.Header().Get("token")
	request,_ = http.NewRequest(http.MethodGet, "/user/list",nil)
	request.Header.Set("token",token)
	s.ServeHTTP(response,request)
	helpers.AssertEqual(t,200,response.Code,"")
}

func TestLogoutUser(t *testing.T){
	saLogin := Login{Email:"sourav241196@gmail.com",Pass:"zamorak"}
	db := mysql.NewMySqlDB("127.0.0.1","root","Zamorak1","3306","Restaurant_Test")
	defer db.Close()
	router := Server.NewRouter(db)
	s := router.Router()
	b, _ := json.Marshal(saLogin)

	request,_ := http.NewRequest(http.MethodPost, "/superadmin/login",bytes.NewReader(b))
	request.Header.Set("Content-Type","application/json")
	response :=httptest.NewRecorder()
	s.ServeHTTP(response,request)
	token := response.Header().Get("token")
	request,_ = http.NewRequest(http.MethodGet,"/logout",nil)
	request.Header.Set("token",token)
	response = httptest.NewRecorder()
	s.ServeHTTP(response,request)
	helpers.AssertEqual(t,200,response.Code,"")
}
