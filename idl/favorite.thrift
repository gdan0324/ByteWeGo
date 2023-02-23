namespace go favoriteservice

struct User {
    1: i64 id
    2: string name
    3: i64 follow_count
    4: i64 follower_count
    5: bool is_follow
    6: string avatar
    7: string background_image
    8: string signature
    9: i64 total_favorited
    10: i64 work_count
    11: i64 favorite_count
}


struct Video {
  1: i64 id
  2: User author
  3: string play_url
  4: string cover_url
  5: i64 favorite_count
  6: i64 comment_count
  7: bool is_favorite
  8: string title
}

struct GetFavoriteListRequest {
    1: i64 user_id
    2: string token
}

struct GetFavoriteListResponse {
    1: i32 status_code
    2: string status_msg
    3: list<Video> video_list
}

struct DoFavoriteRequest {
    1: string token
    2: i64 video_id
    3: i32 action_type
}
struct DoFavoriteResponse {
    1: i32 status_code
    2: string status_msg
}

service FavoriteService {
    DoFavoriteResponse DoFavorite(1: DoFavoriteRequest req)
    GetFavoriteListResponse GetFavoriteList(1: GetFavoriteListRequest req)
}
