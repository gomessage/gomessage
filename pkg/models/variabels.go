package models

import (
	"gomessage/pkg/utils/database"
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
	createResult := database.DB.Default.Create(&v)
	return v, createResult.Error
}

func DeleteVariables(id int) (int, error) {
	var vv Variables
	result := database.DB.Default.Delete(&vv, id)
	return int(result.RowsAffected), result.Error
}

func DeleteVariablesByNs(namespace string) (int, error) {
	var vv Variables
	result := database.DB.Default.Where("namespace = ?", namespace).Delete(&vv)
	return int(result.RowsAffected), result.Error
}

func UpdateVariables(id int, v *Variables) (*Variables, error) {
	vv := Variables{}
	readResult := database.DB.Default.First(&vv, id)
	if readResult.Error != nil {
		return &vv, readResult.Error

	} else {
		updateResult := database.DB.Default.Model(&vv).Omit("id").Updates(&v)
		return &vv, updateResult.Error
	}
}

func GetVariablesById(id int) (Variables, error) {
	var vv Variables
	queryResult := database.DB.Default.Where(&Variables{ID: id}).First(&vv)
	return vv, queryResult.Error
}

func ListVariables(ns string) (*[]Variables, error) {
	var vv []Variables
	queryResult := database.DB.Default.Where(Variables{Namespace: ns}).Order("id").Find(&vv)
	return &vv, queryResult.Error
}
