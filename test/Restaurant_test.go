package test_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/restaurant/pkg/database/mysql"
	Server "github.com/restaurant/pkg/server"
	"net/http"
	"net/http/httptest"
	"testing"
)

//todo: write integration tests


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
	assertEqual(t,200,response.Code,"")
}



func TestCreate(t *testing.T) {
	type Login struct {
		Email string `json:"email"`
		Pass string `json:"pass"`
	}
	saLogin := Login{Email:"sourav241196@gmail.com",Pass:"zamorak"}
	type create struct {
		Name string `json:"name"`
		Pass string `json:"pass"`
		Email string `json:"email"`
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
	t.Run("it returns 200 on creating a user", func(t *testing.T) {
		userdata := create{
			Name:  "user2",
			Pass:  "zamorak",
			Email: "user2@gmail.com",
		}
		user,_ := json.Marshal(userdata)
		request,_ := http.NewRequest(http.MethodPost, "/user/create",bytes.NewReader(user))
		request.Header.Set("token",token)
		request.Header.Set("Content-Type","application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response,request)
		assertEqual(t,200,response.Code,"")
		resetDBafterCreate(db)
	})
	t.Run("it returns 200 on creating an admin", func(t *testing.T) {
		userdata := create{
			Name:  "admin2",
			Pass:  "zamorak",
			Email: "admin2@gmail.com",
		}
		admin,_ := json.Marshal(userdata)
		request,_ := http.NewRequest(http.MethodPost, "/admin/create",bytes.NewReader(admin))
		request.Header.Set("token",token)
		request.Header.Set("Content-Type","application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response,request)
		assertEqual(t,200,response.Code,"")
		resetDBafterCreate(db)
	})
	t.Run("it returns 200 on creating a superadmin", func(t *testing.T) {
		userdata := create{
			Name:  "superadmin1",
			Pass:  "zamorak",
			Email: "superadmin1@gmail.com",
		}
		superadmin,_ := json.Marshal(userdata)
		request,_ := http.NewRequest(http.MethodPost, "/superadmin/create",bytes.NewReader(superadmin))
		request.Header.Set("token",token)
		request.Header.Set("Content-Type","application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response,request)
		assertEqual(t,200,response.Code,"")
		resetDBafterCreate(db)
	})
	t.Run("integration test for creating user", func(t *testing.T) {
		userdata := make([]create,5)
		login := make([]Login,3)
		login[0] = Login{
			Email: "sourav241196@gmail.com",
			Pass:  "zamorak",
		}
		login[1] = Login{
			Email: "user1@gmail.com",
			Pass:  "gutthix",
		}
		login[2] = Login{
			Email: "admin1@gmail.com",
			Pass:  "gutthix",
		}
		userdata[0] = create{
			Name:  "user2",
			Pass:  "zamorak",
			Email: "user2@gmail.com",
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
		for i,_ :=range data {

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
			assertEqual(t,data[i].expected,response.Code,"")
			resetDBafterCreate(db)
		}

	})


}

func TestList(t *testing.T) {
	type Login struct {
		Email string `json:"email"`
		Pass string `json:"pass"`
	}
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
	assertEqual(t,200,response.Code,"")
}

func TestGet(t *testing.T) {

	type Login struct {
		Email string `json:"email"`
		Pass string `json:"pass"`
	}
	saLogin := Login{Email:"sourav241196@gmail.com",Pass:"zamorak"}
	type get struct {
		Id	string	`json:"id"`
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
	t.Run("it returns 200 on getting a user", func(t *testing.T) {
		userget := get{"1"}
		user,_ := json.Marshal(userget)
		request,_ := http.NewRequest(http.MethodGet, "/user/get",bytes.NewReader(user))
		request.Header.Set("token",token)
		request.Header.Set("Content-Type","application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response,request)
		assertEqual(t,200,response.Code,"")
	})
	t.Run("it returns 200 on getting an admin", func(t *testing.T) {
		userget := get{"1"}
		admin,_ := json.Marshal(userget)
		request,_ := http.NewRequest(http.MethodGet, "/admin/get",bytes.NewReader(admin))
		request.Header.Set("token",token)
		request.Header.Set("Content-Type","application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response,request)
		assertEqual(t,200,response.Code,"")
	})
	t.Run("it returns 200 on getting an admin", func(t *testing.T) {
		userget := get{"1"}
		superadmin,_ := json.Marshal(userget)
		request,_ := http.NewRequest(http.MethodGet, "/superadmin/get",bytes.NewReader(superadmin))
		request.Header.Set("token",token)
		request.Header.Set("Content-Type","application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response,request)
		assertEqual(t,200,response.Code,"")
	})
}

func TestLogin(t *testing.T) {
	type Login struct {
		Email string `json:"email"`
		Pass string `json:"pass"`
	}


	db := mysql.NewMySqlDB("127.0.0.1","root","Zamorak1","3306","Restaurant_Test")
	defer db.Close()
	router := Server.NewRouter(db)
	s := router.Router()

	t.Run("it returns 200 on user login", func(t *testing.T) {
		saLogin := Login{Email:"user1@gmail.com",Pass:"gutthix"}
		b, _ := json.Marshal(saLogin)

		request,_ := http.NewRequest(http.MethodPost, "/user/login",bytes.NewReader(b))
		request.Header.Set("Content-Type","application/json")
		response :=httptest.NewRecorder()
		s.ServeHTTP(response,request)
		assertEqual(t,200,response.Code,"")
	})
	t.Run("it returns 200 on admin login", func(t *testing.T) {
		saLogin := Login{Email:"admin1@gmail.com",Pass:"gutthix"}
		b, _ := json.Marshal(saLogin)

		request,_ := http.NewRequest(http.MethodPost, "/admin/login",bytes.NewReader(b))
		request.Header.Set("Content-Type","application/json")
		response :=httptest.NewRecorder()
		s.ServeHTTP(response,request)
		assertEqual(t,200,response.Code,"")
	})
	t.Run("it returns 200 on superadmin login", func(t *testing.T) {
		saLogin := Login{Email:"sourav241196@gmail.com",Pass:"zamorak"}
		b, _ := json.Marshal(saLogin)

		request,_ := http.NewRequest(http.MethodPost, "/superadmin/login",bytes.NewReader(b))
		request.Header.Set("Content-Type","application/json")
		response :=httptest.NewRecorder()
		s.ServeHTTP(response,request)
		assertEqual(t,200,response.Code,"")
	})

}


func TestUpdate (t *testing.T) {
	type Login struct {
		Email string `json:"email"`
		Pass string `json:"pass"`
	}
	saLogin := Login{Email:"sourav241196@gmail.com",Pass:"zamorak"}
	type upd struct {
		Id	string	`json:"id"`
		Flag string `json:"flag"`
		Update string `json:"update"`
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
	t.Run("It returns 200 on user update", func(t *testing.T) {
		userget := upd{
			Id:     "1",
			Flag:   "1",
			Update: "gutthix",
		}
		user,_ := json.Marshal(userget)
		request,_ := http.NewRequest(http.MethodPut, "/user/update",bytes.NewReader(user))
		request.Header.Set("token",token)
		request.Header.Set("Content-Type","application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response,request)
		assertEqual(t,200,response.Code,"")
	})
	t.Run("it returns 200 on admin update", func(t *testing.T) {
		userget := upd{
			Id:     "1",
			Flag:   "1",
			Update: "gutthix",
		}
		admin,_ := json.Marshal(userget)
		request,_ := http.NewRequest(http.MethodPut, "/admin/update",bytes.NewReader(admin))
		request.Header.Set("token",token)
		request.Header.Set("Content-Type","application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response,request)
		assertEqual(t,200,response.Code,"")
	})
}

func TestDel(t *testing.T) {

	type Login struct {
		Email string `json:"email"`
		Pass string `json:"pass"`
	}
	saLogin := Login{Email:"sourav241196@gmail.com",Pass:"zamorak"}
	type del struct {
		Id	string	`json:"id"`
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
	t.Run("it returns 200 on deleting a user", func(t *testing.T) {
		userget := del{"22"}
		user,_ := json.Marshal(userget)
		request,_ := http.NewRequest(http.MethodDelete, "/user/del",bytes.NewReader(user))
		request.Header.Set("token",token)
		request.Header.Set("Content-Type","application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response,request)
		assertEqual(t,200,response.Code,"")
	})
	t.Run("it returns 200 on deleting an admin", func(t *testing.T) {
		userget := del{"5"}
		admin,_ := json.Marshal(userget)
		request,_ := http.NewRequest(http.MethodDelete, "/admin/del",bytes.NewReader(admin))
		request.Header.Set("token",token)
		request.Header.Set("Content-Type","application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response,request)
		assertEqual(t,200,response.Code,"")
	})
	t.Run("it returns 200 on deleting superadmin", func(t *testing.T) {
		request,_ :=http.NewRequest(http.MethodDelete,"/superadmin/del",nil)
		request.Header.Set("token",token)
		response = httptest.NewRecorder()
		s.ServeHTTP(response,request)
		assertEqual(t,200,response.Code,"")
	})
	resetDBafterDelete(db,0)
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
	type Login struct {
		Email string `json:"email"`
		Pass string `json:"pass"`
	}
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
		assertEqual(t,200,response.Code,"")
		resetDBafterCreate(db)
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
		assertEqual(t,200,response.Code,"")
	})
	t.Run("It returns 200 on updating a rest", func(t *testing.T) {
		restcreate := rest{
			Id:"1",
			Update1:"user2@gmail.com",
			Update2:"",
			Flag:"2",
		}
		rest,_ := json.Marshal(restcreate)
		request,_ := http.NewRequest(http.MethodPut, "/rest/update",bytes.NewReader(rest))
		request.Header.Set("token",token)
		request.Header.Set("Content-Type","application/json")
		response = httptest.NewRecorder()
		s.ServeHTTP(response,request)
		assertEqual(t,200,response.Code,"")
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
		assertEqual(t,200,response.Code,"")
		resetDBafterDelete(db,1)
	})
	t.Run("It returns 200 on list", func(t *testing.T) {
		request,_ := http.NewRequest(http.MethodGet, "/rest/list",nil)
		request.Header.Set("token",token)
		response = httptest.NewRecorder()
		s.ServeHTTP(response,request)
		assertEqual(t,200,response.Code,"")
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
		assertEqual(t,200,response.Code,"")
	})

}

func TestDish(t *testing.T) {
	type Login struct {
		Email string `json:"email"`
		Pass  string `json:"pass"`
	}
	saLogin := Login{Email: "sourav241196@gmail.com", Pass: "zamorak"}
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
		assertEqual(t, 200, response.Code, "")
		resetDBafterCreate(db)
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
		assertEqual(t, 200, response.Code, "")
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
		assertEqual(t, 200, response.Code, "")
		resetDBafterDelete(db,2)
	})
}

func TestLogout(t *testing.T){
	type Login struct {
		Email string `json:"email"`
		Pass string `json:"pass"`
	}
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
	t.Run("It returns 200 on logout", func(t *testing.T) {
		request,_ := http.NewRequest(http.MethodGet,"/logout",nil)
		request.Header.Set("token",token)
		response := httptest.NewRecorder()
		s.ServeHTTP(response,request)
		assertEqual(t,200,response.Code,"")
	})
}



func assertEqual(t *testing.T, expected interface{}, got interface{}, message string) {
	if expected == got {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", expected, got)
	}
	t.Fatal(message)
}

func resetDBafterCreate(db *mysql.MysqlDB) {
	println("@@@@@@@@@@@@@@")
	db.Query("delete from superadmin where id <> '1'")
	db.Query("ALTER table superadmin AUTO_INCREMENT=2")
	db.Query("delete from admin where id <> '1'")
	db.Query("ALTER table admin AUTO_INCREMENT=2")
	db.Query("delete from user where id <> '1'")
	db.Query("ALTER table user AUTO_INCREMENT=2")
	db.Query("delete from rest where id <> '1'")
	db.Query("ALTER table rest AUTO_INCREMENT=19")
	db.Query("delete from dish where id <> '1'")
	db.Query("ALTER table dish AUTO_INCREMENT=2")

}

func resetDBafterDelete(db *mysql.MysqlDB,flag int) {
	if flag == 0{
		db.Query("ALTER table User AUTO_INCREMENT=1")
		db.Query("INSERT INTO User VALUES ('user1', '$2a$04$y9ExoM60HYRfNxm8N/tE3.YHVS/RhHB/6eaztdwVYhoRPspofsmk2', 'user1@gmail.com', '1', '1', '1', '1', '2019-08-09 14:12:25')")
		db.Query("ALTER table admin AUTO_INCREMENT=1")
		db.Query("INSERT INTO Admin VALUES ('admin1', '$2a$04$AKe7E84SCtQZCjPYpRCN6OrdMS/VnV0Qx93Os.TU8nO71Tt67ysmG', 'admin1@gmail.com', '1', '2', '1', '1', '2019-08-08 15:47:35')")
		db.Query("ALTER table superadmin AUTO_INCREMENT=1")
		db.Query("INSERT INTO superadmin VALUES ('Sourav', '$2a$04$CAIRyD6NVptXbB25PEUFJeYEsBwIouUYLigBhWocbcmvOZrc7.OV.', 'sourav241196@gmail.com', '0', '2', '1', '1', '2019-08-08 14:14:10')")
	} else if flag == 1 {
		db.Query("ALTER table rest AUTO_INCREMEnt=2")
		db.Query("update rest set status = '1' where id = '1'")
	} else if flag == 2 {
		db.Query("ALTER table dish AUTO_INCREMENT=2")
		db.Query("update dish set status = '1' where id = '1'")
	}

}