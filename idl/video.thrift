namespace go videoservice

struct User {
    1: i64 id
    2: string name
    3: i64 follow_count
    4: i64 follower_count
    5: bool is_follow
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
    2: binary data
    3: string title
}

struct CreateVideoResponse {
    1: i32 status_code
    2: string status_msg
}

struct GetFeedRequest {
    1: i64 latest_time
    2: string token
}

struct GetFeedResponse {
    1: i32 status_code
    2: string status_msg
    3: list<Video> video_list
    4: i64 next_time
}

service VideoService {
    CreateVideoResponse CreateVideo(1: CreateVideoRequest req)
    GetVideoListResponse GetVideoList(1: GetVideoListRequest req)
    GetFeedResponse GetFeed(1: GetFeedRequest req)
}