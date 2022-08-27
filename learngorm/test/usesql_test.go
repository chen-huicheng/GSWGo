package test

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"

	db "github.com/chen-huicheng/GSWGo/learngorm/database"
	"github.com/chen-huicheng/GSWGo/learngorm/entity"
)

func TestSave(t *testing.T) {
	db.ConnectDB("192.168.1.6:3306", "blog", "user", "123456")
	nd := db.NewDao()
	users := make([]entity.User, 0)
	for i := 0; i < 10; i++ {
		years := i + 2000
		month := rand.Int()%12 + 1
		day := rand.Int()%28 + 1
		user := entity.User{Name: strconv.Itoa(i), Birthday: fmt.Sprintf("%d-%d-%d", years, month, day)}
		users = append(users, user)
	}
	err := nd.Save(users)
	if err != nil {
		fmt.Println(err)
	}
}

func TestSelect(t *testing.T) {
	db.ConnectDB("192.168.1.6:3306", "blog", "user", "123456")
	nd := db.NewDao()
	users := make([]entity.User, 0)
	err := nd.Select(&users)
	if err != nil {
		fmt.Println(err)
	}
	for _, user := range users {
		fmt.Println(user)
	}
}
