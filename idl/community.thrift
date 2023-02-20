namespace go communityservice

struct User {
    1: i64 id
    2: string name
    3: i64 follow_count
    4: i64 follower_count
    5: bool is_follow
}

struct Follow {
    1: i64 user_id
    2: i64 follow_id
}

struct GetFollowRequest {
    1: i64 user_id
    2: string token
}

struct GetFollowResponse{
    1: i32 status_code
    2: string status_msg
    3: list<User> user_list
}

struct GetFollowerRequest {
    1: i64 user_id
    2: string token
}

struct GetFollowerResponse {
    1: i32 status_code
    2: string status_msg
    3: list<User> user_list
}

struct FollowRequest {
    1: string token
    2: i64 to_user_id
    3: i32 action_type
}

struct FollowResponse {
    1: i32 status_code
    2: string status_msg
}

struct FriendUser {
    1: i64 id
    2: string name
    3: i64 follow_count
    4: i64 follower_count
    5: bool is_follow
    6: string message
    7: i64 msgType
    8: string avatar
}

struct CheckFriendRequest {
    1: i64 user_id
    2: string token
}

struct CheckFriendResponse {
    1: i32 status_code
    2: string status_msg
    3: list<FriendUser> friend_list
}

struct RelationActionRequest {
    1: string token
    2: i64 to_user_id
    3: i32 action_type
    4: string content
}

struct RelationActionResponse {
    1: i32 status_code
    2: string status_msg
}

service CommunityService {
    GetFollowResponse GetFollowList(1: GetFollowRequest req)
    GetFollowerResponse GetFollowerList(1: GetFollowerRequest req)
    FollowResponse Follow(1: FollowRequest req)
    CheckFriendResponse CheckFriend(1: CheckFriendRequest req)
    RelationActionResponse MessageAction(1: RelationActionRequest req)
}