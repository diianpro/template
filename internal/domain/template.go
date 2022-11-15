package domain

type Template struct {
	ID   string `bson:"_id" json:"id"`
	Data []byte `json:"data" bson:"file"`
}

type Templates struct {
	Templates []Template `json:"templates"`
}
