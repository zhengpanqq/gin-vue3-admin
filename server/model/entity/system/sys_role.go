package system

import (
	"github.com/zhengpanone/gin-vue3-admin/global"
)

// SysRole 角色表
type SysRole struct {
	global.BaseModel
	RoleID     string     `json:"uuid" gorm:"column:role_id;not null;unique;primary_key;comment:角色ID;"` // 角色ID
	RoleName   string     `json:"roleName" gorm:"column:role_name;comment:角色名称"`                        // 角色名
	ParentId   string     `json:"parentId" gorm:"column:parent_id;comment:父角色ID"`                       // 父角色ID
	Remark     string     `json:"remark" gorm:"comment:备注"`
	Children   []SysRole  `json:"children" gorm:"-"`
	SysRoleIds []*SysRole `json:"roleIds" gorm:"many2many:role_id;"`
}

// SysUserRole 用户-角色
type SysUserRole struct {
	global.BaseModel
	UserID  string  `json:"userID" gorm:"column:user_id;comment:用户ID"`
	SysUser SysUser `gorm:"foreignkey:UserID"`
	RoleID  string  `json:"roleID" gorm:"column:role_id;comment:角色ID"`
	SysRole SysRole `gorm:"foreignkey:RoleID"`
}

// SysRolePermission 角色-权限关系表
type SysRolePermission struct {
	RoleID string `json:"roleId" gorm:"comment:角色ID"`
	MenuID string `json:"permission" gorm:"comment:"`
}
