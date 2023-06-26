// package main

// import (
// 	"sync"
// )

// type State struct {
// 	mu    sync.Mutex
// 	count int
// }

// func (s *State) setState(i int) {

// 	s.mu.Lock()
// 	defer s.mu.Unlock()
// 	s.count = i

// }

// func main() {
// 	// state := &State{}

// 	// for i := 0; i < 10; i++ {
// 	//  state.count = i

// 	// }
// 	// fmt.Println(state)
// }

// ///----------------Atomic values--------
package main

import (
	"sync/atomic"
)

type State struct {
	count int32
}

func (s *State) setState(i int32) {
	atomic.AddInt32(&s.count, int32(i))

}

func main() {

}
