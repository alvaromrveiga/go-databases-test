package main

import (
	badger "database-benchmarks/db_implementation/badger"
	sqlite "database-benchmarks/db_implementation/sqlite3"
	"testing"
)

func BenchmarkBadger(b *testing.B) {
	db, err := badger.Connect("db/test/badger")
	if err != nil {
		b.Fatalf("Not able to open BadgerDB connection: %v", err)
	}

	badgerDB := badger.BadgerDB{DB: db}

	b.Run("UpdateRead", func(b *testing.B) {
		for b.Loop() {
			updateReadBadger(badgerDB)
		}
	})

	// b.RunParallel(func(pb *testing.PB) {
	// 	for pb.Next() {
	// 		updateRead(badgerDB)
	// 	}
	// })
}

func BenchmarkSQLite(b *testing.B) {
	sql, err := sqlite.Connect("db/test/sqlite3")
	defer sql.Close()
	if err != nil {
		b.Fatalf("Not able to open SQLite3 connection: %v", err)
	}

	sqliteDB := sqlite.SQLiteDB{DB: sql}
	sqliteDB.InitializeDB()
	err = updateReadSqlLite(sqliteDB)
	if err != nil {
		b.Fatalf("Not able to run read-write transactions on SQLite3: %v", err)
	}

	// b.Run("UpdateRead", func(b *testing.B) {
	// 	for b.Loop() {
	// 		updateReadSqlLite(sqliteDB)
	// 	}
	// })

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			updateReadSqlLite(sqliteDB)
		}
	})
}
