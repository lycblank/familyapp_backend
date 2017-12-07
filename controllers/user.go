package controllers

import (
	"context"
	"fmt"
	"gofamily/models/user"
	"gofamily/params"

	"gofamily/httpin"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
	um *user.UserManager
}

func NewUserController() beego.ControllerInterface {
	u := &UserController{
		um: user.NewUserManager(),
	}
	return u
}

// @Title 获取用户动态
// @Description 获取用户动态
// @Success 200 {object} params.Resp
// @Failure 500 服务器错误
// @router /:uid/dynamic [get,options]
func (u *UserController) GetUserDynamic() {
	req := params.GetUserDynamicReq{}
	if err := httpin.FillRequireParam(u.Ctx, &req); err != nil {
		httpin.JSON(u.Ctx, params.Response{
			ErrCode: params.REQUIRE_PARAM_ERROR_ERRCODE,
			ErrMsg:  params.REQUIRE_PARAM_ERROR_ERRMSG,
		})
		return
	}
	ctx := context.Background()
	arg := user.GetUserDynamicParams{
		Uid: req.Uid,
	}
	result := &user.GetUserDynamicResult{}
	u.um.GetUserDynamic(ctx, arg, result)
	resp := params.GetUserDynamicResp{
		Response: params.Response{
			ErrCode: 0,
			ErrMsg:  "success",
		},
		Data: result.Data,
	}
	fmt.Println(resp)
	httpin.JSON(u.Ctx, &resp)
}

// @Title 获取用户的联系人
// @Description 获取用户的联系人
// @Success 200 {object} params.Resp
// @Failure 500 服务器错误
// @router /contacts [get]
func (u *UserController) GetUserContacts() {
	type ReponseData struct {
		Nickname string `json:"nickname"`
		Avatar   string `json:"avatar"`
		Location string `json:"location"`
		Header   string `json:"header"`
	}
	type Response struct {
		ErrCode int           `json:"err_code"`
		ErrMsg  string        `json:"err_msg"`
		Data    []ReponseData `json:"data"`
	}
	resp := Response{
		ErrCode: 0,
		ErrMsg:  "success",
		Data: []ReponseData{
			ReponseData{
				Nickname: "blank",
				Avatar:   "5",
				Location: "china",
				Header:   "B",
			},
		},
	}
	httpin.JSON(u.Ctx, resp)
}

// @Title 获取用户详情
// @Description 获取用户详情
// @Success 200 {object} params.Resp
// @Failure 500 服务器错误
// @router /:uid [get]
func (u *UserController) GetUserDetail() {
	type Require struct {
		Uid string `json:"uid" in:"path" required:"true" description:"用户id"`
	}
	req := Require{}
	httpin.FillRequireParam(u.Ctx, &req)
	type UserInfo struct {
		Uid       string `json:"id"`
		LoginName string `json:"loginName"`
		NickName  string `json:"nickName"`
		Points    string `json:"points"`
		AvatarUrl string `json:"avatarUrl"`
		Gender    string `json:"gender"`
		Location  string `json:"location"`
	}
	type ReponseData struct {
		Sid  string   `json:"sid"`
		User UserInfo `json:"user"`
	}
	type Response struct {
		ErrCode int         `json:"err_code"`
		ErrMsg  string      `json:"err_msg"`
		Data    ReponseData `json:"data"`
	}
	resp := Response{
		ErrCode: 0,
		ErrMsg:  "success",
		Data: ReponseData{
			Sid: "dfasdrfwefcafdsfwefdsaf",
			User: UserInfo{
				Uid:       req.Uid,
				LoginName: "123345678@qq.com",
				NickName:  "blank",
				Points:    "257",
				AvatarUrl: "https://dearb.me/assets/images/avatar.png",
				Gender:    "m",
				Location:  "Chengdu China",
			},
		},
	}
	httpin.JSON(u.Ctx, resp)
}

// @Title 注册用户
// @Description 注册用户
// @Success 200 {object} params.Resp
// @Failure 500 服务器错误
// @router /register [post,options]
func (u *UserController) RegisterUser() {
	if u.Ctx.Input.IsOptions() {
		return
	}
	req := params.RegisterUserReq{}
	if err := httpin.FillRequireParam(u.Ctx, &req); err != nil {
		fmt.Println(err)
		return
	}
	// 注册逻辑
	ctx := context.Background()
	arg := user.RegisterParams{
		Nickname: req.Body.Nickname,
		Email:    req.Body.Email,
		Password: req.Body.Password,
	}
	result := &user.RegisterResult{}
	if err := u.um.Register(ctx, arg, result); err != nil {
		fmt.Println(err)
		return
	}
	// 应答
	resp := &params.RegisterUserResp{
		Response: params.Response{
			ErrCode: 0,
			ErrMsg:  "success",
		},
		Uid: result.Uid,
	}
	httpin.JSON(u.Ctx, resp)
}

// @Title 用户登陆
// @Description 用户登陆
// @Success 200 {object} params.Resp
// @Failure 500 服务器错误
// @router /login [post,options]
func (u *UserController) LoginUser() {
	if u.Ctx.Input.IsOptions() {
		return
	}
	req := params.LoginUserReq{}
	if err := httpin.FillRequireParam(u.Ctx, &req); err != nil {
		httpin.JSON(u.Ctx, params.Response{
			ErrCode: params.REQUIRE_PARAM_ERROR_ERRCODE,
			ErrMsg:  params.REQUIRE_PARAM_ERROR_ERRMSG,
		})
		return
	}
	if req.Body.Email == "" || req.Body.Password == "" {
		httpin.JSON(u.Ctx, params.Response{
			ErrCode: params.USERNAME_PASSWORD_EMPTY_ERRCODE,
			ErrMsg:  params.USERNAME_PASSWORD_EMPTY_ERRMSG,
		})
		return
	}
	// 登陆逻辑
	ctx := context.Background()
	arg := user.LoginParams{
		Email:    req.Body.Email,
		Password: req.Body.Password,
	}
	result := &user.LoginResult{}
	if err := u.um.Login(ctx, arg, result); err != nil {
		httpin.JSON(u.Ctx, params.Response{
			ErrCode: params.USERNAME_PASSWORD_INVALIDATE_ERRCODE,
			ErrMsg:  params.USERNAME_PASSWORD_INVALIDATE_ERRMSG,
		})
		return
	}
	// 应答
	resp := &params.LoginUserResp{
		Response: params.Response{
			ErrCode: 0,
			ErrMsg:  "success",
		},
		Uid:      result.Uid,
		Token:    result.Token,
		Nickname: result.Nickname,
	}
	httpin.JSON(u.Ctx, resp)
}

// @Title 用户发布帖子
// @Description 用户发布帖子
// @Success 200 {object} params.Resp
// @Failure 500 服务器错误
// @router /:uid/publish/topic [post,options]
func (u *UserController) PublishTopic() {
	req := params.UserPublishTopicReq{}
	if err := httpin.FillRequireParam(u.Ctx, &req); err != nil {
		httpin.JSON(u.Ctx, params.Response{
			ErrCode: params.REQUIRE_PARAM_ERROR_ERRCODE,
			ErrMsg:  params.REQUIRE_PARAM_ERROR_ERRMSG,
		})
		return
	}
	if req.Body.Content == "" || req.Uid == "" {
		httpin.JSON(u.Ctx, params.Response{
			ErrCode: params.TOPIC_CONTENT_EMPTY_ERRCODE,
			ErrMsg:  params.TOPIC_CONTENT_EMPTY_ERRMSG,
		})
		return
	}
	// 发布逻辑
	ctx := context.Background()
	arg := user.PublishTopicParams{
		Content: req.Body.Content,
		Uid:     req.Uid,
	}
	result := &user.PublishTopicResult{}
	if err := u.um.PublishTopic(ctx, arg, result); err != nil {
		httpin.JSON(u.Ctx, params.Response{
			ErrCode: params.PUBLISH_TOPIC_ERROR_ERRCODE,
			ErrMsg:  params.PUBLISH_TOPIC_ERROR_ERRMSG,
		})
		return
	}
	resp := params.UserPublishTopicResp{
		Response: params.Response{
			ErrCode: 0,
			ErrMsg:  "success",
		},
		Tid:        result.Tid,
		CreateTime: result.CreateTime,
	}
	httpin.JSON(u.Ctx, &resp)
}

// @Title 用户发布评论
// @Description 用户发布评论
// @Success 200 {object} params.Resp
// @Failure 500 服务器错误
// @router /:uid/topic/:tid/comment [post,options]
func (u *UserController) PublishTopicComment() {
	req := params.PublishTopicCommentReq{}
	if err := httpin.FillRequireParam(u.Ctx, &req); err != nil {
		httpin.JSON(u.Ctx, params.Response{
			ErrCode: params.REQUIRE_PARAM_ERROR_ERRCODE,
			ErrMsg:  params.REQUIRE_PARAM_ERROR_ERRMSG,
		})
		return
	}
	// 发布评论逻辑
	ctx := context.Background()
	arg := user.PublishTopicCommentParams{
		Content: req.Body.Content,
		Uid:     req.Uid,
		Tid:     req.Tid,
	}
	result := &user.PublishTopicCommentResult{}
	if err := u.um.PublishTopicComment(ctx, arg, result); err != nil {
		httpin.JSON(u.Ctx, params.Response{
			ErrCode: params.PUBLISH_TOPIC_COMMENT_ERROR_ERRCODE,
			ErrMsg:  params.PUBLISH_TOPIC_COMMENT_ERROR_ERRMSG,
		})
		return
	}
	resp := params.PublishTopicCommentResp{
		Response: params.Response{
			ErrCode: 0,
			ErrMsg:  "success",
		},
		Cid:        result.Cid,
		CreateTime: result.CreateTime,
	}
	httpin.JSON(u.Ctx, &resp)
}
