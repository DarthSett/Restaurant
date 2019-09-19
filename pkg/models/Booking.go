package models

type Booking struct {
	ID  	int 		`json:"cid" binding:"required"`
	CID 	int 		`json:"rid" binding:"required"`
	RID 	int 		`json:"tid" binding:"required"`
	TID   	int 		`json:"id"`
	BTime	int64
}
