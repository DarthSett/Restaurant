package helpers

import (
	"fmt"
	"github.com/restaurant/pkg/database/mysql"
	"testing"
)

func AssertEqual(t *testing.T, expected interface{}, got interface{}, message string) {
	if expected == got {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", expected, got)
	}
	t.Fatal(message)
}

func ResetDBafterCreate(db *mysql.MysqlDB) {
	println("@@@@@@@@@@@@@@")
	db.Query("delete from superadmin where id <> '1'")
	db.Query("ALTER table superadmin AUTO_INCREMENT=2")
	db.Query("delete from admin where id <> '1'")
	db.Query("ALTER table admin AUTO_INCREMENT=2")
	db.Query("delete from user where id > '2'")
	db.Query("ALTER table user AUTO_INCREMENT=3")
	db.Query("delete from rest where id <> '1'")
	db.Query("ALTER table rest AUTO_INCREMENT=19")
	db.Query("delete from dish where id <> '1'")
	db.Query("ALTER table dish AUTO_INCREMENT=2")

}

func ResetDBafterDelete(db *mysql.MysqlDB,flag int) {
	if flag == 0{
		db.Query("ALTER table User AUTO_INCREMENT=1")
		db.Query("INSERT INTO User VALUES ('user1', '$2a$04$y9ExoM60HYRfNxm8N/tE3.YHVS/RhHB/6eaztdwVYhoRPspofsmk2', 'user1@gmail.com', '1', '1', '1', '1', '2019-08-09 14:12:25')")
		db.Query("ALTER table admin AUTO_INCREMENT=1")
		db.Query("INSERT INTO Admin VALUES ('admin1', '$2a$04$AKe7E84SCtQZCjPYpRCN6OrdMS/VnV0Qx93Os.TU8nO71Tt67ysmG', 'admin1@gmail.com', '1', '2', '1', '1', '2019-08-08 15:47:35')")
		db.Query("ALTER table superadmin AUTO_INCREMENT=1")
		db.Query("INSERT INTO superadmin VALUES ('Sourav', '$2a$04$CAIRyD6NVptXbB25PEUFJeYEsBwIouUYLigBhWocbcmvOZrc7.OV.', 'sourav241196@gmail.com', '0', '2', '1', '1', '2019-08-08 14:14:10')")
	} else if flag == 1 {
		db.Query("ALTER table rest AUTO_INCREMEnt=2")
		db.Query("update rest set status = '1' where id = '1'")
	} else if flag == 2 {
		db.Query("ALTER table dish AUTO_INCREMENT=2")
		db.Query("update dish set status = '1' where id = '1'")
	}

}
