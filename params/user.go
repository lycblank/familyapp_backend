package params

type RegisterUserReq struct {
	Body RegisterUserBody `json:"body" in:"body"`
}

type RegisterUserBody struct {
	Nickname string `json:"nickname" description:"昵称"`
	Email    string `json:"email" description:"邮箱 即登陆时的用户名" required:"true"`
	Password string `json:"password" description:"密码 即登陆时的密码" required:"true"`
}

type RegisterUserResp struct {
	Response
	Uid string `json:"uid,omitempty"`
}

// ========== 登陆用户请求 ================
type LoginUserReq struct {
	Body LoginUserBody `json:"body" in:"body"`
}

type LoginUserBody struct {
	Email    string `json:"username" description:"邮箱 即登陆时的用户名" required:"true"`
	Password string `json:"password" description:"密码 即登陆时的密码" required:"true"`
}

type LoginUserResp struct {
	Response
	Uid      string `json:"uid,omitempty"`
	Token    string `json:"token,omitempty"`
	Nickname string `json:"nickname"`
}

// 用户发布帖子
type UserPublishTopicReq struct {
	Uid  string               `json:"uid" in:"path" required:"true"`
	Body UserPublishTopicBody `json:"body" in:"body"`
}

type UserPublishTopicBody struct {
	Content string `json:"content" description:"帖子内容" required:"true"`
}

type UserPublishTopicResp struct {
	Response
	Tid        string `json:"tid,omitempty"`
	CreateTime string `json:"createTime,omitempty"`
}

// 获取用户动态
type GetUserDynamicReq struct {
	Uid string `json:"uid" description:"用户id" in:"path"`
}

type UserDynamicData struct {
	Tid          uint64 `json:"id"`
	Nickname     string `json:"nickname"`
	Avatar       string `json:"avatar"`
	Content      string `json:"text"`
	OriginalPic  string `json:"original_pic"`
	CommentCount int    `json:"comment_count"`
	LikeCount    int    `json:"like_count"`
	CreateTime   string `json:"created_at"`
}

type GetUserDynamicResp struct {
	Response
	Data []UserDynamicData `json:"data"`
}

type PublishTopicCommentReq struct {
	Uid  string                  `json:"uid" description:"用户id" in:"path" required:"true"`
	Tid  string                  `json:"tid" description:"topic id" in:"path" required:"true"`
	Body PublishTopicCommentBody `json:"body" in:"body"`
}

type PublishTopicCommentBody struct {
	Content string `json:"content" description:"评论的内容"`
}

type PublishTopicCommentResp struct {
	Response
	Cid        string `json:"cid" description:"评论id"`
	CreateTime string `json:"createTime" description:"创建时间"`
}
