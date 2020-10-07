package dao

import "Go-details/internal/model"

/**
* @Author: super
* @Date: 2020-10-07 14:01
* @Description:
**/

type Question struct {
	ID         uint32 `json:"id"`
	Source     string `json:"source"`
	Question   string `json:"question"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      uint8  `json:"state"`
}

func (d *Dao) CreateQuestion(param *Question) (*model.Question, error) {
	question := model.Question{
		Source:   param.Source,
		Question: param.Question,
		State:    param.State,
		Model:    &model.Model{CreatedBy: param.CreatedBy},
	}
	return question.Create(d.engine)
}

func (d *Dao) UpdateQuestion(param *Question) error {
	question := model.Question{Model: &model.Model{ID: param.ID}}
	values := map[string]interface{}{
		"modified_by": param.ModifiedBy,
		"state":       param.State,
	}

	if param.Question != "" {
		values["question"] = param.Question
	}

	return question.Update(d.engine, values)
}

func (d *Dao) GetQuestion(id uint32, state uint8) (model.Question, error) {
	question := model.Question{Model: &model.Model{ID: id}, State: state}
	return question.Get(d.engine)
}

func (d *Dao) DeleteQuestion(id uint32) error {
	question := model.Question{Model: &model.Model{ID: id}}
	return question.Delete(d.engine)
}
