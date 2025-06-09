package service

import (
	"context"
	"time"

	"github.com/not4sure/news-service-test-task/internal/app"

	article_mongo "github.com/not4sure/news-service-test-task/internal/domain/article/mongo"
)

func NewApplication(ctx context.Context) app.Application {
	config, err := Load("./")
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*20)
	defer cancel()

	articleRepo := article_mongo.New(ctx, config.MongoURL, config.MongoDatabaseName)

	return app.New(articleRepo)
}
