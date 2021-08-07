package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// DBConn PostgreSQLに接続
func DBConn() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("envファイルの読み込みに失敗しました。 \n", err)
	}
	// dsn := fmt.Sprintf("%s:%s@%s/%s?sslmode=disable", os.Getenv("PSQL_USER"), os.Getenv("PSQL_PASS"), os.Getenv("PSQL_PORT"), os.Getenv("PSQL_DBNAME"))
	dsn := fmt.Sprintf("host=db port=%s user=%s dbname=%s password=%s sslmode=disable", os.Getenv("PSQL_PORT"), os.Getenv("PSQL_USER"), os.Getenv("PSQL_DBNAME"), os.Getenv("PSQL_PASS"))
	log.Print(dsn)
	log.Print("PostgreSQL DBに接続しています...")
	db, err := gorm.Open("postgres", dsn)
	// log.Println("接続しました。")

	if err != nil {
		log.Fatal("接続に失敗しました。 \n", err)
	}
	// 赤文字のログを出す
	// db.LogMode(true)
	return db
}

// User テーブル名はusers
type User struct {
	ID        string `gorm:"primary_key"`
	Name      string
	ImageURL  string
	UID       string
	CreatedAt time.Time
}

// Comment テーブル名はcomments
type Comment struct {
	ID        string `gorm:"primary_key"`
	Content   string
	UserID    int
	CreatedAt time.Time
}

type Image struct {
	ID        string `gorm:"primary_key"`
	URL       string
	UserID    int
	CreatedAt time.Time
}

type Category struct {
	ID        string `gorm:"primary_key"`
	Content   string `gorm:"ForeignKey:Category"`
	CreatedAt time.Time
}

type CategoriesOfComment struct {
	ID         string `gorm:"primary_key"`
	CommentID  int
	CategoryID int
	CreatedAt  time.Time
}

func getUserAll(w http.ResponseWriter, r *http.Request) {
	var users []User
	db := DBConn()
	defer db.Close()

	db.Order("id").Find(&users)
	res, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(res)
	return
}

func getUserByUID(w http.ResponseWriter, r *http.Request) {
	var user User
	db := DBConn()
	defer db.Close()

	params := mux.Vars(r)
	db.Where("uid = ?", params["uid"]).First(&user)
	res, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(res)
	return
}

func getUserByUserID(w http.ResponseWriter, r *http.Request) {
	var user User
	db := DBConn()
	defer db.Close()

	params := mux.Vars(r)
	db.Where("id = ?", params["id"]).First(&user)
	res, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(res)
	return
}

func createUser(w http.ResponseWriter, r *http.Request) {
	user := User{}
	db := DBConn()
	defer db.Close()

	name := r.FormValue("name")
	uid := r.FormValue("uid")
	user.Name = name
	user.UID = uid
	db.Create(&user)
}

func editUser(w http.ResponseWriter, r *http.Request) {
	userBefore := User{}
	db := DBConn()
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]
	userBefore.ID = id
	name := r.FormValue("name")
	userAfter := userBefore
	db.First(&userAfter)
	userAfter.Name = name
	db.Model(&userBefore).Update(&userAfter)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	users := User{}
	db := DBConn()
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]
	users.ID = id
	db.First(&users)
	db.Delete(&users)
}

func getAllComment(w http.ResponseWriter, r *http.Request) {
	var comment []Comment
	db := DBConn()
	defer db.Close()

	db.Order("created_at desc").Find(&comment)
	res, err := json.Marshal(comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(res)
	return
}

func getUserCommentByUserID(w http.ResponseWriter, r *http.Request) {
	var comment []Comment
	db := DBConn()
	defer db.Close()

	params := mux.Vars(r)
	db.Order("created_at desc").Where("user_id = ?", params["id"]).Find(&comment)
	res, err := json.Marshal(comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(res)
	return
}

func postComment(w http.ResponseWriter, r *http.Request) {
	comment := Comment{}
	db := DBConn()
	defer db.Close()

	content := r.FormValue("content")
	// UserID := r.FormValue("UserID")
	userID, _ := strconv.Atoi(r.FormValue("user_id"))
	comment.Content = content
	comment.UserID = userID
	db.Create(&comment)
}

func deleteUserCommentByUserID(w http.ResponseWriter, r *http.Request) {
	comment := Comment{}
	db := DBConn()
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]
	comment.ID = id
	db.First(&comment)
	db.Delete(&comment)
}

func getUserImages(w http.ResponseWriter, r *http.Request) {
	var image []Image
	db := DBConn()
	defer db.Close()

	params := mux.Vars(r)
	db.Order("created_at desc").Where("user_id = ?", params["id"]).Find(&image)
	res, err := json.Marshal(image)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(res)
	return
}

func deleteUserImage(w http.ResponseWriter, r *http.Request) {
	image := Image{}
	db := DBConn()
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]
	image.ID = id
	db.First(&image)
	db.Delete(&image)
}

func getCategoriesAll(w http.ResponseWriter, r *http.Request) {
	var category []Category
	db := DBConn()
	defer db.Close()

	db.Order("id").Find(&category)
	res, err := json.Marshal(category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(res)
	return
}

func getAttachedCategoriesByCommentID(w http.ResponseWriter, r *http.Request) {
	var categories_of_comment []CategoriesOfComment
	var category []Category
	db := DBConn()
	defer db.Close()

	params := mux.Vars(r)
	db.Order("created_at desc").Where("comment_id = ?", params["comment_id"]).Find(&categories_of_comment)

	categories_id := categories_of_comment[0].CategoryID

	db.Order("id").Find(&category).Where("id = ?", categories_id).Find(&category)

	res, err := json.Marshal(category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(res)
	return
}

func main() {
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT", "OPTIONS"})
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	r := mux.NewRouter()

	// ルート(エンドポイント)
	r.HandleFunc("/api/users", getUserAll).Methods("GET")
	r.HandleFunc("/api/users/{uid}", getUserByUID).Methods("GET")
	r.HandleFunc("/api/comment_users/{id}", getUserByUserID).Methods("GET")
	r.HandleFunc("/api/users", createUser).Methods("POST")
	r.HandleFunc("/api/users/edit/{id}", editUser).Methods("POST")
	r.HandleFunc("/api/users/delete/{id}", deleteUser).Methods("DELETE")

	r.HandleFunc("/api/comments", getAllComment).Methods("GET")
	r.HandleFunc("/api/comments/{id}", getUserCommentByUserID).Methods("GET")
	r.HandleFunc("/api/comments", postComment).Methods("POST")
	r.HandleFunc("/api/comments/{id}", deleteUserCommentByUserID).Methods("DELETE")

	r.HandleFunc("/api/images/{id}", getUserImages).Methods("GET")
	r.HandleFunc("/api/images/{id}", deleteUserImage).Methods("DELETE")

	r.HandleFunc("/api/categories", getCategoriesAll).Methods("GET")

	r.HandleFunc("/api/attachedCategories/{comment_id}", getAttachedCategoriesByCommentID).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(r)))
}
