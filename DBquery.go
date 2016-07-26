import "log"

func DBquery() {
	var str string
	rows, err := db.Query("select userpass from T")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&str)
		if err != nil {
			log.Fatal(err)
		}
	}
}
