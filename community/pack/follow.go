package pack

import (
	"github.com/gdan0324/ByteWeGo/community/dal/db"
	"github.com/gdan0324/ByteWeGo/community/kitex_gen/communityservice"
)

func User(u *db.FollowInfo, isFollow bool) *communityservice.User {
	if u == nil {
		return nil
	}
	return &communityservice.User{
		Id:            int64(u.UserId),
		Name:          u.UserName,
		FollowCount:   int64(u.FollowCount),
		FollowerCount: int64(u.FollowerCount),
		IsFollow:      isFollow,
	}
}

func Users(us []*db.FollowInfo, isFollow []bool) []*communityservice.User {
	users := make([]*communityservice.User, 0)
	for i := range us {
		if temp := User(us[i], isFollow[i]); temp != nil {
			users = append(users, temp)
		}
	}
	return users
}

func Friend(u *db.FollowInfo, isFollow bool, lastMessage string) *communityservice.FriendUser {
	if u == nil {
		return nil
	}
	return &communityservice.FriendUser{
		Id:            int64(u.UserId),
		Name:          u.UserName,
		FollowCount:   int64(u.FollowCount),
		FollowerCount: int64(u.FollowerCount),
		IsFollow:      isFollow,
		Message:       lastMessage,
	}
}

func Friends(us []*db.FollowInfo, isFollow []bool, lastMessage []string) []*communityservice.FriendUser {
	friendusers := make([]*communityservice.FriendUser, 0)
	for i := range us {
		if temp := Friend(us[i], isFollow[i], lastMessage[i]); temp != nil {
			friendusers = append(friendusers, temp)
		}
	}
	return friendusers
}
