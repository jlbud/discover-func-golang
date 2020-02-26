package github

import (
	"fmt"
	"log"
	"testing"

	//"github.com/garyburd/redigo/redis"
	"github.com/mediocregopher/radix.v2/redis"
)

// test common
func TestOnOperate(t *testing.T) {
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
}


