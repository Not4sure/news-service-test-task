package mongo

import (
	"context"

	"github.com/google/uuid"
	"github.com/not4sure/news-service-test-task/internal/domain/article"
	"github.com/not4sure/news-service-test-task/internal/domain/article/mongo/filter"
	"github.com/not4sure/news-service-test-task/internal/domain/article/mongo/sort"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	articlesCollection = "articles"
)

// MongoRepository is a mongodb implementation of Article Repository.
type MongoRepository struct {
	db *mongo.Database
}

// New creates MongoRepository
func New(ctx context.Context, uri string, dbName string) article.Repository {
	opts := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}

	db := client.Database(dbName)

	return MongoRepository{
		db: db,
	}
}

// articles returns pointer to articles collection
func (m MongoRepository) articles() *mongo.Collection {
	return m.db.Collection(articlesCollection)
}

// Delete implements article.Repository.
func (mr MongoRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := mr.articles().DeleteOne(ctx, filter.ByID(id))
	return err
}

// GetByID implements article.Repository.
func (mr MongoRepository) GetByID(ctx context.Context, id uuid.UUID) (article.Article, error) {
	res := mr.articles().FindOne(ctx, filter.ByID(id))

	err := res.Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return article.Article{}, article.ErrArticleNotFound
		}

		return article.Article{}, err
	}

	var m articleModel
	err = res.Decode(&m)
	if err != nil {
		return article.Article{}, err
	}

	a, err := domainArticleFromModel(m)
	if err != nil {
		return article.Article{}, err
	}

	return a, nil
}

// List implements article.Repository.
func (mr MongoRepository) List(ctx context.Context) (aa []article.Article, err error) {
	opts := options.Find()
	opts.Sort = sort.ByCreatedAt(sort.GreaterFirst)

	cur, err := mr.articles().Find(ctx, filter.All(), opts)
	if err != nil {
		return
	}

	for cur.Next(ctx) {
		var m articleModel
		if err = cur.Decode(&m); err != nil {
			return
		}

		var a article.Article
		a, err = domainArticleFromModel(m)
		if err != nil {
			return
		}

		aa = append(aa, a)
	}

	return
}

// Store implements article.Repository.
func (m MongoRepository) Store(ctx context.Context, a article.Article) error {
	mdl := modelFromDomainArticle(a)

	_, err := m.articles().InsertOne(ctx, mdl, options.InsertOne())
	return err
}

// Update implements article.Repository.
func (mr MongoRepository) Update(ctx context.Context, id uuid.UUID, update func(*article.Article) error) error {
	a, err := mr.GetByID(ctx, id)
	if err != nil {
		return err
	}

	err = update(&a)
	if err != nil {
		return err
	}

	_, err = mr.articles().ReplaceOne(ctx, filter.ByID(a.ID()), modelFromDomainArticle(a))
	return err
}
