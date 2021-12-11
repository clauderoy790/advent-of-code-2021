package main

import (
	"fmt"
	"os"
	"path"
	"runtime"
)

func main() {
	for i := 13; i <= 25; i++ {
		writeFile(i)
	}
}

func writeFile(day int) {
	_, cmdFolder, _, _ := runtime.Caller(0)
	cmdFolder = path.Join(cmdFolder, "../..")
	dir := path.Join(cmdFolder, fmt.Sprintf("/day%v", day))
	fmt.Println("DIR: ", dir)
	err := os.Mkdir(dir, 0755)
	if err != nil {
		panic("error creating dir")
	}

	m := path.Join(dir, "main.go")
	fmt.Println("main file: ", m)
	_, err = os.Create(m)
	if err != nil {
		panic("error creating main file")
	}

	content := []byte(`package main

func main() {
	
}
	`)
	err = os.WriteFile(m, content, 0644)
	if err != nil {
		panic("failed to write file content")
	}

	_, err = os.Create(path.Join(dir, "input.txt"))

	if err != nil {
		panic("error creating input file")
	}
}
