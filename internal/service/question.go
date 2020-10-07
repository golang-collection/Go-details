package service

import (
	"Go-details/internal/dao"
	"Go-details/pkg/util"
)

/**
* @Author: super
* @Date: 2020-10-07 14:32
* @Description:
**/

type QuestionRequest struct {
	ID    uint32 `form:"id" binding:"required,gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateQuestionRequest struct {
	Question  string `form:"question" binding:"required,min=2,max=4294967295"`
	Source    string `form:"source" binding:"required,min=2,max=4294967295"`
	CreatedBy string `form:"created_by" binding:"required,min=2,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateQuestionRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Question   string `form:"question " binding:"min=2,max=4294967295"`
	ModifiedBy string `form:"modified_by" binding:"required,min=2,max=100"`
	State      uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type DeleteQuestionRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

type Question struct {
	ID       uint32 `json:"id"`
	Question string `json:"question"`
	Source   string `json:"source"`
	State    uint8  `json:"state"`
}

func (svc *Service) GetQuestion(param *QuestionRequest) (*Question, error) {
	question, err := svc.dao.GetQuestion(param.ID, param.State)
	if err != nil {
		return nil, err
	}

	//将base64编码后的字符串返回给前端，提升传输效率
	content, err := util.EncodeBase64(question.Question)
	if err != nil {
		return nil, err
	}
	return &Question{
		ID:       question.ID,
		Question: content,
		Source:   question.Source,
		State:    question.State,
	}, nil
}

func (svc *Service) CreateQuestion(param *CreateQuestionRequest) error {
	//TODO 做内容校验，相同的内容就不再插入

	//前端传递过来是base64编码后的字符串
	content := util.DecodeBase64(param.Question)
	_, err := svc.dao.CreateQuestion(&dao.Question{
		Question:  content,
		Source:    param.Source,
		State:     param.State,
		CreatedBy: param.CreatedBy,
	})
	if err != nil {
		return err
	}

	return nil
}

func (svc *Service) UpdateStory(param *UpdateQuestionRequest) error {
	//TODO 做内容校验，相同的内容就不再插入
	//前端传递过来是base64编码后的字符串
	content := util.DecodeBase64(param.Question)
	err := svc.dao.UpdateQuestion(&dao.Question{
		ID:         param.ID,
		Question:   content,
		State:      param.State,
		ModifiedBy: param.ModifiedBy,
	})
	if err != nil {
		return err
	}

	return nil
}

func (svc *Service) DeleteStory(param *DeleteQuestionRequest) error {
	err := svc.dao.DeleteQuestion(param.ID)
	if err != nil {
		return err
	}

	return nil
}
