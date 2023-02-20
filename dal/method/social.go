package method

import "gorm.io/gen"

type FollowMethod interface {
	//sql(insert into follow (uid, follow_uid) values (@uid, @touid))
	InsertFollow(uid, touid int64) error
	//sql(delete from follow where uid=@uid and follow=@uid limit 1)
	DeleteFollow(uid, touid int64) error
	//sql(select follow, ctime from follow where uid=@uid)
	FindUserFollow(uid int64) ([]*gen.T, error)
}

type FollowerMethod interface {
	//sql(insert into follower (uid, follower_uid) values (@uid, @touid))
	InsertFollower(uid, touid int64) error
	//sql(delete from follower where uid=@uid and follower_uid=@touid limit 1)
	DeleteFollower(uid, touid int64) error
	//sql(select follower_uid, ctime from follower where uid=@uid order by ctime desc limit @limit)
	FindUserFollower(uid, limit int64) ([]*gen.T, error)
	//sql(select follower_uid, ctime from follower where uid=@uid and id < (select id from follower where uid=@uid and follower_uid=@lastid) order by ctime desc limit @limit)
	FindUserFollowerByLastID(uid, lastid, limit int64) ([]*gen.T, error)
}

type FollowCountMethod interface {
	//sql(insert into follow_count (uid, follow_count) values (@uid, 1) on duplicate key update follow_count=follow_count+1)
	IncrFollowCount(uid int64) error
	//sql(update follow_count set follow_count=follow_count-1 where uid=@uid limit 1)
	DecrFollowCount(uid int64) error
	//sql(insert into follow_count (uid, follower_count) values (@uid, @cnt) on duplicate key update follower_count=follower_count+@cnt)
	IncrByFollowerCount(uid, cnt int64) error
	//sql(update follow_count set follower_count=follower_count-@cnt where uid=@uid limit 1)
	DecrByFollowerCount(uid, cnt int64) error
	//sql(select uid, follow_count, follower_count follow_count where uid in (@uids))
	FindUsersRelationCount(uids []int64) ([]*gen.T, error)
}

type ExtraFollowerMethod interface {
	//sql(insert into extra_follower (uid) values (@uid))
	InsertFollower(uid int64) error
	//sql(insert into extra_follower (uid, stats) values (@uid, 1))
	DeleteFollower(uid int64) error
	//sql(select id, uid, stats from extra_follower limit @limit)
	FindUnSyncRecord(limit int64) ([]*gen.T, error)
	//sql(select id, uid, stats from extra_follower where uid=@uid)
	FindUnSyncRecordByUID(uid int64) ([]*gen.T, error)
	//sql(delete from extra_follower where id in (@ids))
	DeleteRecord(ids []int64) error
}
