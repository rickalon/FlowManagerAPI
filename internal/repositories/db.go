package repositories

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type PqDB struct {
	DB *sql.DB
}

func NewPqDriver(cfg string) *PqDB {
	//preparing the db connection(creating a pool of connections), db driver postgres, connectino details
	log.Println("Creating a pool of connections to pq")
	db, err := sql.Open("postgres", cfg)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Ping to de db...")
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("Ping accepted")
	return &PqDB{DB: db}
}

func (db *PqDB) SetUpDatabases() {

	log.Println("Setting up Users table")
	if err := db.createUserTable(); err != nil {
		log.Fatal(err)
	}
	log.Println("Setting up Proyects table")
	if err := db.createProyectTable(); err != nil {
		log.Fatal(err)
	}
	log.Println("Setting up Tasks table")
	if err := db.createTaskTable(); err != nil {
		log.Fatal(err)
	}
}

func (db *PqDB) createUserTable() error {
	_, err := db.DB.Exec(
		`
		CREATE TABLE IF NOT EXISTS USERS(
		USER_ID SERIAL PRIMARY KEY,
		FULL_NAME VARCHAR(100) NOT NULL,
		PASSWORD VARCHAR(100) NOT NULL,
		EMAIL VARCHAR(200) UNIQUE NOT NULL,
		CREATED_AT TIMESTAMP DEFAULT NOW()
		)
		`,
	)
	return err
}

func (db *PqDB) createProyectTable() error {
	_, err := db.DB.Exec(
		`
		CREATE TABLE IF NOT EXISTS PROYECTS(
		PROYECT_ID SERIAL PRIMARY KEY,
		NAME VARCHAR(255) NOT NULL,
		CREATED_AT TIMESTAMP NOT NULL DEFAULT NOW()
		)
		`,
	)
	return err
}

func (db *PqDB) createTaskTable() error {
	var statusEnum *bool
	db.DB.QueryRow(`SELECT EXISTS (
		SELECT 1 
		FROM pg_type 
		WHERE typname = 'status_enum' 
		AND typtype = 'e'
	);`).Scan(&statusEnum)

	tx, err := db.DB.Begin()
	if err != nil {
		log.Println("Failed to start transaction")
		return err
	}
	if !(*statusEnum) {
		if _, err := tx.Exec(`CREATE TYPE STATUS_ENUM AS ENUM ('TODO', 'PROGRESS', 'DONE');`); err != nil {
			tx.Rollback()
			log.Println("Failed to create the enum")
			return err
		}
	}
	if _, err := tx.Exec(`
	CREATE TABLE IF NOT EXISTS TASKS(
	TASK_ID SERIAL PRIMARY KEY,
	CONTENT VARCHAR(200) NOT NULL,
	STATUS STATUS_ENUM NOT NULL DEFAULT 'TODO',
	PROYECT_ID INT REFERENCES PROYECTS(PROYECT_ID),
	USER_ID INT REFERENCES USERS(USER_ID),
	CREATED_AT TIMESTAMP NOT NULL DEFAULT NOW()
	)
	;`); err != nil {
		tx.Rollback()
		log.Println("Failed to create the table")
		return err
	}
	err = tx.Commit()
	return err
}
