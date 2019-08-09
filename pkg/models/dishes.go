package models

type Dish struct {
	Name		string
	Price		int
	Rid			int
	id			int
	Adder		int
	AdderRole	int
}

func  NewDish(name string, price int, rid int, id int, adder int, adderRole int) *Dish  {
	return &Dish{
		Name:       name,
		Price:      price,
		Rid: 		rid,
		id :        id,
		Adder:      adder,
		AdderRole:	adderRole,
	}
}





