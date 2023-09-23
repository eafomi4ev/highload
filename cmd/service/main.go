package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"net/http"
	"os"
	"otus_highload/internal/api/rest/login"
	"otus_highload/internal/api/rest/user_register"
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
	userRegisterUC := uc_user_register.New(pgStore)

	// хендлеры
	loginHandler := login.New()
	userRegister := user_register.New(ctx, userRegisterUC)

	router := mux.NewRouter()
	router.HandleFunc("/login", loginHandler.Handle).Methods(http.MethodPost)
	router.HandleFunc("/user/register", userRegister.Handle).Methods(http.MethodPost)

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
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), `postgres://localhost:5432/social_db`) // todo: унести в env переменные
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}
