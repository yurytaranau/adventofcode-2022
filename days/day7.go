package days

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type dir struct {
	size int
	name string
	path []string
}

type cursor []dir

func (c *cursor) Push(v string) {
	d := dir{name: v, path: c.getPath()}
	*c = append(*c, d)
}

func (c *cursor) getPath() []string {
	path := []string{}
	temp := *c
	for _, t := range temp {
		path = append(path, t.name)
	}
	return path
}

func (c *cursor) Pop() dir {
	l := len(*c)
	new := *c
	pop := new[l-1]
	// pop.path = new.getPath()
	new = new[:l-1]
	*c = new
	return pop
}

func (c *cursor) AddSize(s int) {
	new := *c
	for i := range *c {
		new[i].size += s
	}
	*c = new
}

func parseOutput(fs cursor, s string) (cursor, []dir) {
	isCommand := func(s string) bool {
		return strings.HasPrefix(s, "$ ")
	}

	getCommand := func(s string) []string {
		return strings.Split(s, " ")[1:]
	}
	isFile := func(s string) bool {
		matched, err := regexp.MatchString("[0-9]+ [a-zA-Z.]+", s)
		if err != nil {
			panic("regex is faulty")
		}
		return matched
	}
	getFileSize := func(s string) int {
		size, err := strconv.Atoi(strings.Split(s, " ")[0])
		if err != nil {
			panic("cannot get size")
		}
		return size
	}

	var scannedDirs []dir
	switch true {
	case isCommand(s):
		cmd := getCommand(s)
		switch cmd[0] {
		case "cd":
			if cmd[1] == ".." {
				scannedDir := fs.Pop()
				scannedDirs = append(scannedDirs, scannedDir)
			} else {
				fs.Push(cmd[1])
			}
		case "ls":
			// do nothing
		}
	case isFile(s):
		size := getFileSize(s)
		fs.AddSize(size)
	}

	return fs, scannedDirs
}

func (d Days) Day7p1(source io.Reader) (sop int) {
	var fs cursor
	var ackFolders []dir
	var ackFolderIter []dir
	lines := bufio.NewScanner(source)
	for lines.Scan() {
		fs, ackFolderIter = parseOutput(fs, lines.Text())
		ackFolders = append(ackFolders, ackFolderIter...)
	}
	var result int
	ackFolders = append(ackFolders, fs...)
	for _, f := range ackFolders {
		if f.size <= 100000 {
			result += f.size
		}
	}
	fmt.Printf("(Part 1) Sum: %d\n", result)
	return result
}

func (d Days) Day7p2(source io.Reader) (som int) {
	var fs cursor
	var ackFolders []dir
	var ackFolderIter []dir
	lines := bufio.NewScanner(source)
	for lines.Scan() {
		fs, ackFolderIter = parseOutput(fs, lines.Text())
		ackFolders = append(ackFolders, ackFolderIter...)
	}
	ackFolders = append(ackFolders, fs...)

	var targetSize int = 30000000
	var fileSystemSize int = 70000000

	foldersSize := map[string]int{}
	for _, f := range ackFolders {
		foldersSize[f.name] = f.size
	}

	needToDelete := targetSize - (fileSystemSize - foldersSize["/"])
	smallestDirSize := foldersSize["/"]
	for _, size := range foldersSize {
		if size <= smallestDirSize && size >= needToDelete {
			smallestDirSize = size
		}
	}
	fmt.Printf("(Part 2) Smallest dir to delete: %d\n", smallestDirSize)
	return smallestDirSize
}
