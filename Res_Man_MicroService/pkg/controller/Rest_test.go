package controller_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	helpers2 "github.com/restaurant/Res_Man_MicroService/pkg/helpers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDist(t *testing.T) {
	type Distdata struct {
		Lat  string `json:"lat"`
		Long string `json:"long"`
		Dist string `json:"dist"`
	}
	data := Distdata{
		Lat:  "122.23",
		Long: "127.233",
		Dist: "3000",
	}


	b, _ := json.Marshal(data)

	request, _ := http.NewRequest(http.MethodGet, "/rest/dist", bytes.NewReader(b))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	s.ServeHTTP(response, request)
	//
	//var msg Distdata
	//c,_:=ioutil.ReadAll(request.Body)
	helpers2.AssertEqual(t, 200, response.Code, "")
}

func TestRest(t *testing.T) {
	type rest struct {
		Name    string `json:"name"`
		Id      string `json:"rid"`
		Lat     string `json:"lat"`
		Long    string `json:"long"`
		Owner   string `json:"owner"`
		Update1 string `json:"update1"`
		Update2 string `json:"update2"`
		Flag    string `json:"flag"`
	}
	saLogin := Login{Email: "sourav241196@gmail.com", Pass: "zamorak", Rank:"2"}
	Alogin := Login{
		Email: "admin1@gmail.com",
		Pass:  "gutthix",
		Rank:"1",
	}

	b, _ := json.Marshal(saLogin)

	request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewReader(b))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	s.ServeHTTP(response, request)
	token := response.Header().Get("token")
	t.Run("It returns 200 on creating a rest", func(t *testing.T) {
		restcreate := rest{
			Name:  "Barista",
			Lat:   "121.24",
			Long:  "112.23",
			Owner: "user1@gmail.com",
		}
		rest, _ := json.Marshal(restcreate)
		request, _ := http.NewRequest(http.MethodPost, "/rest/create", bytes.NewReader(rest))
		request.Header.Set("token", token)
		request.Header.Set("Content-Type", "application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response, request)
		helpers2.AssertEqual(t, 200, response.Code, "")
		u,err := db.GetRestaurant(2)
		if err != nil {
			panic(err)
		}
		println("Name: " + u.Name)
		helpers2.AssertEqual(t,restcreate.Name,u.Name,"")
		res := response.Body.String()
		helpers2.AssertEqual(t, "Restaurant Saved",res,"")
		helpers2.ResetDB(db)
	})
	t.Run("It returns 200 on getting a rest", func(t *testing.T) {
		restcreate := rest{
			Id: "1",
		}
		rest, _ := json.Marshal(restcreate)
		request, _ := http.NewRequest(http.MethodGet, "/rest/get", bytes.NewReader(rest))
		request.Header.Set("token", token)
		request.Header.Set("Content-Type", "application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response, request)
		helpers2.AssertEqual(t, 200, response.Code, "")
		r, err := db.GetRestaurant(1)
		if err != nil {
			panic("Error getting rest from db: " + err.Error())
		}
		ids, names, prices, err := db.GetMenu(1)
		if err != nil {
			panic("Error getting menu from db: " + err.Error())
		}
		menu := make([]struct {
			Id    int
			Name  string
			Price int
		}, len(ids))
		for i := range ids {
			menu[i].Id = ids[i]
			menu[i].Name = names[i]
			menu[i].Price = prices[i]
		}
		m := make(map[string]interface{})
		m["Restaurant"] = r
		m["Menu"] = menu
		res := response.Body.String()
		jsonU,err := json.Marshal(m)
		if err != nil {
			t.Fatal(err)
		}
		helpers2.AssertEqual(t,res,string(jsonU),"")
	})
	t.Run("It returns 200 on updating a rest", func(t *testing.T) {
		restcreate := rest{
			Id:      "1",
			Update1: "user1@gmail.com",
			Update2: "",
			Flag:    "2",
		}
		rest, _ := json.Marshal(restcreate)
		request, _ := http.NewRequest(http.MethodPut, "/rest/update", bytes.NewReader(rest))
		request.Header.Set("token", token)
		request.Header.Set("Content-Type", "application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response, request)
		helpers2.AssertEqual(t, 200, response.Code, "")
		u,err := db.GetRestaurant(1)
		if err != nil {
			panic(err)
		}
		println("Name: " + u.Name)
		helpers2.AssertEqual(t,restcreate.Update1,u.Owner,"")
		res := response.Body.String()
		helpers2.AssertEqual(t,"Rest Updated",res,"")
	})
	t.Run("It returns 200 on updating a rest from admin login", func(t *testing.T) {
		b, _ := json.Marshal(Alogin)

		request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewReader(b))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		s.ServeHTTP(response, request)
		token := response.Header().Get("token")
		restcreate := rest{
			Id:      "1",
			Update1: "user1@gmail.com",
			Update2: "",
			Flag:    "2",
		}
		rest, _ := json.Marshal(restcreate)
		request, _ = http.NewRequest(http.MethodPut, "/rest/update", bytes.NewReader(rest))
		request.Header.Set("token", token)
		request.Header.Set("Content-Type", "application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response, request)
		helpers2.AssertEqual(t, 200, response.Code, "")
		u,err := db.GetRestaurant(1)
		if err != nil {
			panic(err)
		}
		println("Name: " + u.Name)
		helpers2.AssertEqual(t,restcreate.Update1,u.Owner,"")
		res := response.Body.String()
		helpers2.AssertEqual(t,"Rest Updated",res,"")
	})
	t.Run("It returns 200 on deleting a rest", func(t *testing.T) {
		restcreate := rest{
			Id: "1",
		}
		rest, _ := json.Marshal(restcreate)
		request, _ := http.NewRequest(http.MethodDelete, "/rest/del", bytes.NewReader(rest))
		request.Header.Set("token", token)
		request.Header.Set("Content-Type", "application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response, request)
		helpers2.AssertEqual(t, 200, response.Code, "")
		u,err := db.GetRestaurant(1)
		println("Name: " + u.Name)
		e := fmt.Errorf("no such record in database")
		helpers2.AssertEqual(t,e.Error(),err.Error(),"")
		res := response.Body.String()
		helpers2.AssertEqual(t,"rid: " + restcreate.Id + " Deleted from db",res,"")
		helpers2.ResetDB(db)

	})
	t.Run("It returns 200 on list", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/rest/list", nil)
		request.Header.Set("token", token)
		response = httptest.NewRecorder()
		s.ServeHTTP(response, request)
		helpers2.AssertEqual(t, 200, response.Code, "")
		id, name, err := db.RestList()
		if err != nil {
			t.Fatal("Error getting list from db: " + err.Error())
		}
		o := make([]struct {
			Id   int
			Name string
		}, len(id))
		for i := range id {
			o[i].Id = id[i]
			o[i].Name = name[i]
		}
		jsonU,err := json.Marshal(o)
		if err != nil {
			t.Fatal(err.Error())
		}
		res := response.Body.String()
		helpers2.AssertEqual(t,string(jsonU),res,"")
	})
	t.Run("It return 200 on getting menu", func(t *testing.T) {
		restcreate := rest{
			Id: "1",
		}
		rest, _ := json.Marshal(restcreate)
		request, _ := http.NewRequest(http.MethodGet, "/rest/menu", bytes.NewReader(rest))
		request.Header.Set("token", token)
		request.Header.Set("Content-Type", "application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response, request)
		helpers2.AssertEqual(t, 200, response.Code, "")
		ids, names, prices, err := db.GetMenu(1)
		if err != nil {
			panic("Error getting menu from db: " + err.Error())
		}
		menu := make([]struct {
			Id    int
			Name  string
			Price int
		}, len(ids))
		for i := range ids {
			menu[i].Id = ids[i]
			menu[i].Name = names[i]
			menu[i].Price = prices[i]
		}
		jsonU,err := json.Marshal(menu)
		if err != nil {
			t.Fatal(err)
		}
		res := response.Body.String()
		helpers2.AssertEqual(t,string(jsonU),res,"")
	})

}

func TestDish(t *testing.T) {

	saLogin := Login{Email: "sourav241196@gmail.com", Pass: "zamorak",Rank:"2"}
	Alogin := Login{
		Email: "admin1@gmail.com",
		Pass:  "gutthix",
		Rank:"1",
	}
	Ulogin := Login{
		Email: "user1@gmail.com",
		Pass:  "zamorak",
		Rank: "0",
	}
	type dish struct {
		Name   string `json:"dish"`
		Id     string `json:"id"`
		Rid    string `json:"rid"`
		Update string `json:"update"`
		Price  string `json:"price"`
		Flag   string `json:"flag"`
	}

	b, _ := json.Marshal(saLogin)

	request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewReader(b))
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
		helpers2.AssertEqual(t, 200, response.Code, "")
		q := fmt.Sprintf("select Name from dish where Id = '2' and status = '1'")
		row,err := db.Query(q)
		if err != nil {
			t.Fatal(err)
		}
		var u string
		for row.Next(){
			err = row.Scan(&u)
		}
		if err != nil {
			t.Fatal(err)
		}
		println("Name: " + u)
		helpers2.AssertEqual(t,dishcreate.Name,u,"")
		res := response.Body.String()
		helpers2.AssertEqual(t,"dish created",res,"")
		helpers2.ResetDB(db)
	})
	t.Run("It return 200 on creating a dish by admin login", func(t *testing.T) {
		dishcreate := dish{
			Name:  "burger",
			Rid:   "1",
			Price: "150",
		}
		b, _ := json.Marshal(Alogin)
		request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewReader(b))
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
		helpers2.AssertEqual(t, 200, response.Code, "")
		q := fmt.Sprintf("select Name from dish where Id = '2' and status = '1'")
		row,err := db.Query(q)
		if err != nil {
			t.Fatal(err)
		}
		var u string
		for row.Next(){
			err = row.Scan(&u)
		}
		if err != nil {
			t.Fatal(err)
		}
		println("Name: " + u)
		helpers2.AssertEqual(t,dishcreate.Name,u,"")
		res := response.Body.String()
		helpers2.AssertEqual(t,"dish created",res,"")
		helpers2.ResetDB(db)
	})
	t.Run("It return 200 on creating a dish by user login", func(t *testing.T) {
		dishcreate := dish{
			Name:  "burger",
			Rid:   "1",
			Price: "150",
		}
		b, _ := json.Marshal(Ulogin)
		request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewReader(b))
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
		helpers2.AssertEqual(t, 200, response.Code, "")

		q := fmt.Sprintf("select Name from dish where id = '2' and status = '1'")
		row,err := db.Query(q)
		if err != nil {
			t.Fatal(err)
		}
		var u string
		for row.Next(){
			err = row.Scan(&u)
		}
		if err != nil {
			t.Fatal(err)
		}
		println("Name: " + u)
		helpers2.AssertEqual(t,dishcreate.Name,u,"")
		res := response.Body.String()
		helpers2.AssertEqual(t,"dish created",res,"")
		helpers2.ResetDB(db)
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
		helpers2.AssertEqual(t, 200, response.Code, "")
		q := fmt.Sprintf("select price from dish where Id = '1' and status = '1'")
		row,err := db.Query(q)
		if err != nil {
			t.Fatal(err)
		}
		var u string
		for row.Next(){
			err = row.Scan(&u)
		}
		if err != nil {
			t.Fatal(err)
		}
		println("Price: " + u)
		helpers2.AssertEqual(t,dishcreate.Update,u,"")
	})
	t.Run("It return 200 on updating a dish by admin login", func(t *testing.T) {
		dishcreate := dish{
			Id:     "1",
			Rid:    "1",
			Flag:   "1",
			Update: "100",
		}
		b, _ := json.Marshal(Alogin)
		request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewReader(b))
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
		helpers2.AssertEqual(t, 200, response.Code, "")
		q := fmt.Sprintf("select price from dish where Id = '1' and status = '1'")
		row,err := db.Query(q)
		if err != nil {
			t.Fatal(err)
		}
		var u string
		for row.Next(){
			err = row.Scan(&u)
		}
		if err != nil {
			t.Fatal(err)
		}
		println("Price: " + u)
		helpers2.AssertEqual(t,dishcreate.Update,u,"")
		res := response.Body.String()
		helpers2.AssertEqual(t,"dish updated",res,"")
	})
	t.Run("It return 200 on updating a dish by user login", func(t *testing.T) {
		dishcreate := dish{
			Id:     "1",
			Rid:    "1",
			Flag:   "1",
			Update: "100",
		}
		b, _ := json.Marshal(Ulogin)
		request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewReader(b))
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
		helpers2.AssertEqual(t, 200, response.Code, "")
		q := fmt.Sprintf("select price from dish where Id = '1' and status = '1'")
		row,err := db.Query(q)
		if err != nil {
			t.Fatal(err)
		}
		var u string
		for row.Next(){
			err = row.Scan(&u)
		}
		if err != nil {
			t.Fatal(err)
		}
		println("Price: " + u)
		helpers2.AssertEqual(t,dishcreate.Update,u,"")
		res := response.Body.String()
		helpers2.AssertEqual(t,"dish updated",res,"")
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
		helpers2.AssertEqual(t, 200, response.Code, "")
		q := fmt.Sprintf("select name from dish where Id = '1' and status = '1'")
		row,err := db.Query(q)
		if err != nil {
			t.Fatal(err)
		}
		var u string
		for row.Next(){
			err = row.Scan(&u)
		}
		if err != nil {
			t.Fatal(err)
		}
		println("Name: " + u)
		helpers2.AssertEqual(t,"",u,"")
		res := response.Body.String()
		helpers2.AssertEqual(t,"dish deleted",res,"")
		helpers2.ResetDB(db)
	})
	t.Run("It return 200 on deleting a dish on admin login", func(t *testing.T) {
		dishcreate := dish{
			Id:  "1",
			Rid: "1",
		}
		b, _ := json.Marshal(Alogin)
		request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewReader(b))
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
		helpers2.AssertEqual(t, 200, response.Code, "")
		q := fmt.Sprintf("select name from dish where Id = '1' and status = '1'")
		row,err := db.Query(q)
		if err != nil {
			t.Fatal(err)
		}
		var u string
		for row.Next(){
			err = row.Scan(&u)
		}
		if err != nil {
			t.Fatal(err)
		}
		println("Name: " + u)
		helpers2.AssertEqual(t,"",u,"")
		res := response.Body.String()
		helpers2.AssertEqual(t,"dish deleted",res,"")
		helpers2.ResetDB(db)
	})
	t.Run("It return 200 on deleting a dish on user login", func(t *testing.T) {
		dishcreate := dish{
			Id:  "1",
			Rid: "1",
		}
		b, _ := json.Marshal(Ulogin)
		request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewReader(b))
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
		helpers2.AssertEqual(t, 200, response.Code, "")
		q := fmt.Sprintf("select name from dish where Id = '1' and status = '1'")
		row,err := db.Query(q)
		if err != nil {
			t.Fatal(err)
		}
		var u string
		for row.Next(){
			err = row.Scan(&u)
		}
		if err != nil {
			t.Fatal(err)
		}
		println("Name: " + u)
		helpers2.AssertEqual(t,"",u,"")
		res := response.Body.String()
		helpers2.AssertEqual(t,"dish deleted",res,"")
		helpers2.ResetDB(db)
	})
}
