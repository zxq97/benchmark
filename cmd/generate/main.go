package main

import (
	"bench/dal"
	"bench/dal/method"

	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "dal/query",
		ModelPkgPath: "dal/model",
	})
	g.UseDB(dal.SocialDB)

	g.ApplyInterface(func(method.FollowMethod) {}, g.GenerateModel("follow"))
	g.ApplyInterface(func(followerMethod method.FollowerMethod) {}, g.GenerateModel("follower"))
	g.ApplyInterface(func(method.FollowCountMethod) {}, g.GenerateModel("follow_count"))
	g.ApplyInterface(func(method.ExtraFollowerMethod) {}, g.GenerateModel("extra_follower"))

	g.Execute()
}
