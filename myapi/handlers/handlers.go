package handlers

import (
	"io"
	"log"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yourname/reponame/models"

)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {

	var reqArticle models.Article

  if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
    http.Error(w, "fail to decode json\n", http.StatusBadRequest)
  }
	//fmt.Printf("%+v\n", reqArticle)

	article := reqArticle

  json.NewEncoder(w).Encode(article)
}

func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return 
		}
	} else {
		page = 1
	}

  log.Println(page)

  articleList := []models.Article{models.Article1, models.Article2}
  json.NewEncoder(w).Encode(articleList)

  json.NewEncoder(w).Encode(articleList)
}

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
  articleID, err := strconv.Atoi(mux.Vars(req)["id"])
  if err != nil {
    http.Error(w, "Invalid query parameter", http.StatusBadRequest)
    return 
  }
  // コンパイルエラーを回避
  log.Println(articleID)

  article := models.Article1
  
  json.NewEncoder(w).Encode(article)
}

func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
  var reqArticle models.Article

  if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
    http.Error(w, "fail to decode json\n", http.StatusBadRequest)
  }

  article := models.Article1
  json.NewEncoder(w).Encode(article)
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
  var reqComment models.Comment
  if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
    http.Error(w, "fail to decode json\n", http.StatusBadRequest)
  }
  comment := reqComment

  json.NewEncoder(w).Encode(comment)
}
