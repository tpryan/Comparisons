/*
   Copyright 2015, Google, Inc.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/
package main

import (
	"database/sql"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"text/template"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	sqldriver = "mysql"
	workers   = 8
)

var (
	db    *sql.DB
	query string
	count = flag.Int("count", 1, "The number of times to create all of the text")
)

type Entry struct {
	Title   string
	Excerpt string
	Name    string
	GUID    string
	Date    time.Time
	Content string
	DateF   string
}

func main() {

	flag.Parse()
	var err error
	_ = runtime.GOMAXPROCS(runtime.NumCPU())

	t := template.New("template.html")
	t = t.Funcs(template.FuncMap{
		"fdate":     fdate,
		"repairurl": repairURL,
	})

	t, err = t.ParseFiles("textout/go/template.html")

	if err != nil {
		log.Fatalf("Could not parse template: %v", err)
	}

	SQLhost := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") + "@tcp(" + os.Getenv("DB_HOST") + ":3306)/" + os.Getenv("DB_NAME")

	db, err = sql.Open(sqldriver, SQLhost+"?parseTime=true")
	if err != nil {
		log.Fatalf("Could not open connection to database: %v", err)
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	outdir := wd + "/textout/output/go/"

	if err = cleanDir(outdir); err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadFile(wd + "/textout/sql/entries.sql")
	if err != nil {
		log.Fatalf("Could not open SQL file: %v", err)
	}

	query = string(b)

	entries, err := entries()
	if err != nil {
		log.Fatalf("Could not get entries from database: %v", err)
	}

	if err = writePar(entries, outdir, *count, *t); err != nil {
		log.Fatalf("Could write entry hmtl to disk: %v", err)
	}

}

func writePar(entries []Entry, outdir string, count int, t template.Template) error {
	var wg sync.WaitGroup
	wg.Add(count)

	for i := 1; i <= count; i++ {

		go func(i int) {
			defer wg.Done()
			err := writeEntries(entries, outdir+strconv.Itoa(i), t)
			if err != nil {
				log.Fatal(err)
			}
		}(i)

	}
	wg.Wait()

	return nil
}

func cleanDir(dir string) error {

	err := os.RemoveAll(dir)
	if err != nil {
		return err
	}

	return os.Mkdir(dir, 0777)
}

func entries() ([]Entry, error) {
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []Entry
	for rows.Next() {
		e := Entry{}
		if err := rows.Scan(&e.Title, &e.Excerpt, &e.Name, &e.GUID, &e.Date, &e.Content, &e.DateF); err != nil {
			return nil, err
		}

		res = append(res, e)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func writeEntries(entries []Entry, path string, t template.Template) error {

	if err := os.Mkdir(path, 0777); err != nil {
		return err
	}

	for _, entry := range entries {

		p := path + "/" + entry.Name + ".html"
		f, err := os.Create(p)

		if err != nil {
			return err
		}
		defer f.Close()

		err = t.Execute(f, entry)
		if err != nil {
			return err
		}
		f.Close()
	}
	return nil

}

func repairURL(URL string) string {
	out := strings.Replace(URL, "blog//blog/index.php/", "", -1)
	out = strings.Replace(out, "http://http://", "http://", -1)
	return out
}

func fdate(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
