package must

func Must[A any](a A, err error) A {
	if err != nil {
		panic(err)
	}
	return a
}

func Must2[A, B any](a A, b B, err error) (A, B) {
	if err != nil {
		panic(err)
	}
	return a, b
}

func Must3[A, B, C any](a A, b B, c C, err error) (A, B, C) {
	if err != nil {
		panic(err)
	}
	return a, b, c
}
