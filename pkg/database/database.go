package database

import "github.com/restaurant/pkg/models"

type Database interface {
	CreateUser			(u *models.User) 													error
	CreateDish			(d *models.Dish)													error
	CreateRestaurant 	(r *models.Restaurant)												error
	DeleteDish			(id int)															error
	DeleteRestaurant	(id int, adder string, rank int)									error
	DeleteUser			(email string,adder string)											error
	UpdateDish			(id int, update string, flag int) 									error
	UpdateRest			(id int, update1 string, update2 string, flag int)					error
	GetUser				(email string)														(*models.User,error)
	GetRestaurant		(id int) 															(*models.Restaurant,error)
	//GetDish				(id int)															(*models.Dish,error)
	CreateAdmin			(u *models.User) 													error
	DeleteAdmin			(email string) 														error
	GetAdmin			(email string)														(*models.User,error)
	CreateSuperAdmin	(u *models.User) 													error
	DeleteSuperAdmin	(email string) 														error
	GetSuperAdmin		(email string)														(*models.User,error)
	GetMenu				(rid int,menu string)												(map[string]int,error)
	GetbyDistance 		(Lat float64, Long float64,dist float64)							[]string
}
