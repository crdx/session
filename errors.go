package session

func must[T any](value T, err error) T {
	must0(err)
	return value
}

func must0(err error) {
	if err != nil {
		panic(err)
	}
}
