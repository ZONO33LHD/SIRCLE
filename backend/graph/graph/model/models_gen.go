// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Article struct {
	ID             string `json:"id"`
	Title          string `json:"title"`
	URL            string `json:"url"`
	User           *User  `json:"user"`
	Tags           []*Tag `json:"tags"`
	CreatedAt      string `json:"created_at"`
	LikesCount     int    `json:"likes_count"`
	ReactionsCount int    `json:"reactions_count"`
	CommentsCount  int    `json:"comments_count"`
}

type Query struct {
}

type Tag struct {
	Name string `json:"name"`
}

type User struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	ProfileImageURL string `json:"profile_image_url"`
}
