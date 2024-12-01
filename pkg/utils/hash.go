package utils

import "github.com/zeebo/xxh3"

func GetShardIndex(shardKey string, shards int) (int, error) {
	hash := xxh3.New()
	_, err := hash.Write([]byte(shardKey))
	return int(hash.Sum64() % uint64(shards)), err
}
