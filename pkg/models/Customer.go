package models

type Customer struct {
	Name  	string 		`json:"name" binding:"required"`
	Email 	string 		`json:"email" binding:"required"`
	Pass  	string 		`json:"pass" binding:"required"`
	Id    	int 		`json:"id"`
	BRest 	int64
	Btable	int64
	BTime	int64
}

type Credentials struct {
	Email string	`json:"email" binding:"required"`
	Pass  string 	`json:"pass" binding:"required"`
}

func NewCust(name string, email string, pass string, id int, rid int64, tid int64, btime int64) *Customer {
	return &Customer{
		Name:   	name,
		Email:  	email,
		Pass:   	pass,
		BRest:  	rid,
		Btable: 	tid,
		BTime:		btime,
		Id:     	id,
	}
}
