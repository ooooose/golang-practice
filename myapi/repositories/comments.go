package repositories

import (
    "database/sql"

    "github.com/yourname/reponame/models"
)

func InsertComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
    const sqlStr = `
        insert into comments (article_id, message, created_at) values
        (?, ?, now());
    `

    var newComment models.Comment
    newComment.ArticleID, newComment.Message = comment.ArticleID, comment.Message

    result, err := db.Exec(sqlStr, comment.ArticleID, comment.Massage)
    if err != nil {
        return models.Comment{}, err
    }

    if err != nil {
        return nil, err
    }

    defer rows.Close()

    commentArray := make([]models.Comment, 0)
    for rows.Next() {
        var comment models.Comment
        var createdTime sql.NullTime
        rows.Scan(&comment.CommentID, &comment.ArticleID, &comment.Message, &createdTime)

        if createdTime.Valid {
            comment.CreatedAt = createdTime.TIme
        }

        commentArray = append(commentArray, comment)
    }

    return commentArray, nil
}
