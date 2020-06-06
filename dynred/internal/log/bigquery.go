package log

import (
	"cloud.google.com/go/bigquery"
	"context"
)

type BigQueryRepository struct {
	ctx      context.Context
	inserter *bigquery.Inserter
}

func NewBigQueryRepository(projectId string, datasetId string, tableId string) (Repository, error) {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectId)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	inserter := client.Dataset(datasetId).Table(tableId).Inserter()
	return &BigQueryRepository{ctx: ctx, inserter: inserter}, nil
}

func (r *BigQueryRepository) Dump(log Log) error {
	err := r.inserter.Put(r.ctx, []*Log{&log})
	if err != nil {
		return err
	}
	return nil
}
