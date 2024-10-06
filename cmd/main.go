package main

import (
	"fmt"
	"os"
	"path"
	"sort"
)

type Counter struct {
	dirs  int
	files int
}

var (
	reset  = "\033[0m"
	yellow = "\033[33;1m"
	green  = "\033[32;1m"
)

func (counter *Counter) index(path string) {
	stat, _ := os.Stat(path)
	if stat.IsDir() {
		counter.dirs += 1
	} else {
		counter.files += 1
	}
}

func (counter *Counter) output() string {
	return fmt.Sprintf("\n%d directories, %d files", counter.dirs, counter.files)
}

func dirnamesFrom(base string) []string {
	file, err := os.Open(base)
	if err != nil {
		fmt.Println(err)
	}

	names, _ := file.Readdirnames(0)
	file.Close()

	sort.Strings(names)
	return names
}

func tree(counter *Counter, base string, prefix string) {
	names := dirnamesFrom(base)
	//print(names)

	for index, name := range names {
		if name[0] == '.' {
			continue
		}
		subpath := path.Join(base, name)
		counter.index(subpath)

		if index == len(names)-1 {
			fmt.Println(prefix+yellow+"  └──" , name+reset)
			tree(counter, subpath, prefix+"    ")
		} else {
			fmt.Println(prefix+"  ├──", name)
			tree(counter, subpath, prefix+"  │   ")
		}
	}
}

func main() {
	var directory string
	if len(os.Args) > 1 {
		directory = os.Args[1]
	} else {
		dir, _ := os.Getwd()
		fmt.Println()
		print(green+" ", dir+reset)

		directory = "."
	}

	counter := new(Counter)
	fmt.Println(directory)

	tree(counter, directory, "")
	fmt.Println(counter.output())
}

// func check(base string) (bool) {
//     dir, err := os.Stat(base)
//     if err != nil {
//         print("No working directory.")
//     }

//     if dir.Mode().IsDir() {
//         return false
//     }
//     
//     return true
// }
