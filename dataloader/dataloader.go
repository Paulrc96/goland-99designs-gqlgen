package dataloader

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/paul/go-server/graph/model"
)

const loadersKey = "dataloaders"

type Loaders struct {
	PostsByUserId     PostLoader
	CommentsByPostIds CommentLoader
}

func Middleware(conn *sqlx.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "X-FirstUsers,Content-Type,access-control-allow-origin, access-control-allow-headers")
		// IMPORTANT: Set a good value for waitMilliseconds so that it does not create too many DB connections
		var waitMilliseconds = 5 * time.Millisecond
		firstUsers := r.Header.Get("X-FirstUsers")
		if firstUsers != "" {
			fmt.Println("No X-FirstUsers header received: ", firstUsers)
			firstUserInt, err := strconv.Atoi(firstUsers)
			if err != nil {
				log.Fatalln("Error getting X-FirstUsers header")
			}
			waitMilliseconds = time.Duration(firstUserInt/50) * time.Millisecond

			if firstUserInt <= 1000 {
				waitMilliseconds = 5 * time.Millisecond
			} else if firstUserInt > 1000 && firstUserInt <= 5000 {
				waitMilliseconds = 100 * time.Millisecond
			} else {
				waitMilliseconds = 1000 * time.Millisecond
			}
		} else {
			fmt.Println("No X-FirstUsers header!!!")
		}
		fmt.Println("Waiting", waitMilliseconds)
		ctx := context.WithValue(r.Context(), loadersKey, &Loaders{
			PostsByUserId: PostLoader{
				maxBatch: 50000,
				wait:     waitMilliseconds,
				fetch: func(ids []int) ([][]*model.Post, []error) {
					idsLength := len(ids)

					var postByUserMap = make(map[int][]*model.Post)

					stringIds := make([]string, idsLength)
					for j := 0; j < idsLength; j++ {
						stringIds[j] = fmt.Sprintf("%s", strconv.Itoa(ids[j]))
					}

					fmt.Println("WAITING FOR POSTS...", idsLength)

					posts := []*model.Post{}

					err := conn.Select(&posts, fmt.Sprintf("SELECT * FROM posts WHERE user_id IN (%s)", strings.Join(stringIds, ",")))
					if err != nil {
						fmt.Println("Error getting posts!!!!", len(ids), err)
						log.Fatalln(err)
					}

					for _, postRow := range posts {
						postByUserMap[*postRow.UserID] = append(postByUserMap[*postRow.UserID], postRow)
					}

					dbPosts := [][]*model.Post{}
					for _, userId := range ids {
						dbPosts = append(dbPosts, postByUserMap[userId])
					}

					return dbPosts, nil
				},
			},
			CommentsByPostIds: CommentLoader{
				maxBatch: 60000,
				wait:     waitMilliseconds,
				fetch: func(ids []int) ([][]*model.Comment, []error) {
					idsLength := len(ids)

					var commentsByPostMap = make(map[int][]*model.Comment)

					stringIds := make([]string, idsLength)
					for j := 0; j < idsLength; j++ {
						stringIds[j] = fmt.Sprintf("%s", strconv.Itoa(ids[j]))
					}

					comments := []*model.Comment{}
					fmt.Println("WAITING FOR COMMENTS...", idsLength)
					err := conn.Select(&comments, fmt.Sprintf("SELECT * FROM comments WHERE post_id IN (%s)", strings.Join(stringIds, ",")))
					if err != nil {
						fmt.Println("Error getting comments!!!!!", len(ids), err)
						log.Fatalln(err)
					}

					for _, comment := range comments {
						commentsByPostMap[*comment.PostID] = append(commentsByPostMap[*comment.PostID], comment)
					}

					dbCommentsByPost := [][]*model.Comment{}
					for _, userId := range ids {
						dbCommentsByPost = append(dbCommentsByPost, commentsByPostMap[userId])
					}

					return dbCommentsByPost, nil
				},
			},
		})
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
