package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/chen-huicheng/GSWGo/learngorm/entity"
	"github.com/go-sql-driver/mysql"
)

type Dao struct {
	DB *sql.DB
}

var db *sql.DB

func NewDao() *Dao {
	return &Dao{db}
}

func ConnectDB(addr, dbname, user, passwd string) *sql.DB {
	cfg := mysql.Config{
		User:      user,
		Passwd:    passwd,
		Net:       "tcp",
		Addr:      addr,
		DBName:    dbname,
		ParseTime: true,
	}
	// Get a database handle.
	var err error
	fmt.Println(cfg.FormatDSN())
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	return db
}

type User struct {
	Id         int64     `gorm:"column:id;primary_key" json:"id"`
	AddTime    time.Time `gorm:"column:add_time" json:"add_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	Name       string    `gorm:"column:name" json:"name"`
	Birthday   string    `gorm:"column:birthday" json:"birthday"`
}

// select query all User
func (dao *Dao) Select(users *[]entity.User) error {
	res := make([]entity.User, 0)
	rows, err := dao.DB.Query("select * from user")
	if err != nil {
		return fmt.Errorf("Select %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var user entity.User
		// var at, ut string
		if err := rows.Scan(&user.Id, &user.Name, &user.Birthday, &user.AddTime, &user.UpdateTime); err != nil {
			return fmt.Errorf("Select %v", err)
		}
		// fmt.Println(user, at, ut)
		res = append(res, user)
	}
	*users = res
	if err := rows.Err(); err != nil {
		return fmt.Errorf("Select %v", err)
	}
	return nil
}

func (dao *Dao) Save(users []entity.User) error {

	// dao.DB.ExecContext()
	// dao.DB.PrepareContext()
	// dao.DB.QueryContext()
	// dao.DB.QueryRowContext()
	for _, user := range users {
		sql := fmt.Sprintf("insert into user(name,birthday) VALUES('%s','%s')", user.Name, user.Birthday)
		_, err := dao.DB.Exec(sql)
		if err != nil {
			return fmt.Errorf("Save %v", err)
		}
	}
	return nil
}

func RunSql() {
	ConnectDB("192.168.1.6:3306", "blog", "user", "123456")
	// nd := NewDao()
	// users := make([]entity.User, 0)
	// for i := 0; i < 10; i++ {
	// 	years := i + 2000
	// 	month := rand.Int()%12 + 1
	// 	day := rand.Int()%28 + 1
	// 	user := entity.User{Name: strconv.Itoa(i), Birthday: fmt.Sprintf("%d-%d-%d", years, month, day)}
	// 	users = append(users, user)
	// }
	// err := nd.Save(users)
	// if err != nil {
	// 	fmt.Println(err)
	// }
}
