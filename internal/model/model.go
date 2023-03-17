package model

type Person struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type Data struct {
	People []Person `json:"people"`
}
