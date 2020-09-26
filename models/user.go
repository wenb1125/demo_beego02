package models

//定义人的结构体
type Person struct {
	Name string
	Age int
	Sex string
}

//定义用户的结构体
type User struct {
	Name string `json:"name"`
	Birthday string `json:"birthday"`
	Address string `json:"address"`
	Nick string `json:"nick"`
	Psaaword string `json:"psaaword"`
}

