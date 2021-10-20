package main
//Address 地址结构体
type Address struct {
Province   string
City       string
CreateTime string
}

//Email 邮箱结构体
type Email struct {
Account    string
CreateTime string
}

//User 用户结构体
type User struct {
Name   string
Gender string
Address
Email
}

func main() {
var user3 User
user3.Name = "沙河娜扎"
user3.Gender = "男"
// user3.CreateTime = "2019" //ambiguous selector user3.CreateTime
user3.Address.CreateTime = "2000" //指定Address结构体中的CreateTime
user3.Email.CreateTime = "2000"   //指定Email结构体中的CreateTime
}