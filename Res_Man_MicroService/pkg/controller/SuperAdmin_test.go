package controller_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	helpers2 "github.com/restaurant/Res_Man_MicroService/pkg/helpers"
	"strconv"

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



	superadmindata := make([]create, 5)
	login := Login{
		Email: "sourav241196@gmail.com",
		Pass:  "zamorak",
		Rank: "2",
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
	for i := range data {
		println(i)
		b, _ := json.Marshal(data[i].login)
		request, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewReader(b))
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
		helpers2.AssertEqual(t, data[i].expected, response.Code, "")
		if data[i].expected == 200 {
			u,err := db.GetSuperAdmin(data[i].c.Email,0)
			if err != nil {
				panic(err)
			}
			println("Name: " + u.Name)
			helpers2.AssertEqual(t,data[i].c.Name,u.Name,"")
			res := response.Body.String()
			helpers2.AssertEqual(t, "SuperAdmin Saved",res,"")
		}
		helpers2.ResetDB(db)
	}

}

func TestGetsuperadmin(t *testing.T) {

	saLogin := Login{Email: "sourav241196@gmail.com", Pass: "zamorak",Rank:"2"}
	type get struct {
		Id string `json:"id"`
	}

	b, _ := json.Marshal(saLogin)

	request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewReader(b))
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
	for i := range data {
		superadmin, _ := json.Marshal(data[i].g)
		request, _ = http.NewRequest(http.MethodGet, "/superadmin/get", bytes.NewReader(superadmin))
		request.Header.Set("token", token)
		request.Header.Set("Content-Type", "application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response, request)
		helpers2.AssertEqual(t, data[i].expected, response.Code, "")
		if data[i].expected == 200 {
			id, _ := strconv.Atoi(data[i].g.Id)
			u, err := db.GetSuperAdmin("", id)
			if err != nil {
				t.Fatal(err)
			}
			u.Pass = ""
			res := response.Body.String()
			println("Name: " + u.Name)
			jsonU,err := json.Marshal(u)
			if err != nil {
				t.Fatal(err)
			}
			helpers2.AssertEqual(t,string(jsonU),res,"")
		}
	}
}

func TestDelsuperadmin(t *testing.T) {
	login := Login{
		Email: "sourav241196@gmail.com",
		Pass:  "zamorak",
		Rank: "2",
	}

	b, _ := json.Marshal(login)

	request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewReader(b))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	s.ServeHTTP(response, request)
	token := response.Header().Get("token")

	request, _ = http.NewRequest(http.MethodDelete, "/superadmin/del", nil)
	request.Header.Set("token", token)
	request.Header.Set("Content-Type", "application/json")
	response = httptest.NewRecorder()
	s.ServeHTTP(response, request)
	helpers2.AssertEqual(t, 200, response.Code, "")
	id := 1
	u,err := db.GetSuperAdmin("",id)
	println("Name: " + u.Name)
	e := fmt.Errorf("no such record in database")
	helpers2.AssertEqual(t,e.Error(),err.Error(),"")
	res := response.Body.String()
	helpers2.AssertEqual(t,  "1 Deleted from db",res,"")
	helpers2.ResetDB(db)
}
