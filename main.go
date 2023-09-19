package main

import (
	"BikeCH/src/data"
	"database/sql"
	"fmt"
	"os"

	_ "BikeCH/testing_init"

	_ "github.com/lib/pq"
)

var db *sql.DB

// Connect to the postgresql db and create the nodes table
func main() {
	var err error

	// Postgresql connection
	db_url := os.Getenv("POSTGRES_URL")
	connStr := db_url
	db, err = sql.Open("postgres", connStr)
	checkError(err)

	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("The database is connected")

	// Create a table
	createTb := `CREATE TABLE IF NOT EXISTS nodes(
		id 			int,
		ch_e 		double precision,
		ch_n 		double precision
	);`

	_, err = db.Exec(createTb)
	checkError(err)
	fmt.Println("Successfully created table")

	// Load nodes data from binary files
	graph, err := data.LoadFrom("data/ch_full")
	checkError(err)
	count := graph.NodeCount()
	fmt.Println("Node count: ", count)

	queryLast10 := `SELECT id, ch_e, ch_n FROM nodes ORDER BY id DESC LIMIT 10;`
	rows, err := db.Query(queryLast10)
	checkError(err)
	fmt.Println("Last 10 nodes:")
	for rows.Next() {
		var id int
		var e float64
		var n float64
		err = rows.Scan(&id, &e, &n)
		checkError(err)
		fmt.Println(id, e, n)
	}

	// Start from last checkpoint
	queryMaxId := `SELECT id, ch_e, ch_n FROM nodes ORDER BY id DESC LIMIT 1;`
	row := db.QueryRow(queryMaxId)
	var id int
	var e float64
	var n float64
	err = row.Scan(&id, &e, &n)
	checkError(err)
	fmt.Println("Last checkpoint: ", id, e, n)

	num_iters := 50
	for it := 0; it < num_iters; it++ {
		// Prepare a slice to hold the values
		var values []interface{}
		max_batch_length := 20000

		start := (id + 1) + it*max_batch_length
		end := start + max_batch_length

		if start > count {
			break
		}

		for i := start; i < end && i < count; i++ {
			point := graph.NodePoint(i)
			e := point.E
			n := point.N

			// Append values to the slice
			values = append(values, i, e, n)
		}

		// Build the bulk insert query
		insertData := `INSERT INTO nodes (id, ch_e, ch_n) VALUES `

		for i := 0; i < len(values)/3; i++ {
			if i > 0 {
				insertData += ", "
			}
			insertData += fmt.Sprintf("($%d, $%d, $%d)", i*3+1, i*3+2, i*3+3)
		}

		println("Bulk insert from", start, "to", end-1)

		// Execute the bulk insert
		_, err = db.Exec(insertData, values...)
		checkError(err)
	}

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
