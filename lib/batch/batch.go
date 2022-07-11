package batch

import (
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	ch := make(chan user)
	sem := make(chan struct{}, pool)

	for i := int64(0); i < n; i++ {
		go func(i int64) {
			sem <- struct{}{}
			ch <- getOne(i)
			<-sem
		}(i)
	}

	for i := int64(0); i < n; i++ {
		user := <-ch
		res = append(res, user)
	}

	return res
}
