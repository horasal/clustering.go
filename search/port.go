package search

type IBinarySearchableData interface {
	Equal(t IBinarySearchableData) int
}

type IBinarySearch interface {
	Len() int
	Add(t IBinarySearchableData)
	Search(t IBinarySearchableData) int
	Value(i int) IBinarySearchableData
}
