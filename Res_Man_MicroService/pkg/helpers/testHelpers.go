package helpers

import (
	"fmt"
	"github.com/restaurant/Res_Man_MicroService/pkg/database/mysql"
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

func ResetDB(db *mysql.MysqlDB) {

	println("@@@@@@@@@@@@@@")

	println("@@@@@@@@@@@@@@")
	_,err := db.Exec("delete from superadmin")
	if err != nil {println(err)}
	_,err=db.Exec("INSERT INTO superadmin VALUES ('Sourav', '$2a$04$CAIRyD6NVptXbB25PEUFJeYEsBwIouUYLigBhWocbcmvOZrc7.OV.', 'sourav241196@gmail.com', '0', '2', '1', '1', '2019-08-08 14:14:10')")
	if err != nil {println(err)}
	_,err =db.Exec("ALTER table superadmin AUTO_INCREMENT=2")
	if err != nil {println(err)}

	_,err=db.Exec("delete from admin")
	if err != nil {println(err)}
	_,err=db.Exec("INSERT INTO Admin VALUES ('admin1', '$2a$05$TmfhiQ73jCaIhSBAmXBBb.Ram40COKQvHIgA.3znT0B.ZuhFc33Oe', 'admin1@gmail.com', '1', '2', '1', '1', '2019-08-08 15:47:35')")
	if err != nil {println(err)}

	_,err=db.Exec("ALTER table admin AUTO_INCREMENT=2")
	if err != nil {println(err)}

	_,err=db.Exec("delete from user")
	if err != nil {println(err)}
	_,err=db.Exec("INSERT INTO User VALUES ('user1', '$2a$04$y9ExoM60HYRfNxm8N/tE3.YHVS/RhHB/6eaztdwVYhoRPspofsmk2', 'user1@gmail.com', '1', '2', '1', '1', '2019-08-09 14:12:25')")
	if err != nil {println(err)}
	_,err=db.Exec("INSERT INTO User VALUES ('user2', '$2a$04$y9ExoM60HYRfNxm8N/tE3.YHVS/RhHB/6eaztdwVYhoRPspofsmk2', 'user2@gmail.com', '1', '1', '2', '1', '2019-08-09 14:12:25')")
	if err != nil {println(err)}

	_,err=db.Exec("ALTER table user AUTO_INCREMENT=3")
	if err != nil {println(err)}

	_,err=db.Exec("delete from rest where id <> '1'")
	if err != nil {println(err)}
	_,err=db.Exec("update rest set status = '1' where id = '1'")
	if err != nil {println(err)}

	_,err=db.Exec("ALTER table rest AUTO_INCREMENT=2")
	if err != nil {println(err)}

	_,err=db.Exec("delete from dish where id <> '1'")
	if err != nil {println(err)}

	_,err=db.Exec("update dish set status = '1' where id = '1'")
	if err != nil {println(err)}

	_,err=db.Exec("ALTER table dish AUTO_INCREMENT=2")
	if err != nil {println(err)}

}