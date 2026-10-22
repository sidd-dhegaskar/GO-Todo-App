package database

import(
	"github.com/jackc/pgx/v5/pgxpool"
	"context"
	"log"
)

func Connect(databaseURL string)(*pgxpool.Pool, error){
	var ctx context.Context = context.Background()
	var config *pgxpool.Config
	var err error 
	config, err = pgxpool.ParseConfig(databaseURL)

	if err !=nil{
		log.Printf("Unable to parse DATABASE_URL: %v", err)
		return nil,err
	}

	var pool *pgxpool.Pool
	pool,err = pgxpool.NewWithConfig(ctx, config)
	if err != nil{
		log.Printf("Unable to create connection pool: %v", err)
		return nil,err
	}

	err=pool.Ping(ctx)
	if err != nil{
		log.Printf("Unable to ping db: %v", err)
		pool.Close()
		return nil, err
	}

	log.Println("Successfully connected to database")
	return pool,nil
}