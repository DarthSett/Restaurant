package database

import "github.com/restaurant/pkg/models"

type Database interface {
	CreateCust(c *models.Customer) error
	//UpdCust(id int, update string, flag int) error
	//DelCust(id int) error
	GetCust(id int,email string) (*models.Customer, error)
	//GetRestTable(rid int)	([]int, error)
	//BookTable(id int,rid int,tid int) error
	RestList() ([]int, []string, error)
	Checktoken(token string) (bool, error)
	LogoutUser(token string) error
}
