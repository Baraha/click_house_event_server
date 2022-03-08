package main

import (
	"log"

	"github.com/Baraha/clickHouse_event_server.git/pkg/db"
)

func main() {
	client, err := db.NewSQLRepository("tcp://127.0.0.1:9000?debug=true")
	if err != nil {
		log.Fatal(err)
	}
	client.Find("name", "users")
	// 	r := router.New()
	// 	r.POST("/load", api.LoadFile)
	// 	fmt.Println("server is start!")
	// 	log.Fatal(fasthttp.ListenAndServe(":8081", r.Handler))

}
