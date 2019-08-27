package models

type User struct {
	Name      string
	Email     string
	Pass      string
	Rank      int
	Adder     int
	AdderRole int
	Id        int
}

func NewUser(name string, email string, pass string, rank int, adder int, adderRole int, id int) *User {

	return &User{
		Name:      name,
		Email:     email,
		Pass:      pass,
		Rank:      rank,
		Adder:     adder,
		AdderRole: adderRole,
		Id:        id,
	}
}
