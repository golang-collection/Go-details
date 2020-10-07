package errcode

/**
* @Author: super
* @Date: 2020-09-22 09:49
* @Description: 统一错误代码
**/

var (
	ErrorGetStoryFail    = NewError(20020001, "获取单个问题失败")
	ErrorCreateStoryFail = NewError(20020003, "创建问题失败")
	ErrorUpdateStoryFail = NewError(20020004, "更新问题失败")
	ErrorDeleteStoryFail = NewError(20020005, "删除问题失败")
)
