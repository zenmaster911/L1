package main

import (
	"cmp"
	"fmt"

	"sync"
)

func main() {
	var wg sync.WaitGroup
	safemap1 := new(sync.Map)
	var mt sync.Mutex

	map1 := make(map[string]string)
	map2 := make(map[int]int)

	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(i int) {
			mt.Lock()
			map2[3] = i
			mt.Unlock()
		}(i)
	}

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			SafeWriteToMap(map1, &mt)
		}(i)
		go func(i int) {
			defer wg.Done()
			safemap1.Store(i, struct{}{})

			mt.Lock()
			map2[2] = i
			mt.Unlock()
		}(i)
	}
	wg.Wait()
	fmt.Printf("safemap: %v\nmap1: %v\nmap2: %v", safemap1, map1, map2)
}

func SafeWriteToMap[T cmp.Ordered, V any](mp map[T]V, mt *sync.Mutex) {

	mt.Lock()
	var key T
	var value V
	fmt.Scan(&key, &value)
	mp[key] = value
	mt.Unlock()
}
