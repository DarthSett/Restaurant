package controllers_test

import (
	"bytes"
	"encoding/json"
	"github.com/restaurant/Customer_MicroService/pkg/database/mysql"
	"github.com/restaurant/Customer_MicroService/pkg/server"
	"github.com/restaurant/Res_Man_MicroService/pkg/helpers"
	"github.com/restaurant/pkg/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	db1 = mysql.NewMySqlDB("localhost", "root", "password", "3306", "Restaurant_Test")
	router = server.NewRouter(db1)
	s = router.Router()
)
func TestCustController_CustMake(t *testing.T) {
	t.Run("It returns 200 on saving a customer", func(t *testing.T) {
		c := models.Customer{
			Name:   "user1",
			Email:  "user1@gmail.com",
			Pass:   "gutthix",
		}
		res := SendReq(t,http.MethodPost,"/create",c,"")
		//cust,err := json.Marshal(c)
		//if err != nil {
		//	t.Fatal(err)
		//}
		//req := httptest.NewRequest(http.MethodPost,"/create",bytes.NewReader(cust))
		//res := httptest.NewRecorder()
		//s.ServeHTTP(res,req)
		code := res.Code
		helpers.AssertEqual(t,200,code,"")
		response := res.Body.String()
		helpers.AssertEqual(t,"Customer saved",response,"")
	})
	t.Run("It returns 500 on saving a customer without Name", func(t *testing.T) {
		c := models.Customer{
			Name:   "",
			Email:  "user1@gmail.com",
			Pass:   "gutthix",
		}
		res := SendReq(t,http.MethodPost,"/create",c,"")
		//cust,err := json.Marshal(c)
		//if err != nil {
		//	t.Fatal(err)
		//}
		//req := httptest.NewRequest(http.MethodPost,"/create",bytes.NewReader(cust))
		//res := httptest.NewRecorder()
		//s.ServeHTTP(res,req)
		code := res.Code
		helpers.AssertEqual(t,400,code,"")
	})


}

func TestCustController_CustLogin(t *testing.T) {
	t.Run("It returns 200 on logging in", func(t *testing.T) {
		c := models.Credentials{
			Email: "user1@gmail.com",
			Pass:  "gutthix",
		}
		res := SendReq(t,http.MethodPost,"/login",c,"")
		code := res.Code
		helpers.AssertEqual(t,200,code,"")
		response := res.Body.String()
		helpers.AssertEqual(t,"Customer logged in. Token generated",response,"")
	})
	t.Run("It returns 500 on wrong email", func(t *testing.T) {
		c := models.Credentials{
			Email: "user2@gmail.com",
			Pass:  "gutthix",
		}
		res := SendReq(t,http.MethodPost,"/login",c,"")
		code := res.Code
		helpers.AssertEqual(t,500,code,"")
	})
	t.Run("It returns 400 on wrong password", func(t *testing.T) {
		c := models.Credentials{
			Email: "user1@gmail.com",
			Pass:  "gutthix1",
		}
		res := SendReq(t,http.MethodPost,"/login",c,"")
		code := res.Code
		helpers.AssertEqual(t,400,code,"")
	})
}

func TestCustController_CustGetResList(t *testing.T) {
	t.Run("It returns 200 with the list", func(t *testing.T) {
		c := models.Credentials{
			Email: "user1@gmail.com",
			Pass:  "gutthix",
		}
		res := SendReq(t,http.MethodPost,"/login",c,"")
		token := res.Header().Get("token")
		res = SendReq(t,http.MethodGet,"/list","",token)
		code := res.Code
		helpers.AssertEqual(t,200,code,"")
		response := res.Body.String()
		helpers.AssertEqual(t,"[{\"Id\":1,\"Name\":\"McDonalds\"}]",response,"")
	})

}

func TestCustController_Logout(t *testing.T) {
	c := models.Credentials{
		Email: "user1@gmail.com",
		Pass:  "gutthix",
	}
	res := SendReq(t,http.MethodPost,"/login",c,"")
	token := res.Header().Get("token")
	res = SendReq(t,http.MethodGet,"/logout","",token)
	code := res.Code
	helpers.AssertEqual(t,200,code,"")
	response := res.Body.String()
	helpers.AssertEqual(t,"User Logged Out",response,"")
}

func SendReq(t *testing.T,method string,path string,c interface{},token string) *httptest.ResponseRecorder {
	cust,err := json.Marshal(c)
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(method,path,bytes.NewReader(cust))
	req.Header.Set("token",token)
	res := httptest.NewRecorder()
	s.ServeHTTP(res,req)
	return res
}
