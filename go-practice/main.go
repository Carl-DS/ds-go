package main

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"github.com/drone/routes/exp/router"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql" // MySQL Database driver
)

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPwd := "root"
	dbName := "gs_internationalization"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPwd+"@/"+dbName)
	checkErr(err)
	// Return db object to be used by other functions
	return db
}

type I18n struct {
	Id        int
	Code      string
	ZhCn      string
	UsEn      string
	Project   string
	CreatedAt string
	UpdatedAt string
}

var tmpl = template.Must(template.ParseGlob("tmpl/*"))

func Index(w http.ResponseWriter, r *http.Request) {

	param := r.URL.Query()
	condition := param.Get("condition")
	res := GetAllI18N(condition)

	tmpl.ExecuteTemplate(w, "Index", res)

}

func GetAllI18N(condition string) []I18n {
	// Open database connection
	db := dbConn()
	// Close database connection
	defer db.Close()

	rows, err := db.Query("select id, code, zh_cn as zhCn, us_en as usEn, project, created_at as createdAt, updated_at as updatedAt from pub_internationalization order by id DESC ")
	checkErr(err)

	if condition != "" {
		fmt.Println(condition)
	}
	n := I18n{}

	res := []I18n{}

	for rows.Next() {
		var id int
		var code, zhCn, usEn, project, createdAt, updatedAt string

		err = rows.Scan(&id, &code, &zhCn, &usEn, &project, &createdAt, &updatedAt)
		checkErr(err)

		n.Id = id
		n.Code = code
		n.ZhCn = zhCn
		n.UsEn = usEn
		n.Project = project
		n.CreatedAt = createdAt
		n.UpdatedAt = updatedAt

		res = append(res, n)
	}
	return res
}

func Show(w http.ResponseWriter, r *http.Request) {
	// Open database connection
	db := dbConn()

	nId := r.URL.Query().Get("id")

	rows, err := db.Query("select id, code, zh_cn as zhCn, us_en as usEn, project, created_at as createdAt, updated_at as updatedAt from pub_internationalization where id = ?", nId)
	checkErr(err)

	n := I18n{}

	for rows.Next() {
		var id int
		var code, zhCn, usEn, project, createdAt, updatedAt string
		err := rows.Scan(&id, &code, &zhCn, &usEn, &project, &createdAt, &updatedAt)
		checkErr(err)

		n.Id = id
		n.Code = code
		n.ZhCn = zhCn
		n.UsEn = usEn
		n.Project = project
		n.CreatedAt = createdAt
		n.UpdatedAt = updatedAt
	}

	tmpl.ExecuteTemplate(w, "Show", n)
	// Close database connection
	defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	nId := r.URL.Query().Get("id")

	rows, err := db.Query("select id, code, zh_cn as zhCn, us_en as usEn, project, created_at as createdAt, updated_at as updatedAt from pub_internationalization where id = ?", nId)
	checkErr(err)

	n := I18n{}

	for rows.Next() {
		var id int
		var code, zhCn, usEn, project, createdAt, updatedAt string
		err := rows.Scan(&id, &code, &zhCn, &usEn, &project, &createdAt, &updatedAt)
		checkErr(err)

		n.Id = id
		n.Code = code
		n.ZhCn = zhCn
		n.UsEn = usEn
		n.Project = project
		n.CreatedAt = createdAt
		n.UpdatedAt = updatedAt
	}

	tmpl.ExecuteTemplate(w, "Edit", n)

	defer db.Close()
}

func UploadPage(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "UploadPage", nil)
}

func InsertI18n(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	if r.Method == "POST" {
		code := r.FormValue("code")
		zhCn := r.FormValue("zhCn")
		usEn := r.FormValue("usEn")
		project := r.FormValue("project")

		stmt, err := db.Prepare("insert into pub_internationalization(code, zh_cn, us_en, project) values (?, ?, ?, ?)")
		checkErr(err)
		stmt.Exec(code, zhCn, usEn, project)
	}

	defer db.Close()

	http.Redirect(w, r, "/", 301)
}

func UpdateI18n(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	if r.Method == "POST" {
		zhCn := r.FormValue("zhCn")
		usEn := r.FormValue("usEn")
		project := r.FormValue("project")
		id := r.FormValue("uid")

		stmt, err := db.Prepare("update pub_internationalization set zh_cn = ?, us_en = ?, project = ? where id = ?")
		checkErr(err)
		stmt.Exec(zhCn, usEn, project, id)
	}

	defer db.Close()

	http.Redirect(w, r, "/", 301)
}

func DeleteI18n(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	nId := r.URL.Query().Get("id")

	stmt, err := db.Prepare("delete from pub_internationalization where id = ?")
	checkErr(err)

	stmt.Exec(nId)

	defer db.Close()

	http.Redirect(w, r, "/", 301)
}

/**
处理/upload 逻辑
*/
func Upload(w http.ResponseWriter, r *http.Request) {
	// 获取请求的方法
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		currentTime := time.Now().Unix()
		hash := md5.New()
		io.WriteString(hash, strconv.FormatInt(currentTime, 10))
		token := fmt.Sprintf("%x", hash.Sum(nil))

		t, _ := tmpl.ParseFiles("UploadPage.tmpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, header, e := r.FormFile("uploadfile")
		checkErr(e)
		defer file.Close()
		fmt.Fprint(w, "%v", header.Header)
		openFile, err := os.OpenFile("./test/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		checkErr(err)
		defer openFile.Close()
		io.Copy(openFile, file)
	}
}

/**
生成国际化文件
*/
func Generate(w http.ResponseWriter, r *http.Request) {
	// Open database connection
	db := dbConn()
	defer db.Close()

	rows, err := db.Query("SELECT DISTINCT project FROM pub_internationalization")
	checkErr(err)
	for rows.Next() {
		var project string
		err := rows.Scan(&project)
		checkErr(err)
		fmt.Println("project：", project)

		process(project+"-i18n-cn.json", "CN", project)
		process(project+"-i18n-en.json", "EN", project)
	}

}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func process(filename string, typestr string, project string) {
	filepath := "E:\\IdeaProjects\\workspace\\gs-i18n\\doc\\i18n\\"
	exist, err := PathExists(filepath)
	if err != nil {
		fmt.Println("get dir error! [%v]", err)
		return
	}
	if !exist {
		fmt.Printf("no dir! [%v]\n", filepath)
		// 创建文件夹
		err := os.Mkdir(filepath, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed! [%v]\n", err)
		} else {
			fmt.Printf("mkdir success!\n")
		}
	}

	pathname := filepath + filename
	pathExists, err := PathExists(pathname)
	if err != nil {
		fmt.Println("get file error! [%v]", err)
		return
	}

	if !pathExists {
		os.Create(pathname)
	}

	inputFile, inputError := os.Open(pathname)
	defer inputFile.Close()
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return // exit the function on error
	}
	inputFile.WriteString("i18nStr\n")
}

func main() {

	log.Println("Server start on: http-server://localhost:9000")

	mux := router.New()
	mux.Get("/i18n/generate", Generate)

	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/uploadPage", UploadPage)

	http.HandleFunc("/insert", InsertI18n)
	http.HandleFunc("/update", UpdateI18n)
	http.HandleFunc("/delete", DeleteI18n)

	http.HandleFunc("/upload", Upload)

	http.ListenAndServe(":9000", nil)
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
