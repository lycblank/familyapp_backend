package cache

const APP_ID = "family"
const (
	REDIS_KEY_USER_AUTH          = "user.in:auth:"
	REDIS_KEY_USER_DATA_IN_ID    = "user.in:data_in_id:"
	REDIS_KEY_USER_ATTR_EMAIL    = "user.in:data_attr_email"
	REDIS_KEY_USER_ATTR_NICKNAME = "user.in:data_attr_nickname"

	REDIS_KEY_TOPIC_DATA_IN_ID       = "topic.in:data_in_id:"
	REDIS_KEY_TOPIC_ATTR_CONTENT     = "topic.in:data_attr_content"
	REDIS_KEY_TOPIC_ATTR_UID         = "topic.in:data_attr_uid"
	REDIS_KEY_TOPIC_ATTR_CREATE_TIME = "topic.in:data_attr_create_time"

	REDIS_KEY_COMMENT_DATA_IN_ID       = "comment.in:data_in_id:"
	REDIS_KEY_COMMENT_ATTR_UID         = "comment.in:data_attr_uid"
	REDIS_KEY_COMMENT_ATTR_CONTENT     = "comment.in:data_attr_content"
	REDIS_KEY_COMMENT_ATTR_PARENT_TYPE = "comment.in:data_attr_parent_type"
	REDIS_KEY_COMMENT_ATTR_PARENT_ID   = "comment.in:data_attr_parent_id"
	REDIS_KEY_COMMENT_ATTR_CREATE_TIME = "comment.in:data_attr_create_time"

	// 全局feed流
	REDIS_KEY_GLOBAL_FEED      = "feed.in:global_in_feed"
	REDIS_KEY_FEED_CIDS_IN_TID = "feed.in:cids_in_tid:"
)
