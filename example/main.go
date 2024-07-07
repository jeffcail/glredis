package main

import (
	"context"
	"fmt"
	"github.com/jeffcail/glredis"
	"github.com/redis/go-redis/v9"
	"log"
)

var (
	rdb *redis.Client
	ctx = context.Background()
)

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:7379",
		Password: "123456",
		DB:       1,
	})
	if err := rdb.Ping(ctx).Err(); err != nil {
		fmt.Printf("err: %v", err)
		panic(err)
	}
}

func main() {
	publish, err := glredis.NewPubSub().Publish(rdb, []string{"go"}, "123456")
	if err != nil {
		log.Fatal(fmt.Sprintf("err: %v\n", err))
	}

	fmt.Println(fmt.Sprintf("publish: %d\n", publish))

	//glredis.NewGLRSortedSet().SetZAdd(rdb, []string{"tt"}, 1, "abc")
	//zRand, err := glredis.NewGLRSortedSet().SetZRange(rdb, []string{"myzset"}, 0, 10)
	//if err != nil {
	//	log.Fatal(fmt.Sprintf("err: %v\n", err))
	//}
	//
	//fmt.Println(fmt.Sprintf("zRand: %v\n", zRand))

	//res := glredis.NewGLRSet().SetSAdd(rdb, []string{"score"}, "90")
	//_ = glredis.NewGLRSet().SetSAdd(rdb, []string{"score"}, "100")
	//
	//fmt.Printf("sadd set key success: %v\n", res)
	//
	//res, _ = glredis.NewGLRSet().SMembers(rdb, []string{"score"})
	//
	//fmt.Printf("sadd set key success: %v\n", res)

	//res1 := glredis.NewGLRList().ListLen(rdb, []string{"zk"})
	//
	//fmt.Printf("get list len success: %v\n", res1)
	//
	//res2 := glredis.NewGLRList().ListPop(rdb, []string{"zk"})
	//fmt.Printf("pop element success, element: %v\n", res2)
	//
	//res1 = glredis.NewGLRList().ListLen(rdb, []string{"zk"})
	//
	//fmt.Printf("get list len success: %v\n", res1)

	//res1 := glredis.NewGLRList().ListLen(rdb, []string{"jj"})
	//
	//fmt.Printf("get list len success: %v\n", res1)

	//res := glredis.NewGLRList().ListPush(rdb, []string{"zk"}, "golang")
	////if err != nil {
	////	log.Fatalf("err: %v", err)
	////}
	//fmt.Printf("push lish success: %v\n", res)
	//
	//res1 := glredis.NewGLRList().ListLen(rdb, []string{"jj"})
	//
	//fmt.Printf("get list len success: %v\n", res1)

	//err := glredis.NewGLRHash().SetHashKey(rdb, []string{"book"}, "golang", "abc")
	//if err != nil {
	//	log.Fatalf("err: %v", err)
	//}
	//fmt.Printf("set hash key success\n")
	//
	//ret, err := glredis.NewGLRHash().GetHashFiled(rdb, []string{"book"}, "golang")
	//if err != nil {
	//	log.Fatalf("err: %v", err)
	//}
	//fmt.Printf("get hash field success; %v\n", ret)

	//err = glredis.NewGLRHash().DelHashField(rdb, []string{"book"}, "golang")
	//if err != nil {
	//	log.Fatalf("err: %v", err)
	//}
	//fmt.Printf("del hash field success\n")

	//res, err := glredis.NewGLRedis().SetWithExpire(rdb, []string{"name1"}, "golang", 60)
	//if err != nil {
	//	log.Fatalf("err: %v", err)
	//}
	//fmt.Printf("res: %s\n", res)
	//
	//res1, err := glredis.NewGLRedis().PtlKey(rdb, []string{"name1"})
	//if err != nil {
	//	log.Fatalf("err: %v", err)
	//}
	//fmt.Printf("res1: %d\n", res1)
	//
	//res2, err := glredis.NewGLRedis().MoveKeyToDb(rdb, []string{"name1"}, 2)
	//if err != nil {
	//	log.Fatalf("err: %v", err)
	//}
	//fmt.Printf("res2: %d\n", res2)
	//
	//res, err = glredis.NewGLRedis().SetWithExpire(rdb, []string{"name2"}, "golang", 380)
	//if err != nil {
	//	log.Fatalf("err: %v", err)
	//}
	//fmt.Printf("res: %s\n", res)
	//
	//res, err = glredis.NewGLRedis().SetWithExpire(rdb, []string{"name3"}, "golang", 380)
	//if err != nil {
	//	log.Fatalf("err: %v", err)
	//}
	//fmt.Printf("res: %s\n", res)
	//
	//keys, err := glredis.NewGLRedis().PatternKeys(rdb, []string{"name*"})
	//if err != nil {
	//	log.Fatalf("err: %v", err)
	//}
	//fmt.Printf("keys: %v\n", keys)
	//
	//ok, err := glredis.NewGLRedis().ExpireKey(rdb, []string{"name"}, 60)
	//if err != nil {
	//	log.Fatalf("err: %v", err)
	//}
	//fmt.Printf("ok: %d\n", ok)
	//
	//has, err := glredis.NewGLRedis().ExistsKey(rdb, []string{"name"})
	//if err != nil {
	//	log.Fatalf("err: %v", err)
	//}
	//fmt.Printf("has: %d\n", has)
	//
	//row, err := glredis.NewGLRedis().DelKey(rdb, []string{"name"})
	//if err != nil {
	//	log.Fatalf("err: %v", err)
	//}
	//fmt.Printf("ok: %v\n", row)
	//
	//res, err = glredis.NewGLRedis().SetWithoutExpire(rdb, []string{"score"}, "80")
	//if err != nil {
	//	log.Fatalf("err: %v", err)
	//}
	//fmt.Printf("res: %s\n", res)
	//
	//res, err = glredis.NewGLRedis().GetKey(rdb, []string{"name"})
	//if err != nil && errors.Is(err, redis.Nil) {
	//	log.Fatalf("err: %v", "key不存在")
	//} else {
	//	log.Fatalf("err: %v", err)
	//}
	//fmt.Printf("res: %s\n", res)
	//
	//res, err = glredis.NewGLRedis().GetOrSetWithExpire(rdb, []string{"name"}, "golang", 60)
	//if err != nil {
	//	log.Fatalf("err: %v", err)
	//}
	//fmt.Printf("res: %s\n", res)
	//
	//res, err = glredis.NewGLRedis().GetOrSetWithoutExpire(rdb, []string{"age"}, "20")
	//if err != nil {
	//	log.Fatalf("err: %v", err)
	//}
	//fmt.Printf("res: %s\n", res)
}
