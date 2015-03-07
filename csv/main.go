package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"
)

const DATA_DIR = "/Users/ujanssen/Dropbox/Apps/MyUploader1842846218191/"

type AllData struct {
	Date       string
	Dates      []string
	Categories string
	Data       string
}

func toString(items []string) string {
	str := "["
	for i, item := range items {
		if i > 0 {
			str += ","
		}
		str += `"` + item + `"`
	}
	return str + "]"
}
func fToString(items []string) string {
	str := "["
	for i, item := range items {
		if i > 0 {
			str += ","
		}
		str += item
	}
	return str + "]"
}

func readFrame() *template.Template {
	pattern := "*.html"
	log.Println("parse templates with pattern ", pattern)
	template, err := template.New("frame").ParseGlob(pattern)
	if err != nil {
		log.Println(err)
	}

	return template
}

func RenderPage(r *http.Request, w http.ResponseWriter, page string) {
	date := r.URL.Query().Get("date")
	if len(date) == 0 {
		date = "2015-03-04"
	}
	err := readFrame().ExecuteTemplate(w, "hst.html", read(date))
	if err != nil {
		log.Println(err)
	}
}

var RequestHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		RenderPage(r, w, "index")
	} else {
		RenderPage(r, w, r.URL.Path[1:len(r.URL.Path)])
	}
})

func main() {
	port := 9090
	//serve static files
	for _, dir := range []string{"/js/", "/css/", "/images/", "/fonts/"} {
		http.Handle(dir, http.StripPrefix(dir, http.FileServer(http.Dir("static/"+dir))))
	}
	http.HandleFunc("/", RequestHandler)
	log.Println("Starting web server on port: ", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}

func read(date string) AllData {
	csvfile, err := os.Open(DATA_DIR + date + "-iperf.csv")
	checkErr(err)

	defer csvfile.Close()

	reader := csv.NewReader(csvfile)

	reader.FieldsPerRecord = -1 // see the Reader struct information below

	rawCSVdata, err := reader.ReadAll()
	checkErr(err)

	c := make([]string, 0)
	d := make([]string, 0)
	const layout = "15:04"

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
		c = append(c, tm.Format(layout))
		d = append(d, fmt.Sprintf("%.1f", bpsf))
	}
	return AllData{
		Date:       date,
		Categories: toString(c),
		Data:       fToString(d),
		Dates:      listDates()}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func listDates() (dates []string) {
	dates = make([]string, 0)
	fileInfos, err := ioutil.ReadDir(DATA_DIR)
	if err != nil {
		log.Println("error:", err)
		return nil
	}
	for _, info := range fileInfos {
		if strings.HasPrefix(info.Name(), "2015") {
			dates = append(dates, info.Name()[0:10])
		}
	}
	return dates
}
