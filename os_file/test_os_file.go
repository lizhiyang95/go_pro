package main

import (
	"fmt"
	"os"
)

func createFile() {
	f, err := os.Create("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("f.Name(): %v\n", f.Name())
}

func mkDir() {
	err := os.Mkdir("a", os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	err2 := os.MkdirAll("a/b/c", os.ModePerm)
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
	}
}

func remove() {
	err := os.Remove("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	err2 := os.RemoveAll("a")
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
	}
}

func workDir() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("dir: %v\n", dir)

	err2 := os.Chdir("d:/")
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
	}
	dir2, _ := os.Getwd()
	fmt.Printf("dir2: %v\n", dir2)

	s := os.TempDir()
	fmt.Printf("s: %v\n", s)
}

func renameFile() {
	err := os.Rename("a.txt", "b.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func readFile() {
	b, err := os.ReadFile("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	s := string(b)
	fmt.Printf("s: %v\n", s)

}

func writeFile() {
	var s = "hello word"
	b := []byte(s)
	os.WriteFile("a.txt", b, os.ModePerm)
}

func write() {
	u := "hello golang"
	b := []byte(u)
	f, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_APPEND, os.ModePerm)
	n, _ := f.Write(b)
	fmt.Printf("n: %v\n", n)
}

func writeString() {
	u := "hello java"
	f, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_APPEND, os.ModePerm)
	n, _ := f.WriteString(u)
	fmt.Printf("n: %v\n", n)
}

func writeAt() {
	f, _ := os.OpenFile("a.txt", os.O_RDWR, os.ModePerm)
	u := "hello pyhon"
	b := []byte(u)
	n, _ := f.WriteAt(b, 4)
	fmt.Printf("n: %v\n", n)
}

func openCloseFile() {
	//不存在报错
	/* 	f, err := os.Open("a.txt")
	   	if err != nil {
	   		fmt.Printf("err: %v\n", err)
	   	}
	   	fmt.Printf("f.Name(): %v\n", f.Name())

	   	f.Close() */

	//不存在创建
	f, err := os.OpenFile("a.txt", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("f.Name(): %v\n", f.Name())
	f.Close()
}

func createTemp() {
	f, _ := os.CreateTemp("", "prefix")
	fmt.Printf("f.Name(): %v\n", f.Name())
}

func readOps() {
	f, _ := os.Open("a.txt")
	//buf := make([]byte, 3)
	/* 	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			fmt.Println("文件末尾")
			break
		}
		fmt.Printf("n: %v\n", n)
		fmt.Printf("buf: %v\n", buf)
	} */

	/* 	buf := make([]byte, 20)
	   	fmt.Printf("buf: %T\n", buf)
	   	n, _ := f.ReadAt(buf, 5)
	   	fmt.Printf("n: %v\n", n)
	   	s := string(buf)
	   	fmt.Printf("s: %v\n", s) */

	ret, _ := f.Seek(3, 0)
	fmt.Printf("ret: %v\n", ret)
	//b := []byte{}
	var b []byte
	fmt.Printf("(b == nil): %v\n", (b == nil))
	fmt.Printf("b: %v\n", b)
	n, _ := f.Read(b)
	fmt.Printf("n: %v\n", n)
	fmt.Printf("string(b): %v\n", string(b))

}

func readDir() {
	de, _ := os.ReadDir("./")
	for _, v := range de {
		fmt.Printf("v.IsDir(): %v\n", v.IsDir())
		fmt.Printf("v.Name(): %v\n", v.Name())
	}
}

func main() {
	//createFile()
	// writeFile()
	// readFile()
	//openCloseFile()
	//createTemp()
	//readOps()
	//readDir()
	workDir()
	// write()
	//writeString()
	//writeAt()
}
