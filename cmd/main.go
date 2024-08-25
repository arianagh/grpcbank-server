package main

import (
	"database/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
	mydb "grpcbank/src/adapter/database"
	"grpcbank/src/adapter/grpc"
	"grpcbank/src/application"
	"grpcbank/src/application/domain"
	dbmigration "grpcbank/src/db"
	"log"
	"math/rand"
	"time"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(logWriter{})

	sqlDB, err := sql.Open("pgx", "postgres://postgres:postgres@localhost:5432/grpc")

	if err != nil {
		log.Fatalln("Can't connect database :", err)
	}

	dbmigration.Migrate(sqlDB)

	databaseAdapter, err := mydb.NewDatabaseAdapter(sqlDB)

	if err != nil {
		log.Fatalln("Can't create database adapter :", err)
	}

	bankService := application.NewBankService(databaseAdapter)

	go generateExchangeRates(bankService, "USD", "IDR", 5*time.Second)

	grpcAdapter := grpc.NewGrpcAdapter(bankService, 9000)

	grpcAdapter.Run()
}

func generateExchangeRates(bs *application.BankService, fromCurrency, toCurrency string, duration time.Duration) {
	ticker := time.NewTicker(duration)

	for range ticker.C {
		now := time.Now()
		validFrom := now.Truncate(time.Second).Add(3 * time.Second)
		validTo := validFrom.Add(duration).Add(-1 * time.Millisecond)

		dummyRate := domain.ExchangeRate{
			FromCurrency:       fromCurrency,
			ToCurrency:         toCurrency,
			ValidFromTimestamp: validFrom,
			ValidToTimestamp:   validTo,
			Rate:               2000 + float64(rand.Intn(300)),
		}

		bs.CreateExchangeRate(dummyRate)
	}
}
