package main

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/otium/queue"
)
import (
	_ "github.com/go-sql-driver/mysql"
)

const (
	sqldriver = "mysql"
	workers   = 8
)

var (
	db       *sql.DB
	entrySQL string
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
	var err error
	_ = runtime.GOMAXPROCS(runtime.NumCPU())

	method := os.Args[2]
	loopcount, err := strconv.Atoi(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	SQLhost := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") + "@tcp(" + os.Getenv("DB_HOST") + ":3306)/" + os.Getenv("DB_NAME")

	db, err = sql.Open(sqldriver, SQLhost+"?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	outdir := wd + "/textout/output/go/"

	err = cleanDir(outdir)
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadFile(wd + "/textout/sql/entries.sql")
	if err != nil {
		log.Fatal(err)
	}

	entrySQL = string(b)

	entries, err := getEntries()
	if err != nil {
		log.Fatal(err)
	}

	if method == "p" {
		if err = writePar(entries, outdir, loopcount); err != nil {
			log.Fatal(err)
		}
	} else if method == "q" {
		if err = writeQueue(entries, outdir, loopcount); err != nil {
			log.Fatal(err)
		}
	} else {
		if err = writeSeq(entries, outdir, loopcount); err != nil {
			log.Fatal(err)
		}
	}

}

func writeQueue(entries []Entry, outdir string, count int) error {
	defer un(trace("writeQueue\t\t"))
	q := queue.NewQueue(func(val interface{}) {

	}, 20)

	for i := 0; i < count; i++ {
		q.Push(writeEntries(entries, outdir+strconv.Itoa(i)))
	}
	q.Wait()

	return nil
}

func writePar(entries []Entry, outdir string, count int) error {
	defer un(trace("writePar\t\t"))
	var wg sync.WaitGroup
	wg.Add(count)

	for i := 1; i <= count; i++ {

		go func(entries []Entry, i int) {
			defer wg.Done()
			err := writeEntries(entries, outdir+strconv.Itoa(i))
			if err != nil {
				log.Fatal(err)
			}
		}(entries, i)

	}
	wg.Wait()

	return nil
}

func writeSeq(entries []Entry, outdir string, count int) error {
	defer un(trace("writeSeq\t\t"))
	for i := 1; i <= count; i++ {
		err := writeEntries(entries, outdir+strconv.Itoa(i))
		if err != nil {
			return err
		}
	}
	return nil
}

func cleanDir(dir string) error {

	err := os.RemoveAll(dir)
	if err != nil {
		return err
	}

	err = os.Mkdir(dir, 0777)

	return err
}

func getEntries() ([]Entry, error) {
	defer un(trace("getEntries\t\t"))
	rows, err := db.Query(entrySQL)
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

func writeEntries(entries []Entry, path string) error {

	if err := os.Mkdir(path, 0777); err != nil {
		return err
	}

	for _, entry := range entries {
		item := `<article>` + "\n" +
			`	<h1><a href="` + repairURL(entry.GUID) + `">` + entry.Title + `</a></h1>` + "\n" +
			`	<time datetime="` + entry.Date.Format("2006-01-02 15:04:05") + `">` + entry.DateF + `</time>` + "\n" +
			`	<div>` + "\n" +
			entry.Content + "\n" +
			`	</div>` + "\n" +
			`</article>` + "\n"

		f := path + "/" + entry.Name + ".html"

		if err := ioutil.WriteFile(f, []byte(item), 0777); err != nil {
			return err
		}
	}
	return nil

}

func repairURL(URL string) string {
	out := strings.Replace(URL, "blog//blog/index.php/", "", -1)
	out = strings.Replace(out, "http://http://", "http://", -1)
	return out

}

func trace(s string) (string, time.Time) {
	return s, time.Now()
}

func un(s string, startTime time.Time) {
	endTime := time.Now()
	log.Println(s, "ElapsedTime in seconds:", endTime.Sub(startTime).Seconds())
}
