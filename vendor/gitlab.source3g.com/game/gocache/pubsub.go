/*
Package cache

Copyright(c) 2000-2016 Quqian Technology Company Limit.
All rights reserved.
创建人：  mark
创建日期：2016-09-18 14:31:34
修改记录：
----------------------------------------------------------------
修改人        |  修改日期    				|    备注
----------------------------------------------------------------
longyongchun  | 2016-09-18 14:31:34		|    创建文件
----------------------------------------------------------------
*/
package cache

func (cache *Cache) Publish(channel, value interface{}) {
	conn := cache.RedisPool().Get()
	defer conn.Close()

	conn.Do("PUBLISH", channel, value)
}
