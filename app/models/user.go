package models

type User struct {
	Name string
	Idcard string
	Email  string
	StuId  string
}

func (u *User)CheckUser() bool{
	var user []User
	DB.Table("student").Debug().Where("name=? and idcard=? and email=? and stu_id=?",u.Name,u.Idcard,u.Email,u.StuId).Find(&user)
	if len(user)==1{
		return true
	}
	return false
}

func GetUser(stu string) User {
	u:=User{}
    DB.Table("student").Where("stu_id=?",stu).First(&u)
    return u
}
