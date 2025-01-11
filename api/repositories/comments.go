package repositories

import (
	"database/sql"

	"github.com/yourname/reponame/models"
)

func InsertComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
	const sqlStr = `
		insert into comments (article_id, username, comment, created_at) values
		(?, ?, ?, now());
	`

	var newComment models.Comment
	newComment.ArticleID, newComment.Message = comment.ArticleID, comment.Message

	result, err := db.Exec(sqlStr, comment.ArticleID, comment.Message)
	if err != nil {
		return models.Comment{}, err
	}

	id, _ := result.LastInsertId()

	newComment.CommentID = int(id)

	return newComment, nil
}

func SElectCommentList(db *sql.DB, articleID int) ([]models.Comment, error) {
	const sqlStr = `
		select comment_id, article_id, comment, created_at
		from comments
		where article_id = ?;
	`

	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		return nil, err
	}

	commentArray := make([]models.Comment, 0)

	for rows.Next() {
		var comment models.Comment
	  var createdTime sql.NullTime		
		rows.Scan(&comment.CommentID, &comment.ArticleID, &comment.Message, &createdTime)

		if createdTime.Valid {
			comment.CreatedAt = createdTime.Time
		}

		commentArray = append(commentArray, comment)
	}

	return commentArray, nil
}