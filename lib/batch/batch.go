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
	for i := int64(0); i < n; i++ {
		res = append(res, getOne(i))
	}
	return res
}
