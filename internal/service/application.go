package service

import (
	"context"

	"github.com/not4sure/news-service-test-task/internal/app"
	"github.com/not4sure/news-service-test-task/internal/domain/article"
	article_memory "github.com/not4sure/news-service-test-task/internal/domain/article/memory"
)

func NewApplication(ctx context.Context) app.Application {
	// In-memory Article repository with one article pre-inserted
	// for dev purposes.
	articleRepo := article_memory.New()
	articleRepo.Store(ctx, *article.MustNewArticle(
		"Test article",
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur ornare velit odio, ac sodales tortor tincidunt id. Mauris nec fermentum augue, sit amet aliquam urna. Ut nisl purus, aliquet quis felis ac, accumsan varius tortor. Mauris diam turpis, blandit sed metus quis, porta auctor odio. Mauris metus eros, pharetra sed mauris vel, rhoncus luctus sem. Maecenas justo sem, aliquam nec massa vitae, efficitur cursus ligula. Suspendisse et libero mauris. Donec libero neque, ultrices vel justo ac, fermentum accumsan diam. Donec eget metus gravida, commodo velit sed, egestas leo. Fusce ornare eros non purus aliquet, eu bibendum quam convallis. Mauris libero ligula, scelerisque lacinia odio facilisis, maximus dictum magna.",
	))

	return app.New(articleRepo)
}
