package models

import (
	"errors"
	"gomessage/utils/database"
	"gomessage/utils/runLog"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type Namespace struct {
	ID          int            `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	IsActive    bool           `json:"is_active" gorm:"default:false"`
	Name        string         `json:"name" gorm:"unique;not null"` //字段约束：唯一键、且不能为空
	Description string         `json:"description"`
}

func (*Namespace) TableName() string {
	return "namespace"
}

func AddNamespace(n *Namespace) (*Namespace, error) {
	createResult := database.DB.DefaultClient.Create(&n)
	if createResult.Error != nil {
		return nil, createResult.Error
	} else {
		if createResult.RowsAffected != 1 {
			return nil, errors.New("创建数据失败：记录影响行数不为1")
		}
	}
	return n, createResult.Error
}

func DeleteNamespace(id int) (int, error) {
	var ns Namespace
	result := database.DB.DefaultClient.Delete(&ns, id)
	return int(result.RowsAffected), result.Error
}

func UpdateNamespace(id int, newData *Namespace) (*Namespace, error) {
	var ns Namespace
	database.DB.DefaultClient.First(&ns, id)
	updateResult := database.DB.DefaultClient.Model(&ns).Omit("id").Updates(&newData)
	return &ns, updateResult.Error
}

func ListNamespace(isActive string) (*[]Namespace, error) {
	var nsList []Namespace

	//如果这里的查询参数is_active中是空字符串，则查询全部返回出去
	if len(isActive) == 0 {
		result := database.DB.DefaultClient.Find(&nsList)
		return &nsList, result.Error

	} else { //如果不是空字符串，则把字符串转换为布尔值，然后再进行查询.
		status, _ := strconv.ParseBool(isActive) //字符串转布尔值（1=True，0=False）
		result := database.DB.DefaultClient.Where("is_active", status).Find(&nsList)
		return &nsList, result.Error
	}
}

func GetNamespaceById(id int) (*Namespace, error) {
	var ns Namespace
	result := database.DB.DefaultClient.Where(&Namespace{ID: id}).First(&ns)
	return &ns, result.Error
}

func GetNamespaceByName(name string) (*Namespace, error) {
	var ns Namespace
	result := database.DB.DefaultClient.Where(&Namespace{Name: name}).First(&ns)
	return &ns, result.Error
}

func IsNamespaceExist(nsName string) bool {
	var ns Namespace
	result := database.DB.DefaultClient.Where(&Namespace{Name: nsName}).First(&ns)
	if result.RowsAffected == 1 {
		return true
	} else {
		return false
	}
}

// InitNamespace 在main函数中被调用，全局只被调用一次，创建一些默认的Namespace
func InitNamespace() {
	var queryNamespace Namespace
	result := database.DB.DefaultClient.Where(&Namespace{Name: "default"}).First(&queryNamespace)
	if result.Error != nil {
		newNamespace := Namespace{IsActive: true, Name: "default", Description: "This is the default namespace created by GoMessage."}
		database.DB.DefaultClient.Create(&newNamespace)
		runLog.Log.Info("创建default命名空间...")
	} else {
		runLog.Log.Info("default命名空间已存在...")
	}

}
