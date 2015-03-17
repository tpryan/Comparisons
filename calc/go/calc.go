package main

import (
	"bytes"
	"database/sql"
	"flag"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db  *sql.DB
	max = flag.String("max", "1", "the max number of records to process.")
)

type Route struct {
	Airline  string
	SCode    string  //Source Airport Code
	SName    string  //Source Airport Name
	SLat     float64 //Source Airport Latitude
	SLon     float64 //Source Airport Longitude
	DCode    string  //Destination Airport Code
	DName    string  //Destination Airport Name
	DLat     float64 //Destination Airport Latitude
	DLon     float64 //Destination Airport Longitude
	Distance float64 //Distance
}

func main() {
	var err error
	flag.Parse()

	SQLhost := os.Getenv("OF_USER") + ":" + os.Getenv("OF_PASS") + "@tcp(" + os.Getenv("OF_HOST") + ":3306)/" + os.Getenv("OF_NAME")

	if db, err = sql.Open("mysql", SQLhost); err != nil {
		log.Fatal(err)
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	outdir := wd + "/calc/output/go/"

	if err = cleanDir(outdir); err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadFile(wd + "/calc/sql/prepstatement.sql")
	if err != nil {
		log.Fatal(err)
	}

	RouteSQL := string(b)
	RouteSQL += "\n" + "Limit 0," + *max + "\n"

	Routes, err := getRoutes(RouteSQL)
	if err != nil {
		log.Fatal(err)
	}

	Routes = processRoutes(Routes)
	if err = writeRoutes(Routes, outdir+"/1"); err != nil {
		log.Fatal(err)
	}

}

func writeRoutes(routes []Route, path string) error {
	if err := os.Mkdir(path, 0777); err != nil {
		return err
	}

	var b bytes.Buffer

	b.WriteString("<table>" + "\n")
	b.WriteString("	<tr>" + "\n")
	b.WriteString("		<th>Airline</th>" + "\n")
	b.WriteString("		<th>Origin Aiport Code</th>" + "\n")
	b.WriteString("		<th>Origin Aiport Name</th>" + "\n")
	b.WriteString("		<th>Origin Latitude</th>" + "\n")
	b.WriteString("		<th>Origin Longitude</th>" + "\n")
	b.WriteString("		<th>Destination Aiport Code</th>" + "\n")
	b.WriteString("		<th>Destination Aiport Name</th>" + "\n")
	b.WriteString("		<th>Destination Latitude</th>" + "\n")
	b.WriteString("		<th>Destination Longitude</th>" + "\n")
	b.WriteString("		<th>Distance</th>" + "\n")
	b.WriteString("	</tr>" + "\n")

	for _, r := range routes {

		b.WriteString("	<tr>" + "\n")
		b.WriteString("		<td>" + r.Airline + "</td>" + "\n")
		b.WriteString("		<td>" + r.SCode + "</td>" + "\n")
		b.WriteString("		<td>" + r.SName + "</td>" + "\n")
		b.WriteString("		<td>" + strconv.FormatFloat(r.SLat, 'f', 8, 64) + "</td>" + "\n")
		b.WriteString("		<td>" + strconv.FormatFloat(r.SLon, 'f', 8, 64) + "</td>" + "\n")
		b.WriteString("		<td>" + r.DCode + "</td>" + "\n")
		b.WriteString("		<td>" + r.DName + "</td>" + "\n")
		b.WriteString("		<td>" + strconv.FormatFloat(r.DLat, 'f', 8, 64) + "</td>" + "\n")
		b.WriteString("		<td>" + strconv.FormatFloat(r.DLon, 'f', 8, 64) + "</td>" + "\n")
		b.WriteString("		<td>" + strconv.FormatFloat(r.Distance, 'f', 10, 64) + "</td>" + "\n")
		b.WriteString("	</tr>" + "\n")
	}

	b.WriteString("</table>" + "\n")
	f := path + "/table.html"

	if err := ioutil.WriteFile(f, b.Bytes(), 0777); err != nil {
		return err
	}

	return nil

}

func processRoutes(routes []Route) []Route {
	for i, r := range routes {
		r.Distance = getDistance(r.SLat, r.SLon, r.DLat, r.DLon)
		routes[i] = r
	}
	return routes
}

func getDistance(la1, lo1, la2, lo2 float64) float64 {
	earth_radius := float64(3963)

	dLat := deg2rad(la2 - la1)
	dLon := deg2rad(lo2 - lo1)

	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(deg2rad(la1))*math.Cos(deg2rad(la2))*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Asin(math.Sqrt(a))
	d := earth_radius * c

	return d
}

func deg2rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func getRoutes(RouteSQL string) ([]Route, error) {
	rows, err := db.Query(RouteSQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []Route
	for rows.Next() {
		r := Route{}
		if err := rows.Scan(&r.Airline, &r.SCode, &r.SName, &r.SLat, &r.SLon, &r.DCode, &r.DName, &r.DLat, &r.DLon); err != nil {
			return nil, err
		}

		res = append(res, r)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return res, nil
}

func cleanDir(dir string) error {
	if err := os.RemoveAll(dir); err != nil {
		return err
	}

	err := os.Mkdir(dir, 0777)
	return err
}
