package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/restaurant/pkg/models"
	"math"
	"strconv"
)

type MysqlDB struct {
	*sql.DB
}

func NewMySqlDB(ip string, username string, password string, port string,schema string) *MysqlDB {

	conn :=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",username,password,ip,port,schema)
	DB,err := sql.Open("mysql", conn )
	if err != nil {
		panic("Can't connect to db" + err.Error())
	}
	return &MysqlDB{
		DB,
	}
}




//  all database interface methods should be implemented below



//User Functions

// CreateUser creates a user in db
func (db *MysqlDB) CreateUser(u *models.User) error{
	q := fmt.Sprintf("INSERT INTO User VALUES ('%s', '%s', '%s','%s')",u.Name,u.Pass,u.Email,u.Adder)
	_,err := db.Query(q)
	return err
}


//GetUser fetches the data of a user from db
func(db *MysqlDB) GetUser(Email string) (*models.User,error) {
	row, err := db.Query("select * from user where email = ?", Email)
	if err != nil {
		return &models.User{}, err
	}
	defer row.Close()
	var (
		name  string
		pass  string
		email string
		adder string
	)
	for row.Next() {
		err = row.Scan(&name, &pass, &email, &adder)
		if err != nil {
			return &models.User{},err
		}
	}

	if email == "" {
		err = fmt.Errorf("no such record in database")
		return &models.User{},err
	}

	return models.NewUser(name,email,pass,0,adder), nil
}


//DeleteUser deletes a user in db
func (db *MysqlDB) DeleteUser(email string, adder string) error{
	var q string
	if adder == "" {
		q = fmt.Sprintf("Delete from User where email ='%s'",email)

	} else {
		q = fmt.Sprintf("Delete from User where email ='%s' AND adder = '%s'", email, adder)
	}
	_,err := db.Query(q)
	return err
}


//
func (db *MysqlDB) UserList() ([]string,[]string,error) {
	row,err := db.Query("select name, email from user")
	var (
		name 	string
		email 	string
		names[] string
		emails[]string
	)
	if err != nil {
		return []string{},[]string{},err
	}
	for row.Next() {
		err = row.Scan(&name,&email)
		if err != nil {
			return []string{},[]string{},err
		}
		names = append(names, name)
		emails = append(emails, email)
	}
	return names,emails,nil
}


//Admin Functions

// CreateAdmin creates a admin in db
func (db *MysqlDB) CreateAdmin(u *models.User) error{
	q := fmt.Sprintf("INSERT INTO Admin VALUES ('%s', '%s', '%s', '%s')",u.Name,u.Pass,u.Email,u.Adder)
	_,err := db.Query(q)
	return err
}

//GetAdmin fetches the data of a admin from db
func(db *MysqlDB) GetAdmin(email string) (*models.User,error){
	row,err := db.Query("select * from admin where email = ?",email)
	if err != nil{
		return &models.User{},err
	}
	defer row.Close()
	var (
		name string
		mail string
		pass string
		adder string
	)
	for row.Next() {
		err := row.Scan(&name, &pass, &mail, &adder)
		if err != nil {
			return &models.User{},err
		}
	}
	println("Name: " + name + "\nemail: "  + mail)
	return models.NewUser(name,mail,pass,1,adder), nil
}

//DeleteAdmin deletes a Admin in db
func (db *MysqlDB) DeleteAdmin(email string) error{
	q := fmt.Sprintf("Delete from Admin where email ='%s'",email)
	_,err := db.Query(q)
	return err
}



//SuperAdminFunctions

// CreateSuperAdmin creates a SuperAdmin in db
func (db *MysqlDB) CreateSuperAdmin(u *models.User) error{
	q := fmt.Sprintf("INSERT INTO SuperAdmin VALUES ('%s', '%s', '%s')",u.Name,u.Pass,u.Email)
	_,err := db.Query(q)
	return err
}

//GetSuperAdmin fetches the data of a SuperAdmin from db
func(db *MysqlDB) GetSuperAdmin(Email string) (*models.User,error){
	row,err := db.Query("select * from SuperAdmin where email = ?",Email)
	if err != nil{
		return &models.User{},err
	}
	defer row.Close()
	var (
		name string
		email string
		pass string
	)
	for row.Next() {
		err := row.Scan(&name, &pass, &email)
		if err != nil {
			return &models.User{},err
		}
	}
	return models.NewUser(name,email,pass,2,""), nil
}

//DeleteSuperAdmin deletes a SuperAdmin in db
func (db *MysqlDB) DeleteSuperAdmin(email string) error{
	_,err := db.Query("Delete from SuperAdmin where email =?",email)
	return err
}



//Dish Functions

// CreateDish creates a dish in db
func (db *MysqlDB) CreateDish(d *models.Dish) error{
	q := fmt.Sprintf("INSERT INTO Dish VALUES ('%s', '%d', '%s', '%d','%d')",d.Name,d.Price,d.Menu,d.Rid,0)
	_,err := db.Query(q)
	return err
}

//GetDish fetches the data of a Dish from db
//func(db *MysqlDB) GetDish(Id int) (*models.Dish,error){
//	q := fmt.Sprintf("select * from Dish where id = '%d'",Id)
//	row,err := db.Query(q)
//	if err != nil{
//		return &models.Dish{},err
//	}
//	row.Close()
//	var (
//		name string
//		price int
//		menu string
//		rid int
//		id int
//	)
//	for row.Next() {
//		err := row.Scan(&name,&price,&menu,&rid,&id)
//		if err != nil {
//			return &models.Dish{},err
//		}
//	}
//	defer row.Close()
//	return models.NewDish(name,price,rid,menu,id), nil
//}


//UpdateDish updates the dish. flag = 0 for updating name, 1 for updating price, 2 for updating menu
func(db *MysqlDB) UpdateDish(id int, update string, flag int) error {
	var q string
	if flag == 0 {
		q = fmt.Sprintf("update dish set Name = '%s' id = '%d'",update,id)
	} else if flag == 1 {
		price,_ := strconv.Atoi(update)
		q = fmt.Sprintf("update dish set Price = '%d' where id = '%d'",price,id)
	} else if flag == 2 {
		q = fmt.Sprintf("update dish set Menu = '%s' where id = '%d'",update,id)
	} else {
		return fmt.Errorf("enter a valid flag")
	}
	_,err := db.Query(q)
	return err
}

//DeleteDish deletes a dish in a db
func(db *MysqlDB) DeleteDish(id int) error{
	 q := fmt.Sprintf("delete from dish where id = '%d'",id)
	_,err := db.Query(q)
	return err
}



//Restaurant Functions

//CreateRestaurant creates a restaurant in db
func (db *MysqlDB) CreateRestaurant(r *models.Restaurant) error{
	q := fmt.Sprintf("INSERT INTO Rest VALUES ('%s', '%s', '%s', '%s','%s','%d')",r.Name,r.Lat,r.Long,r.Owner,r.AddedBy,0)
	_,err := db.Query(q)
	return err
}

// UpdateRest updates the restaurant. flag = 0 for updating name, 1 for updating location.
// keep update2 empty in case of flag=0
func(db *MysqlDB) UpdateRest(id int, update1 string, update2 string, flag int) error {

	var q string
	if flag == 0 {
		q = fmt.Sprintf("update Rest set Name = '%s' where id = '%d'",update1,id)
	} else if flag == 1 {
		lat := update1
		long:= update2

		q = fmt.Sprintf("update Rest set latitude = '%s', longitude = '%v' where id = '%d'",lat,long,id)
	} else if flag == 2 {
		q = fmt.Sprintf("update Rest set Owner = '%s' where id = '%d'",update1,id)
	}
	_,err := db.Query(q)
	return err
}

//GetRestaurant fetches the data of a restaurant from db
func(db *MysqlDB) GetRestaurant(Id int) (*models.Restaurant,error){
	row,err := db.Query("select * from Rest where id = ?",Id)
	if err != nil{
		return &models.Restaurant{},err
	}
	defer row.Close()
	var (
		name 	string
		lat 	string
		long 	string
		owner	string
		adder	string
		id		int
	)
	for row.Next() {
		err := row.Scan(&name, &lat, &long, &owner, &adder, &id)
		if err != nil {
			return &models.Restaurant{},err
		}
	}

	return models.NewRestaurant(name,lat,long,owner,adder,id), nil
}

//DeleteRestaurant deletes a restaurant in db
func (db *MysqlDB) DeleteRestaurant(id int,adder string,rank int) error{
	var q string
	if rank == 2 {
		q = fmt.Sprintf("Delete from Rest where id = '%d'",id)
	} else if rank == 1 {
		q = fmt.Sprintf("Delete from Rest where id = '%d' AND adder = '%s'",id,adder)
	} else {
		return fmt.Errorf("rank not valid")
	}
	_,err := db.Query(q)
	return err
}


//GetMenu fetches the menu
func (db *MysqlDB) GetMenu (rid int,menu string) (map[string]int,error) {
	q := fmt.Sprintf("select name, price from Dish where rid = '%d' AND menu = '%s'",rid,menu)
	rows,err := db.Query(q)
	defer rows.Close()
	Menu := make(map[string]int)
	var (
		name string
		price int
	)
	for rows.Next() {
		err = rows.Scan(&name,&price)
		if err != nil {
			return Menu,err
		}
		Menu[name] = price
	}


	return Menu,nil
}

func (db *MysqlDB) GetbyDistance (Lat float64, Long float64,dist float64) []string{
	q := fmt.Sprintf("SELECT name,latitude,longitude FROM rest")
	rows,err := db.Query(q)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var (
		lat float64
		long float64
		name string
		names []string
	)
	for rows.Next() {
		err := rows.Scan(&name,&lat,&long)
		if err != nil {
			panic(err)
		}

		x := math.Acos(math.Sin(Lat*math.Pi/180) * math.Sin(lat*math.Pi/180) + math.Cos(Lat*math.Pi/180) * math.Cos(lat*math.Pi/180) * math.Cos((long - Long)*math.Pi/180)) * 6371
		if x<dist {
			names = append(names,name)
		}
	}
	return names

}
