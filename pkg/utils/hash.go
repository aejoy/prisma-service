package utils

import (
	"math"

	"github.com/aejoy/prisma-service/pkg/consts"
	"github.com/zeebo/xxh3"
)

func GetShardIndex(shardKey string, shards int) (int, error) {
	hash := xxh3.New()
	_, err := hash.Write([]byte(shardKey))

	if shards < 0 {
		return 0, consts.ErrOverflowOccurred
	}

	sum := hash.Sum64() % uint64(shards)
	if sum > math.MaxInt {
		return 0, consts.ErrOverflowOccurred
	}

	return int(sum), err
}
