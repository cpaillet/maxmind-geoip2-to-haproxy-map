package main

import (
	"fmt"
	maxminddb "github.com/oschwald/maxminddb-golang"
	"log"
	"os"
	"path"
	"strings"
)

var db string
var destDir string

func main() {
	getFlags()
	mmdb, err := maxminddb.Open(db)
	if err != nil {
		log.Fatal(err)
	}
	defer mmdb.Close()

	record := struct {
		Country struct {
			ISOCode string `maxminddb:"iso_code"`
		} `maxminddb:"country"`
		Location struct {
			Latitude  float64 `maxminddb:"latitude"`
			Longitude float64 `maxminddb:"longitude"`
		} `maxminddb:"location"`
		Subdivisions []struct {
			IsoCode string `maxminddb:"iso_code"`
		} `maxminddb:"subdivisions"`
	}{}

	if err = os.Mkdir(destDir, os.ModePerm); err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	file, err := os.Create(path.Join(destDir, "geoip.map"))
	if err != nil {
			log.Fatal(err)
		}

	networks := mmdb.Networks()
	i := 0
	for networks.Next() {
		subnet, err := networks.Network(&record)
		if err != nil {
			log.Fatal(err)
		}
		key := subnet.String()
		values := []string{record.Country.ISOCode,
			record.Subdivisions[0].IsoCode,
			fmt.Sprintf("%f", record.Location.Latitude),
			fmt.Sprintf("%f", record.Location.Longitude)}
		_, err = file.WriteString(key + " " + strings.Join(values, "|") + "\n")
		i++
	}

	if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	fmt.Printf("Wrote %d records from %s\n", i, db)
}