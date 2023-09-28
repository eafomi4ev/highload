package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"

	"otus_highload/internal/api/rest/login"
	"otus_highload/internal/api/rest/user_get_by_id"
	"otus_highload/internal/api/rest/user_register"
	"otus_highload/internal/app/usecase/uc_login"
	"otus_highload/internal/app/usecase/uc_user_get_by_id"
	"otus_highload/internal/app/usecase/uc_user_register"
	"otus_highload/internal/storage/db_pg"
)

func main() {
	ctx := context.Background()

	connect := connectToDB(ctx)
	defer connect.Close(ctx)

	// БД
	pgStore := db_pg.New(connect)

	// юзкейсы
	loginUC := uc_login.New(pgStore)
	userRegisterUC := uc_user_register.New(pgStore)
	userGetByIDUC := uc_user_get_by_id.New(pgStore)

	// хендлеры
	loginHandler := login.New(ctx, loginUC)
	userRegister := user_register.New(ctx, userRegisterUC)
	getUserByID := user_get_by_id.New(ctx, userGetByIDUC)

	// роутер
	router := mux.NewRouter()
	router.HandleFunc("/login", loginHandler.Handle).Methods(http.MethodPost)
	router.HandleFunc("/user/register", userRegister.Handle).Methods(http.MethodPost)
	router.HandleFunc("/user/get/{id}", getUserByID.Handle).Methods(http.MethodGet)

	// сервер
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", 8000),
		Handler: router,
	}
	if err := server.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("error running http server: %s\n", err)
		}
	}
}

func connectToDB(ctx context.Context) *pgx.Conn {
	const (
		POSTGRES_HOST        = "POSTGRES_HOST"
		POSTGRES_PORT        = "POSTGRES_PORT"
		POSTGRES_USER_NAME   = "POSTGRES_USER_NAME"
		POSTGRES_DB_PASSWORD = "POSTGRES_DB_PASSWORD"
		POSTGRES_DB_NAME     = "POSTGRES_DB_NAME"
	)

	pgHost := shouldBePresent(POSTGRES_HOST)(os.LookupEnv(POSTGRES_HOST))
	pgPort := shouldBePresent(POSTGRES_PORT)(os.LookupEnv(POSTGRES_PORT))
	pgUserName := shouldBePresent(POSTGRES_USER_NAME)(os.LookupEnv(POSTGRES_USER_NAME))
	pgDBPassword := os.Getenv(POSTGRES_DB_PASSWORD) // для локальной разработки пароля нет; для прода, конечно, должен быть пароль.
	pgDBName := shouldBePresent(POSTGRES_DB_NAME)(os.LookupEnv(POSTGRES_DB_NAME))

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", pgUserName, pgDBPassword, pgHost, pgPort, pgDBName)
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}

func shouldBePresent(envName string) func(string, bool) string {
	return func(env string, isPresent bool) string {
		if isPresent && env != "" {
			return env
		}

		log.Fatal(fmt.Sprintf("env %s is not present", envName))
		return ""
	}
}
