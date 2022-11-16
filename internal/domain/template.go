package domain

type Template struct {
	ID   string `json:"id" bson:"id"`
	Data []byte `json:"data" bson:"data"`
}

type Templates struct {
	Templates []Template `json:"templates"`
}
