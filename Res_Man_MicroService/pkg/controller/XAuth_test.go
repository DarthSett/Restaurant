package controller_test

import (
	"bytes"
	"encoding/json"
	_ "github.com/restaurant/Customer_MicroService/pkg/database/mysql"
	"github.com/restaurant/Res_Man_MicroService/pkg/database/mysql"
	"github.com/restaurant/Res_Man_MicroService/pkg/helpers"
	"github.com/restaurant/Res_Man_MicroService/pkg/server"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	db = mysql.NewMySqlDB("localhost", "root", "password", "3306", "Restaurant_Test")
	router = server.NewRouter(db)
	s = router.Router()
)

func TestLogin(t *testing.T) {

	login := make([]Login, 6)
	login[0] = Login{
		Email: "admin1@gmail.com",
		Pass:  "gutthix",
		Rank: "1",
	}
	login[1] = Login{
		Email: "",
		Pass:  "gutthix",
		Rank: "1",
	}
	login[2] = Login{
		Email: "admin1@gmail.com",
		Pass:  "",
		Rank: "1",
	}
	login[3] = Login{
		Email: "admin1@gmail.com",
		Pass:  "gutthix",
		Rank: "",
	}
	login[4] = Login{
		Email: "admin1@gmail.com",
		Pass:  "zamorak",
		Rank: "1",
	}
	login[5] = Login{
		Email: "admin3@gmail.com",
		Pass:  "gutthix",
		Rank: "1",
	}
	type testData struct {
		l        Login
		expected int
	}
	data := make([]testData, 6)
	data[0] = testData{login[0], 200}
	data[1] = testData{login[1], http.StatusBadRequest}
	data[2] = testData{login[2], http.StatusBadRequest}
	data[3] = testData{login[2], http.StatusBadRequest}
	data[4] = testData{login[4], http.StatusInternalServerError}
	data[5] = testData{login[5], http.StatusInternalServerError}
	for i := range data {
		admin, _ := json.Marshal(data[i].l)
		request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewReader(admin))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		s.ServeHTTP(response, request)
		helpers.AssertEqual(t, data[i].expected, response.Code, "")
		if data[i].expected == 200 {
			res := response.Body.String()
			helpers.AssertEqual(t,"User logged in. Token generated",res,"")
		}
	}
}

func TestLogout(t *testing.T) {
	saLogin := Login{Email: "sourav241196@gmail.com", Pass: "zamorak",Rank:"2"}

	b, _ := json.Marshal(saLogin)

	request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewReader(b))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	s.ServeHTTP(response, request)
	token := response.Header().Get("token")
	request, _ = http.NewRequest(http.MethodGet, "/logout", nil)
	request.Header.Set("token", token)
	response = httptest.NewRecorder()
	s.ServeHTTP(response, request)
	helpers.AssertEqual(t, 200, response.Code, "")
	flag,err := db.Checktoken(token)
	if err != nil {
		t.Fatal(err)
	}
	helpers.AssertEqual(t,false,flag,"")
	res := response.Body.String()
	helpers.AssertEqual(t,"User Logged Out",res,"")
}


