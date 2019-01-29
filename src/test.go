package main
import (
"fmt"
"net/http"
"strings"
"log"
"./myhttp"
"database/sql"
_ "github.com/go-sql-driver/mysql"
)
func sayhelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm() //解析参数，默认是不会解析的
    fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
    // fmt.Println("path", r.URL.Path)
    // fmt.Println("scheme", r.URL.Scheme)
    // fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func main() {
    http.HandleFunc("/", sayhelloName) //设置访问的路由
    http.HandleFunc("/login", login) //设置访问的路由
    err := http.ListenAndServe(":9090", nil) //设置监听的端口
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

func login(w http.ResponseWriter,r *http.Request){
    mysqltest()
    if myhttp.GetMethod(r) == "GET" {
        var data []string
        myhttp.View(w,"views/login.html",data)
    } else {
    //请求的是登陆数据，那么执行登陆的逻辑判断
        //r.ParseForm()
        fmt.Println("username:", myhttp.GetVal(r,"username"))
        fmt.Println("password:", myhttp.GetVal(r,"password"))
    }
}

func mysqltest(){
    db, _ := sql.Open("mysql", "root:root@tcp(120.78.183.163:3306)/novel_test?charset=utf8")
    rows, _ := db.Query("SELECT content FROM novel_content limit 5")
    fmt.Println(rows)
    for rows.Next() {
        var content string
        rows.Scan(&content)
        //checkErr(err)
        fmt.Println(content)
    }
}