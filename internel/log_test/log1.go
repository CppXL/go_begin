package log_test

import (
	"fmt"
	"log"
)

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}
func TLog() {
	fmt.Println("call TLog")
	log.Println("message")

	// log.Fatalln("fatal message")

}

func TLog1() {
	//log.Panicln("panic message")

}
