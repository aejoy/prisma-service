package db

import (
	"context"

	"github.com/aejoy/prisma-service/pkg/consts"

	"github.com/aejoy/prisma-service/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

func (db *DB) GetPhotos() ([]*models.Photo, error) {
	photos := make([]*models.Photo, 0)

	for _, bucket := range db.buckets {
		rows, err := bucket.Query(context.TODO(), "SELECT id, creator, \"to\", url, blur_hash, height, width, size, published, updated, archived FROM photo LIMIT 25;")
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			photo := new(models.Photo)

			if err := rows.Scan(
				&photo.ID, &photo.Creator, &photo.To, &photo.URL, &photo.BlurHash,
				&photo.Height, &photo.Width, &photo.Size,
				&photo.Published, &photo.Updated, &photo.Archived,
			); err != nil {
				return nil, err
			}

			kb := float32(photo.Size) / consts.Kilobyte
			mb := float32(photo.Size) / consts.Megabyte

			if mb >= 1 {
				photo.SizeInMB = mb
			} else if kb >= 1 {
				photo.SizeInKB = kb
			}

			photos = append(photos, photo)
		}
	}

	return photos, nil
}

func (db *DB) GetPhotosByIDs(photoIDs []string) ([]*models.Photo, error) {
	buckets := map[*pgxpool.Pool][]string{}
	photos := make([]*models.Photo, 0)

	for _, photoID := range photoIDs {
		bucket, _, err := db.GetBucket(photoID)
		if err != nil {
			return nil, err
		}

		if buckets[bucket] == nil {
			buckets[bucket] = make([]string, 0)
		}

		buckets[bucket] = append(buckets[bucket], photoID)
	}

	for bucket, photoIDs := range buckets {
		rows, err := bucket.Query(context.TODO(), "SELECT id, creator, \"to\", url, blur_hash, height, width, size, published, updated, archived FROM photo WHERE id = ANY($1::text[]);", photoIDs)
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			photo := new(models.Photo)
			if err := rows.Scan(&photo.ID, &photo.Creator, &photo.To, &photo.URL, &photo.BlurHash, &photo.Height, &photo.Width, &photo.Size, &photo.Published, &photo.Updated, &photo.Archived); err != nil {
				return nil, err
			}

			kb := float32(photo.Size) / consts.Kilobyte
			mb := float32(photo.Size) / consts.Megabyte

			if mb >= 1 {
				photo.SizeInMB = mb
			} else if kb >= 1 {
				photo.SizeInKB = kb
			}

			photos = append(photos, photo)
		}
	}

	return photos, nil
}

func (db *DB) CreatePhoto(photoID, creator, url, blurHash string, height, width, size int) error {
	bucket, _, err := db.GetBucket(photoID)
	if err != nil {
		return err
	}

	_, err = bucket.Exec(context.TODO(),
		"INSERT INTO photo(id, creator, url, blur_hash, height, width, size) VALUES($1, $2, $3, $4, $5, $6, $7)",
		photoID, creator, url, blurHash, height, width, size,
	)

	return err
}
