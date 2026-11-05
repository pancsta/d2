package loops

import "github.com/pancsta/d2/lib/geo"

func nonNilEquals(first, second *geo.Point) bool {
	return first != nil && second != nil && *first == *second
}
