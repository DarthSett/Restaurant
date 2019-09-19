package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"log"
	"os"
	"time"
)

type MysqlDB struct {
	*sql.DB
}

func NewMySqlDB(ip string, Username string, password string, port string, schema string) *MysqlDB {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", Username, password, ip, port, schema)
	DB := rConnectDB(conn)

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

func rConnectDB (conn string) *sql.DB {

	DB, err := sql.Open("mysql", conn)
	if err != nil {
		println(err.Error())
		println("trying to rc")
		time.Sleep(5 * time.Second)
		return rConnectDB(conn)
	}
	return DB
}