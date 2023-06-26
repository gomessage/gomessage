package models

import (
	"errors"
	"gomessage/pkg/database"
	"gomessage/pkg/log/loggers"
	"gorm.io/gorm"
	"strconv"
	"time"
)

// Namespace 这里虽然名字交过namespace有点命名空间的意思，但是实际上就是用户看到的"通道"的概念
type Namespace struct {
	ID          int            `json:"id" gorm:"primarykey"`           //gorm自带字段，这里粘贴过来，显式的声明出来
	CreatedAt   time.Time      `json:"-"`                              //gorm自带字段，这里粘贴过来，显式的声明出来
	UpdatedAt   time.Time      `json:"-"`                              //gorm自带字段，这里粘贴过来，显式的声明出来
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`                 //gorm自带字段，这里粘贴过来，显式的声明出来
	Name        string         `json:"name" gorm:"unique;not null"`    //通道名称：唯一键、且不能为空
	IsActive    bool           `json:"is_active" gorm:"default:true"`  //是否激活（默认激活）
	IsRenders   bool           `json:"is_renders" gorm:"default:true"` //是否开启渲染模式（默认开启）
	Description string         `json:"description"`                    //通道描述
}

func (*Namespace) TableName() string {
	return "namespaces"
}

func AddNamespace(n *Namespace) (*Namespace, error) {
	createResult := database.DB.Default.Create(&n)
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
	result := database.DB.Default.Delete(&ns, id)
	return int(result.RowsAffected), result.Error
}

func UpdateNamespace(id int, newData *Namespace) (*Namespace, error) {
	var ns Namespace
	database.DB.Default.First(&ns, id)

	ns.Name = newData.Name
	ns.IsActive = newData.IsActive
	ns.IsRenders = newData.IsRenders
	ns.Description = newData.Description
	result := database.DB.Default.Save(&ns)

	return &ns, result.Error
}

func ListNamespace(isActive string) (*[]Namespace, error) {
	var nsList []Namespace

	//如果这里的查询参数is_active中是空字符串，则查询全部返回出去
	if len(isActive) == 0 {
		result := database.DB.Default.Find(&nsList)
		return &nsList, result.Error

	} else { //如果不是空字符串，则把字符串转换为布尔值，然后再进行查询.
		status, _ := strconv.ParseBool(isActive) //字符串转布尔值（1=True，0=False）
		result := database.DB.Default.Where("is_active", status).Find(&nsList)
		return &nsList, result.Error
	}
}

// GetNamespaceById 根据id查询namespace
func GetNamespaceById(id int) (*Namespace, error) {
	var ns Namespace
	result := database.DB.Default.Where(&Namespace{ID: id}).First(&ns)
	return &ns, result.Error

}

// GetNamespaceByName 根据name查询namespace
func GetNamespaceByName(name string) (*Namespace, error) {
	var ns Namespace
	result := database.DB.Default.Where(&Namespace{Name: name}).First(&ns)
	return &ns, result.Error
}

// IsNamespaceExist 判断namespace是否存在
func IsNamespaceExist(nsName string) bool {
	var ns Namespace
	result := database.DB.Default.Where(&Namespace{Name: nsName}).First(&ns)
	if result.RowsAffected == 1 {
		return true
	} else {
		return false
	}
}

// InitNamespace 在main函数中被调用，全局只被调用一次，创建一些默认的Namespace
func InitNamespace() {
	var queryNamespace Namespace
	result := database.DB.Default.Where(&Namespace{Name: "default"}).First(&queryNamespace)
	if result.Error != nil {
		newNamespace := Namespace{
			IsActive:    true,
			Name:        "default",
			Description: "系统自动创建的\"默认通道\"，可通过 /go/message 或 /go/default 接收消息推送。",
		}
		database.DB.Default.Create(&newNamespace)
		loggers.DefaultLogger.Info("创建default命名空间...")
	} else {
		loggers.DefaultLogger.Info("default命名空间已存在...")
	}

}
