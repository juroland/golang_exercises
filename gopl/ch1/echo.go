package main

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
	"time"
)

func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}

func exo1() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}

func exo2() {
	for i, arg := range os.Args[1:] {
		fmt.Printf("%v : %v\n", i+1, arg)
	}
}

func benchmark(f func(), n int) {
	// Redirect stdout to /dev/null
	devNull, _ := os.Open(os.DevNull)
	stdout := os.Stdout
	os.Stdout = devNull
	defer devNull.Close()

	pc := reflect.ValueOf(f).Pointer()
	fName := runtime.FuncForPC(pc).Name()

	start := time.Now()
	for i := 1; i <= n; i++ {
		f()
	}
	elapsed := time.Since(start).Nanoseconds() / int64(n)

	os.Stdout = stdout
	fmt.Printf("%v : %v nano seconds elapsed\n", fName, elapsed)
}

func exo3() {
	benchmark(echo1, 30)
	benchmark(echo2, 30)
	benchmark(echo3, 30)
}

func main() {
	exo3()
}
