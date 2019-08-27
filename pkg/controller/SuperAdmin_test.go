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

func TestCreatesuperadmin(t *testing.T) {
	//saLogin := Login{Email:"sourav241196@gmail.com",Pass:"zamorak"}
	type create struct {
		Name  string `json:"name"`
		Pass  string `json:"pass"`
		Email string `json:"email"`
	}

	db := mysql.NewMySqlDB("127.0.0.1", "root", "Zamorak1", "3306", "Restaurant_Test")
	//db.SetMaxOpenConns(5)
	defer db.Close()
	router := Server.NewRouter(db)
	s := router.Router()

	superadmindata := make([]create, 5)
	login := Login{
		Email: "sourav241196@gmail.com",
		Pass:  "zamorak",
	}
	superadmindata[0] = create{
		Name:  "superadmin2",
		Pass:  "zamorak",
		Email: "superadmin2@gmail.com",
	}
	superadmindata[1] = create{
		Name:  "superadmin2",
		Pass:  "",
		Email: "superadmin2@gmail.com",
	}
	superadmindata[2] = create{
		Name:  "",
		Pass:  "zamorak",
		Email: "superadmin2@gmail.com",
	}
	superadmindata[3] = create{
		Name:  "superadmin2",
		Pass:  "zamorak",
		Email: "",
	}
	superadmindata[4] = create{
		Name:  "superadmin2",
		Pass:  "zamorak",
		Email: "sourav241196@gmail.com",
	}

	type testData struct {
		login    Login
		c        create
		expected int
	}
	data := make([]testData, 5)
	data[0] = testData{login, superadmindata[0], 200}
	data[1] = testData{login, superadmindata[1], 400}
	data[2] = testData{login, superadmindata[2], 400}
	data[3] = testData{login, superadmindata[3], 400}
	data[4] = testData{login, superadmindata[4], 500}

	var request *http.Request
	for i, _ := range data {
		println(i)

		if i < 5 {
			b, _ := json.Marshal(data[i].login)
			request, _ = http.NewRequest(http.MethodPost, "/superadmin/login", bytes.NewReader(b))
		} else {
			b, _ := json.Marshal(data[i].login)
			request, _ = http.NewRequest(http.MethodPost, "/admin/login", bytes.NewReader(b))
		}
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		s.ServeHTTP(response, request)
		token := response.Header().Get("token")
		superadmin, _ := json.Marshal(data[i].c)
		request, _ = http.NewRequest(http.MethodPost, "/superadmin/create", bytes.NewReader(superadmin))
		request.Header.Set("token", token)
		request.Header.Set("Content-Type", "application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response, request)
		helpers.AssertEqual(t, data[i].expected, response.Code, "")
		helpers.ResetDBafterCreate(db)
	}

}

func TestGetsuperadmin(t *testing.T) {

	saLogin := Login{Email: "sourav241196@gmail.com", Pass: "zamorak"}
	type get struct {
		Id string `json:"id"`
	}
	db := mysql.NewMySqlDB("127.0.0.1", "root", "Zamorak1", "3306", "Restaurant_Test")
	//db.SetMaxOpenConns(5)
	defer db.Close()
	router := Server.NewRouter(db)
	s := router.Router()
	b, _ := json.Marshal(saLogin)

	request, _ := http.NewRequest(http.MethodPost, "/superadmin/login", bytes.NewReader(b))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	s.ServeHTTP(response, request)
	token := response.Header().Get("token")
	superadmindata := make([]get, 3)
	superadmindata[0] = get{Id: "1"}
	superadmindata[1] = get{Id: ""}
	superadmindata[2] = get{Id: "one"}

	type testData struct {
		g        get
		expected int
	}
	data := make([]testData, 3)
	data[0] = testData{g: superadmindata[0], expected: 200}
	data[1] = testData{g: superadmindata[1], expected: 400}
	data[2] = testData{g: superadmindata[2], expected: 500}
	for i, _ := range data {
		superadmin, _ := json.Marshal(data[i].g)
		request, _ = http.NewRequest(http.MethodGet, "/superadmin/get", bytes.NewReader(superadmin))
		request.Header.Set("token", token)
		request.Header.Set("Content-Type", "application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response, request)
		helpers.AssertEqual(t, data[i].expected, response.Code, "")
	}
}

func TestLoginsuperadmin(t *testing.T) {
	db := mysql.NewMySqlDB("127.0.0.1", "root", "Zamorak1", "3306", "Restaurant_Test")
	//db.SetMaxOpenConns(5)
	defer db.Close()
	router := Server.NewRouter(db)
	s := router.Router()
	login := make([]Login, 5)
	login[0] = Login{
		Email: "sourav241196@gmail.com",
		Pass:  "zamorak",
	}
	login[1] = Login{
		Email: "",
		Pass:  "gutthix",
	}
	login[2] = Login{
		Email: "superadmin1@gmail.com",
		Pass:  "",
	}
	login[3] = Login{
		Email: "sourav241196@gmail.com",
		Pass:  "gutthix",
	}
	login[4] = Login{
		Email: "superadmin3@gmail.com",
		Pass:  "gutthix",
	}
	type testData struct {
		l        Login
		expected int
	}
	data := make([]testData, 5)
	data[0] = testData{login[0], 200}
	data[1] = testData{login[1], http.StatusBadRequest}
	data[2] = testData{login[2], http.StatusBadRequest}
	data[3] = testData{login[3], http.StatusInternalServerError}
	data[4] = testData{login[4], http.StatusInternalServerError}
	for i, _ := range data {
		superadmin, _ := json.Marshal(data[i].l)
		request, _ := http.NewRequest(http.MethodPost, "/superadmin/login", bytes.NewReader(superadmin))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		s.ServeHTTP(response, request)
		helpers.AssertEqual(t, data[i].expected, response.Code, "")
	}
}

func TestDelsuperadmin(t *testing.T) {
	type del struct {
		Id string `json:"id"`
	}
	login := Login{
		Email: "sourav241196@gmail.com",
		Pass:  "zamorak",
	}
	db := mysql.NewMySqlDB("127.0.0.1", "root", "Zamorak1", "3306", "Restaurant_Test")
	defer db.Close()
	router := Server.NewRouter(db)
	s := router.Router()
	b, _ := json.Marshal(login)

	request, _ := http.NewRequest(http.MethodPost, "/superadmin/login", bytes.NewReader(b))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	s.ServeHTTP(response, request)
	token := response.Header().Get("token")

	request, _ = http.NewRequest(http.MethodDelete, "/superadmin/del", nil)
	request.Header.Set("token", token)
	request.Header.Set("Content-Type", "application/json")
	response = httptest.NewRecorder()
	s.ServeHTTP(response, request)
	helpers.AssertEqual(t, 200, response.Code, "")
	helpers.ResetDBafterDelete(db, 0)
}
