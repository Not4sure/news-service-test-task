package app

import (
	"context"

	"github.com/google/uuid"
	"github.com/not4sure/news-service-test-task/internal/domain/article"
)

type Application struct {
	repo article.Repository
}

func New(articleRepo article.Repository) Application {
	return Application{
		repo: articleRepo,
	}
}

func (app *Application) CreateArticle(
	ctx context.Context,
	title string,
	content string,
) (vvm ArticleViewModel, err error) {
	a, err := article.NewArticle(title, content)
	if err != nil {
		return ArticleViewModel{}, err
	}

	err = app.repo.Store(ctx, *a)
	if err != nil {
		return ArticleViewModel{}, err
	}

	return viewModelFromDomainArticle(*a), nil
}

func (app *Application) ListArticles(ctx context.Context) (vvm []ArticleViewModel, err error) {
	aa, err := app.repo.List(ctx)
	if err != nil {
		return
	}

	for _, a := range aa {
		vvm = append(vvm, viewModelFromDomainArticle(a))
	}

	return
}

func (app *Application) GetByID(ctx context.Context, id uuid.UUID) (ArticleViewModel, error) {
	a, err := app.repo.GetByID(ctx, id)
	if err != nil {
		return ArticleViewModel{}, err
	}

	return viewModelFromDomainArticle(a), nil
}

func (app *Application) UpdateArticle(
	ctx context.Context,
	id uuid.UUID,
	title string,
	content string,
) (ArticleViewModel, error) {
	err := app.repo.Update(ctx, id, func(a *article.Article) error {
		err := a.SetTitle(title)
		if err != nil {
			return err
		}

		err = a.SetContent(content)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return ArticleViewModel{}, err
	}

	a, err := app.repo.GetByID(ctx, id)
	if err != nil {
		return ArticleViewModel{}, err
	}

	return viewModelFromDomainArticle(a), nil
}

func (app *Application) Delete(ctx context.Context, id uuid.UUID) error {
	return app.repo.Delete(ctx, id)
}
