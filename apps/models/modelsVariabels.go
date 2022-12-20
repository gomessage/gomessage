package models

import (
	"gomessage/utils/database"
	"gorm.io/gorm"
	"time"
)

type Variables struct {
	ID          int            `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	Namespace   string         `json:"namespace"`
	Key         string         `json:"key"`
	Value       string         `json:"value"`
	Description string         `json:"description"`
}

func (*Variables) TableName() string {
	return "variables"
}

func AddVariables(v *Variables) (*Variables, error) {
	createResult := database.DB.DefaultClient.Create(&v)
	return v, createResult.Error
}

func DeleteVariables(id int) (int, error) {
	var vv Variables
	result := database.DB.DefaultClient.Delete(&vv, id)
	return int(result.RowsAffected), result.Error
}

func UpdateVariables(id int, v *Variables) (*Variables, error) {
	vv := Variables{}
	readResult := database.DB.DefaultClient.First(&vv, id)
	if readResult.Error != nil {
		return &vv, readResult.Error

	} else {
		updateResult := database.DB.DefaultClient.Model(&vv).Omit("id").Updates(&v)
		return &vv, updateResult.Error
	}
}

func GetVariablesById(id int) (Variables, error) {
	var vv Variables
	queryResult := database.DB.DefaultClient.Where(&Variables{ID: id}).First(&vv)
	return vv, queryResult.Error
}

func ListVariables(ns string) (*[]Variables, error) {
	var vv []Variables
	queryResult := database.DB.DefaultClient.Where(Variables{Namespace: ns}).Order("id").Find(&vv)
	return &vv, queryResult.Error
}
