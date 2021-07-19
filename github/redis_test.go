package github

import (
	"encoding/base64"
	"log"
	"testing"
	"time"

	//"github.com/garyburd/redigo/redis"
	//"github.com/mediocregopher/radix.v2/redis"
	"github.com/go-redis/redis"
)

// test common
/*func TestOnOperate(t *testing.T) {
	// connect redis server
	c, err := redis.Dial("tcp", "192.168.1.174:6379")
	if err != nil {
		fmt.Println(err)
		log.Fatal("connect redis-server error: ", err)
		return
	}
	defer c.Close()

	// SET
	//v, err := c.Cmd("SET", "test_name", "red")
	resp := c.Cmd("SET", "test_333", "fdsafdsafdsa")
	if resp.Err != nil {
		fmt.Println(err)
		return
	}

	// GET
	//v, err = redis.String(c.Do("GET", "test_name"))
	resp = c.Cmd("GET", "test_333")
	if resp.Err != nil {
		fmt.Println(err)
		return
	}
	b, err := resp.Bytes()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}*/

func TestStandardRedis(t *testing.T) {
	// connect redis server
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.5.216:6379",
		Password: "123456",
		DB:       0,
	})

	data := `AAAAAA0AAAAAAAAAAQAAAAAJIDkBAAAAAgAAAAAAAAACAAAAAwAAAAAAAAADAAAABAAAAAAAAAAEAAAABQAAAAAAAAAFAAAABgAAAAB/AAAGAAAABwAAAAAhvzkHAAAACAAAAAB/AAAIAAAACQAAAAAAAAAJAAAACgAAAAB/AAAKAAAACwAAAABQvzkLAAAADAAAAAAAAAAMAAAADAAAAAEAAAAMAAAA5AcgOQAAAAABAAAAAAAAALoAAAAAAAAAAgAAAAAAAACMNxgCAAAAAAMAAAAAAAAA5TK/OQAAAAAEAAAAAAAAAI0xvzkAAAAABQAAAAAAAACBAQAAAAAAAAYAAAAAAAAA5QAAAAAAAAAHAAAAAAAAAIX///8AAAAACAAAAAAAAACtIE1pAAAAAAkAAAAAAAAA5QpUIAAAAAAKAAAAAAAAALogMiAAAAAACwAAAAAAAACmIDQgAQAAAAwAAAAAAAAAAgAAAA4AAAAA5LqM5Y2B5YWt5bqmAA==`
	dec, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		log.Printf("decode base64 that user input rules err:%v\n", err)
		return
	}

	client.Set("post-user-rules", dec, time.Hour*1000)
	client.HMSet("map-post-user-rules", map[string]interface{}{
		"data": dec,
	})

	byt, _ := client.Get("post-user-rules").Bytes()
	byt1, _ := client.HGet("map-post-user-rules", "data").Bytes()
	t.Log(string(byt))
	t.Log(string(byt1))
}
