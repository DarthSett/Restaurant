package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/restaurant/pkg/database"
	"github.com/restaurant/pkg/models"
	"strconv"
)




type RestController struct{
	database.Database

}

func NewRestController(db database.Database) *RestController {
	return &RestController{db}
}



func (u *RestController) RestMake(c *gin.Context){

	//name := c.PostForm("name")
	//lat := c.PostForm("lat")
	//long := c.PostForm("long")
	//owner := c.PostForm("owner")
	input := decodeJson(c)
	claims,err := GetTokenClaims(c.GetHeader("token"))
	if err != nil {
		panic("Error getting claims from token: " + err.Error())
	}
	adder := fmt.Sprintf("%v",claims["email"])
	rest := models.AddRestaurant(input["name"],input["lat"],input["long"],input["owner"],adder)
	err = u.CreateRestaurant(rest)
	if err != nil {
		panic("Error while saving Restaurant to db: " + err.Error())
	}
	c.Writer.Write([]byte("Restaurant Saved"))

}

func (u *RestController) Restget(c *gin.Context){
	//name := c.PostForm("name")
	input := decodeJson(c)
	rid,err := strconv.Atoi(input["rid"])
	if err != nil {
		panic("Error getting rid from input: " + err.Error())
	}
	rest,err := u.GetRestaurant(rid)
	if err != nil {
		panic("Error getting rest from db: " + err.Error())
	}
	//lat,long := rest.GetLoc()
	//loc := fmt.Sprintf("(%v,%v)",lat,long)
	encodeJson(c,rest)
	//c.Writer.Write([]byte("Name: " + rest.Name + "\nLoc: " + loc + "\nAdder: " + rest.GetAdder() + "\nOwner" + rest.GetOwner() ))
}

func (u *RestController) RestDel(c *gin.Context){
	//name := c.PostForm("name")
	input := decodeJson(c)
	rid,err := strconv.Atoi(input["rid"])
	if err != nil {
		panic("Error getting rid from input: " + err.Error())
	}
	claims,err := GetTokenClaims(c.GetHeader("token"))
	if err != nil {
		panic("Error getting claims from token: " + err.Error())
	}
	rank,err := strconv.Atoi(fmt.Sprintf("%v",claims["rank"]))
	if err != nil {
		panic("Error getting rank of admin: "+ err.Error())
	}
	adder := fmt.Sprintf("%v",claims["email"])
	err = u.DeleteRestaurant(rid,adder,rank)
	if err != nil {
		panic("Error getting user from db: "+ err.Error())
	}
	c.Writer.Write([]byte("rid: " + input["rid"] + " Deleted from db"))
}

func (u *RestController) RestUpdate(c *gin.Context){
	//name := c.PostForm("name")
	//f := c.PostForm("flag")
	input := decodeJson(c)
	rid,err := strconv.Atoi(input["rid"])
	if err != nil {
		panic("Error getting rid from input: " + err.Error())
	}
	flag,err := strconv.Atoi(input["flag"])
	if err != nil {
		panic("error converting flag to int: "+err.Error())
	}

	claims,err := GetTokenClaims(c.GetHeader("token"))
	if err != nil {
		panic("Error getting claims from token: " + err.Error())
	}

	rank,err := strconv.Atoi(fmt.Sprintf("%v",claims["rank"]))
	if err != nil {
		panic("Error getting claims from token: " + err.Error())
	}

	rest,err := u.GetRestaurant(rid)
	if err != nil {
		panic("Error getting rest from db: " + err.Error())
	}

	owner := rest.GetOwner()
	adder := rest.GetAdder()


	if rank == 2 {
		err = u.UpdateRest(rid,input["update1"],input["update2"],flag)

	} else if rank == 1 {
		if adder != fmt.Sprintf("%v",claims["email"]) || owner != fmt.Sprintf("%v",claims["email"]) {
			panic("Admin is not the adder or owner of the restaurant")
		} else {
			err = u.UpdateRest(rid,input["update1"],input["update2"],flag)
		}
	}

	if err != nil {
		panic("Error while updating in db: " + err.Error())
	}
}

func (u *RestController) AddDish(c *gin.Context) {
	//restName := c.PostForm("name")
	//dishName := c.PostForm("dish")
	//menu := c.PostForm("menu")
	//p := c.PostForm("price")
	input := decodeJson(c)
	claims,err := GetTokenClaims(c.GetHeader("token"))
	if err != nil {
		panic("Error getting claims from token: " + err.Error())
	}
	rank,err := strconv.Atoi(fmt.Sprintf("%v",claims["rank"]))
	if err != nil {
		panic("Error getting rank from claims: " + err.Error())
	}
	price,err := strconv.Atoi(fmt.Sprintf("%v",input["price"]))
	if err != nil {
		panic("Error getting price from input: " + err.Error())
	}
	rid,err := strconv.Atoi(fmt.Sprintf("%v",input["rid"]))
	if err != nil {
		panic("Error getting rid from input: " + err.Error())
	}
	rest,err := u.GetRestaurant(rid)
	if err != nil {
		panic("Error getting rest from db: " + err.Error())
	}
	owner := rest.GetOwner()
	adder := rest.GetAdder()

	if rank == 2 {
		err = u.CreateDish(models.NewDish(input["dish"],price,rid,input["menu"],0))

	} else if rank == 1 {
		if adder != fmt.Sprintf("%v",claims["email"]) || owner != fmt.Sprintf("%v",claims["email"]) {
			panic("Admin is not the adder or owner of the restaurant")
		} else {
			err = u.CreateDish(models.NewDish(input["dish"],price,rid,input["menu"],0))
		}
	} else if rank == 0 {
		if owner != fmt.Sprintf("%v",claims["email"]) {
			panic("user is not the owner of the restaurant")
		} else {
			err = u.CreateDish(models.NewDish(input["dish"],price,rid,input["menu"],0))
		}
	}
	if err != nil {
		panic("error adding dish to db: " + err.Error())
	}
}

func (u *RestController) DelDish(c *gin.Context) {
	input := decodeJson(c)
	rid,err := strconv.Atoi(input["rid"])
	if err != nil {
		panic("Error getting rid from input: " + err.Error())
	}
	id,err := strconv.Atoi(input["id"])
	if err != nil {
		panic("Error getting rid from input: " + err.Error())
	}
	claims,err := GetTokenClaims(c.GetHeader("token"))
	if err != nil {
		panic("Error getting claims from token: " + err.Error())
	}
	rank,err := strconv.Atoi(fmt.Sprintf("%v",claims["rank"]))
	if err != nil {
		panic("Error getting rank from claims: " + err.Error())
	}
	rest,err := u.GetRestaurant(rid)
	if err != nil {
		panic("Error getting rest from db: " + err.Error())
	}
	owner := rest.GetOwner()
	adder := rest.GetAdder()

	if rank == 2 {
		err = u.DeleteDish(id)

	} else if rank == 1 {
		if adder != fmt.Sprintf("%v",claims["email"]) || owner != fmt.Sprintf("%v",claims["email"]) {
			panic("Admin is not the adder or owner of the restaurant")
		} else {
			err = u.DeleteDish(id)

		}
	} else if rank == 0 {
		if owner != fmt.Sprintf("%v",claims["email"]) {
			panic("user is not the owner of the restaurant")
		} else {
			err = u.DeleteDish(id)

		}
	}
	if err != nil {
		panic("error deleting dish from db: " + err.Error())
	}
}



func (u *RestController) UpdDish(c *gin.Context) {
	input := decodeJson(c)
	rid,err := strconv.Atoi(input["rid"])
	if err != nil {
		panic("Error getting rid from input: " + err.Error())
	}
	id,err := strconv.Atoi(input["id"])
	if err != nil {
		panic("Error getting rid from input: " + err.Error())
	}
	claims,err := GetTokenClaims(c.GetHeader("token"))
	if err != nil {
		panic("Error getting claims from token: " + err.Error())
	}
	flag,err := strconv.Atoi(input["flag"])
	if err != nil {
		panic("Error getting flag: " + err.Error())
	}
	rank,err := strconv.Atoi(fmt.Sprintf("%v",claims["rank"]))
	if err != nil {
		panic("Error getting rank from claims: " + err.Error())
	}
	rest,err := u.GetRestaurant(rid)
	if err != nil {
		panic("Error getting rest from db: " + err.Error())
	}
	owner := rest.GetOwner()
	adder := rest.GetAdder()

	if rank == 2 {
		err = u.UpdateDish(id,input["update"],flag)

	} else if rank == 1 {
		if adder != fmt.Sprintf("%v",claims["email"]) || owner != fmt.Sprintf("%v",claims["email"]) {
			panic("Admin is not the adder or owner of the restaurant")
		} else {
			err = u.UpdateDish(id,input["update"],flag)

		}
	} else if rank == 0 {
		if owner != fmt.Sprintf("%v",claims["email"]) {
			panic("user is not the owner of the restaurant")
		} else {
			err = u.UpdateDish(id,input["update"],flag)

		}
	}
	if err != nil {
		panic("error deleting dish from db: " + err.Error())
	}
}

func (u *RestController) MenuGet(c *gin.Context) {
	input := decodeJson(c)
	rid,err := strconv.Atoi(input["rid"])
	if err != nil {
		panic("Error getting rid from input: " + err.Error())
	}
	menu := make(map[string]int)
	menu,err = u.GetMenu(rid,input["menu"])
	if err != nil {
		panic("Error getting menu from db: " + err.Error())
	}
	encodeJson(c,menu)


}

func (u *RestController) Getbydist(c *gin.Context) {

	input := decodeJson(c)
	lat,err := strconv.ParseFloat(input["lat"],64)
	if err != nil {
		panic(err)
	}
	long,err := strconv.ParseFloat(input["long"],64)
	if err != nil {
		panic(err)
	}
	dist,err := strconv.ParseFloat(input["dist"],64)
	if err != nil {
		panic(err)
	}
	names := u.GetbyDistance(lat,long,dist)
	encodeJson(c,names)

}


func (u *RestController) ListRest (c *gin.Context) {
	id,name,err := u.RestList()
	if err != nil {
		panic("Error getting list from db: "+err.Error())
	}
	o := make(map[int]string)
	for i,v := range id {
		o[v] = name[i]
	}
	encodeJson(c,o)
}




