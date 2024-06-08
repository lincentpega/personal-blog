package main

import (
	"content/app"
	"content/post"
	"content/post/adapters"
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

const dbName = "content"

func main() {
	r := chi.NewRouter()
	config := app.LoadConfiguration()
	ctx := context.Background()
	mongoRepo, err := adapters.NewMongoPostRepository(getMongoDatabase(ctx, config))
	if err != nil {
		panic(err.Error())
	}
	r.Route("/posts", func(r chi.Router) {
		r.Get("/{postId}", func(w http.ResponseWriter, r *http.Request) {
			postId := chi.URLParam(r, "postId")
			ctx := r.Context()
			p, err := mongoRepo.GetPostById(ctx, postId)
			if err != nil {
				return
			}
			_, err = w.Write([]byte(fmt.Sprintf("%v, %v, %v, %v", p.Id(), p.Text(), p.CreatedAt(), p.UpdatedAt())))
			if err != nil {
				return
			}
		})
		r.Post("/create", func(w http.ResponseWriter, r *http.Request) {
			err := mongoRepo.AddPost(r.Context(), post.NewWithText("Hello, world"))
			if err != nil {
				return
			}
		})
	})

	http.ListenAndServe(":3333", r)
}

func getMongoDatabase(ctx context.Context, config *app.Configuration) *mongo.Database {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MongoUrl))
	if err != nil {
		panic(err.Error())
	}
	db := client.Database(dbName)
	return db
}
