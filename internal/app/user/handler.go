package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	sq "github.com/Masterminds/squirrel"
	"github.com/golang_example/internal/pkg/db/mysql"
	_ "github.com/golang_example/internal/pkg/db/mysql"
	"github.com/golang_example/internal/pkg/http/respond"
)

type (
	service interface {
		GetAllUser()
		UpdateUser()
		CreateUser()
	}
	User struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
	UserCreate struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	Handler struct {
	}
)

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) GetAllUser(w http.ResponseWriter, r *http.Request) {
	db := mysql.Connect()
	sql, _, _ := sq.Select("*").From("users").ToSql()
	result, err := db.Query(sql)

	if err != nil {
		panic(err.Error())
	}
	var users []User
	for result.Next() {
		var user User
		err := result.Scan(&user.ID, &user.Username, &user.Password)
		if err != nil {
			panic(err.Error())
		}
		users = append(users, user)
	}

	res := map[string]interface{}{
		"success": true,
		"users":   users,
	}
	respond.JSON(w, 200, res)
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	db := mysql.Connect()

	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err.Error())
	}

	sqlUpdate, args, _ := sq.Update("users").Set("username", user.Username).Set("password", user.Password).Where(sq.Eq{"id": user.ID}).ToSql()

	fmt.Print(user)

	userUpdate, err2 := db.Exec(sqlUpdate, args...)

	if err2 != nil {
		panic(err2.Error())
	}

	lastID, _ := userUpdate.LastInsertId()

	res := map[string]interface{}{
		"success": true,
		"id":      lastID,
	}
	respond.JSON(w, 200, res)
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	db := mysql.Connect()
	decoder := json.NewDecoder(r.Body)
	var userCreate UserCreate
	err := decoder.Decode(&userCreate)
	if err != nil {
		panic(err.Error())
	}

	sqlCreate, args, _ := sq.Insert("users").Columns("username", "password").Values(userCreate.Username, userCreate.Password).ToSql()
	user, err := db.Exec(sqlCreate, args...)

	if err != nil {
		panic(err.Error())
	}

	id, _ := user.LastInsertId()

	row, _ := user.RowsAffected()

	res := map[string]interface{}{
		"success":        true,
		"lastID":         id,
		"totalRowUpdate": row,
	}
	respond.JSON(w, 200, res)
}
