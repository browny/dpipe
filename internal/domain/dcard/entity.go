package dcard

import "time"

type Post struct {
	ID           int           `json:"id"            bigquery :"id"`
	Title        string        `json:"title"         bigquery :"title"`
	Excerpt      string        `json:"excerpt"       bigquery :"excerpt"`
	CreatedAt    time.Time     `json:"createdAt"     bigquery :"created_at"`
	UpdatedAt    time.Time     `json:"updatedAt"     bigquery :"updated_at"`
	CommentCount int           `json:"commentCount"  bigquery :"comment_count"`
	LikeCount    int           `json:"likeCount"     bigquery :"like_count"`
	Tags         []interface{} `json:"tags"          bigquery :"tags"`
	Topics       []string      `json:"topics"        bigquery :"topics"`
	ForumName    string        `json:"forumName"     bigquery :"forum_name"`
	ForumAlias   string        `json:"forumAlias"    bigquery :"forum_alias"`
	Gender       string        `json:"gender"        bigquery :"gender"`
	School       string        `json:"school"        bigquery :"school"`
	WithImages   bool          `json:"withImages"    bigquery :"with_images"`
	WithVideos   bool          `json:"withVideos"    bigquery :"with_videos"`
}

type PostMeta struct {
	ForumAlias  string `json:"forumAlias"  bigquery :"forum_alias"`
	StartPostID int    `json:"startPostId" bigquery :"start_post_id"`
	EndPostID   int    `json:"endPostId"   bigquery :"end_post_id"`
	CreatedAt   string `json:"createdAt"   bigquery :"created_at"`
}

type Comment struct {
	ID        string    `json:"id"        bigquery :"id"`
	PostID    int       `json:"postId"    bigquery :"post_id"`
	CreatedAt time.Time `json:"createdAt" bigquery :"created_at"`
	UpdatedAt time.Time `json:"updatedAt" bigquery :"updated_at"`
	Floor     int       `json:"floor"     bigquery :"floor"`
	Content   string    `json:"content"   bigquery :"content"`
	Gender    string    `json:"gender"    bigquery :"gender"`
}
