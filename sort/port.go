package csort

type SortableSource interface {
	At(i int) float64
	Dim() int
}
