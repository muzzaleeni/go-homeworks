func merge(channels ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(channels))
	for _, c := range channels {
		go func(c <-chan int) {
			defer wg.Done()
			for n := range c {
				out <- n
			}
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}