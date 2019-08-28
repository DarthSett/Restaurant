package controller

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/restaurant/pkg/database"
	"github.com/restaurant/pkg/models"
	"strconv"
)

type RestController struct {
	database.Database
}

func NewRestController(db database.Database) *RestController {
	return &RestController{db}
}

//todo: replace panics with errors

func (u *RestController) RestMake(c *gin.Context) {

	//name := c.PostForm("name")
	//lat := c.PostForm("lat")
	//long := c.PostForm("long")
	//owner := c.PostForm("owner")
	input := make(map[string]string)
	err := c.BindJSON(&input)
	if err != nil {
		panic("Error getting inputs: " + err.Error())
	}

	if input["name"] == "" {
		panic("no name sent")
	}
	if input["lat"] == "" {
		panic("no lat sent")
	}
	if input["long"] == "" {
		panic("no long sent")
	}
	if input["owner"] == "" {
		panic("no owner sent")
	}

	clms, _ := c.Get("claims")
	claims := clms.(jwt.MapClaims)
	adder, _ := strconv.Atoi(fmt.Sprintf("%v", claims["id"]))
	if err != nil {
		panic("Error getting id of admin: " + err.Error())
	}
	rank, _ := strconv.Atoi(fmt.Sprintf("%v", claims["rank"]))
	if err != nil {
		panic("Error getting rank of admin: " + err.Error())
	}
	owner := input["owner"]
	rest := models.AddRestaurant(input["name"], input["lat"], input["long"], owner, adder, rank)
	err = u.CreateRestaurant(rest)
	if err != nil {
		panic("Error while saving Restaurant to db: " + err.Error())
	}
	c.Writer.Write([]byte("Restaurant Saved"))

}

func (u *RestController) Restget(c *gin.Context) {
	//name := c.PostForm("name")
	input := make(map[string]string)
	err := c.BindJSON(&input)
	if err != nil {
		panic("error getting inputs: " + err.Error())
	}

	if input["rid"] == "" {
		panic("no rid sent")
	}

	rid, err := strconv.Atoi(input["rid"])
	if err != nil {
		panic("Error getting rid from input: " + err.Error())
	}
	rest, err := u.GetRestaurant(rid)
	if err != nil {
		panic("Error getting rest from db: " + err.Error())
	}
	ids, names, prices, err := u.GetMenu(rid)
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
	//lat,long := rest.GetLoc()
	//loc := fmt.Sprintf("(%v,%v)",lat,long)
	c.JSON(200, gin.H{
		"Restaurant": rest,
		"Menu":       menu,
	})
	//c.Writer.Write([]byte("Name: " + rest.Name + "\nLoc: " + loc + "\nAdder: " + rest.GetAdder() + "\nOwner" + rest.GetOwner() ))
}

func (u *RestController) RestDel(c *gin.Context) {
	//name := c.PostForm("name")
	input := make(map[string]string)
	err := c.BindJSON(&input)
	if err != nil {
		panic("Error getting inputs: " + err.Error())
	}

	if input["rid"] == "" {
		panic("no rid sent")
	}

	rid, err := strconv.Atoi(input["rid"])
	if err != nil {
		panic("Error getting rid from input: " + err.Error())
	}
	clms, _ := c.Get("claims")
	claims := clms.(jwt.MapClaims)
	rank, err := strconv.Atoi(fmt.Sprintf("%v", claims["rank"]))
	if err != nil {
		panic("Error getting rank of admin: " + err.Error())
	}
	adder, _ := strconv.Atoi(fmt.Sprintf("%v", claims["id"]))
	if err != nil {
		panic("Error getting id of admin: " + err.Error())
	}
	err = u.DeleteRestaurant(rid, adder, rank)
	if err != nil {
		panic("Error Deleting rest from db: " + err.Error())
	}
	c.Writer.Write([]byte("rid: " + input["rid"] + " Deleted from db"))
}

func (u *RestController) RestUpdate(c *gin.Context) {
	//name := c.PostForm("name")
	//f := c.PostForm("flag")
	input := make(map[string]string)
	err := c.BindJSON(&input)
	if err != nil {
		panic("Error getting inputs: " + err.Error())
	}

	if input["rid"] == "" {
		panic("no rid sent")
	}
	if input["flag"] == "" {
		panic("no flag sent")
	}
	if input["update1"] == "" {
		panic("no update sent")
	}

	rid, err := strconv.Atoi(input["rid"])
	if err != nil {
		panic("Error getting rid from input: " + err.Error())
	}
	flag, err := strconv.Atoi(input["flag"])
	if err != nil {
		panic("error converting flag to int: " + err.Error())
	}

	clms, _ := c.Get("claims")
	claims := clms.(jwt.MapClaims)

	rank, _ := strconv.Atoi(fmt.Sprintf("%v", claims["rank"]))
	user, _ := strconv.Atoi(fmt.Sprintf("%v", claims["id"]))

	rest, err := u.GetRestaurant(rid)
	if err != nil {
		panic("Error getting rest from db: " + err.Error())
	}

	RestOwner := rest.GetOwner()
	RestAdder := rest.GetAdder()
	RestAdderRole := rest.GetAdderRole()

	if rank == 2 {
		err = u.UpdateRest(rid, input["update1"], input["update2"], flag)

	} else if (RestAdder == user && RestAdderRole == rank) || (RestOwner == claims["email"]) {
		err = u.UpdateRest(rid, input["update1"], input["update2"], flag)
	} else {
		panic("Admin is not the adder or owner of the restaurant")
	}

	if err != nil {
		panic("Error while updating in db: " + err.Error())
	}
	c.Writer.Write([]byte("Rest Updated"))
}

func (u *RestController) AddDish(c *gin.Context) {
	//restName := c.PostForm("name")
	//dishName := c.PostForm("dish")
	//menu := c.PostForm("menu")
	//p := c.PostForm("price")
	input := make(map[string]string)
	err := c.BindJSON(&input)
	if err != nil {
		panic("Error getting inputs: " + err.Error())
	}
	clms, _ := c.Get("claims")
	claims := clms.(jwt.MapClaims)
	rank, err := strconv.Atoi(fmt.Sprintf("%v", claims["rank"]))
	if err != nil {
		panic("Error getting rank from claims: " + err.Error())
	}
	price, err := strconv.Atoi(fmt.Sprintf("%v", input["price"]))
	if err != nil {
		panic("Error getting price from input: " + err.Error())
	}
	rid, err := strconv.Atoi(fmt.Sprintf("%v", input["rid"]))
	if err != nil {
		panic("Error getting rid from input: " + err.Error())
	}
	rest, err := u.GetRestaurant(rid)
	if err != nil {
		panic("Error getting rest from db: " + err.Error())
	}
	RestOwner := rest.GetOwner()
	RestAdder := rest.GetAdder()
	RestAdderRole := rest.GetAdderRole()
	adder, _ := strconv.Atoi(fmt.Sprintf("%v", claims["id"]))
	if err != nil {
		panic("Error getting rank of admin: " + err.Error())
	}

	if rank == 2 {
		err = u.CreateDish(models.NewDish(input["dish"], price, rid, 0, adder, rank))

	} else if rank == 1 {
		if (RestAdder == adder && RestAdderRole == rank) || (RestOwner == claims["email"]) {
			err = u.CreateDish(models.NewDish(input["dish"], price, rid, 0, adder, rank))
		} else {
			panic("Admin is not the adder or owner of the restaurant")

		}
	} else if rank == 0 {
		if RestOwner != claims["email"] {
			panic("user is not the owner of the restaurant")
		} else {
			err = u.CreateDish(models.NewDish(input["dish"], price, rid, 0, adder, rank))
		}
	}
	if err != nil {
		panic("error adding dish to db: " + err.Error())
	}
	c.Writer.Write([]byte("dish created"))
}

func (u *RestController) DelDish(c *gin.Context) {
	input := make(map[string]string)
	err := c.BindJSON(&input)
	if err != nil {
		panic("Error getting inputs: " + err.Error())
	}
	rid, err := strconv.Atoi(input["rid"])
	if err != nil {
		panic("Error getting rid from input: " + err.Error())
	}
	id, err := strconv.Atoi(input["id"])
	if err != nil {
		panic("Error getting rid from input: " + err.Error())
	}
	clms, _ := c.Get("claims")
	claims := clms.(jwt.MapClaims)
	rank, err := strconv.Atoi(fmt.Sprintf("%v", claims["rank"]))
	if err != nil {
		panic("Error getting rank from claims: " + err.Error())
	}
	user, _ := strconv.Atoi(fmt.Sprintf("%v", claims["id"]))
	if err != nil {
		panic("Error getting id of admin: " + err.Error())
	}
	rest, err := u.GetRestaurant(rid)
	if err != nil {
		panic("Error getting rest from db: " + err.Error())
	}
	owner := rest.GetOwner()
	adder := rest.GetAdder()
	adderRole := rest.GetAdderRole()

	if rank == 2 {
		err = u.DeleteDish(id)

	} else if rank == 1 {
		if (adder == user && adderRole == rank) || owner == claims["email"] {
			err = u.DeleteDish(id)
		} else {
			panic("Admin is not the adder or owner of the restaurant")

		}
	} else if rank == 0 {
		if owner != claims["email"] {
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
	input := make(map[string]string)
	err := c.BindJSON(&input)
	if err != nil {
		panic("Error getting inputs: " + err.Error())
	}
	if input["id"] == "" {
		panic("no id sent")
	}
	if input["rid"] == "" {
		panic("no rid sent")
	}
	if input["flag"] == "" {
		panic("no flag sent")
	}
	if input["update"] == "" {
		panic("no update sent")
	}

	rid, err := strconv.Atoi(input["rid"])
	if err != nil {
		panic("Error getting rid from input: " + err.Error())
	}
	id, err := strconv.Atoi(input["id"])
	if err != nil {
		panic("Error getting rid from input: " + err.Error())
	}
	clms, _ := c.Get("claims")
	claims := clms.(jwt.MapClaims)
	flag, err := strconv.Atoi(input["flag"])
	if err != nil {
		panic("Error getting flag: " + err.Error())
	}
	rank, _ := strconv.Atoi(fmt.Sprintf("%v", claims["rank"]))
	user, _ := strconv.Atoi(fmt.Sprintf("%v", claims["id"]))
	rest, err := u.GetRestaurant(rid)
	if err != nil {
		panic("Error getting rest from db: " + err.Error())
	}
	owner := rest.GetOwner()
	adder := rest.GetAdder()
	adderRole := rest.GetAdderRole()

	if rank == 2 {
		err = u.UpdateDish(id, input["update"], flag)
	} else if adder == user && adderRole == rank {
		err = u.UpdateDish(id, input["update"], flag)
	} else if owner == claims["email"] {
		err = u.UpdateDish(id, input["update"], flag)
	} else {
		panic("user is not the owner or adder of the restaurant")
	}
	if err != nil {
		panic("error deleting dish from db: " + err.Error())
	}
	c.Writer.Write([]byte("dish updated"))
}

func (u *RestController) MenuGet(c *gin.Context) {
	input := make(map[string]string)
	err := c.BindJSON(&input)
	if err != nil {
		panic("Error getting inputs: " + err.Error())
	}
	rid, err := strconv.Atoi(input["rid"])
	if err != nil {
		panic("Error getting rid from input: " + err.Error())
	}
	ids, names, prices, err := u.GetMenu(rid)
	if err != nil {
		panic("Error getting menu from db: " + err.Error())
	}
	o := make([]struct {
		Id    int
		Name  string
		Price int
	}, len(ids))
	for i := range ids {
		o[i].Id = ids[i]
		o[i].Name = names[i]
		o[i].Price = prices[i]
	}
	c.JSON(200, o)

}

func (u *RestController) Getbydist(c *gin.Context) {

	input := make(map[string]string)
	err := c.BindJSON(&input)
	if err != nil {
		panic("Error getting inputs: " + err.Error())
	}
	lat, err := strconv.ParseFloat(input["lat"], 64)
	if err != nil {
		panic(err)
	}
	long, err := strconv.ParseFloat(input["long"], 64)
	if err != nil {
		panic(err)
	}
	dist, err := strconv.ParseFloat(input["dist"], 64)
	if err != nil {
		panic(err)
	}
	names, id := u.GetbyDistance(lat, long, dist)
	o := make([]struct {
		Id   int
		Name string
	}, len(id))
	for i := range id {
		o[i].Name = names[i]
		o[i].Id = id[i]
	}
	println(o[0].Name)
	c.JSON(200, o)

}

func (u *RestController) ListRest(c *gin.Context) {
	id, name, err := u.RestList()
	if err != nil {
		panic("Error getting list from db: " + err.Error())
	}
	o := make([]struct {
		Id   int
		Name string
	}, len(id))
	for i := range id {
		o[i].Id = id[i]
		o[i].Name = name[i]
	}
	c.JSON(200, o)
}

//MySQL relational
//A<->B
//
//id int-autoincrement
//status - 1/0
//created - time
//updated - time
//deleted - time
//
//Resturant
//id, name, lat, lng, owner, added_by, created, status,role
//
//Owner
//id, name, email, password, created_by, status, role
//
//Menu/dish
//id, name, price, resurant_id, status, created
//
