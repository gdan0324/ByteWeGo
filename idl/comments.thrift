namespace go commentservice
struct Comment{
    i64 id
    User user
    string comment_text
    string comment_date
}
struct User {
    1: i64 id
    2: string name
    3: i64 follow_count
    4: i64 follower_count
    5: bool is_follow
}
struct CommentActionRequest{
    i64 user_id
    string token
    i64 video_id
    i32 action_type
    i64 comment_id
    string comment_text
}

struct CommentActionResponse{
    i32 status_code
    string status_msg
    Comment comment
}

struct GetCommentsRequest{
    string token
    i64 video_id
}

struct GetCommentsResponse{
    i32 status_code
    string status_msg
    list<Comment> comment_list
}

service CommentService{
    CommentActionResponse CommentAction(1: CommentActionRequest req)
    GetCommentsResponse GetComments(1: GetCommentsRequest req)
}