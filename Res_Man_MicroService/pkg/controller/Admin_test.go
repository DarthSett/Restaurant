package controller_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/restaurant/Res_Man_MicroService/pkg/helpers"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestCreateAdmin(t *testing.T) {
	//saLogin := Login{Email:"sourav241196@gmail.com",Pass:"zamorak"}
	type create struct {
		Name  string `json:"name"`
		Pass  string `json:"pass"`
		Email string `json:"email"`
	}
	db.SetMaxOpenConns(151)



	admindata := make([]create, 5)
	login := make([]Login, 2)
	login[0] = Login{
		Email: "sourav241196@gmail.com",
		Pass:  "zamorak",
		Rank: "2",
	}
	login[1] = Login{
		Email: "admin1@gmail.com",
		Pass:  "gutthix",
		Rank: "1",
	}
	admindata[0] = create{
		Name:  "admin2",
		Pass:  "zamorak",
		Email: "admin2@gmail.com",
	}
	admindata[1] = create{
		Name:  "admin2",
		Pass:  "",
		Email: "admin2@gmail.com",
	}
	admindata[2] = create{
		Name:  "",
		Pass:  "zamorak",
		Email: "admin2@gmail.com",
	}
	admindata[3] = create{
		Name:  "admin2",
		Pass:  "zamorak",
		Email: "",
	}
	admindata[4] = create{
		Name:  "admin2",
		Pass:  "zamorak",
		Email: "admin1@gmail.com",
	}

	type testData struct {
		login    Login
		c        create
		expected int
	}
	data := make([]testData, 6)
	data[0] = testData{login[0], admindata[0], 200}
	data[1] = testData{login[0], admindata[1], 400}
	data[2] = testData{login[0], admindata[2], 400}
	data[3] = testData{login[0], admindata[3], 400}
	data[4] = testData{login[0], admindata[4], 500}
	data[5] = testData{login[1], admindata[0], 401}

	var request *http.Request
	for i := range data {
		println(i)
		b, _ := json.Marshal(data[i].login)
		request, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewReader(b))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		s.ServeHTTP(response, request)
		token := response.Header().Get("token")
		admin, _ := json.Marshal(data[i].c)
		request, _ = http.NewRequest(http.MethodPost, "/admin/create", bytes.NewReader(admin))
		request.Header.Set("token", token)
		request.Header.Set("Content-Type", "application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response, request)
		helpers.AssertEqual(t, data[i].expected, response.Code, "")
		if data[i].expected == 200 {
			u,err := db.GetAdmin(data[i].c.Email,0)
			if err != nil {
				panic(err)
			}
			println("Name: " + u.Name)
			helpers.AssertEqual(t,data[i].c.Name,u.Name,"")
			res := response.Body.String()
			helpers.AssertEqual(t, "Admin Saved",res,"")
		}
		helpers.ResetDB(db)
	}

}

func TestGetAdmin(t *testing.T) {

	saLogin := Login{Email: "sourav241196@gmail.com", Pass: "zamorak", Rank: "2"}
	type get struct {
		Id string `json:"id"`
	}

	b, _ := json.Marshal(saLogin)

	request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewReader(b))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	s.ServeHTTP(response, request)
	token := response.Header().Get("token")
	admindata := make([]get, 3)
	admindata[0] = get{Id: "1"}
	admindata[1] = get{Id: ""}
	admindata[2] = get{Id: "one"}

	type testData struct {
		g        get
		expected int
	}
	data := make([]testData, 3)
	data[0] = testData{g: admindata[0], expected: 200}
	data[1] = testData{g: admindata[1], expected: 400}
	data[2] = testData{g: admindata[2], expected: 500}
	for i := range data {
		admin, _ := json.Marshal(data[i].g)
		request, _ = http.NewRequest(http.MethodGet, "/admin/get", bytes.NewReader(admin))
		request.Header.Set("token", token)
		request.Header.Set("Content-Type", "application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response, request)
		helpers.AssertEqual(t, data[i].expected, response.Code, "")
		if data[i].expected == 200 {
			id, _ := strconv.Atoi(data[i].g.Id)
			u, err := db.GetAdmin("", id)
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
			helpers.AssertEqual(t,string(jsonU),res,"")
		}
	}
}



func TestUpdateAdmin(t *testing.T) {

	type upd struct {
		Id     string `json:"id"`
		Flag   string `json:"flag"`
		Update string `json:"update"`
	}

	login := Login{
		Email: "sourav241196@gmail.com",
		Pass:  "zamorak",
		Rank:"2",
	}
	admindata := make([]upd, 8)
	admindata[0] = upd{Id: "1", Flag: "0", Update: "admin1"}
	admindata[1] = upd{Id: "1", Flag: "1", Update: "gutthix"}
	admindata[2] = upd{Id: "1", Flag: "2", Update: "admin1@gmail.com"}
	admindata[3] = upd{Id: "", Flag: "0", Update: "admin11"}
	admindata[4] = upd{Id: "1", Flag: "", Update: "admin11"}
	admindata[5] = upd{Id: "1", Flag: "0", Update: ""}
	admindata[6] = upd{Id: "one", Flag: "0", Update: "admin11"}
	admindata[7] = upd{Id: "1", Flag: "one", Update: "admin11"}
	type testData struct {
		admin    upd
		Expected int
	}
	data := make([]testData, 8)
	data[0] = testData{admin: admindata[0], Expected: 200}
	data[1] = testData{admin: admindata[1], Expected: 200}
	data[2] = testData{admin: admindata[2], Expected: 200}
	data[3] = testData{admin: admindata[3], Expected: 400}
	data[4] = testData{admin: admindata[4], Expected: 400}
	data[5] = testData{admin: admindata[5], Expected: 400}
	data[6] = testData{admin: admindata[6], Expected: 500}
	data[7] = testData{admin: admindata[7], Expected: 500}

	b, _ := json.Marshal(login)
	request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewReader(b))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	s.ServeHTTP(response, request)
	token := response.Header().Get("token")
	for i := range data {

		admin, _ := json.Marshal(data[i].admin)
		println(i)
		request, _ = http.NewRequest(http.MethodPut, "/admin/update", bytes.NewReader(admin))
		request.Header.Set("token", token)
		request.Header.Set("Content-Type", "application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response, request)
		helpers.AssertEqual(t, data[i].Expected, response.Code, "")
		if data[i].Expected == 200 {
			id,_ := strconv.Atoi(data[i].admin.Id)
			u,err := db.GetAdmin("",id)
			if err != nil {
				t.Fatal(err)
			}
			println("Name: " + u.Name)
			if data[i].admin.Flag == "0"{
				helpers.AssertEqual(t,data[i].admin.Update,u.Name,"")
			} else if data[i].admin.Flag == "1" {
				err = bcrypt.CompareHashAndPassword([]byte(u.Pass),[]byte(data[i].admin.Update) )
				if err != nil {
					t.Fatal(err)
				}
			} else {
				helpers.AssertEqual(t,data[i].admin.Update,u.Email,"")
			}
			res := response.Body.String()
			helpers.AssertEqual(t,"admin Updated",res,"")
		}
	}
	helpers.ResetDB(db)
}

func TestDelAdmin(t *testing.T) {
	type del struct {
		Id string `json:"id"`
	}
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
	admin := make([]del, 3)
	admin[0] = del{Id: "1"}
	admin[1] = del{Id: ""}
	admin[2] = del{Id: "one"}
	type testdata struct {
		D        del
		Expected int
	}
	data := make([]testdata, 3)
	data[0] = testdata{D: admin[0], Expected: 200}
	data[1] = testdata{D: admin[1], Expected: 400}
	data[2] = testdata{D: admin[2], Expected: 500}

	for i := range admin {
		println(i)
		admin, _ := json.Marshal(data[i].D)
		request, _ := http.NewRequest(http.MethodDelete, "/admin/del", bytes.NewReader(admin))
		request.Header.Set("token", token)
		request.Header.Set("Content-Type", "application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response, request)
		helpers.AssertEqual(t, data[i].Expected, response.Code, "")
		if data[i].Expected==200 {
			id,_ := strconv.Atoi(data[i].D.Id)
			u,err := db.GetAdmin("",id)
			println("Name: " + u.Name)
			e := fmt.Errorf("no such record in database")
			helpers.AssertEqual(t,e.Error(),err.Error(),"")
			res := response.Body.String()
			helpers.AssertEqual(t,data[i].D.Id + " Deleted from db",res,"")
		}

	}
	helpers.ResetDB(db)
}
