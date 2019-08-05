package models

type Restaurant struct {
	Name	string
	Lat		string
	Long	string
	Owner	string
	AddedBy	string
	ID		int
}

func NewRestaurant(name string, lat string, long string, owner string, adder string,id int) *Restaurant {
	return &Restaurant{
		Name:    name,
		Lat:     lat,
		Long:    long,
		Owner:   owner,
		AddedBy: adder,
		ID:		 id,
	}
}

func AddRestaurant(name string, lat string, long string, owner string, adder string) *Restaurant {

	return NewRestaurant(name,lat,long,owner,adder,0)
}



func (r *Restaurant) GetOwner() string {
	return r.Owner
}

func (r *Restaurant) GetAdder() string {
	return r.AddedBy
}

//func (r *Restaurant) ChangeLoc(lat string, long string) {
//	r.lat = lat
//	r.long = long
//}
//
//func (r *Restaurant) ChangeName(name string){
//	r.Name = name
//}
