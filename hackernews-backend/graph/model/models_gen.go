// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Article struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url"`
	User  *User  `json:"user"`
}

type ArticleVotes struct {
	ID        string `json:"id"`
	ArticleID string `json:"article_id"`
	UserID    string `json:"user_id"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type NewArticle struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type NewUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type RefreshTokenInput struct {
	Token string `json:"token"`
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
