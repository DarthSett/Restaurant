package models

type User struct {
	Name string
	Email string
	Pass string
	Rank int
	Adder string
}

func NewUser (name string, email string, pass string, rank int, adder string) *User {

	return &User{
		Name: name,
		Email: email,
		Pass:  pass,
		Rank:  rank,
		Adder: adder,
	}
}



