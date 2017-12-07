package topic

import (
	"context"
	"fmt"
	"gofamily/params"

	familycache "gofamily/cache"

	"github.com/garyburd/redigo/redis"
	cache "gitlab.source3g.com/game/gocache"
)

type TopicManager struct {
}

type GetCommentListParams struct {
	Tid string
}

type GetCommentListResult struct {
	Data []params.TopicCommentData
}

func (tm *TopicManager) GetCommentList(ctx context.Context, arg GetCommentListParams, result *GetCommentListResult) error {
	conn := cache.GetRedisConn()
	defer conn.Close()
	result.Data = []params.TopicCommentData{}
	if mems, err := redis.Strings(conn.Do("ZREVRANGE", fmt.Sprintf("%s:%s%s", familycache.APP_ID, familycache.REDIS_KEY_FEED_CIDS_IN_TID, arg.Tid), 0, -1)); err == nil && len(mems) > 0 {
		fmt.Println(mems, err)
		result.Data = make([]params.TopicCommentData, 0, len(mems))
		uids := make([]string, 0, len(mems))
		for _, mem := range mems {
			commentKey := fmt.Sprintf("%s:%s%s", familycache.APP_ID, familycache.REDIS_KEY_COMMENT_DATA_IN_ID, mem)
			conn.Send("HMGET", commentKey,
				familycache.REDIS_KEY_COMMENT_ATTR_UID,
				familycache.REDIS_KEY_COMMENT_ATTR_CONTENT,
				familycache.REDIS_KEY_COMMENT_ATTR_CREATE_TIME)
		}
		conn.Flush()
		for range mems {
			if attrs, err := redis.Strings(conn.Receive()); err == nil && len(attrs) == 3 {
				uids = append(uids, attrs[0])
				result.Data = append(result.Data, params.TopicCommentData{
					Content:    attrs[1],
					CreateTime: attrs[2],
					Avatar:     "1",
				})
			}
		}
		for _, uid := range uids {
			userKey := fmt.Sprintf("%s:%s%s", familycache.APP_ID, familycache.REDIS_KEY_USER_DATA_IN_ID, uid)
			conn.Send("HMGET", userKey, familycache.REDIS_KEY_USER_ATTR_NICKNAME)
		}
		conn.Flush()
		for idx := range uids {
			if attrs, err := redis.Strings(conn.Receive()); err == nil && len(attrs) == 1 {
				result.Data[idx].Nickname = attrs[0]
			}
		}
	} else {
		fmt.Println(mems, err)
	}
	return nil
}
