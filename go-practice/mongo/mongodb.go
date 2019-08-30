package mongo

import (
	"gopkg.in/mgo.v2"
	"log"
	"time"
)

var globalSession *mgo.Session

const (
	dbhost    = "127.0.0.1:27017"
	authdb    = "gs"
	authuser  = "gs"
	authpass  = "gs123"
	timeout   = 60 * time.Second
	poollimit = 4096
)

func init() {
	dialInfo := &mgo.DialInfo{
		Addrs:     []string{dbhost}, //数据库地址 dbhost: mongodb://user@123456:127.0.0.1:27017
		Timeout:   timeout,          // 连接超时时间 timeout: 60 * time.Second
		Source:    authdb,           // 设置权限的数据库 authdb: admin
		Username:  authuser,         // 设置的用户名 authuser: user
		Password:  authpass,         // 设置的密码 authpass: 123456
		PoolLimit: poollimit,        // 连接池的数量 poollimit: 100
	}

	s, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Fatalf("Create Session: %s\n", err)
	}
	globalSession = s
}

func connect(db, collection string) (*mgo.Session, *mgo.Collection) {
	// 每次操作都Copy 一份Session, 避免每次创建Session 导致连接数量超过设置的最大值
	session := globalSession.Copy()
	c := session.DB(db).C(collection)
	session.SetMode(mgo.Monotonic, true)
	return session, c
}

func getDb(db string) (*mgo.Session, *mgo.Database) {
	ms := globalSession.Copy()
	return ms, ms.DB(db)
}

func IsEmpty(db, collection string) bool {
	session, c := connect(db, collection)
	defer session.Close()
	count, err := c.Count()
	if err != nil {
		log.Fatal(err)
	}
	return count == 0
}

func Count(db, collection string, query interface{}) (int, error) {
	session, c := connect(db, collection)
	defer session.Close()
	return c.Find(query).Count()
}

func Insert(db, collection string, docs ...interface{}) error {
	session, c := connect(db, collection)
	defer session.Close()
	return c.Insert(docs...)
}

func FindOne(db, collection string, query, selector, result interface{}) error {
	session, c := connect(db, collection)
	defer session.Close()
	return c.Find(query).Select(selector).One(result)
}

func FindAll(db, collection string, query, selector, result interface{}) error {
	session, c := connect(db, collection)
	defer session.Close()
	return c.Find(query).Select(selector).All(result)
}

func FindPage(db, collection string, page, limit int, query, selector, result interface{}) error {
	session, c := connect(db, collection)
	defer session.Close()
	return c.Find(query).Select(selector).Skip(page * limit).Limit(limit).All(result)
}

func FindIter(db, collection string, query interface{}) *mgo.Iter {
	session, c := connect(db, collection)
	defer session.Close()
	return c.Find(query).Iter()
}

func Update(db, collection string, selector, update interface{}) error {
	session, c := connect(db, collection)
	defer session.Close()
	return c.Update(selector, update)
}

func Upsert(db, collection string, selector, update interface{}) error {
	session, c := connect(db, collection)
	defer session.Close()
	_, err := c.Upsert(selector, update)
	return err
}

func UpdateAll(db, collection string, selector, update interface{}) error {
	session, c := connect(db, collection)
	defer session.Close()
	_, err := c.UpdateAll(selector, update)
	return err
}

func Remove(db, collection string, selector interface{}) error {
	session, c := connect(db, collection)
	defer session.Close()
	return c.Remove(selector)
}

func RemoveAll(db, collection string, selector interface{}) error {
	session, c := connect(db, collection)
	defer session.Close()
	_, err := c.RemoveAll(selector)
	return err
}

//insert one or multi documents
func BulkInsert(db, collection string, docs ...interface{}) (*mgo.BulkResult, error) {
	session, c := connect(db, collection)
	defer session.Close()
	bulk := c.Bulk()
	bulk.Insert(docs...)
	return bulk.Run()
}

func BulkRemove(db, collection string, selector ...interface{}) (*mgo.BulkResult, error) {
	session, c := connect(db, collection)
	defer session.Close()
	bulk := c.Bulk()
	bulk.Remove(selector...)
	return bulk.Run()
}

func BulkRemoveAll(db, collection string, selector ...interface{}) (*mgo.BulkResult, error) {
	session, c := connect(db, collection)
	defer session.Close()
	bulk := c.Bulk()
	bulk.RemoveAll(selector...)
	return bulk.Run()
}

func BulkUpsert(db, collection string, pairs ...interface{}) (*mgo.BulkResult, error) {
	session, c := connect(db, collection)
	defer session.Close()
	bulk := c.Bulk()
	bulk.Upsert(pairs...)
	return bulk.Run()
}

func PipeAll(db, collection string, pipeline, result interface{}, allowDiskUse bool) error {
	session, c := connect(db, collection)
	defer session.Close()
	var pipe *mgo.Pipe
	if allowDiskUse {
		pipe = c.Pipe(pipeline).AllowDiskUse()
	} else {
		pipe = c.Pipe(pipeline)
	}
	return pipe.All(result)
}

func PipeOne(db, collection string, pipeline, result interface{}, allowDiskUse bool) error {
	session, c := connect(db, collection)
	defer session.Close()
	var pipe *mgo.Pipe
	if allowDiskUse {
		pipe = c.Pipe(pipeline).AllowDiskUse()
	} else {
		pipe = c.Pipe(pipeline)
	}
	return pipe.One(result)
}

func PipeIter(db, collection string, pipeline interface{}, allowDiskUse bool) *mgo.Iter {
	session, c := connect(db, collection)
	defer session.Close()
	var pipe *mgo.Pipe
	if allowDiskUse {
		pipe = c.Pipe(pipeline).AllowDiskUse()
	} else {
		pipe = c.Pipe(pipeline)
	}
	return pipe.Iter()
}

func Explain(db, collection string, pipeline, result interface{}) error {
	session, c := connect(db, collection)
	defer session.Close()
	pipe := c.Pipe(pipeline)
	return pipe.Explain(result)
}

func GridFSCreate(db, prefix, name string) (*mgo.GridFile, error) {
	session, d := getDb(db)
	defer session.Close()
	gridFs := d.GridFS(prefix)
	return gridFs.Create(name)
}

func GridFSFindOne(db, prefix string, query, result interface{}) error {
	session, d := getDb(db)
	defer session.Close()
	gridFs := d.GridFS(prefix)
	return gridFs.Find(query).One(result)
}

func GridFSFindAll(db, prefix string, query, result interface{}) error {
	session, d := getDb(db)
	defer session.Close()
	gridFs := d.GridFS(prefix)
	return gridFs.Find(query).All(result)
}

func GridFSOpen(db, prefix, name string) (*mgo.GridFile, error) {
	session, d := getDb(db)
	defer session.Close()
	gridFs := d.GridFS(prefix)
	return gridFs.Open(name)
}

func GridFSRemove(db, prefix, name string) error {
	session, d := getDb(db)
	defer session.Close()
	gridFs := d.GridFS(prefix)
	return gridFs.Remove(name)
}
