package test

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"

	db "github.com/chen-huicheng/GSWGo/learngorm/database"
	"github.com/chen-huicheng/GSWGo/learngorm/entity"
)

func TestConnectGormDB(t *testing.T) {
	db.ConnectGormDB()
}

func TestGormSave(t *testing.T) {
	nd := db.NewGormDao()
	users := make([]entity.User, 0)
	for i := 0; i < 10; i++ {
		years := i + 2000
		month := rand.Int()%12 + 1
		day := rand.Int()%28 + 1
		user := entity.User{Name: strconv.Itoa(i), Birthday: fmt.Sprintf("%d-%d-%d", years, month, day)}
		user.Name += "_hello"
		users = append(users, user)
	}
	err := nd.GormSave(users)
	if err != nil {
		fmt.Println(err)
	}
}

func TestGormSelect(t *testing.T) {
	nd := db.NewGormDao()
	users, err := nd.GormSelect()
	if err != nil {
		fmt.Println(err)
	}
	for i := range users {
		fmt.Println(users[i])
		users[i].Name += "_test"
	}
	err = nd.GormSave(users)
	if err != nil {
		fmt.Println(err)
	}
}
