package domain

import (
	"math/rand"
	"synergybench/entity"
	"testing"
	"time"
)

var tags = generateSliceInt64(10)
var allowed = generateTags(5000)

func BenchmarkCheckTags(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = CheckTagsMap(allowed, tags)
	}
}

func BenchmarkCheckTagsParallel(b *testing.B) {
	b.ReportAllocs()

	b.RunParallel(
		func(pb *testing.PB) {
			for pb.Next() {
				_ = CheckTagsMap(allowed, tags)
			}
		},
	)

}

func BenchmarkCheckLoopTags(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = CheckTagsLoop(allowed, tags)
	}
}

func BenchmarkCheckTagsLoopParallel(b *testing.B) {
	b.ReportAllocs()

	b.RunParallel(
		func(pb *testing.PB) {
			for pb.Next() {
				_ = CheckTagsLoop(allowed, tags)
			}
		},
	)

}

func generateTags(n int) []entity.Tag {
	var tags []entity.Tag
	var now = time.Now()

	for i := 0; i < n; i++ {
		tags = append(
			tags, entity.Tag{
				ID:        int64(i),
				Title:     "books",
				CreatedAt: now,
				UpdatedAt: now,
			},
		)
	}

	return tags
}

func generateSliceInt64(n int64) []int64 {
	rand.Seed(time.Now().Unix())

	var s []int64

	for i := n; i > 0; i-- {
		s = append(s, rand.Int63n(n))
	}

	return s
}
