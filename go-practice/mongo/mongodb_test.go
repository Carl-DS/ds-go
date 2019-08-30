package mongo

import (
	"encoding/json"
	"fmt"
	"github.com/lexkong/log"
	"gopkg.in/mgo.v2/bson"
	"testing"
	"time"
)

const (
	db         = "gs"
	collection = "user"
)

type User struct {
	Id        bson.ObjectId `bson:"_id"`
	Username  string        `bson:"username"`
	Password  string        `bson:"password"`
	CreatedAt time.Time     `bson:"created_at"`
}

type UserVO struct {
	Username  string    `bson:"username"`
	Password  string    `bson:"password"`
	CreatedAt time.Time `bson:"created_at"`
}

// 新增数据
func TestInsert(t *testing.T) {
	user := &User{
		Id:        bson.NewObjectId(),
		Username:  "carl4",
		Password:  "1304",
		CreatedAt: time.Now(),
	}

	err := Insert(db, collection, user)
	if err != nil {
		fmt.Print("插入成功")
	}
}

func TestFindOne(t *testing.T) {
	// find one with all fields
	var result User
	err := FindOne(db, collection, bson.M{"_id": bson.ObjectIdHex("5c629917ffe7911188f470ad")}, nil, &result)
	if err != nil {
		log.Fatalf(err, "查询单个失败")
	}
	resultJson, _ := json.Marshal(result)
	fmt.Printf("find one without id field, %s\n", resultJson)

	// find one without id field
	var result1 UserVO
	err1 := FindOne(db, collection, bson.M{"_id": bson.ObjectIdHex("5c629917ffe7911188f470ad")}, bson.M{"_id": 0}, &result1)
	if err1 != nil {
		log.Fatalf(err1, "查询单个失败")
	}

	resultJson1, err := json.Marshal(result1)

	fmt.Printf("find one with all fields, %s\n", resultJson1)
}
