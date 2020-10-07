package v1

import (
	"Go-details/global"
	"Go-details/internal/service"
	"Go-details/pkg/app"
	"Go-details/pkg/convert"
	"Go-details/pkg/errcode"
	"github.com/gin-gonic/gin"
)

/**
* @Author: super
* @Date: 2020-10-07 15:03
* @Description:
**/

type Question struct {

}

func NewQuestion() Question{
	return Question{}
}

// @Summary 获取问题
// @Produce json
// @Param id path int true "问题ID"
// @Success 200 {object} model.Question "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/question/{id} [get]
func (q Question)Get(c *gin.Context){
	param := service.QuestionRequest{ID:convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	question, err := svc.GetQuestion(&param)

	if err != nil {
		global.Logger.Errorf(c, "svc.GetQuestion err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetStoryFail)
		return
	}

	response.ToResponse(question)
	return
}

// @Summary 创建问题
// @Produce json
// @Param content body string true "问题内容"
// @Param author body string true "作者"
// @Param created_by body int true "创建者"
// @Param state body int false "状态"
// @Success 200 {object} model.Question "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/question [post]
func (q Question) Create(c *gin.Context) {
	param := service.CreateQuestionRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreateQuestion(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.CreateQuestion err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateStoryFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// @Summary 更新问题
// @Produce json
// @Param content body string false "问题内容"
// @Param modified_by body string true "修改者"
// @Success 200 {object} model.Question "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/question/{id} [put]
func (q Question) Update(c *gin.Context) {
	param := service.UpdateQuestionRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateQuestion(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.UpdateQuestion err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateStoryFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// @Summary 删除问题
// @Produce  json
// @Param id path int true "问题ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/question/{id} [delete]
func (q Question) Delete(c *gin.Context) {
	param := service.DeleteQuestionRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteQuestion(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.DeleteQuestion err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteStoryFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}