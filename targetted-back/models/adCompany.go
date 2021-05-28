package models

type AdCompany struct {
	Name            string
	Purpose         string
	FieldOfActivity string
	BusinessAdress  string
	Creatives       []Creative
}
