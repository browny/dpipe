package cloud

import (
	"context"
	"fmt"

	"github.com/browny/dpipe/internal/domain/dcard"
	"github.com/browny/dpipe/pkg/api"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

type BqClient struct {
	client *bigquery.Client
}

func Query(projectID string) error {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		return err
	}

	q := client.Query(`
    	SELECT year, SUM(number) as num
    	FROM ` + "`bigquery-public-data.usa_names.usa_1910_2013`" + `
    	WHERE name = "William"
    	GROUP BY year
    	ORDER BY year`)
	it, err := q.Read(ctx)
	if err != nil {
		return err
	}

	for {
		var values []bigquery.Value
		err := it.Next(&values)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		fmt.Println(values)
	}

	return nil
}

func CreateTable(projectID string) error {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		return err
	}

	schema, err := bigquery.InferSchema(dcard.Post{})
	if err != nil {
		return err
	}

	table := client.Dataset("dcard").Table("posts")
	err = table.Create(ctx, &bigquery.TableMetadata{Schema: schema})
	if err != nil {
		return err
	}

	return nil
}

func WritePosts(projectID string) error {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	table := client.Dataset("dcard").Table("posts")
	u := table.Uploader()

	posts, err := api.GetPostsUntil("dressup", 230955909)
	if err != nil {
		return err
	}
	err = u.Put(ctx, posts)
	if err != nil {
		return err
	}
	return nil
}
