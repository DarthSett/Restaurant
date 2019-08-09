package database

import "github.com/restaurant/pkg/models"




type Database interface {
	CreateUser			(u *models.User) 													error
	GetUser				(email string,id int)												(*models.User,error)
	UpdateUser			(id int,update string,flag int) 									error
	DeleteUser			(id int)															error
	GetUserRests 		(email string) 														([]int,[]string,error)
	UserList			()																	([]string,[]string,[]int,error)
	CreateAdmin			(u *models.User) 													error
	GetAdmin			(email string,id int)												(*models.User,error)
	DeleteAdmin			(id int) 															error
	UpdateAdmin			(id int,update string,flag int) 									error
	CreateSuperAdmin	(u *models.User) 													error
	GetSuperAdmin		(email string,id int)												(*models.User,error)
	DeleteSuperAdmin	(id int) 															error
	CreateRestaurant 	(r *models.Restaurant)												error
	GetRestaurant		(id int) 															(*models.Restaurant,error)
	UpdateRest			(id int, update1 string, update2 string, flag int)					error
	GetMenu				(rid int)															([]int,[]string,[]int,error)
	DeleteRestaurant	(id int, adder int, rank int)										error
	GetbyDistance 		(Lat float64, Long float64,dist float64)							([]string,[]int)
	RestList			() 																	([]int,[]string,error)
	CreateDish			(d *models.Dish)													error
	UpdateDish			(id int, update string, flag int) 									error
	DeleteDish			(id int)															error
	LogoutUser			(token string) 														error
	Checktoken			(token string) 														(bool,error)
	//GetDish				(id int)														(*models.Dish,error)
}
