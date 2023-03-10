package models

import "time"

// RoleTable 角色
type RoleTable struct {
	ID        int64     `gorm:"primarykey" json:"role_id"`
	RoleName  string    `gorm:"column:role_name;type:varchar(120);comment:角色名称" json:"role_name"`
	CreatedAt time.Time `gorm:"created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"updated_at" json:"updatedAt"`
}

func (t *RoleTable) TableName() string {
	return "t_role"
}
