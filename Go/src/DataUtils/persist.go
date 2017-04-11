// package contains utility functions for manipulating the data we want
package DataUtils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"bufio"
	"os"
	"fmt"
	"log"
)

// persists the data on the server
func PersistData(db *sql.DB, src, dest string) {

	// opening the file for writing, with append flag
	file, err := os.OpenFile(fmt.Sprintf(FILE_LOC, src, dest), os.O_RDONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var line string

	for err == nil{
		fmt.Printf("reading lines and trying to send stuff across the wire\n");
		line, err = reader.ReadString('\n')
		if line != "\n" {
			db.Exec(line)
		}
	}


  // // Prepare statement for reading data
  //     stmtOut, err := db.Prepare("SELECT Price FROM BHX_MAD")
  //     if err != nil {
  //            fmt.Printf("tried to read the data opened the socket and errored");
  //         panic(err.Error()) // proper error handling instead of panic in your app
  //     }
  //     defer stmtOut.Close()
	//
  // rows, err := db.Query("SELECT * FROM BHX_MAD;")
  // if err != nil {
  //         panic(err.Error())
  // }
	//
  // for rows.Next() {
  //         var id int
  //         var dest string
  //         var src string
  //         var first string
  //         var second string
  //         var price int
  //         if err := rows.Scan(&id,&dest,&src,&first,&second,&price); err != nil {
  //                 panic(err.Error())
  //         }
  //         fmt.Printf("%d %s %s %s %s %d\n",id,dest,src,first,second,price)
  // }
	//
  // if err := rows.Err(); err != nil {
  //                 panic(err.Error())
  // }
}
