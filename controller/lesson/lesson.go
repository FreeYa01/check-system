package lesson

import (
	"github.com/gin-gonic/gin"
	"qiandao/pkg/app"
	"qiandao/service"
	"qiandao/viewmodel"
)

// CreateLesson 创建课程
func CreateLesson(ctx *gin.Context)  {
//	 1.绑定参数
	lesson := new(viewmodel.Lesson)
	err := ctx.ShouldBindJSON(lesson)
	if err != nil {
		app.SendResponse(ctx,app.ErrBind,nil)
		return
	}
//	 2.判断参数是否合理
	if lesson.LessonName == "" ||  len(lesson.ClassList) == 0 {
		app.SendResponse(ctx,app.ErrParamNull,nil)
		return
	}
//  3.调用逻辑处理
	err = service.CreateLesson(lesson)
	if err != nil {
		app.SendResponse(ctx,err,nil)
		return
	}

//	 3.返回响应
	app.SendResponse(ctx,app.OK,nil)
}

// GetCreateLessonList 获取创建的课程列表
func GetCreateLessonList(ctx *gin.Context)  {
//	1.绑定参数
	userId := ctx.Query("user_id")
	if userId == "" {
		app.SendResponse(ctx,app.ErrBind,nil)
		return
	}
//	2.调用业务逻辑
	dataList,err := service.GetCreateLessonList(userId)
	if err != nil {
		app.SendResponse(ctx,err,nil)
		return
	}
//	3.返回响应
	app.SendResponse(ctx,app.OK,dataList)
}

// GetJoinLessonList 获取加入的课程列表
func GetJoinLessonList(ctx *gin.Context){
//	1.绑定参数
	classId := ctx.Query("class_id")
	if classId == ""{
		app.SendResponse(ctx,app.ErrBind,nil)
		return
	}
//	2.调用业务逻辑
	joinList,err:= service.GetJoinLessonList(classId)
	if err != nil{
		app.SendResponse(ctx,err,nil)
		return
	}
//	3.返回响应
	app.SendResponse(ctx,app.OK,joinList)

}
// EditorLesson 编辑课程信息
func EditorLesson(ctx *gin.Context){
// 	1.绑定参数
	var lesson *viewmodel.LessonEditor
	err := ctx.ShouldBindJSON(&lesson)
	if err != nil {
		app.SendResponse(ctx,app.ErrBind,nil)
		return
	}
	if lesson.LessonID == "" || lesson.LessonName == ""{
		app.SendResponse(ctx,app.ErrParamNull,nil)
		return
	}
//	2.业务处理
	err = service.EditorLesson(lesson)
	if err != nil {
		app.SendResponse(ctx,err,nil)
		return
	}
//	3.返回响应
	app.SendResponse(ctx,app.OK,nil)
}
// RemoveLesson 移除课程
func RemoveLesson(ctx *gin.Context)  {
// 1.绑定参数
	lesson := new(viewmodel.LessonRemove)
	err := ctx.ShouldBind(&lesson)
	if err != nil {
		app.SendResponse(ctx,app.ErrBind,nil)
		return
	}
	if lesson.LessonID == "" || lesson.LessonCreator == ""{
		app.SendResponse(ctx,app.ErrParamNull,nil)
		return
	}
// 2.调用业务逻辑
	err = service.RemoveLesson(lesson)
	if err != nil {
		app.SendResponse(ctx,err,nil)
		return
	}
// 3.返回响应
	app.SendResponse(ctx,app.OK,nil)
}


