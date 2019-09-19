package models

type Booking struct {
	CID  	int 		`json:"cid" binding:"required"`
	RID 	int 		`json:"rid" binding:"required"`
	TID 	int 		`json:"tid" binding:"required"`
	Id    	int 		`json:"id"`
	BTime	int64
}
