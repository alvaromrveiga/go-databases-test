package main

import (
	badger "database-benchmarks/db_implementation/badger"
	sqlite "database-benchmarks/db_implementation/sqlite3"
	"log"
	"strconv"
)

func main() {
	db, err := badger.Connect("db/badger")
	if err != nil {
		log.Fatalf("Not able to open BadgerDB connection: %v", err)
	}
	defer db.Close()

	badgerDB := badger.BadgerDB{DB: db}

	err = updateReadBadger(badgerDB)
	if err != nil {
		log.Fatalf("Not able to run read-write transactions on BadgerDB: %v", err)
	}

	sql, err := sqlite.Connect("db/sqlite3")
	defer sql.Close()
	if err != nil {
		log.Fatalf("Not able to open SQLite3 connection: %v", err)
	}

	sqliteDB := sqlite.SQLiteDB{DB: sql}
	sqliteDB.InitializeDB()
	err = updateReadSqlLite(sqliteDB)
	if err != nil {
		log.Fatalf("Not able to run read-write transactions on SQLite3: %v", err)
	}
}

func updateReadBadger(badgerDB badger.BadgerDB) error {
	for i := 0; i < 1000; i++ {
		key := []byte(strconv.Itoa(i))
		value := []byte(strconv.Itoa(i))

		err := badgerDB.UpdateOne(key, value)
		if err != nil {
			return err
		}
	}

	for i := 0; i < 1000; i++ {
		key := []byte(strconv.Itoa(i))

		item, err := badgerDB.FindOne(key)
		if err != nil {
			return err
		}

		log.Println(string(item))
	}

	log.Println()

	for i := 0; i < 1000; i++ {
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

		log.Println(string(item))
	}
	return nil
}

func updateReadSqlLite(sqliteDB sqlite.SQLiteDB) error {
	for i := 1; i <= 1000; i++ {
		sqliteDB.InsertOneTestTable(strconv.Itoa(1000 - i))
	}

	for i := 1; i <= 1000; i++ {
		item, err := sqliteDB.FindByIdTestTable(i)
		if err != nil {
			return err
		}

		log.Println(item)
	}

	log.Println()

	for i := 1; i <= 1000; i++ {
		sqliteDB.UpdateOneTestTable(i, strconv.Itoa(i))
		item, err := sqliteDB.FindByIdTestTable(i)
		if err != nil {
			return err
		}
		log.Println(item)
	}

	return nil
}
