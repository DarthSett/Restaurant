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

func TestDist(t *testing.T){
	type Distdata struct {
		Lat string	`json:"lat"`
		Long string	`json:"long"`
		Dist string	 `json:"dist"`
	}
	data := Distdata{
		Lat:  "112.23",
		Long: "30.233",
		Dist: "3000",
	}

	db := mysql.NewMySqlDB("127.0.0.1","root","Zamorak1","3306","Restaurant_Test")
	defer db.Close()
	router := Server.NewRouter(db)
	s := router.Router()
	b, _ := json.Marshal(data)

	request,_ := http.NewRequest(http.MethodGet, "/rest/dist",bytes.NewReader(b))
	request.Header.Set("Content-Type","application/json")
	response :=httptest.NewRecorder()
	s.ServeHTTP(response,request)
	//
	//var msg Distdata
	//c,_:=ioutil.ReadAll(request.Body)
	helpers.AssertEqual(t,200,response.Code,"")
}

func TestRest(t *testing.T) {
	type rest struct {
		Name string `json:"name"`
		Id	string	`json:"rid"`
		Lat string `json:"lat"`
		Long string `json:"long"`
		Owner string `json:"owner"`
		Update1 string `json:"update1"`
		Update2 string `json:"update2"`
		Flag string `json:"flag"`
	}
	saLogin := Login{Email:"sourav241196@gmail.com",Pass:"zamorak"}
	Alogin := Login{
		Email: "admin1@gmail.com",
		Pass:  "zamorak",
	}
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
	t.Run("It returns 200 on creating a rest", func(t *testing.T) {
		restcreate := rest{
			Name:"Barista",
			Lat: "121.24",
			Long: "112.23",
			Owner: "user1@gmail.com",
		}
		rest,_ := json.Marshal(restcreate)
		request,_ := http.NewRequest(http.MethodPost, "/rest/create",bytes.NewReader(rest))
		request.Header.Set("token",token)
		request.Header.Set("Content-Type","application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response,request)
		helpers.AssertEqual(t,200,response.Code,"")
		helpers.ResetDBafterCreate(db)
	})
	t.Run("It returns 200 on getting a rest", func(t *testing.T) {
		restcreate := rest{
			Id:"1",
		}
		rest,_ := json.Marshal(restcreate)
		request,_ := http.NewRequest(http.MethodGet, "/rest/get",bytes.NewReader(rest))
		request.Header.Set("token",token)
		request.Header.Set("Content-Type","application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response,request)
		helpers.AssertEqual(t,200,response.Code,"")
	})
	t.Run("It returns 200 on updating a rest", func(t *testing.T) {
		restcreate := rest{
			Id:"1",
			Update1:"user1@gmail.com",
			Update2:"",
			Flag:"2",
		}
		rest,_ := json.Marshal(restcreate)
		request,_ := http.NewRequest(http.MethodPut, "/rest/update",bytes.NewReader(rest))
		request.Header.Set("token",token)
		request.Header.Set("Content-Type","application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response,request)
		helpers.AssertEqual(t,200,response.Code,"")
	})
	t.Run("It returns 200 on updating a rest from admin login", func(t *testing.T) {
		b, _ := json.Marshal(Alogin)

		request,_ := http.NewRequest(http.MethodPost, "/admin/login",bytes.NewReader(b))
		request.Header.Set("Content-Type","application/json")
		response :=httptest.NewRecorder()
		s.ServeHTTP(response,request)
		token := response.Header().Get("token")
		restcreate := rest{
			Id:"1",
			Update1:"user1@gmail.com",
			Update2:"",
			Flag:"2",
		}
		rest,_ := json.Marshal(restcreate)
		request,_ = http.NewRequest(http.MethodPut, "/rest/update",bytes.NewReader(rest))
		request.Header.Set("token",token)
		request.Header.Set("Content-Type","application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response,request)
		helpers.AssertEqual(t,200,response.Code,"")
	})
	t.Run("It returns 200 on deleting a rest", func(t *testing.T) {
		restcreate := rest{
			Id:"1",
		}
		rest,_ := json.Marshal(restcreate)
		request,_ := http.NewRequest(http.MethodDelete, "/rest/del",bytes.NewReader(rest))
		request.Header.Set("token",token)
		request.Header.Set("Content-Type","application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response,request)
		helpers.AssertEqual(t,200,response.Code,"")
		helpers.ResetDBafterDelete(db,1)
	})
	t.Run("It returns 200 on list", func(t *testing.T) {
		request,_ := http.NewRequest(http.MethodGet, "/rest/list",nil)
		request.Header.Set("token",token)
		response = httptest.NewRecorder()
		s.ServeHTTP(response,request)
		helpers.AssertEqual(t,200,response.Code,"")
	})
	t.Run("It return 200 on getting menu", func(t *testing.T) {
		restcreate:= rest{
			Id: "1",
		}
		rest,_ := json.Marshal(restcreate)
		request,_ := http.NewRequest(http.MethodGet, "/rest/menu",bytes.NewReader(rest))
		request.Header.Set("token",token)
		request.Header.Set("Content-Type","application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response,request)
		helpers.AssertEqual(t,200,response.Code,"")
	})

}

func TestDish(t *testing.T) {
	type Login struct {
		Email string `json:"email"`
		Pass  string `json:"pass"`
	}
	saLogin := Login{Email: "sourav241196@gmail.com", Pass: "zamorak"}
	Alogin := Login{
		Email: "admin1@gmail.com",
		Pass:  "zamorak",
	}
	Ulogin := Login{
		Email: "user1@gmail.com",
		Pass:  "zamorak",
	}
	type dish struct {
		Name   string `json:"dish"`
		Id     string `json:"id"`
		Rid    string `json:"rid"`
		Update string `json:"update"`
		Price  string `json:"price"`
		Flag   string `json:"flag"`
	}
	db := mysql.NewMySqlDB("127.0.0.1", "root", "Zamorak1", "3306", "Restaurant_Test")
	defer db.Close()
	router := Server.NewRouter(db)
	s := router.Router()
	b, _ := json.Marshal(saLogin)

	request, _ := http.NewRequest(http.MethodPost, "/superadmin/login", bytes.NewReader(b))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	s.ServeHTTP(response, request)
	token := response.Header().Get("token")
	t.Run("It return 200 on creating a dish", func(t *testing.T) {
		dishcreate := dish{
			Name:  "burger",
			Rid:   "1",
			Price: "150",
		}
		dish, _ := json.Marshal(dishcreate)
		request, _ := http.NewRequest(http.MethodPost, "/rest/dish/create", bytes.NewReader(dish))
		request.Header.Set("token", token)
		request.Header.Set("Content-Type", "application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response, request)
		helpers.AssertEqual(t, 200, response.Code, "")
		helpers.ResetDBafterCreate(db)
	})
	t.Run("It return 200 on creating a dish by admin login", func(t *testing.T) {
		dishcreate := dish{
			Name:  "burger",
			Rid:   "1",
			Price: "150",
		}
		b, _ := json.Marshal(Alogin)
		request, _ := http.NewRequest(http.MethodPost, "/admin/login", bytes.NewReader(b))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		s.ServeHTTP(response, request)
		token := response.Header().Get("token")
		dish, _ := json.Marshal(dishcreate)
		request, _ = http.NewRequest(http.MethodPost, "/rest/dish/create", bytes.NewReader(dish))
		request.Header.Set("token", token)
		request.Header.Set("Content-Type", "application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response, request)
		helpers.AssertEqual(t, 200, response.Code, "")
		helpers.ResetDBafterCreate(db)
	})
	t.Run("It return 200 on creating a dish by user login", func(t *testing.T) {
		dishcreate := dish{
			Name:  "burger",
			Rid:   "1",
			Price: "150",
		}
		b, _ := json.Marshal(Ulogin)
		request, _ := http.NewRequest(http.MethodPost, "/user/login", bytes.NewReader(b))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		s.ServeHTTP(response, request)
		token := response.Header().Get("token")
		dish, _ := json.Marshal(dishcreate)
		request, _ = http.NewRequest(http.MethodPost, "/rest/dish/create", bytes.NewReader(dish))
		request.Header.Set("token", token)
		request.Header.Set("Content-Type", "application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response, request)
		helpers.AssertEqual(t, 200, response.Code, "")
		helpers.ResetDBafterCreate(db)
	})
	t.Run("It return 200 on updating a dish", func(t *testing.T) {
		dishcreate := dish{
			Id:     "1",
			Rid:    "1",
			Flag:   "1",
			Update: "100",
		}

		dish, _ := json.Marshal(dishcreate)
		request, _ := http.NewRequest(http.MethodPut, "/rest/dish/update", bytes.NewReader(dish))
		request.Header.Set("token", token)
		request.Header.Set("Content-Type", "application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response, request)
		helpers.AssertEqual(t, 200, response.Code, "")
	})
	t.Run("It return 200 on updating a dish by admin login", func(t *testing.T) {
		dishcreate := dish{
			Id:     "1",
			Rid:    "1",
			Flag:   "1",
			Update: "100",
		}
		b, _ := json.Marshal(Alogin)
		request, _ := http.NewRequest(http.MethodPost, "/admin/login", bytes.NewReader(b))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		s.ServeHTTP(response, request)
		token := response.Header().Get("token")
		dish, _ := json.Marshal(dishcreate)
		request, _ = http.NewRequest(http.MethodPut, "/rest/dish/update", bytes.NewReader(dish))
		request.Header.Set("token", token)
		request.Header.Set("Content-Type", "application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response, request)
		helpers.AssertEqual(t, 200, response.Code, "")
	})
	t.Run("It return 200 on updating a dish by user login", func(t *testing.T) {
		dishcreate := dish{
			Id:     "1",
			Rid:    "1",
			Flag:   "1",
			Update: "100",
		}
		b, _ := json.Marshal(Ulogin)
		request, _ := http.NewRequest(http.MethodPost, "/user/login", bytes.NewReader(b))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		s.ServeHTTP(response, request)
		token := response.Header().Get("token")
		dish, _ := json.Marshal(dishcreate)
		request, _ = http.NewRequest(http.MethodPut, "/rest/dish/update", bytes.NewReader(dish))
		request.Header.Set("token", token)
		request.Header.Set("Content-Type", "application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response, request)
		helpers.AssertEqual(t, 200, response.Code, "")
	})
	t.Run("It return 200 on deleting a dish", func(t *testing.T) {
		dishcreate := dish{
			Id:  "1",
			Rid: "1",
		}
		dish, _ := json.Marshal(dishcreate)
		request, _ := http.NewRequest(http.MethodDelete, "/rest/dish/del", bytes.NewReader(dish))
		request.Header.Set("token", token)
		request.Header.Set("Content-Type", "application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response, request)
		helpers.AssertEqual(t, 200, response.Code, "")
		helpers.ResetDBafterDelete(db,2)
	})
	t.Run("It return 200 on deleting a dish on admin login", func(t *testing.T) {
		dishcreate := dish{
			Id:  "1",
			Rid: "1",
		}
		b, _ := json.Marshal(Alogin)
		request, _ := http.NewRequest(http.MethodPost, "/admin/login", bytes.NewReader(b))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		s.ServeHTTP(response, request)
		token := response.Header().Get("token")
		dish, _ := json.Marshal(dishcreate)
		request, _ = http.NewRequest(http.MethodDelete, "/rest/dish/del", bytes.NewReader(dish))
		request.Header.Set("token", token)
		request.Header.Set("Content-Type", "application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response, request)
		helpers.AssertEqual(t, 200, response.Code, "")
		helpers.ResetDBafterDelete(db,2)
	})
	t.Run("It return 200 on deleting a dish on user login", func(t *testing.T) {
		dishcreate := dish{
			Id:  "1",
			Rid: "1",
		}
		b, _ := json.Marshal(Ulogin)
		request, _ := http.NewRequest(http.MethodPost, "/user/login", bytes.NewReader(b))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		s.ServeHTTP(response, request)
		token := response.Header().Get("token")
		dish, _ := json.Marshal(dishcreate)
		request, _ = http.NewRequest(http.MethodDelete, "/rest/dish/del", bytes.NewReader(dish))
		request.Header.Set("token", token)
		request.Header.Set("Content-Type", "application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response, request)
		helpers.AssertEqual(t, 200, response.Code, "")
		helpers.ResetDBafterDelete(db,2)
	})
}






