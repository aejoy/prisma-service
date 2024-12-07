package consts

import (
	"errors"

	"github.com/aejoy/prisma-service/pkg/types"
)

const (
	AvatarImageType types.ImageType = iota
	BannerImageType types.ImageType = iota
	PhotoImageType  types.ImageType = iota
)

const (
	Kibibyte  = 1024
	Mebibyte  = 1024 * Kibibyte
	BodyLimit = 30 * Mebibyte

	Kilobyte = 1000
	Megabyte = 1000 * Kilobyte

	NanoAlphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	NanoLength   = 16

	BlurHashXComponents = 4
	BlurHashYComponents = 3

	AvatarHeight = 256
	AvatarWidth  = 256

	BannerHeight = 382
	BannerWidth  = 728

	AVIFQuality = 85
)

var (
	ErrNotFoundBucket = errors.New("bucket is null")

	ErrPostgresConnection = errors.New("postgres connect error")
	ErrPostgresPing       = errors.New("postgres ping error")

	ErrSQLOpen            = errors.New("sql open error")
	ErrMigrate            = errors.New("migrate error")
	ErrMigrateInstance    = errors.New("migrate with instance sql error")
	ErrMigrateOpenFile    = errors.New("open migrations file error")
	ErrMigrateNewInstance = errors.New("new migration instance error")
	ErrMigrateUp          = errors.New("up migrate error")

	ErrOverflowOccurred = errors.New("overflow occurred")
)
