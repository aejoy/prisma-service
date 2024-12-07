package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/aejoy/prisma-service/pkg/consts"
	"github.com/aejoy/prisma-service/pkg/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	buckets []*pgxpool.Pool
}

func NewPostgres(urls []string) (*DB, error) {
	buckets := make([]*pgxpool.Pool, len(urls))

	for i, url := range urls {
		bucket, err := pgxpool.New(context.TODO(), url)
		if err != nil {
			return nil, fmt.Errorf("%w: %v", consts.ErrPostgresConnection, err)
		}

		if err := bucket.Ping(context.TODO()); err != nil {
			return nil, fmt.Errorf("%w: %v", consts.ErrPostgresPing, err)
		}

		db, err := sql.Open("postgres", url)
		if err != nil {
			return nil, fmt.Errorf("%w: %v", consts.ErrSQLOpen, err)
		}

		if err := utils.PostgresMigrate(db); err != nil {
			return nil, fmt.Errorf("%w: %v", consts.ErrMigrate, err)
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

	return nil, shardIndex, consts.ErrNotFoundBucket
}
