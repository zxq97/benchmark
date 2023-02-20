package social

func serviceFollow(uid, touid int64) {
	followch <- &asyncFollow{uid: uid, touid: touid}
}
