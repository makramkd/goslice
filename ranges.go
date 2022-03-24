package goslice

// Zip zips the two given slices into a single slice.
func Zip[T any, U any](r1 []T, r2 []U) (r []Pair[T, U]) {
	r = make([]Pair[T, U], len(r1))
	for i := range r1 {
		r[i] = MakePair(r1[i], r2[i])
	}
	return
}

func ZipLongest[T any, U any](r1 []T, r2 []U) (r []Pair[T, U]) {
	r = make([]Pair[T, U], Max(len(r1), len(r2)))
	if len(r1) >= len(r2) {
		for i := range r1 {
			if i >= len(r2) {
				r[i] = Pair[T, U]{
					First: r1[i],
				}
			} else {
				r[i] = MakePair(r1[i], r2[i])
			}
		}
	} else {
		for i := range r2 {
			if i >= len(r1) {
				r[i] = Pair[T, U]{
					Second: r2[i],
				}
			} else {
				r[i] = MakePair(r1[i], r2[i])
			}
		}
	}
	return
}
