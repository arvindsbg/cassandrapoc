package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gocql/gocql"
	"gopkg.in/inf.v0"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Expected address argument")
	}
	addr := os.Args[1]
	cluster := gocql.NewCluster(addr)
	cluster.Keyspace = "blackjack"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	defer session.Close()

	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/insert", func(w http.ResponseWriter, r *http.Request) {
		insert(session)
	})

	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		delete(session)
	})

	http.HandleFunc("/select", func(w http.ResponseWriter, r *http.Request) {
		selectRow(session)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("error running http: %v", err)
	}

	fmt.Print("service running on port 8080")
}

func insert(session *gocql.Session) {

	n := inf.NewDec(4999, 2)

	insertQuery := session.Query(`INSERT INTO facts (accountid, requestid, tid, spinextid, gameid, amount,ts ) VALUES (?, ?, ?,?,?,?,toTimeStamp(toDate(now())))`,
		"testaccount4", "4123", "4456", "4789", "4111", n)

	fmt.Print(insertQuery)

	if err := insertQuery.Exec(); err != nil {
		log.Fatal(err)
	}
}

func delete(session *gocql.Session) {

	deleteQuery := session.Query(`delete from facts where accountid = ? and requestid = ?`,
		"testaccount4", "4123")

	fmt.Print(deleteQuery)

	if err := deleteQuery.Exec(); err != nil {
		log.Fatal(err)
	}
}

func selectRow(session *gocql.Session) {

	var result map[string]interface{}

	result = make(map[string]interface{})

	if err := session.Query(`SELECT * FROM facts LIMIT 1`).Consistency(gocql.One).MapScan(result); err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
