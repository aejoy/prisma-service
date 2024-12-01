package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/aejoy/prisma-service/pkg/consts"
	"github.com/aejoy/prisma-service/pkg/utils"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

type DB struct {
	buckets []*pgxpool.Pool
}

func NewPostgres(urls []string) (*DB, error) {
	buckets := make([]*pgxpool.Pool, len(urls))

	for i, url := range urls {
		bucket, err := pgxpool.New(context.TODO(), url)
		if err != nil {
			return nil, fmt.Errorf("postgres.connect error: %v\n", err)
		}

		if err := bucket.Ping(context.TODO()); err != nil {
			return nil, fmt.Errorf("postgres.ping error: %v\n", err)
		}

		db, err := sql.Open("postgres", url)
		if err != nil {
			return nil, errors.Wrap(err, "sql.open")
		}

		if err := utils.PostgresMigrate(db); err != nil {
			return nil, errors.Wrap(err, "migrate")
		}

		buckets[i] = bucket
	}

	return &DB{buckets}, nil
}

func (db *DB) GetBucket(shardKey string) (*pgxpool.Pool, int, error) {
	bucketsCount := len(db.buckets)

	shardIndex, err := utils.GetShardIndex(shardKey, bucketsCount)
	if err != nil {
		return nil, shardIndex, err
	}

	if shardIndex <= bucketsCount {
		return db.buckets[shardIndex], shardIndex, nil
	}

	return nil, shardIndex, consts.NotFoundBucketErr
}
