package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/restaurant/pkg/models"
	"log"
	"os"
)

type MysqlDB struct {
	*sql.DB
}

func NewMySqlDB(ip string, Username string, password string, port string, schema string) *MysqlDB {

	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", Username, password, ip, port, schema)
	DB, err := sql.Open("mysql", conn)
	if err != nil {
		panic("Can't connect to db" + err.Error())
	}
	return &MysqlDB{
		DB,
	}
}

func MigrateDatabase(db *MysqlDB) error {
	driver, err := mysql.WithInstance(db.DB, &mysql.Config{})
	if err != nil {
		return err
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	migration, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s/DB/dumps", dir),
		"mysql",
		driver,
	)
	if err != nil {
		return err
	}
	err = migration.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

//  all database interface methods should be implemented below

func (db *MysqlDB) CreateCust(c *models.Customer) error{
	q := fmt.Sprintf("insert into customers values ('%s','%s','%s','0',NULL,NULL,NULL)",c.Name,c.Pass,c.Email)
	_,err := db.Exec(q)

	return err
}
//func (db *MysqlDB) UpdCust(id int, update string, flag int) error {
//	return nil
//}
//func (db *MysqlDB) DelCust(id int) error {
//	return nil
//}
func (db *MysqlDB) GetCust(id int, email string) (*models.Customer, error) {
	var q string
	if email == "" {
		q = fmt.Sprintf("select * from customers where id = '%d'",id)
	} else {
		q = fmt.Sprintf("select * from customers where email = '%s'",email)
	}
	row,err := db.Query(q)
	if err != nil {
		println("@@@@")
		return &models.Customer{},err
	}
	u := &models.Customer{}
	var (
		rid sql.NullInt64
		tid sql.NullInt64
		time sql.NullInt64
	)
	if row.Next(){
		err = row.Scan(&u.Name,&u.Pass,&u.Email,&u.Id,&rid,&tid,&time)
		if err != nil {
			return &models.Customer{},err
		}
		//err = row.Scan(&u)
		println("@@@@")
		if rid.Valid {u.BRest = rid.Int64}
		if tid.Valid {u.BRest = tid.Int64}
		if time.Valid {u.BRest = time.Int64}
	} else {
		err = fmt.Errorf("No such customer")
	}

	return u,err
}
//func (db *MysqlDB) GetRestTable(rid int)	([]int, error) {
//	return nil, nil
//}
//func (db *MysqlDB) BookTable(id int,rid int,tid int) error {
//	return nil
//}


func (db *MysqlDB) LogoutUser(token string) error {
	_, err := db.Query("INSERT INTO DeletedTokens values (?)", token)
	return err
}

func (db *MysqlDB) Checktoken(token string) (bool, error) {
	row, err := db.Query("select token from DeletedTokens where token =?", token)
	if err != nil {
		return false, err
	}
	flag := true
	if row.Next() {
		flag = false
	}
	return flag, err
}

func (db *MysqlDB) RestList() ([]int, []string, error) {

	row, err := db.Query("SELECT id,name from Rest where status = '1'")
	if err != nil {
		return []int{}, []string{}, err
	}
	var (
		name  string
		id    int
		names []string
		ids   []int
	)
	for row.Next() {
		err = row.Scan(&id, &name)
		if err != nil {
			return []int{}, []string{}, err
		}
		names = append(names, name)
		ids = append(ids, id)
	}
	return ids, names, err
}