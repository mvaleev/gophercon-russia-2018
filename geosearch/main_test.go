package geosearch

import (
	"fmt"
	"testing"
)

func prepare() *Index {
	i := NewIndex(20 /* ~ 1km */)
	i.AddUser(1, 37.604178, 55.608279)
	i.AddUser(2, 37.612375, 55.611211)
	i.AddUser(3, 37.581741, 55.602520)
	return i
}

func TestSearch(t *testing.T) {
	indx := prepare()

	found, _ := indx.Search(37.608821, 55.609525, 315)
	fmt.Println("found:", found)
	if len(found) != 2 {
		t.Fatal("error while searching with radius 315")
	}

	found, _ = indx.Search(37.608821, 55.609525, 1870)
	fmt.Println("found:", found)
	if len(found) != 3 {
		t.Fatal("error while searching with radius 1870")
	}
}

func TestSearchFaster(t *testing.T) {
	indx := prepare()

	found, _ := indx.SearchFaster(37.608821, 55.609525, 210)
	if len(found) != 2 {
		t.Fatal("error while searching with radius 210")
	}

	found, _ = indx.SearchFaster(37.608821, 55.609525, 1000)
	if len(found) != 3 {
		t.Fatal("error while searching with radius 1000")
	}
}

var res []uint32

func BenchmarkSearch(b *testing.B) {
	indx := prepare()

	for i := 0; i < b.N; i++ {
		res, _ = indx.Search(14.1313, 14.1313, 50000)
	}
}

func BenchmarkSearchFaster(b *testing.B) {
	indx := prepare()

	for i := 0; i < b.N; i++ {
		res, _ = indx.SearchFaster(14.1313, 14.1313, 50000)
	}
}
