/**
 * @Author: yinjie_luo
 * @Description:
 * @File:  project_mgm
 * @Date: 2023/07/25 11:57
 */

package http_server

import (
	"demo/internal/rest"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AddProject  @Summary 项目管理
// @Description 项目管理-新建项目
// @Param addProjectRequest query rest.AddProjectRequest true "新建项目参数"
// @Success 200 {object} rest.Result1{data=rest.AddProjectResp}
// @Router /edapi/bankbills/project/add [get]
func AddProject(context *gin.Context) {

	context.JSON(http.StatusOK, rest.Result{
		Data: rest.AddProjectResp{},
	})
}

// ProjectList  @Summary 项目管理
// @Description 项目管理-项目列表
// @Param projectListRequest query rest.ProjectListRequest true "获取项目列表参数"
// @Success 200 {object} rest.Result1{data=rest.ProjectListResp}
// @Router /edapi/bankbills/project/query [get]
func ProjectList(context *gin.Context) {

	context.JSON(http.StatusOK, rest.Result{
		Data: rest.ProjectListResp{},
	})
}

// UpdateProject  @Summary 项目管理
// @Description 项目管理-更新项目
// @Param updateProjectRequest query rest.UpdateProjectRequest true "获取项目列表参数"
// @Success 200 {object} rest.Result1{}
// @Router /edapi/bankbills/project/update [get]
func UpdateProject(context *gin.Context) {

	context.JSON(http.StatusOK, rest.Result{})
}