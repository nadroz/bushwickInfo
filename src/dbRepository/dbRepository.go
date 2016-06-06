package dbRepository

import (
    "fmt"
    "database/sql"
    _ "github.com/nakagami/firebirdsql"
    "strconv"
    "config"
    "dbSelect"
)

type DbRepository struct{
	db *sql.DB
}

type actor struct{
	Id int
	FirstName string
	LastName string
	Middlename string
	Single string
}

type DataFrame struct  {
	Headers []string  `json:"headers"`
	Rows [][]string `json:"rows"`
}

func Initialize(config *config.DbConfig) (*DbRepository, error){
	bdb, err := sql.Open("firebirdsql", config.BushwickConn)
	if err != nil {
		fmt.Println("connection error!")
		return nil, err
	}
	repo := DbRepository{ db: bdb }
	return &repo, nil
}

func (repo DbRepository) HealthCheck() (int, error){
	fmt.Println("checking health!")

	var actors int
	qErr := repo.db.QueryRow(dbSelect.CountActor).Scan(&actors)
	if qErr != nil {
		fmt.Println("Cannot ping!")
		return 0, qErr
	}
	
	fmt.Println(strconv.Itoa(actors))
	return actors, nil
}

func (repo DbRepository) GetActors() (*DataFrame, error) {
	rows, qErr := repo.db.Query(dbSelect.Actor)
	if qErr != nil {
		fmt.Println("Query error!")
		return nil, qErr
	}

	headers := make([]string, 2)
	headers[0] = "id"
	headers [1]  = "FullName"

	records := [][]string{}
	for rows.Next() {
		var id int
		var fullName string
		sErr := rows.Scan(&id, &fullName)
		if  sErr != nil {
			fmt.Println("Query problems, yo.")
		}
		values := []string{ strconv.Itoa(id), fullName }

		records = append(records, values)
	}

	frame := DataFrame{ Headers: headers, Rows: records }
	return &frame, nil
}
