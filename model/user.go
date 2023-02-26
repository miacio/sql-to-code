package model

import (
	"time"
)

// User 用户表
type User struct {
	ID         uint64    `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id" toml:"id"`             // ID id
	UserID     int       `gorm:"column:user_id;not null" json:"userID" toml:"userID"`                         // UserID user_id 用户 id
	Email      string    `gorm:"column:email;not null;default:''" json:"email" toml:"email"`                  // Email email 用户邮箱
	Phone      string    `gorm:"column:phone" json:"phone" toml:"phone"`                                      // Phone phone 手机号
	Role       int8      `gorm:"column:role;not null" json:"role" toml:"role"`                                // Role role 用户角色  1:超级管理员 2:其他
	Sex        IBool     `gorm:"column:sex;default:b'0'" json:"sex" toml:"sex"`                               // Sex sex 用户性别 0 男 1 女
	WebsiteURL string    `gorm:"column:website_url" json:"websiteURL" toml:"websiteURL"`                      // WebsiteURL website_url 个人主页
	Remark     string    `gorm:"column:remark" json:"remark" toml:"remark"`                                   // Remark remark 备注
	UserSeat   IPoint    `gorm:"column:user_seat" json:"userSeat" toml:"userSeat"`                            // UserSeat user_seat 用户位置
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime;not null" json:"createdAt" toml:"createdAt"` // CreatedAt created_at 创建时间
	UpdatedAt  time.Time `gorm:"column:updated_at;autoUpdateTime;not null" json:"updatedAt" toml:"updatedAt"` // UpdatedAt updated_at 更新时间
	DeletedAt  time.Time `gorm:"column:deleted_at" json:"deletedAt" toml:"deletedAt"`                         // DeletedAt deleted_at 移除时间
}

// TableName User user
func (User) TableName() string {
	return "user"
}
