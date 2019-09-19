package controller_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	helpers2 "github.com/restaurant/Res_Man_MicroService/pkg/helpers"
	"golang.org/x/crypto/bcrypt"
	"strconv"

	"net/http"
	"net/http/httptest"
	"testing"
)

type Login struct {
	Email string `json:"email"`
	Pass  string `json:"pass"`
	Rank  string `json:"rank"`
}

func TestCreateUser(t *testing.T) {
	//saLogin := Login{Email:"sourav241196@gmail.com",Pass:"zamorak"}
	type create struct {
		Name  string `json:"name"`
		Pass  string `json:"pass"`
		Email string `json:"email"`
	}

	userdata := make([]create, 5)
	login := make([]Login, 3)
	login[0] = Login{
		Email: "sourav241196@gmail.com",
		Pass:  "zamorak",
		Rank:"2",
	}
	login[1] = Login{
		Email: "user1@gmail.com",
		Pass:  "zamorak",
		Rank:"0",
	}
	login[2] = Login{
		Email: "admin1@gmail.com",
		Pass:  "gutthix",
		Rank:"1",
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
		login    Login
		c        create
		expected int
	}
	data := make([]testData, 15)
	data[0] = testData{login[0], userdata[0], 200}
	data[1] = testData{login[0], userdata[1], 400}
	data[2] = testData{login[0], userdata[2], 400}
	data[3] = testData{login[0], userdata[3], 400}
	data[4] = testData{login[0], userdata[4], 500}
	data[5] = testData{login[1], userdata[0], 401}
	data[6] = testData{login[1], userdata[1], 401}
	data[7] = testData{login[1], userdata[2], 401}
	data[8] = testData{login[1], userdata[3], 401}
	data[9] = testData{login[1], userdata[4], 401}
	data[10] = testData{login[2], userdata[0], 200}
	data[11] = testData{login[2], userdata[1], 400}
	data[12] = testData{login[2], userdata[2], 400}
	data[13] = testData{login[2], userdata[3], 400}
	data[14] = testData{login[2], userdata[4], 500}

	var request *http.Request
	for i := range data {
		println("i: ", i)
		b, _ := json.Marshal(data[i].login)
		request, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewReader(b))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		s.ServeHTTP(response, request)
		token := response.Header().Get("token")
		user, _ := json.Marshal(data[i].c)
		request, _ = http.NewRequest(http.MethodPost, "/user/create", bytes.NewReader(user))
		request.Header.Set("token", token)
		request.Header.Set("Content-Type", "application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response, request)
		helpers2.AssertEqual(t, data[i].expected, response.Code, "")
		if data[i].expected == 200 {
			u,err := db.GetUser(data[i].c.Email,0)
			if err != nil {
				panic(err)
			}
			println("Name: " + u.Name)
			helpers2.AssertEqual(t,data[i].c.Name,u.Name,"")
			res := response.Body.String()
			helpers2.AssertEqual(t, "User Saved",res,"")
		}
		helpers2.ResetDB(db)
	}

}

func TestGetUser(t *testing.T) {

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
	userdata := make([]get, 3)
	userdata[0] = get{Id: "1"}
	userdata[1] = get{Id: ""}
	userdata[2] = get{Id: "one"}

	type testData struct {
		g        get
		expected int
	}
	data := make([]testData, 3)
	data[0] = testData{g: userdata[0], expected: 200}
	data[1] = testData{g: userdata[1], expected: 400}
	data[2] = testData{g: userdata[2], expected: 500}
	for i := range data {
		user, _ := json.Marshal(data[i].g)
		request, _ = http.NewRequest(http.MethodGet, "/user/get", bytes.NewReader(user))
		request.Header.Set("token", token)
		request.Header.Set("Content-Type", "application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response, request)
		helpers2.AssertEqual(t, data[i].expected, response.Code, "")
		if data[i].expected == 200 {
			id, _ := strconv.Atoi(data[i].g.Id)
			u, err := db.GetUser("", id)
			if err != nil {
				t.Fatal(err)
			}
			u.Pass = ""
			res := response.Body.String()
			println("Name: " + u.Name)
			rids, names, err := db.GetUserRests(u.Email)
			if err != nil {
				t.Fatal(err)
			}
			o := make([]struct {
				RID      int
				RestName string
			}, len(rids))
			for i := range rids {
				o[i].RID = rids[i]
				o[i].RestName = names[i]
			}
			m := make(map[string]interface{})
			m["user"]= u
			m["Restaurants owned"] = o
			jsonU,err := json.Marshal(m)
			if err != nil {
				t.Fatal(err)
			}
			helpers2.AssertEqual(t,string(jsonU),res,"")
		}
	}
}

func TestUpdateUser(t *testing.T) {

	type upd struct {
		Id     string `json:"id"`
		Flag   string `json:"flag"`
		Update string `json:"update"`
	}
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
	userdata := make([]upd, 9)
	userdata[0] = upd{Id: "1", Flag: "0", Update: "user1"}
	userdata[1] = upd{Id: "1", Flag: "1", Update: "zamorak"}
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
		request, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewReader(b))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		s.ServeHTTP(response, request)
		token := response.Header().Get("token")
		user, _ := json.Marshal(data[i].User)
		println(i)
		request, _ = http.NewRequest(http.MethodPut, "/user/update", bytes.NewReader(user))
		request.Header.Set("token", token)
		request.Header.Set("Content-Type", "application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response, request)
		helpers2.AssertEqual(t, data[i].Expected, response.Code, "")
		if data[i].Expected == 200 {
			id,_ := strconv.Atoi(data[i].User.Id)
			u,err := db.GetUser("",id)
			if err != nil {
				t.Fatal(err)
			}
			println("Name: " + u.Name)
			if data[i].User.Flag == "0"{
				helpers2.AssertEqual(t,data[i].User.Update,u.Name,"")
			} else if data[i].User.Flag == "1" {
				err = bcrypt.CompareHashAndPassword([]byte(u.Pass),[]byte(data[i].User.Update) )
				if err != nil {
					t.Fatal(err)
				}
			} else {
				helpers2.AssertEqual(t,data[i].User.Update,u.Email,"")
			}
			res := response.Body.String()
			helpers2.AssertEqual(t,"User Updated",res,"")
		}
	}
	helpers2.ResetDB(db)
}

func TestDelUser(t *testing.T) {
	type del struct {
		Id string `json:"id"`
	}
	login := Login{
		Email: "sourav241196@gmail.com",
		Pass:  "zamorak",
		Rank:"2",
	}

	b, _ := json.Marshal(login)

	request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewReader(b))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	s.ServeHTTP(response, request)
	token := response.Header().Get("token")
	user := make([]del, 3)
	user[0] = del{Id: "1"}
	user[1] = del{Id: ""}
	user[2] = del{Id: "one"}
	type testdata struct {
		D        del
		Expected int
	}
	data := make([]testdata, 3)
	data[0] = testdata{D: user[0], Expected: 200}
	data[1] = testdata{D: user[1], Expected: 400}
	data[2] = testdata{D: user[2], Expected: 500}

	for i := range user {
		println(i)
		user, _ := json.Marshal(data[i].D)
		request, _ := http.NewRequest(http.MethodDelete, "/user/del", bytes.NewReader(user))
		request.Header.Set("token", token)
		request.Header.Set("Content-Type", "application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response, request)
		helpers2.AssertEqual(t, data[i].Expected, response.Code, "")
		if data[i].Expected==200 {
			id,_ := strconv.Atoi(data[i].D.Id)
			u,err := db.GetUser("",id)
			println("Name: " + u.Name)
			e := fmt.Errorf("no such record in database")
			helpers2.AssertEqual(t,e.Error(),err.Error(),"")
			res := response.Body.String()
			helpers2.AssertEqual(t,data[i].D.Id + " Deleted from db",res,"")
		}
	}
	helpers2.ResetDB(db)
}
func TestListUser(t *testing.T) {

	saLogin := Login{Email: "sourav241196@gmail.com", Pass: "zamorak",Rank:"2"}


	b, _ := json.Marshal(saLogin)

	request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewReader(b))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	s.ServeHTTP(response, request)
	token := response.Header().Get("token")
	request, _ = http.NewRequest(http.MethodGet, "/user/list", nil)
	request.Header.Set("token", token)
	s.ServeHTTP(response, request)
	helpers2.AssertEqual(t, 200, response.Code, "")
	name, email, id, err := db.UserList()
	x := make([]struct {
		Name  string
		Email string
		Id    int
	}, len(name))
	if err != nil {
		t.Fatal(err)
	}
	//o := make(map[string]string)
	for i := range name {
		x[i].Name = name[i]
		x[i].Email = email[i]
		x[i].Id = id[i]
	}
	jsonu,err := json.Marshal(x)
	if err != nil {
		t.Fatal(err)
	}
	res := response.Body.String()
	helpers2.AssertEqual(t,"User logged in. Token generated" + string(jsonu),res,"")
}

