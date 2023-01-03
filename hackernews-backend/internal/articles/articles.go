package articles

import (
	database "github.com/fajrizulfikar/hackernews-backend/internal/pkg/db/migrations/mysql"
	"github.com/fajrizulfikar/hackernews-backend/internal/users"

	"log"
)

type Article struct {
	Id    string
	Title string
	Url   string
	User  *users.User
}

func (article Article) Save() int64 {
	stmt, err := database.DB.Prepare("INSERT INTO article(title,url,user_id) VALUES(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(article.Title, article.Url, article.User.Id)
	if err != nil {
		log.Fatal(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("Row inserted!")
	return id
}

func GetAll() []Article {
	stmt, err := database.DB.Prepare("select L.id, L.title, L.address, L.UserID, U.Username from Links L inner join Users U on L.UserID = U.ID")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var articles []Article
	var username string
	var id string
	for rows.Next() {
		var article Article
		err := rows.Scan(&article.Id, &article.Title, &article.Url, &id, &username)
		if err != nil {
			log.Fatal(err)
		}
		article.User = &users.User{
			Id:       id,
			Username: username,
		}
		articles = append(articles, article)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return articles
}
