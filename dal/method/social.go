package method

import "gorm.io/gen"

type FollowMethod interface {
	//sql(insert into @@table (uid, follow_uid) values (@uid, @touid))
	InsertFollow(uid, touid int64) error
	//sql(delete from @@table where uid=@uid and follow_uid=@touid limit 1)
	DeleteFollow(uid, touid int64) error
	//sql(select uid, follow_uid, ctime from @@table where uid=@uid)
	FindUserFollow(uid int64) ([]*gen.T, error)
}

type FollowerMethod interface {
	//sql(insert into @@table (uid, follower_uid) values (@uid, @touid))
	InsertFollower(uid, touid int64) error
	//sql(delete from @@table where uid=@uid and follower_uid=@touid limit 1)
	DeleteFollower(uid, touid int64) error
	////sql(select uid, follower_uid, ctime from @@table where uid=@uid {{if lasstid != 0}} and id < (select id from @@table where uid=@uid and follower_uid=@touid limit 1) {{end}} order by ctime desc limit @limit)
	//FindUserFollower(uid, lastid, limit int64) ([]*gen.T, error)
}

type FollowCountMethod interface {
	//sql(insert into @@table (uid, follow_count) values (@uid, 1) on duplicate key update follow_count=follow_count+1)
	IncrFollowCount(uid int64) error
	//sql(update @@table set follow_count=follow_count-1 where uid=@uid limit 1)
	DecrFollowCount(uid int64) error
	//sql(insert into @@table (uid, follower_count) values (@uid, @cnt) on duplicate key update follower_count=follower_count+@cnt)
	IncrByFollowerCount(uid, cnt int64) error
	//sql(update @@table set follower_count=follower_count-@cnt where uid=@uid limit @cnt)
	DecrByFollowerCount(uid, cnt int64) error
	//sql(select uid, follow_count, follower_count from @@table where uid in (@uids))
	FindUsersRelationCount(uids []int64) ([]*gen.T, error)
}

type ExtraFollowerMethod interface {
	//sql(insert into extra_follower (uid) values (@uid))
	InsertFollower(uid int64) error
	//sql(insert into extra_follower (uid, stats) values (@uid, 1))
	DeleteFollower(uid int64) error
	//sql(select id, uid, stats from extra_follower limit @limit)
	FindUnSyncRecord(limit int64) ([]*gen.T, error)
	//sql(select id, uid, stats from extra_follower where uid=@uid limit @limit)
	FindUnSyncRecordByUID(uid, limit int64) ([]*gen.T, error)
	//sql(delete from extra_follower where id in (@ids))
	DeleteRecord(ids []int64) error
}
