package must

func Do(err error) {
	if err != nil {
		panic(err)
	}
}
