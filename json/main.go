package main

import (
    "fmt"
    "time"
)

type Comment struct {
	CommentID int
	ArticleID int
	Message   string
	CreatedAt time.Time
}

type Article struct {
	ID 		   		int
	Title     	string
	Contents  	string
	UserName    string
	NiceNum     int
	CommentList []Comment
	CreatedAt   time.Time
}

func main() {
	comment1 := Comment{
		CommentID: 1,
		ArticleID: 1,
		Message: "test comment1",
		CreatedAt: time.Now(),
	}

	comment2 := Comment{
		CommentID: 2,
		ArticleID: 1,
		Message: "second comment",
		CreatedAt: time.Now(),
	}

	article := Article{
		ID: 1,
		Title: "first article",
		Contents: "this is the test article.",
		UserName: "oose",
		NiceNum: 1,
		CommentList: []Comment{comment1, comment2},
		CreatedAt:   time.Now(),
	}

	fmt.Printf("%+v\n", article)
}