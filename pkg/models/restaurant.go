package models

type Restaurant struct {
	Name	string
	Lat		string
	Long	string
	Owner	string
	AddedBy	int
	ID		int
	AdderRole int
}

func NewRestaurant(name string, lat string, long string, owner string, adder int,id int,adderrole int) *Restaurant {
	return &Restaurant{
		Name:    name,
		Lat:     lat,
		Long:    long,
		Owner:   owner,
		AddedBy: adder,
		ID:		 id,
		AdderRole: adderrole,
	}
}

func AddRestaurant(name string, lat string, long string, owner string, adder int,adderrole int) *Restaurant {

	return NewRestaurant(name,lat,long,owner,adder,0,adderrole)
}



func (r *Restaurant) GetOwner() string {
	return r.Owner
}

func (r *Restaurant) GetAdder() int {
	return r.AddedBy
}

func (r *Restaurant) GetAdderRole() int {
	return r.AdderRole
}

//func (r *Restaurant) ChangeLoc(lat string, long string) {
//	r.lat = lat
//	r.long = long
//}
//
//func (r *Restaurant) ChangeName(name string){
//	r.Name = name
//}
