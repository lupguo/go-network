package mutext

import (
	"sync"
)

func SimulatedDec(count int, lock bool) int {
	m := new(sync.Mutex)
	wg := new(sync.WaitGroup)
	wg.Add(count)
	limit := count
	for i := 0; i < limit; i++ {
		go func(i int) {
			if lock {
				m.Lock()
				defer m.Unlock()
			}
			count--
			//decrease(&count, m, lock)
			wg.Done()
		}(i)
	}
	wg.Wait()
	//fmt.Scanln(new(string))
	//fmt.Println("count=", count)
	return count
}

func decrease(num *int, m *sync.Mutex, lock bool) {
	if lock {
		m.Lock()
		defer m.Unlock()
	}
	*num--
}
