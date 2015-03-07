package main

import (
	//	"database/sql"
	"encoding/csv"
	"fmt"
	"github.com/cznic/ql"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const DATA_DIR = "/Users/ujanssen/Dropbox/Apps/MyUploader1842846218191/"

func main() {
	db, err := ql.OpenMem()
	if err != nil {
		panic(err)
	}
	_, _, err = db.Run(ql.NewRWCtx(), `
        BEGIN TRANSACTION;
        CREATE TABLE data (t time, h uint, bps float);
        COMMIT;
        `)
	if err != nil {
		panic(err)
	}
	for _, date := range listDates() {
		addToDB(date, db)
	}
	rss, _, err := db.Run(ql.NewRWCtx(), "SELECT h, min(bps),avg(bps),max(bps), count(t) FROM data GROUP BY h;")
	if err != nil {
		panic(err)
	}

	for _, rs := range rss {
		if err := rs.Do(false, func(data []interface{}) (bool, error) {
			fmt.Printf("hour %2d, min %02.1f, avg %02.1f, max %02.1f, count %3d\n", data[0], data[1], data[2], data[3], data[4])
			return true, nil
		}); err != nil {
			panic(err)
		}
		fmt.Println("----")
	}
}

func addToDB(date string, db *ql.DB) {
	csvfile, err := os.Open(DATA_DIR + date + "-iperf.csv")
	checkErr(err)

	defer csvfile.Close()

	reader := csv.NewReader(csvfile)

	reader.FieldsPerRecord = -1 // see the Reader struct information below

	rawCSVdata, err := reader.ReadAll()
	checkErr(err)

	// sanity check, display to standard output
	for _, each := range rawCSVdata {
		bps, err := strconv.ParseUint(each[len(each)-1], 10, 64)
		if err != nil {
			fmt.Printf("err parsing bsb: %v in line %v\n", err, each)
			bps = 0
		}
		bpsf := float64(bps) / 1000000.0
		t, err := strconv.ParseInt(each[len(each)-2], 10, 64)
		if err != nil {
			fmt.Printf("err parsing time: %v in line %v\n", err, each)
			continue
		}
		tm := time.Unix(t, 0)
		_, _, err = db.Run(ql.NewRWCtx(), fmt.Sprintf("BEGIN TRANSACTION;INSERT INTO data VALUES(parseTime(\"2006-01-02 15:04:05.999999999 -0700 MST\",\"%v\"), %d, %v);COMMIT;", tm, tm.Hour(), bpsf))
		if err != nil {
			panic(err)
		}
	}
}
