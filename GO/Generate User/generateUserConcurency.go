package Generate_User

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

type LogItem struct {
	act string
	t   time.Time
}

var action = []string{
	"loged",
	"outlog",
	"cleaning history",
	"exit",
}

type User struct {
	id    int
	email string
	log   []LogItem
}

func main() {
	wg := sync.WaitGroup{}
	rand.NewSource(time.Now().UnixNano())

	createNewUser := GenerateUsers(rand.Intn(100))

	for _, user := range createNewUser {
		wg.Add(1)
		CreateWrite(user, &wg)
	}

	wg.Wait()
}

func (u User) GetInfo() string {
	out := fmt.Sprintf("id:%d | email:%s\n", u.id, u.email)
	for _, v := range u.log {
		out += fmt.Sprintf("action:%s | time:%s\n", v.act, v.t)
	}

	return out
}

func GenerateUsers(n int) []User {
	u := make([]User, n)

	for i := 1; i < n; i++ {
		u[i] = User{
			id:    i,
			email: fmt.Sprintf("emailUser%d@mail.ru", i),
			log:   RandomLog(rand.Intn(10)),
		}
	}
	return u
}

func RandomLog(count int) []LogItem {
	logy := make([]LogItem, count)

	for i := 1; i < count; i++ {
		logy[i] = LogItem{
			t:   time.Now(),
			act: action[rand.Intn(len(action)-1)],
		}
	}

	return logy
}

func CreateWrite(user User, group *sync.WaitGroup) error {
	filename := fmt.Sprintf("GO/Generate User/uid_%d", user.id)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	file.WriteString(user.GetInfo())

	group.Done()
	return nil
}
