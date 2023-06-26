package authorization

import (
	"gomessage/pkg/database"
	"gorm.io/gorm"
	"time"
)

type Users struct {
	ID           int            `json:"id" gorm:"primarykey"`
	CreatedAt    time.Time      `json:"-"`
	UpdatedAt    time.Time      `json:"-"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
	Username     string         `json:"username" gorm:"unique;not null;comment:账号"` //账号
	Nickname     string         `json:"nickname" gorm:"昵称"`                         //昵称
	PasswordHash string         `json:"-" gorm:"comment:密码哈希"`                      //这里对哈希加密后的密码存库，但是不对json开放
	Active       bool           `json:"-" gorm:"default:true"`                      //账号是否激活的（默认是激活的）
	Email        string         `json:"email" gorm:"邮箱"`                            //邮箱
	Phone        string         `json:"phone" gorm:"电话"`                            //电话号码
	Description  string         `json:"description" gorm:"描述"`                      //描述
}

func (user *Users) TableName() string {
	return "users"
}

func CreateUser(user *Users) (*Users, error) {
	result := database.DB.Default.Create(&user)
	return user, result.Error
}

func DeleteUser(id int) (int, error) {
	var users Users
	result := database.DB.Default.Delete(&users, id)
	return int(result.RowsAffected), result.Error
}

func UpdateUser(id int, newData *Users) (*Users, error) {
	var u Users
	database.DB.Default.First(&u, id)
	u.Username = newData.Username
	result := database.DB.Default.Save(&u)
	return &u, result.Error
}

func ListUser() (*[]Users, error) {
	var userList *[]Users
	result := database.DB.Default.Find(&userList)
	return userList, result.Error
}

func GetUserById(id int) (*Users, error) {
	var user *Users
	result := database.DB.Default.Where(&Users{ID: id}).First(&user)
	return user, result.Error
}

func QueryUserByUsername(username string) (*Users, error) {
	var user *Users
	result := database.DB.Default.Where("username = ?", username).First(&user)
	return user, result.Error
}

type Sessions struct {
	ID        int            `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	Username  string         `json:"-" gorm:"comment:账号"`   //账号
	Token     string         `json:"-" gorm:"comment:令牌凭证"` //令牌
}

func (user *Sessions) TableName() string {
	return "sessions"
}

func CreateSession(username, token string) (*Sessions, error) {
	ss := Sessions{
		Username: username,
		Token:    token,
	}
	result := database.DB.Default.Create(&ss)
	return &ss, result.Error
}

func DeleteSession(token string) {
	var ss *Sessions
	database.DB.Default.Where("token = ?", token).First(&ss)
	database.DB.Default.Delete(&Sessions{}, ss.ID)
}

func QueryToken(token string) []Sessions {
	var list []Sessions
	database.DB.Default.Where("token = ?", token).Find(&list)
	return list
}
