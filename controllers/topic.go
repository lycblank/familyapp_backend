package controllers

import (
	"context"
	"fmt"
	"gofamily/models/topic"
	"gofamily/params"

	"github.com/astaxie/beego"
	"gitlab.source3g.com/game/gorunmanhttp/httpin"
)

type TopicController struct {
	beego.Controller
	tm *topic.TopicManager
}

func NewTopicController() beego.ControllerInterface {
	t := &TopicController{
		tm: &topic.TopicManager{},
	}
	return t
}

// @Title 发布帖子
// @Description 发布帖子
// @Success 200 {object} params.Resp
// @Failure 500 服务器错误
// @router / [post,options]
func (t *TopicController) Publish() {
	if t.Ctx.Input.IsOptions() {
		return
	}
	type Body struct {
		Text string `json:"text"`
	}
	type Require struct {
		Body Body `json:"body" in:"body"`
	}
	req := Require{}
	httpin.FillRequireParam(t.Ctx, &req)
	fmt.Println(req)
	type Response struct {
		ErrCode int    `json:"err_code"`
		ErrMsg  string `json:"err_msg"`
	}
	resp := Response{
		ErrCode: 0,
		ErrMsg:  "success",
	}
	httpin.JSON(t.Ctx, resp)
}

// @Title 获取帖子评论列表
// @Description 获取帖子评论列表
// @Success 200 {object} params.Resp
// @Failure 500 服务器错误
// @router /:tid/comments [get,options]
func (t *TopicController) GetTopicCommentList() {
	req := params.GetTopicCommentListReq{}
	if err := httpin.FillRequireParam(t.Ctx, &req); err != nil {
		httpin.JSON(t.Ctx, params.Response{
			ErrCode: params.USERNAME_PASSWORD_INVALIDATE_ERRCODE,
			ErrMsg:  params.USERNAME_PASSWORD_INVALIDATE_ERRMSG,
		})
		return
	}
	ctx := context.Background()
	arg := topic.GetCommentListParams{
		Tid: req.Tid,
	}
	result := &topic.GetCommentListResult{}
	t.tm.GetCommentList(ctx, arg, result)
	resp := params.GetTopicCommentListResp{
		Response: params.Response{
			ErrCode: 0,
			ErrMsg:  "success",
		},
		Data: result.Data,
	}
	httpin.JSON(t.Ctx, &resp)
}
