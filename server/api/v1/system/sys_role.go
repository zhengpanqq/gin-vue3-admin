package system

import (
	"github.com/gin-gonic/gin"
	"github.com/zhengpanone/gin-vue3-admin/global"
	"github.com/zhengpanone/gin-vue3-admin/model/common/request"
	response "github.com/zhengpanone/gin-vue3-admin/model/common/response"
	sysReq "github.com/zhengpanone/gin-vue3-admin/model/system/request"
	"go.uber.org/zap"
)

type RoleApi struct{}

// CreateRole
// @Tags RoleApi
// @Summary 创建角色
// @Description 创建角色
// @ID /v1/api/role/addRole
// @Accept  json
// @Produce  application/json
// @Param data body request.RoleParam true "body" #[roleName,密码,手机号码] body [string,string,string] [required,required,required] "[system.SysUser]"
// @Router /v1/api/role/addRole [post]
func (r *RoleApi) CreateRole(c *gin.Context) {
	var roleParam sysReq.RoleParam
	_ = c.ShouldBindJSON(&roleParam)
	role, err := roleService.CreateRole(roleParam)
	if err != nil {
		response.ErrorWithMsg(c, err.Error())
		return
	}
	response.OkWithData(c, role)
}

// PageRole
// @Tags RoleApi
// @Summary 分页获取角色
// @Description 分页获取角色
// @ID /v1/api/role/pageRole
// @Accept  json
// @Produce  application/json
// @Param data body request.PageInfo true "body" #[roleName,密码,手机号码] body [string,string,string] [required,required,required] "[system.SysUser]"
// @Router /v1/api/role/pageRole [post]
func (r *RoleApi) PageRole(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.ErrorWithMsg(c, err.Error())
		return
	}
	list, total, err := roleService.PageRole(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败！", zap.Error(err))
		response.ErrorWithMsg(c, err.Error())
		return
	}
	response.OkWithDataAndMsg(c, response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "")
}

// DeleteRole
// @Tags RoleApi
// @Summary 删除角色
// @Description 删除角色
// @ID /v1/api/role/delete/:id
// @Accept  json
// @Produce  application/json
// @Param data body request.PageInfo true "body" #[roleName,密码,手机号码] body [string,string,string] [required,required,required] "[system.SysUser]"
// @Router /v1/api/role/pageRole [post]
func (r *RoleApi) DeleteRole(c *gin.Context) {
	id := c.Param("id")
	response.OkWithData(c, id)
}

// UpdateRole
// @Tags RoleApi
// @Summary 更新角色
// @Description 更新角色
// @ID /v1/api/role/update/:id
// @Accept  json
// @Produce  application/json
// @Param data body request.PageInfo true "body" #[roleName,密码,手机号码] body [string,string,string] [required,required,required] "[system.SysUser]"
// @Router /v1/api/role/pageRole [post]
func (r *RoleApi) UpdateRole(c *gin.Context) {
	id := c.Param("id")
	response.OkWithData(c, id)
}
