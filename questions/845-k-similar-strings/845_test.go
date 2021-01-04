package _45_k_similar_strings

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_kSimilarity(t *testing.T) {
	table := []struct {
		A      string
		B      string
		Except int
	}{
		//{"ab", "ba", 1},
		{"abc", "bca", 2},
		{"abac", "baca", 2},
		{"abccaacceecdeea", "bcaacceeccdeaae", 9},
		{"cdebcdeadedaaaebfbcf", "baaddacfedebefdabecc", 12},
	}
	for _, v := range table {
		Convey("kSimilarity: "+v.A, t, func() {
			So(kSimilarity(v.A, v.B), ShouldEqual, v.Except)
		})
	}

}

func Benchmark_kSimilarity(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		actual := kSimilarity("cdebcdeadedaaaebfbcf", "baaddacfedebefdabecc")
		if actual != 12 {
			b.Fatal(actual)
		}
	}
	b.StopTimer()
}

func Benchmark_kSimilarity2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		actual := kSimilarity2("cdebcdeadedaaaebfbcf", "baaddacfedebefdabecc")
		if actual != 12 {
			b.Fatal(actual)
		}
	}
	b.StopTimer()
}
