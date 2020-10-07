package dao

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

