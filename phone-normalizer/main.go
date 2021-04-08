package main

import (
	"fmt"
	"regexp"

	"github.com/Gabriel2233/gophercises/phone-normalizer/db"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "phone_db"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)
	must(db.Reset("postgres", psqlInfo, dbname))

	psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
	must(db.Migrate("postgres", psqlInfo))

	db, err := db.Open("postgres", psqlInfo)
	must(err)
	defer db.Close()

	if err = db.Seed(); err != nil {
		panic(err)
	}

	phones, err := db.AllPhones()
	must(err)
	for _, p := range phones {
		fmt.Printf("Working on... %+v\n", p)
		number := normalize(p.Number)

		if number != p.Number {
			fmt.Println("Updating or removing...", number)
			existing, err := db.FindPhone(number)
			must(err)
			if existing != nil {
				must(db.DeletePhone(p.ID))
			} else {
				p.Number = number
				must(db.UpdatePhone(&p))
			}
		} else {
			fmt.Println("No changes required")
		}
	}
}

func normalize(n string) string {
	re := regexp.MustCompile("[^0-9]")
	return re.ReplaceAllString(n, "")
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
