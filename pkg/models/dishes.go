package models

type Dish struct {
	Name		string
	Price		int
	Rid			int
	Menu 		string
	id			int
}

func  NewDish(name string, price int, rid int, menu string, id int) *Dish  {
	return &Dish{
		Name:       name,
		Price:      price,
		Rid: 		rid,
		Menu:		menu,
		id :        id,
	}
}





