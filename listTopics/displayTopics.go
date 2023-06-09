package forum

import (
	"database/sql"
	"fmt"
	t "forum/structs"
	t2 "forum/views"
	"net/http"
	"strings"
)

var TOPICSANDSESSION t.TopicsAndSession

func DisplayTopic(r *http.Request, db *sql.DB) {
	//displays all topics of the current category
	var name string
	var likes int
	var creationDate string
	var owner string
	var uuid string
	var nmbPosts int
	var filter string
	var id int
	var searchName string
	var query string
	var firstMessage string
	var username string
	var likeOrDislike int
	var category string
	var admin int
	// var lastPost string

	filter = r.FormValue("filter")
	DESCOASC := "DESC"

	if filter == "" {
		filter = "creationDate"
	}

	categoryUrl := strings.Split(r.URL.Path, "/")
	categoryUrl = strings.Split(categoryUrl[2], "=")
	urlCat := categoryUrl[1]

	if r.FormValue("filter") == "mostRecent" {
		filter = "creationDate"
	} else if r.FormValue("filter") == "oldest" {
		filter = "creationDate"
		DESCOASC = "ASC"
	}

	query = fmt.Sprintf("SELECT id, name, firstMessage, creationDate, owner, likes, nmbPosts, category, uuid FROM topics WHERE category = '%s' ORDER BY %s %s ", urlCat, filter, DESCOASC)

	if r.FormValue("searchbar") != "" {
		searchName = "%" + r.FormValue("searchbar") + "%"
		query = fmt.Sprintf("SELECT id, name, firstMessage, creationDate, owner, likes, nmbPosts, category, uuid FROM topics WHERE name LIKE '%s' AND category = '%s' ORDER BY %s %s", searchName, urlCat, filter, DESCOASC)
	}

	cookie, err := r.Cookie("session")
	if err != nil {
		fmt.Println(err)
		TOPICSANDSESSION.SessionUser = ""
	} else {
		queryGetName := fmt.Sprintf("SELECT username FROM users WHERE uuid = '%s'", cookie.Value)
		row, err := db.Query(queryGetName)
		if err != nil {
			fmt.Println(err)
		} else {
			for row.Next() {
				defer row.Close()
				row.Scan(&username)
			}
			TOPICSANDSESSION.SessionUser = username
		}
	}

	row2, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	} else {
		TOPICSANDSESSION.Topics = nil
		for row2.Next() {
			defer row2.Close()
			err = row2.Scan(&id, &name, &firstMessage, &creationDate, &owner, &likes, &nmbPosts, &category, &uuid)
			if err != nil {
				fmt.Println(err)
			} else {
				topicIndex := len(TOPICSANDSESSION.Topics)
				TOPICSANDSESSION.Topics = append(TOPICSANDSESSION.Topics, t.Topic{})
				TOPICSANDSESSION.Topics[topicIndex].Name = name
				TOPICSANDSESSION.Topics[topicIndex].Likes = likes
				TOPICSANDSESSION.Topics[topicIndex].CreationDate = t2.DisplayTime(creationDate)
				TOPICSANDSESSION.Topics[topicIndex].Owner = owner
				TOPICSANDSESSION.Topics[topicIndex].NmbPosts = nmbPosts
				TOPICSANDSESSION.Topics[topicIndex].Uuid = uuid
				TOPICSANDSESSION.Topics[topicIndex].Id = id
				TOPICSANDSESSION.Topics[topicIndex].FirstMessage = firstMessage

				if TOPICSANDSESSION.SessionUser != "" {

					queryGetIfAdmin := fmt.Sprintf("SELECT admin FROM users WHERE username = '%s'", TOPICSANDSESSION.SessionUser)

					row, err := db.Query(queryGetIfAdmin)
					if err != nil {
						fmt.Println(err)
					} else {
						for row.Next() {
							defer row.Close()
							err = row.Scan(&admin)
							if err != nil {
								fmt.Println(err)
							}
						}
					}

					if owner == TOPICSANDSESSION.SessionUser || admin == 1 {
						TOPICSANDSESSION.Topics[topicIndex].IsOwnerOrAdmin = 1
					}

					TOPICSANDSESSION.Category = category
					if cookie.Value != "" {
						checkIfLiked := fmt.Sprintf("SELECT likeOrDislike FROM likesFromUser WHERE uuidUser = '%s' AND uuidLiked = '%s'", cookie.Value, uuid)
						row, err := db.Query(checkIfLiked)
						if err != nil {
							fmt.Println(err)
						} else {
							for row.Next() {
								defer row.Close()
								err = row.Scan(&likeOrDislike)
								if err != nil {
									fmt.Println(err)
								} else {
									if likeOrDislike == 1 {
										TOPICSANDSESSION.Topics[topicIndex].IsLiked = 1
									} else if likeOrDislike == -1 {
										TOPICSANDSESSION.Topics[topicIndex].IsDisliked = 1
									}
								}
							}
						}
					}
					// if lastPost == "" {
					// 	lastPost = creationDate
					// }
					// TOPICSANDSESSION.Topics[topicIndex].LastPost = lastPost
				}
			}
		}
	}
}
