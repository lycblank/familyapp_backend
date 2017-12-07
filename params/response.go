package params

type Response struct {
	ErrCode int    `json:"errCode" description:"401:没有授权"`
	ErrMsg  string `json:"errMsg,omitempty"`
}

const (
	NOT_AUTH_ERRCODE                     = 401000
	USERNAME_PASSWORD_EMPTY_ERRCODE      = 400000
	REQUIRE_PARAM_ERROR_ERRCODE          = 400001
	USERNAME_PASSWORD_INVALIDATE_ERRCODE = 400002
	TOPIC_CONTENT_EMPTY_ERRCODE          = 400003
	PUBLISH_TOPIC_ERROR_ERRCODE          = 500000
	PUBLISH_TOPIC_COMMENT_ERROR_ERRCODE  = 500001
)

const (
	NOT_AUTH_ERRMSG                     = "请先登陆后继续!!!"
	USERNAME_PASSWORD_EMPTY_ERRMSG      = "用户名与密码不能为空!!!"
	REQUIRE_PARAM_ERROR_ERRMSG          = "请求参数错误!!!"
	USERNAME_PASSWORD_INVALIDATE_ERRMSG = "用户名或密码错误!!!"
	TOPIC_CONTENT_EMPTY_ERRMSG          = "帖子内容不能为空!!!"
	PUBLISH_TOPIC_ERROR_ERRMSG          = "发布帖子出错，请稍后重试!!!"
	PUBLISH_TOPIC_COMMENT_ERROR_ERRMSG  = "发布帖子评论出错，请稍后重试!!!"
)
