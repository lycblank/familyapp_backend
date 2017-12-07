package user

import (
	"context"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	familycache "gofamily/cache"
	"gofamily/params"

	"github.com/astaxie/beego/orm"
	"github.com/garyburd/redigo/redis"
	cache "gitlab.source3g.com/game/gocache"
)

type UserManager struct {
}

func NewUserManager() *UserManager {
	return &UserManager{}
}

type RegisterParams struct {
	Nickname string
	Email    string
	Password string
}

type RegisterResult struct {
	Uid string
}

func (um *UserManager) Register(ctx context.Context, arg RegisterParams, resp *RegisterResult) error {
	// 生成随机种子
	tmp := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, tmp); err != nil {
		return err
	}
	seed := strings.ToUpper(hex.EncodeToString(tmp))
	// md5密码 password+seed
	w := md5.New()
	io.WriteString(w, arg.Password)
	io.WriteString(w, seed)
	password := strings.ToUpper(hex.EncodeToString(w.Sum(nil)))
	nowTime := time.Now().Unix()
	// 插入数据库
	if res, err := orm.NewOrm().Raw("insert into family_user(nickname,email,password,seed,create_time,update_time) values(?,?,?,?,?,?)", arg.Nickname, arg.Email, password, seed, nowTime, nowTime).Exec(); err != nil {
		return err
	} else {
		if id, err := res.LastInsertId(); err == nil {
			resp.Uid = strconv.FormatInt(id, 10)
		}
	}
	// 插入redis
	conn := cache.GetRedisConn()
	defer conn.Close()
	conn.Do("HMSET", fmt.Sprintf("%s:%s%s", familycache.APP_ID, familycache.REDIS_KEY_USER_DATA_IN_ID, resp.Uid),
		familycache.REDIS_KEY_USER_ATTR_EMAIL, arg.Email,
		familycache.REDIS_KEY_USER_ATTR_NICKNAME, arg.Nickname)

	return nil
}

type LoginParams struct {
	Email    string
	Password string
}

type LoginResult struct {
	Uid      string
	Token    string
	Nickname string
}

func (um *UserManager) Login(ctx context.Context, arg LoginParams, res *LoginResult) error {
	//1. 生成数据库能识别的密码
	seed := ""
	nickname := ""
	uid := int64(0)
	password := ""
	if err := orm.NewOrm().Raw("select id, password, seed, nickname from family_user where email = ?", arg.Email).QueryRow(&uid, &password, &seed, &nickname); err != nil {
		return err
	}
	w := md5.New()
	io.WriteString(w, arg.Password)
	io.WriteString(w, seed)
	if password == strings.ToUpper(hex.EncodeToString(w.Sum(nil))) {
		res.Uid = strconv.FormatInt(uid, 10)
		res.Token = um.GenerateToken()
		res.Nickname = nickname
		go func() {
			conn := cache.GetRedisConn()
			conn.Do("SET", fmt.Sprintf("%s:%s%s", familycache.APP_ID, familycache.REDIS_KEY_USER_AUTH, res.Token), time.Now().Unix(), "EX", 2*24*60)
		}()
		return nil
	}
	return errors.New("username or password invalidate")
}

type PublishTopicParams struct {
	Uid     string
	Content string
}

type PublishTopicResult struct {
	Tid        string
	CreateTime string
}

func (um *UserManager) PublishTopic(ctx context.Context, arg PublishTopicParams, resp *PublishTopicResult) error {
	// 插入数据库
	tid := int64(0)
	createTime := time.Now().Unix()
	uid, _ := strconv.ParseUint(arg.Uid, 10, 64)
	if res, err := orm.NewOrm().Raw("insert into family_topic(uid,content,create_time) values(?,?,?)", uid, arg.Content, createTime).Exec(); err == nil {
		tid, _ = res.LastInsertId()
	} else {
		return err
	}
	conn := cache.GetRedisConn()
	defer conn.Close()
	conn.Send("HMSET", fmt.Sprintf("%s:%s%d", familycache.APP_ID, familycache.REDIS_KEY_TOPIC_DATA_IN_ID, tid),
		familycache.REDIS_KEY_TOPIC_ATTR_CONTENT, arg.Content,
		familycache.REDIS_KEY_TOPIC_ATTR_UID, arg.Uid,
		familycache.REDIS_KEY_TOPIC_ATTR_CREATE_TIME, time.Now().Unix())
	conn.Send("ZADD", fmt.Sprintf("%s:%s", familycache.APP_ID, familycache.REDIS_KEY_GLOBAL_FEED), createTime, tid)
	conn.Flush()

	resp.Tid = strconv.FormatInt(tid, 10)
	resp.CreateTime = strconv.FormatInt(createTime, 10)
	return nil
}

func (um *UserManager) GenerateToken() string {
	buf := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, buf); err != nil {
		buf = []byte(time.Now().String())
	}
	w := md5.New()
	w.Write(buf)
	return strings.ToUpper(hex.EncodeToString(w.Sum(nil)))
}

type GetUserDynamicParams struct {
	Uid string
}

type GetUserDynamicResult struct {
	Data []params.UserDynamicData
}

func (um *UserManager) GetUserDynamic(ctx context.Context, arg GetUserDynamicParams, result *GetUserDynamicResult) error {
	result.Data = []params.UserDynamicData{}
	if data := um.getTopicList(); len(data) > 0 {
		result.Data = data
	}
	return nil
}

func (um *UserManager) getTopicList() []params.UserDynamicData {
	res := []params.UserDynamicData{}
	topicListKey := fmt.Sprintf("%s:%s", familycache.APP_ID, familycache.REDIS_KEY_GLOBAL_FEED)
	conn := cache.GetRedisConn()
	defer conn.Close()
	if mems, _ := redis.Strings(conn.Do("ZREVRANGE", topicListKey, 0, -1)); len(mems) > 0 {
		uids := make([]string, 0, len(mems))
		res = make([]params.UserDynamicData, 0, len(mems))
		for _, mem := range mems {
			topicKey := fmt.Sprintf("%s:%s%s", familycache.APP_ID, familycache.REDIS_KEY_TOPIC_DATA_IN_ID, mem)
			conn.Send("HMGET", topicKey, familycache.REDIS_KEY_TOPIC_ATTR_UID,
				familycache.REDIS_KEY_TOPIC_ATTR_CONTENT,
				familycache.REDIS_KEY_TOPIC_ATTR_CREATE_TIME)
		}
		conn.Flush()
		for _, mem := range mems {
			if attrs, err := redis.Strings(conn.Receive()); err == nil && len(attrs) == 3 {
				uids = append(uids, attrs[0])
				tid, _ := strconv.ParseUint(mem, 10, 64)
				res = append(res, params.UserDynamicData{
					Tid:        tid,
					Content:    attrs[1],
					CreateTime: attrs[2],
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
				res[idx].Nickname = attrs[0]
			}
		}
	}
	return res
}

type PublishTopicCommentParams struct {
	Uid     string
	Tid     string
	Content string
}

type PublishTopicCommentResult struct {
	Cid        string
	CreateTime string
}

func (um *UserManager) PublishTopicComment(ctx context.Context, arg PublishTopicCommentParams, result *PublishTopicCommentResult) error {
	// 创建评论
	tid, _ := strconv.ParseUint(arg.Tid, 10, 64)
	uid, _ := strconv.ParseUint(arg.Uid, 10, 64)
	cid := int64(0)
	createTime := time.Now().Unix()
	if res, err := orm.NewOrm().Raw("insert into family_comment(object_type, object_id, uid, content, create_time) values(?, ?, ?, ?, ?)", 1, tid, uid, arg.Content, createTime).Exec(); err == nil {
		cid, _ = res.LastInsertId()
	} else {
		return err
	}
	conn := cache.GetRedisConn()
	defer conn.Close()
	commentDataKey := fmt.Sprintf("%s:%s%d", familycache.APP_ID, familycache.REDIS_KEY_COMMENT_DATA_IN_ID, cid)
	conn.Send("HMSET", commentDataKey, familycache.REDIS_KEY_COMMENT_ATTR_CONTENT, arg.Content,
		familycache.REDIS_KEY_COMMENT_ATTR_PARENT_ID, tid,
		familycache.REDIS_KEY_COMMENT_ATTR_PARENT_TYPE, 1,
		familycache.REDIS_KEY_COMMENT_ATTR_UID, uid,
		familycache.REDIS_KEY_COMMENT_ATTR_CREATE_TIME, createTime)
	conn.Send("ZADD", fmt.Sprintf("%s:%s%d", familycache.APP_ID, familycache.REDIS_KEY_FEED_CIDS_IN_TID, tid), createTime, cid)
	conn.Flush()
	result.Cid = strconv.FormatInt(cid, 10)
	result.CreateTime = strconv.FormatInt(createTime, 10)
	return nil
}
