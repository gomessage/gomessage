package models

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"gomessage/pkg/utils/database"
	"strings"
	"time"

	"gorm.io/gorm"
)

// CrontabNamespace 是为了处理[]string的自定义类型
type CrontabNamespace []string

// Scan 实现方法，定义 `[]string` 如何从数据库读取
func (c *CrontabNamespace) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("对string类型断言失败")
	}
	*c = strings.Split(str, "|||")
	return nil
}

// Value 实现方法，定义 `[]string` 如何写入数据库
func (c CrontabNamespace) Value() (driver.Value, error) {
	return strings.Join(c, "|||"), nil
}

type Crontab struct {
	ID               int              `json:"id" gorm:"primarykey"`
	CreatedAt        time.Time        `json:"-"`
	UpdatedAt        time.Time        `json:"-"`
	DeletedAt        gorm.DeletedAt   `json:"-" gorm:"index"`
	CrontabName      string           `json:"crontab_name"`
	CrontabRule      string           `json:"crontab_rule"`
	CrontabNamespace CrontabNamespace `json:"crontab_namespace"`
	CrontabActivate  bool             `json:"crontab_activate" gorm:"default:false"`
	CrontabContent   string           `json:"crontab_content"`
}

func (*Crontab) TableName() string {
	return "crontab"
}

func AddCrontab(c *Crontab) error {
	if c.CrontabName == "" {
		return errors.New("crontab name 不能为空")
	}
	if c.CrontabRule == "" {
		return errors.New("crontab rule 不能为空")
	}
	if len(c.CrontabNamespace) == 0 {
		return errors.New("crontab namespace 不能为空")
	}
	result := database.DB.Default.Create(c)
	if result.Error != nil {
		// 添加详细的日志打印，以便了解错误原因
		fmt.Printf("添加Crontab失败: %v. 错误: %v", c, result.Error)
		return result.Error
	}
	return nil
}

func DeleteCrontab(id int) error {
	result := database.DB.Default.Delete(&Crontab{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("没有找到对应的Crontab进行删除")
	}
	return nil
}

func UpdateCrontab(id int, c *Crontab) error {
	if c.CrontabName == "" {
		return errors.New("crontab name 不能为空")
	}
	if c.CrontabRule == "" {
		return errors.New("crontab rule 不能为空")
	}
	if len(c.CrontabNamespace) == 0 {
		return errors.New("crontab namespace 不能为空")
	}
	// 使用Select确保所有字段都被考虑更新，包括零值字段
	result := database.DB.Default.Model(&Crontab{}).Where("id = ?", id).Select("*").Updates(c)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("没有找到对应的Crontab进行更新")
	}
	return nil
}

func GetCrontab(id int) (*Crontab, error) {
	var c Crontab
	result := database.DB.Default.First(&c, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &c, nil
}

func ListCrontabs() ([]Crontab, error) {
	var crontabs []Crontab
	result := database.DB.Default.Find(&crontabs)
	if result.Error != nil {
		return nil, result.Error
	}
	return crontabs, nil
}
