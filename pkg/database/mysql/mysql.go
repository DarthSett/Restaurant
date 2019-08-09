package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/restaurant/pkg/models"
	"math"
	"strconv"
	"time"
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
	q := fmt.Sprintf("INSERT INTO User VALUES ('%s', '%s', '%s','%d','%d','%d','%d','%s')",u.Name,u.Pass,u.Email,u.Adder,u.AdderRole,0,1,time.Now().Format("2006-01-02 15:04:05"))
	_,err := db.Query(q)
	return err
}


//GetUser fetches the data of a user from db
func(db *MysqlDB) GetUser(Email string,id int) (*models.User,error) {
	var row *sql.Rows
	var err error
	if Email == "" {
		row, err = db.Query("select * from user where id = ? and status = '1'", id)
		if err != nil {
			return &models.User{}, err
	}
	} else {
		row, err = db.Query("select * from user where email = ? and status = '1'", Email)
		if err != nil {
			return &models.User{}, err
		}
	}
	defer row.Close()
	var (
		name      string
		pass      string
		email     string
		adder     int
		adderRole int
		Id        int
		status    int
		createdAt string
	)
	for row.Next() {
		err = row.Scan(&name, &pass, &email, &adder,&adderRole,&Id,&status,&createdAt)
		if err != nil {
			return &models.User{},err
		}
	}

	if email == "" {
		err = fmt.Errorf("no such record in database")
		return &models.User{},err
	}

	return models.NewUser(name,email,pass,0,adder,adderRole,Id), nil
}

func (db *MysqlDB)GetUserRests (email string) ([]int,[]string,error) {
	row,err := db.Query("select id,name from rest where owner = ?",email)
	if err != nil {
		return []int{},[]string{},err
	}
	defer row.Close()
	var (
		id int
		name string
		ids []int
		names []string
	)
	for row.Next() {
		err = row.Scan(&id,&name)
		if err != nil {
			return []int{},[]string{},err
		}

		ids = append(ids,id)
		names = append(names,name)
	}
	return ids,names,nil
}


//DeleteUser deletes a user in db
func (db *MysqlDB) DeleteUser(id int) error{
	q := fmt.Sprintf("Delete from User where id ='%d'",id)
	_,err := db.Query(q)
	return err
}
//UpdateUser updates a user row in db. flag =0 for name, flag = 1 for password, flag = 2 for email
func (db *MysqlDB) UpdateUser(id int,update string,flag int) error {
	var q string
	if flag == 0 {
		q = fmt.Sprintf("update User set Name = '%s' where id = '%d'",update,id)
	} else if flag == 1 {
		q = fmt.Sprintf("update User set PassHash = '%s' where id = '%d'",update,id)

	} else if flag == 2 {
		q = fmt.Sprintf("update User set email = '%s' where id = '%d'",update,id)
	}
	_,err := db.Query(q)
	return err

}

//UserList fetches the list of all users
func (db *MysqlDB) UserList() ([]string,[]string,[]int,error) {
	row,err := db.Query("select name, email,id from user where status = '1'")
	var (
		name 	string
		email 	string
		id		int
		names   []string
		emails  []string
		ids		[]int
	)
	defer row.Close()
	if err != nil {
		return []string{},[]string{},[]int{},err
	}
	for row.Next() {
		err = row.Scan(&name,&email,&id)
		if err != nil {
			return []string{},[]string{},[]int{},err
		}
		names = append(names, name)
		emails = append(emails, email)
		ids = append(ids,id)
	}
	return names,emails,ids,nil
}


//Admin Functions

// CreateAdmin creates a admin in db
func (db *MysqlDB) CreateAdmin(u *models.User) error{
	q := fmt.Sprintf("INSERT INTO Admin VALUES ('%s', '%s', '%s','%d','%d','%d','%d','%v')",u.Name,u.Pass,u.Email,u.Adder,u.AdderRole,0,1,time.Now().Format("2006-01-02 15:04:05"))
	_,err := db.Query(q)
	return err
}

//GetAdmin fetches the data of a admin from db
func(db *MysqlDB) GetAdmin(Email string,id int) (*models.User,error){
	var row *sql.Rows
	var err error
	if Email == "" {
		row, err = db.Query("select * from admin where id = ? and status = '1'", id)
		if err != nil {
			return &models.User{}, err
		}
	} else {
		row, err = db.Query("select * from admin where email = ? and status = '1'", Email)
		if err != nil {
			return &models.User{}, err
		}
	}
	defer row.Close()
	var (
		name      string
		pass      string
		email     string
		adder     int
		adderRole int
		Id        int
		status    int
		createdAt string
	)
	for row.Next() {
		err = row.Scan(&name, &pass, &email, &adder,&adderRole,&Id,&status,&createdAt)
		if err != nil {
			return &models.User{},err
		}
	}

	if email == "" {
		err = fmt.Errorf("no such record in database")
		return &models.User{},err
	}

	return models.NewUser(name,email,pass,1,adder,adderRole,Id), nil

}

//DeleteAdmin deletes a Admin in db
func (db *MysqlDB) DeleteAdmin(id int) error{
	q := fmt.Sprintf("Delete from admin where id ='%d'",id)
	_,err := db.Query(q)
	return err
}
//UpdateAdmin updates an admin row in db. flag =0 for name, flag = 1 for password, flag = 2 for email
func (db *MysqlDB) UpdateAdmin(id int,update string,flag int) error {
	var q string
	if flag == 0 {
		q = fmt.Sprintf("update Admin set Name = '%s' where id = '%d'",update,id)
	} else if flag == 1 {
		q = fmt.Sprintf("update Admin set Pass = '%s' where id = '%d'",update,id)

	} else if flag == 2 {
		q = fmt.Sprintf("update Admin set email = '%s' where id = '%d'",update,id)
	}
	_,err := db.Query(q)
	return err

}



//SuperAdminFunctions

// CreateSuperAdmin creates a SuperAdmin in db
func (db *MysqlDB) CreateSuperAdmin(u *models.User) error{
	q := fmt.Sprintf("INSERT INTO superadmin VALUES ('%s', '%s', '%s','%d','%d','%d','%d','%v')",u.Name,u.Pass,u.Email,u.Adder,u.AdderRole,0,1,time.Now().Format("2006-01-02 15:04:05"))
	_,err := db.Query(q)
	return err
}

//GetSuperAdmin fetches the data of a SuperAdmin from db
func(db *MysqlDB) GetSuperAdmin(Email string,id int) (*models.User,error){
	var row *sql.Rows
	var err error
	if Email == "" {
		println("@@@")
		row, err = db.Query("select * from superadmin where id = ? and status = '1'", id)
		if err != nil {
			return &models.User{}, err
		}
	} else {
		println("!!!")
		row, err = db.Query("select * from superadmin where email = ? and status = '1'", Email)
		if err != nil {
			return &models.User{}, err
		}
	}
	defer row.Close()
	var (
		name      string
		pass      string
		email     string
		adder     int
		adderRole int
		Id        int
		status    int
		createdAt string
	)
	for row.Next() {
		err = row.Scan(&name, &pass, &email, &adder,&adderRole,&Id,&status,&createdAt)
		if err != nil {
			return &models.User{},err
		}
	}

	if email == "" {
		err = fmt.Errorf("no such record in database")
		return &models.User{},err
	}

	return models.NewUser(name,email,pass,2,adder,adderRole,Id), nil

}


//DeleteSuperAdmin deletes a SuperAdmin in db
func (db *MysqlDB) DeleteSuperAdmin(id int) error{
	q := fmt.Sprintf("delete from superadmin where id ='%d'",id)
	_,err:=db.Query(q)
	return err
}



//Dish Functions

// CreateDish creates a dish in db
func (db *MysqlDB) CreateDish(d *models.Dish) error{
	q := fmt.Sprintf("INSERT INTO Dish VALUES ('%s', '%d', '%d','%d','%d','%d','%d','%s')",d.Name,d.Price,d.Rid,d.Adder,d.AdderRole,0,1,time.Now().Format("2006-01-02 15:04:05"))
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


//UpdateDish updates the dish. flag = 0 for updating name, 1 for updating price
func(db *MysqlDB) UpdateDish(id int, update string, flag int) error {
	var q string
	if flag == 0 {
		q = fmt.Sprintf("update dish set Name = '%s' where id = '%d'",update,id)
	} else if flag == 1 {
		price,_ := strconv.Atoi(update)
		q = fmt.Sprintf("update dish set Price = '%d' where id = '%d'",price,id)
	} else {
		return fmt.Errorf("enter a valid flag")
	}
	_,err := db.Query(q)
	return err
}

//DeleteDish deletes a dish in a db
func(db *MysqlDB) DeleteDish(id int) error{
	 q := fmt.Sprintf("update dish set status = '0' where id = '%d'",id)
	_,err := db.Query(q)
	return err
}



//Restaurant Functions

//CreateRestaurant creates a restaurant in db
func (db *MysqlDB) CreateRestaurant(r *models.Restaurant) error{
	q := fmt.Sprintf("INSERT INTO Rest VALUES ('%s', '%s', '%s', '%s','%d','%d','%d','1','%s')",r.Name,r.Lat,r.Long,r.Owner,r.AddedBy,r.AdderRole,0,time.Now().Format("2006-01-02 15:04:05"))
	_,err := db.Query(q)
	return err
}

// UpdateRest updates the restaurant. flag = 0 for updating name, 1 for updating location, 2 for updating owner.
// keep update2 empty in case of flag=0/2
func(db *MysqlDB) UpdateRest(id int, update1 string, update2 string, flag int) error {

	var q string
	if flag == 0 {
		q = fmt.Sprintf("update Rest set Name = '%s' where id = '%d'",update1,id)
	} else if flag == 1 {
		lat := update1
		long:= update2

		q = fmt.Sprintf("update Rest set latitude = '%s', longitude = '%s' where id = '%d'",lat,long,id)
	} else if flag == 2 {
		q = fmt.Sprintf("update Rest set Owner = '%s' where id = '%d'",update1,id)
	}
	_,err := db.Query(q)
	return err
}

//GetRestaurant fetches the data of a restaurant from db
func(db *MysqlDB) GetRestaurant(Id int) (*models.Restaurant,error){
	row,err := db.Query("select * from Rest where id = ? and status = '1'",Id)
	if err != nil{
		return &models.Restaurant{},err
	}
	defer row.Close()
	var (
		name	 	string
		lat 		string
		long 		string
		owner		string
		adder		int
		adderRole	int
		id			int
		status		int
		created		string
	)
	for row.Next() {
		err := row.Scan(&name, &lat, &long, &owner, &adder, &adderRole, &id,&status,&created)
		if err != nil {
			return &models.Restaurant{},err
		}
	}
	if name == "" {
		return &models.Restaurant{},fmt.Errorf("no such record in database")
	}

	return models.NewRestaurant(name,lat,long,owner,adder,id,adderRole), nil
}

//DeleteRestaurant deletes a restaurant in db
func (db *MysqlDB) DeleteRestaurant(id int,adder int,rank int) error{
	var q string
	if rank == 2 {
		q = fmt.Sprintf("Update Rest set status = '0' where id = '%d'",id)
	} else if rank == 1 {
		q = fmt.Sprintf("Update Rest set status = '0' where id = '%d' AND adder = '%d' AND AdderRole = '1'",id,adder)
	}
	row,err := db.Exec(q)
	if err != nil {
		return err
	}
	affected,err := row.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return fmt.Errorf("error deleting restaurant")
	}
	return err
}


//GetMenu fetches the menu
func (db *MysqlDB) GetMenu (rid int) ([]int,[]string,[]int,error) {
	q := fmt.Sprintf("select name, price,id from Dish where rid = '%d' and status = '1'",rid)
	rows,err := db.Query(q)
	if err != nil {
		return []int{},[]string{},[]int{},err
	}
	defer rows.Close()
	var (
		name string
		price int
		id	int
		names	[]string
		ids		[]int
		prices	[]int
	)
	for rows.Next() {
		err = rows.Scan(&name,&price,&id)
		if err != nil {
			return []int{},[]string{},[]int{},err
		}
		names = append(names,name)
		ids =	append(ids,id)
		prices=	append(prices,price)
	}


	return ids,names,prices,nil
}


//getbyDistance fetches the list of all restaurants in a certain radius from a point
func (db *MysqlDB) GetbyDistance (Lat float64, Long float64,dist float64) ([]string,[]int){
	q := fmt.Sprintf("SELECT name,id,latitude,longitude FROM rest where status = '1'")
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
		id int
		ids []int
	)
	for rows.Next() {
		err := rows.Scan(&name,&id,&lat,&long)
		if err != nil {
			panic(err)
		}

		x := math.Acos(math.Sin(Lat*math.Pi/180) * math.Sin(lat*math.Pi/180) + math.Cos(Lat*math.Pi/180) * math.Cos(lat*math.Pi/180) * math.Cos((long - Long)*math.Pi/180)) * 6371
		if x<dist {
			names = append(names,name)
			ids	= append(ids,id)
		}
	}
	return names,ids

}


//RestList fetches the List of all restaurants
func (db *MysqlDB) RestList() ([]int,[]string,error){

	row,err := db.Query("SELECT id,name from rest where status = '1'")
	if err != nil {
		return []int{},[]string{},err
	}
	var (
		name string
		id int
		names []string
		ids []int
	)
	for row.Next(){
		err = row.Scan(&id,&name)
		if err != nil {
			return []int{},[]string{},err
		}
		names = append(names, name)
		ids = append(ids,id)
	}
	return ids,names,err
}

func (db *MysqlDB) LogoutUser(token string) error {
	_,err:=db.Query("INSERT INTO DeletedTokens values (?)",token)
	return err
}

func (db *MysqlDB) Checktoken(token string) (bool,error) {
	row,err := db.Query("select token from DeletedTokens where token = ?",token)
	if err != nil {
		return false,err
	}
	flag := true
	if row.Next() {
		flag = false
	}
	return flag,err
}
