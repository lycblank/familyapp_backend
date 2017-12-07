package params

type GetTopicCommentListReq struct {
	Tid string `json:"tid" description:"帖子id" in:"path"`
}

type TopicCommentData struct {
	Avatar     string `json:"avatar"`
	Nickname   string `json:"name"`
	Content    string `json:"text"`
	CreateTime string `json:"time"`
}

type GetTopicCommentListResp struct {
	Response
	Data []TopicCommentData `json:"data"`
}
