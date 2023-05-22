package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func ConnectRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	ctx := context.Background()

	//ping the redis server to test the conection
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println("failed to connect to redis:", err)
		return
	}

	fmt.Println("connected to Redis", pong)

	//logic
	// HSET command to categorize data
	err = client.HSet(context.Background(), "categories", "fruit", "apple").Err()
	if err != nil {
		fmt.Println("Failed to execute HSET:", err)
		return
	}

	// Retrieve data from categories
	result, err := client.HGet(context.Background(), "categories", "fruit").Result()
	if err != nil {
		fmt.Println("Failed to execute HGET:", err)
		return
	}

	fmt.Println("Category 'fruit':", result)
	// Add some members to the set

	err = client.Close()
	if err != nil {
		fmt.Println("failed to close redis connection", err)
	}

}

//logic test1

// client.SAdd(ctx, "SAddKey", "data1", "data2", "data3", "data4", "data5", "data6", "data7", "data8", "data9", "data10", "data11", "data12")

// result, cursor, err := client.SScan(ctx, "SAddKey", 0, "data3", 10).Result()
// if err != nil {
// 	fmt.Println(err)
// 	return
// }
// fmt.Println(result)
// fmt.Println(cursor)

// client.SAdd(ctx, "SAddKey2", "data1", "data2")

/////////////////////////////

// members := []string{"member1", "member2", "member3", "member4", "member5"}
// for _, member := range members {
// 	err := client.SAdd(context.Background(), "myset", member).Err()
// 	if err != nil {
// 		fmt.Println("Failed to add member to set:", err)
// 		return
// 	}
// }

// // SSCAN command example
// var cursor uint64 = 0
// var keys []string

// for {
// 	var err error
// 	var result []string

// 	// Execute SSCAN command
// 	result, cursor, err = client.SScan(context.Background(), "myset", cursor, "", 3).Result()
// 	if err != nil {
// 		fmt.Println("Failed to execute SSCAN:", err)
// 		return
// 	}
// 	fmt.Println(cursor)
// 	keys = append(keys, result...)

// 	// Break the loop if the cursor is 0, indicating iteration is finished
// 	if cursor == 0 {
// 		break
// 	}
// }

// fmt.Println("Set Members:", keys)
