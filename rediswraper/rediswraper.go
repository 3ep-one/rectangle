package rediswraper

import (
	"log"
	"strconv"

	"github.com/go-redis/redis/v7"
)

//Makeredisclient set up new redis connection
func Makeredisclient() (client *redis.Client) {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, errID := client.Get("id").Result()
	if errID == redis.Nil {
		err := client.Set("id", "0", 0).Err()
		if err != nil {
			panic(err)
		}
	} else if errID != nil {
		panic(errID)
	}

	_, err := client.Ping().Result()
	if err != nil {
		log.Println(err)
	}
	return client
}

//Setkeyvalue set key-value in redis
func Setkeyvalue(client *redis.Client, rectangleList []string) {

	keyID, err := client.Get("id").Result()
	if err != nil {
		panic(err)
	}
	keyIDint, errSet := strconv.Atoi(keyID)
	if errSet != nil {
		panic(errSet)
	}
	for _, rectangle := range rectangleList {
		errSet = client.Set(strconv.Itoa(keyIDint), rectangle, 0).Err()
		if errSet != nil {
			panic(errSet)
		}
		keyIDint++
	}
	errSet = client.Set("id", strconv.Itoa(keyIDint), 0).Err()
	if err != nil {
		panic(errSet)
	}

}

//Getkeyvalue get key-value from redis
func Getkeyvalue(client *redis.Client) (rectangleList []string) {
	keyID, err := client.Get("id").Result()
	if err != nil {
		panic(err)
	}
	keyIDint, errSet := strconv.Atoi(keyID)
	if errSet != nil {
		panic(errSet)
	}
	for id := 0; id < keyIDint; id++ {
		rectangle, err := client.Get(strconv.Itoa(id)).Result()
		if err != nil {
			panic(err)
		}
		rectangleList = append(rectangleList, rectangle)
	}
	return rectangleList

}

//Closeredisclient close redis connection
func Closeredisclient(client *redis.Client) {
	client.Close()
}
