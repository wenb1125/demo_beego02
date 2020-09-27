package models

//定义人的结构体
type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Sex string `json:"sex"`
}

//定义用户的结构体
type User struct {
	Name string `json:"name"`
	Birthday string `json:"birthday"`
	Email string `json:"email"`
	Address string `json:"address"`
	Nick string `json:"nick"`
	Psaaword string `json:"psaaword"`
}

