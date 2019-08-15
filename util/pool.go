package util

func StartPool(size int, functions chan func()) {
	out := make(chan func())
	for i := 0; i < size; i++ {
		go worker(out)
	}
	go func() {
		for f := range functions {
			out <- f
		}
		close(out)
	}()
	return
}

func worker(incoming chan func()) {
	for f := range incoming {
		f()
	}
}
