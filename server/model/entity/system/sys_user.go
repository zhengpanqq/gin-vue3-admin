package system

import (
	uuid "github.com/satori/go.uuid"
	"github.com/zhengpanone/gin-vue3-admin/global"
)

// SysUser 用户表
type SysUser struct {
	global.BaseModel
	UUID        uuid.UUID   `json:"uuid" gorm:"comment:用户UUID"` // 用户UUID
	Username    string      `json:"username" gorm:"column:username;comment:用户登录名"`
	NickName    string      `json:"nickName" gorm:"column:nickname;type:varchar(20);not null;default:'';comment:用户昵称"`
	Phone       string      `json:"phone" gorm:"type:char(11);unique:un_phone;comment:手机号" validate:"required"`
	Password    string      `json:"password" gorm:"column:password;type:varchar(255);comment:密码" validate:"required"`
	Status      string      `json:"status" gorm:"size:4;default:1;comment:状态 1:正常 2:白名单 3:黑名单"`
	UserInfo    SysUserInfo `json:"userInfo" gorm:"-"`
	Token       string      `json:"token" gorm:"-"`
	Sex         string      `validate:"oneof=female male" gorm:"column:sex;comment:性别"`
	Authorities []SysRole   `json:"roles" gorm:"many2many:sys_role;"`
}

// SysUserInfo 用户信息表
type SysUserInfo struct {
	global.BaseModel
	UserID   string `json:"userID"`
	Birthday string `json:"birthday" gorm:"type:varchar(10);comment:生日"`
	Address  string `json:"address" gorm:"type:text;comment:地址"`
}
