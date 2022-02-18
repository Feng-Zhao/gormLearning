package main

import (
	"booklibrary/service"
	"bufio"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"reflect"
	"strings"

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
	viper.AddConfigPath("./conf")
	viper.SetConfigName("config")
	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {             // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
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
	// 标准输入流
	scanner := bufio.NewScanner(os.Stdin)
	// while true
	for scanner.Scan() {
		cmd := scanner.Text()
		// quit
		if cmd == "quit" {
			os.Exit(0)
		} else if cmd == "help" { // help msg
			printHelpMsg()
		} else { // api调用
			// 切分参数
			args := strings.Split(cmd, " ")
			// api检索
			if fc, ok := s.APIS[args[0]]; !ok {
				log.Println("wrong command, try again")
			} else {
				// 反射调用 method
				callFunc(fc.Func, args[1:], db)
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
	log.Printf("callfunc: %v | args: %v\n", t.Name(), args)
	// 构建 []Value 作为参数
	in := []reflect.Value{reflect.ValueOf(db)}
	for _, v := range args {
		arg := reflect.ValueOf(v)
		in = append(in, arg)
	}
	// 检查参数个数
	if t.NumIn() != len(in) {
		log.Println("wrong args, please check the manual")
		return
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
	fmt.Println("listClass : show all classes information.")
	fmt.Println("listTeacher : show all teacher information.")
	fmt.Println("listBook : show all books information.")
}
