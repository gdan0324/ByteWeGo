namespace go api

struct User {
    1: i64 id
    2: string name
    3: i64 follow_count
    4: i64 follower_count
    5: bool is_follow
}

struct CheckUserRequest {
    1: string username (vt.max_size = "32")
    2: string password (vt.max_size = "32")
}

struct CheckUserResponse {
    1: i32 status_code
    2: string status_msg
    3: i64 user_id
    4: string token
}

struct CreateUserRequest {
    1: string username (vt.max_size = "32")
    2: string password (vt.max_size = "32")
}

struct CreateUserResponse {
    1: i32 status_code
    2: string status_msg
    3: i64 user_id
    4: string token
}

struct GetUserRequest {
    1: i64 user_id
    2: string token
}

struct GetUserResponse {
    1: i32 status_code
    2: string status_msg
    3: User user
}

struct Video {
    1: i64 id
    2: User user
    3: string play_url
    4: string cover_url
    5: i64 favorite_count
    6: i64 comment_count
    7: bool is_favorite
    8: string title
}

struct GetVideoListRequest {
    1: i64 user_id
    2: string token
}

struct GetVideoListResponse {
    1: i32 status_code
    2: string status_msg
    3: list<Video> video_list
}

struct CreateVideoRequest {
    1: string token
    2: list<byte> data
    3: string title
}

struct CreateVideoResponse {
    1: i32 status_code
    2: string status_msg
}

service ApiService {
    CheckUserResponse CheckUser(1: CheckUserRequest req) (api.post="/douyin/user/login")
    CreateUserResponse CreateUser(1: CreateUserRequest req) (api.post="/douyin/user/register")
    GetUserResponse GetUser(1: GetUserRequest req) (api.get="/douyin/user")
    CreateVideoResponse CreateVideo(1: CreateVideoRequest req) (api.post="/douyin/publish/action")
    GetVideoListResponse GetVideoList(1: GetVideoListRequest req) (api.get="/douyin/publish/list")
}