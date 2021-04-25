package model

type Paging struct {
	Page     int
	PrevPage int
	NextPage int
	HasPrev  bool
	HasNext  bool
}
