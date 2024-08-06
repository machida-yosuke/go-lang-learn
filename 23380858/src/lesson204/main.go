package main

import "fmt"

func generator(done <-chan interface{}) <-chan interface{} {
	intChan := make(chan interface{})
	go func() {
		defer fmt.Println("close intChan")
		defer close(intChan)

		for i := 0; i < 100; i++ {
			select {
			case <-done:
				return
			case intChan <- i:
			}
		}
	}()
	return intChan
}

func orDone(done <-chan interface{}, c <-chan interface{}) <-chan interface{} {
	valChan := make(chan interface{})
	go func() {
		defer close(valChan)

		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if !ok {
					return
				}
				select {
				case valChan <- v:
				case <-done:
				}
			}
		}
	}()
	return valChan
}

func tee(done, c <-chan interface{}) (<-chan interface{}, <-chan interface{}) {
	out1 := make(chan interface{})
	out2 := make(chan interface{})

	go func() {
		defer close(out1)
		defer close(out2)

		for v := range orDone(done, c) {
			var Out1, Out2 = out1, out2
			for i := 0; i < 2; i++ {
				select {
				case Out1 <- v:
					Out1 = nil
				case Out2 <- v:
					Out2 = nil
				}
			}
		}
	}()
	return out1, out2
}

func main() {
	done := make(chan interface{})
	defer close(done)

	out1, out2 := tee(done, generator(done))

	for v1 := range out1 {
		fmt.Println(v1, <-out2)
	}
	// for v2 := range out2 {
	// 	fmt.Printf("out2: %v\n", v2)
	// }
}
