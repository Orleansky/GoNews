package main

import (
	"GoNews/pkg/api"
	"GoNews/pkg/storage"
	"GoNews/pkg/storage/memdb"
	"GoNews/pkg/storage/postgres"
	"log"
	"net/http"
)

// Сервер GoNews.
type server struct {
	db  storage.Interface
	api *api.API
}

func main() {
	// Создаём объект сервера.
	var srv server

	// Создаём строку подключения к PostgreSQL.
	connstr := "postgres://postgres:qwerty@localhost/posts?sslmode=disable"

	// Создаём объекты баз данных.
	// БД в памяти.
	db := memdb.New()

	// Реляционная БД PostgreSQL.
	db2, err := postgres.New(connstr)
	if err != nil {
		log.Fatal(err)
	}

	// post1 := storage.Post{
	// 	Title:    "Title",
	// 	Content:  "Content",
	// 	AuthorID: 0,
	// }
	// db2.AddPost(post1)
	// Документная БД MongoDB.
	// db3, err := mongo.New("mongodb://server.domain:27017/")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	_ = db

	// Инициализируем хранилище сервера конкретной БД.
	srv.db = db2

	// Создаём объект API и регистрируем обработчики.
	srv.api = api.New(srv.db)

	// Запускаем веб-сервер на порту 8080 на всех интерфейсах.
	// Предаём серверу маршрутизатор запросов,
	// поэтому сервер будет все запросы отправлять на маршрутизатор.
	// Маршрутизатор будет выбирать нужный обработчик.
	http.ListenAndServe(":8080", srv.api.Router())
}
