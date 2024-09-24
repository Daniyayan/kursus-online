package Response

type users struct {
	name        string  `json:"name"`
	description string  `json:"description"`
	duration    string  `json:"duration"`
	price       float64 `json:"price"`
}
