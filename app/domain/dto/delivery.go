package dto

type Delivery struct {
	Reference  string
	Qty        int32
	ClientID   int32
	UnitID     int32
	ProductIDs []int
}
