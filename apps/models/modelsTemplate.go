package models

import (
    "gomessage/utils/database"
    "gorm.io/gorm"
    "time"
)

type Template struct {
    ID              int            `json:"id" gorm:"primarykey"`
    CreatedAt       time.Time      `json:"-"`
    UpdatedAt       time.Time      `json:"-"`
    DeletedAt       gorm.DeletedAt `json:"-" gorm:"index"`
    Namespace       string         `json:"namespace"`
    TemplateName    string         `json:"template_name"`
    TemplateContent string         `json:"template_content"`
    TemplateIsMerge bool           `json:"template_is_merge"`
}

func (*Template) TableName() string {
    return "templates"
}

func AddTemplate(t *Template) (*Template, error) {
    createResult := database.DB.DefaultClient.Create(&t)
    return t, createResult.Error
}

func DeleteTemplate(id int) (int, error) {
    var template Template
    result := database.DB.DefaultClient.Delete(&template, id)
    return int(result.RowsAffected), result.Error
}

func DeleteTemplateByNs(ns string) (int, error) {
    var template Template
    result := database.DB.DefaultClient.Where("namespace = ?", ns).Delete(&template)
    return int(result.RowsAffected), result.Error
}

func UpdateTemplate(id int, t *Template) (*Template, error) {
    template := Template{}
    readResult := database.DB.DefaultClient.First(&template, id)

    //如果Error不为空
    if readResult.Error != nil {
        return &template, readResult.Error

    } else {
        updateResult := database.DB.DefaultClient.Model(&template).Omit("id").Updates(&t)
        return &template, updateResult.Error
    }
}

func ListTemplate(ns string) (*[]Template, error) {
    var templates []Template
    queryResult := database.DB.DefaultClient.Where(Template{Namespace: ns}).Find(&templates)
    return &templates, queryResult.Error
}

func GetTemplateById(id int) (Template, error) {
    var template Template
    queryResult := database.DB.DefaultClient.Where(&Template{ID: id}).First(&template)
    return template, queryResult.Error
}

func GetTemplateByNamespace(ns string) (Template, error) {
    var template Template
    queryResult := database.DB.DefaultClient.Where(&Template{Namespace: ns}).First(&template)
    return template, queryResult.Error
}
