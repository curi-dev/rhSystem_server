package helpers

func ExecuteServiceOnRoutine(f func(c chan bool)) chan bool {
	c := make(chan bool)

	go func() {
		f(c)
	}()

	return c
}
