package main

import (
	"booklibrary/config"
	"booklibrary/service"
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/spf13/viper"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var s *service.Service
var dsn string

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime | log.Lmsgprefix)
	s = service.NewService()
	s.InitApi()

}

func main() {
	config.InitConfig()
	db := dbLink()
	// 标准输入流
	scanner := bufio.NewScanner(os.Stdin)

	// 主循环
	for scanner.Scan() {
		//  每次读一行
		cmd := scanner.Text()
		// quit
		if cmd == "quit" {
			os.Exit(0)
		} else if cmd == "help" { // help msg
			printHelpMsg()
		} else { // api调用
			// 切分参数
			args := strings.Split(strings.Trim(cmd, " "), " ")
			// api检索
			if fc, ok := s.APIS[args[0]]; !ok {
				log.Println("wrong command, try again")
			} else {
				// 反射调用 method
				callFunc(fc.Func, args, db)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func callFunc(f interface{}, args []string, db *gorm.DB) {
	// 获取 type
	t := reflect.TypeOf(f)
	if t.Kind() != reflect.Func {
		fmt.Println("api's func must be a function")
		return
	}
	log.Printf("callfunc: %v | args: %v\n", args[0], args[1:])

	// 检查参数个数
	if t.NumIn() != len(args) {
		log.Println("wrong number of args, please check the manual")
		return
	}

	// 构建 []Value 作为参数
	in := []reflect.Value{reflect.ValueOf(db)}

	// 转换 json
	for i := 1; i < len(args); i++ {
		req := newInstance(t.In(i))

		err := json.Unmarshal([]byte(args[i]), req)
		if err != nil {
			log.Printf("参数 %s 不符合 json 格式\n cause: %s", args[i], err)
		}
		arg := reflect.ValueOf(req)
		in = append(in, arg)
	}

	log.Printf("%v\n", in)
	// 执行方法
	reflect.ValueOf(f).Call(in)
}

func printHelpMsg() {
	fmt.Println("quit : quit the application.")
	fmt.Println("help : print the help msg.")
	fmt.Printf("\n//------------ list ------------------------\n")
	fmt.Println("listStudent : show all students information.")
	//fmt.Println("listClass : show all classes information.")
	fmt.Println("listTeacher : show all teacher information.")
	//fmt.Println("listBook : show all books information.")
}

func dbLink() *gorm.DB {
	// 数据库URL
	if viper.GetBool("db.needpassword") {
		dsn = fmt.Sprintf("%v:%v@tcp(%v:%v)/test?charset=utf8&parseTime=True&loc=Local",
			viper.GetString("db.username"),
			viper.GetString("db.password"),
			viper.GetString("db.ip"),
			viper.GetString("db.port"))
	} else {
		dsn = fmt.Sprintf("%v@tcp(%v:%v)/test?charset=utf8&parseTime=True&loc=Local",
			viper.GetString("db.username"),
			viper.GetString("db.ip"),
			viper.GetString("db.port"))
	}
	// 链接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	log.Printf("database connected:%v\n,", db)
	return db
}

func newInstance(t reflect.Type) interface{} {
	switch t.Kind() {
	case reflect.Ptr, reflect.Interface:
		return newInstance(t.Elem())
	default:
		return reflect.New(t).Interface()
	}
}
