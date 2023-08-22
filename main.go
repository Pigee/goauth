package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
        "net/http"
        "encoding/json"
 	"crypto/sha512"

)

// 认证数据结构
type Authstr struct {
	Id  string
	Sha string
}

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "123456", // no password set
	DB:       0,        // use default DB
})

func main() {
	at := Authstr{Id: "xx",
		Sha: "fdfdxxfdfdf",
	}
	setHash(at)
	getHash(at)
        fmt.Printf("%x",sha512.Sum512([]byte("fdfdfd")))
        //fmt.Println(string(sha512.Sum512([]byte("fdfdfd"))))
	fmt.Println("\n")
        // WebServerBase()
}

func WebServerBase() {
	fmt.Println("This is webserver base!")
	http.HandleFunc("/authcust", authCust)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("ListenAndServe error: ", err.Error())
	}
}

func authCust(w http.ResponseWriter, req *http.Request) {
        // req.ParseForm()
	// param_Id, found1 := req.Form["Id"]
	// param_Sha, found2 := req.Form["Sha"]
        var auth Authstr
	if err := json.NewDecoder(req.Body).Decode(&auth); err != nil {
		req.Body.Close()
		fmt.Fprint(w, "请勿非法访问")
		return
	}

}
func setHash(a Authstr) {
	err := rdb.Set(ctx, a.Id, a.Sha, 0).Err()
	if err != nil {
		panic(err)
	}
}

func getHash(a Authstr) (v string) {
	val2, err := rdb.Get(ctx, a.Id).Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist,return random sha code")
               // return sha512.Sum512([]byte("Keys does not exist"))
                return "fdfdfd"

	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
         return val2
}

// Output: key value
// key2 does not exist
