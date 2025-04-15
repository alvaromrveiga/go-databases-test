package main

import (
	badger "database-benchmarks/db_implementation/badger"
	sqlite "database-benchmarks/db_implementation/sqlite3"
	"log"
	"os"
	"strconv"
	"time"
)

const shouldLog = true
const goRoutinesCount = 10
const operationsCount = 1000

func main() {
	os.RemoveAll("db")

	log.Println("Starting BadgerDB benchmark")
	start := time.Now()
	db, err := badger.Connect("db/badger")
	if err != nil {
		log.Fatalf("Not able to open BadgerDB connection: %v", err)
	}
	defer db.Close()

	badgerDB := badger.BadgerDB{DB: db}

	runBagerRoutines(badgerDB)

	log.Printf("Badger total time: %v", time.Since(start))

	log.Println("Starting SQLite3 benchmark")
	sqlStart := time.Now()

	sql, err := sqlite.Connect("db/sqlite3")
	defer sql.Close()
	if err != nil {
		log.Fatalf("Not able to open SQLite3 connection: %v", err)
	}

	sqliteDB := sqlite.SQLiteDB{DB: sql}
	sqliteDB.InitializeDB()

	runSQLiteRoutines(sqliteDB)

	log.Println("SQLite3 total time: ", time.Since(sqlStart))
}

func updateReadBadger(badgerDB badger.BadgerDB) error {
	for i := 0; i < operationsCount; i++ {
		key := []byte(strconv.Itoa(i))
		value := []byte(strconv.Itoa(i))

		err := badgerDB.UpdateOne(key, value)
		if err != nil {
			return err
		}
	}

	for i := 0; i < operationsCount; i++ {
		key := []byte(strconv.Itoa(i))

		item, err := badgerDB.FindOne(key)
		if err != nil {
			return err
		}

		if shouldLog {
			log.Println(string(item))
		}
	}

	if shouldLog {
		log.Println()
	}

	for i := 0; i < operationsCount; i++ {
		key := []byte(strconv.Itoa(i))
		value := []byte(strconv.Itoa(i))

		err := badgerDB.UpdateOne(key, value)
		if err != nil {
			return err
		}

		item, err := badgerDB.FindOne(key)
		if err != nil {
			return err
		}

		if shouldLog {
			log.Println(string(item))
		}
	}
	return nil
}

func updateReadSqlLite(sqliteDB sqlite.SQLiteDB) error {
	for i := 1; i <= operationsCount; i++ {
		sqliteDB.InsertOneTestTable(strconv.Itoa(operationsCount - i))
	}

	for i := 1; i <= operationsCount; i++ {
		item, err := sqliteDB.FindByIdTestTable(i)
		if err != nil {
			return err
		}

		if shouldLog {
			log.Println(string(item))
		}
	}

	if shouldLog {
		log.Println()
	}

	for i := 1; i <= operationsCount; i++ {
		sqliteDB.UpdateOneTestTable(i, strconv.Itoa(i))
		item, err := sqliteDB.FindByIdTestTable(i)
		if err != nil {
			return err
		}
		if shouldLog {
			log.Println(string(item))
		}
	}

	return nil
}

func runBagerRoutines(db badger.BadgerDB) {
	channel := make(chan error)
	for i := 0; i < goRoutinesCount; i++ {
		go (func() {
			err := updateReadBadger(db)
			channel <- err
		})()
	}

	for i := 0; i < goRoutinesCount; i++ {
		err := <-channel
		if err != nil {
			log.Printf("Badger error %d: %v", i, err)
		}
	}
}

func runSQLiteRoutines(db sqlite.SQLiteDB) {
	channel := make(chan error)
	for i := 0; i < goRoutinesCount; i++ {
		go (func() {
			err := updateReadSqlLite(db)
			channel <- err
		})()
	}

	for i := 0; i < goRoutinesCount; i++ {
		err := <-channel
		if err != nil {
			log.Printf("SQLite error %d: %v", i, err)
		}
	}
}
