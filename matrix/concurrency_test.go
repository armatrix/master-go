package matrix

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// User for test
type User struct {
	Name string
	Age  int
}

func (user User) call() {
	time.Sleep(time.Second * 2)
	fmt.Printf("user: %v, name: %v\n", user.Name, user.Age)
}

func TestConcurrent(t *testing.T) {
	var users []User
	for i := 0; i < 10; i++ {
		user := User{
			Name: fmt.Sprintf("no.%v", i),
			Age:  i,
		}
		users = append(users, user)
	}

	var concurrency = 100

	var newUsers []User
	var lock sync.Mutex
	t.Run("concurrency", func(t *testing.T) {
		Concurrent(concurrency, users, func(user User) {
			lock.Lock()
			user.Age++
			newUsers = append(newUsers, user)
			lock.Unlock()
		})
	})

	var totalAge int
	for _, user := range newUsers {
		totalAge += user.Age
	}
	wantAge := 55
	if totalAge != wantAge {
		t.Errorf("\ntotal age: %v\nwant age: %v\nusers: %v\nnew users: %v", totalAge, wantAge, users, newUsers)
	}

}
