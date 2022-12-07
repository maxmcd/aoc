package main

import (
	"fmt"
	"strings"

	"github.com/maxmcd/aoc"
)

type File struct {
	Name    string
	Size    int
	IsDir   bool
	Entries []*File
	Parent  *File
}

func print(files []*File, depth int) {
	for _, file := range files {
		var sb strings.Builder
		fmt.Fprint(&sb, strings.Repeat(" ", depth), "- ", file.Name)
		if file.IsDir {
			fmt.Fprint(&sb, " (dir)")
		} else {
			fmt.Fprintf(&sb, " (file, size=%d)", file.Size)
		}
		fmt.Println(sb.String())
		print(file.Entries, depth+1)
	}
}

func size(file *File, cb func(int)) int {
	return _size(file.Name, file.Entries, cb)
}

var dirTotal int

func _size(name string, files []*File, cb func(int)) int {
	total := 0
	for _, file := range files {
		if file.IsDir {
			dirSize := _size(file.Name, file.Entries, cb)
			fmt.Println("sizeof", file.Name, dirSize)
			total += dirSize
			cb(dirSize)
		} else {
			total += file.Size
		}
	}
	return total
}

func main() {
	fileTree := &File{Name: "/", IsDir: true}
	cd := fileTree
	_ = cd
	aoc.ReadInput(func(line string) {
		switch line {
		case "$ ls", "$ cd /":
			// Skip, we infer these
			return
		}
		var name string
		if _, err := fmt.Sscanf(line, "dir %s", &name); err == nil {
			cd.Entries = append(cd.Entries, &File{
				Name:   name,
				IsDir:  true,
				Parent: cd,
			})
			return
		}
		var size int
		if _, err := fmt.Sscanf(line, "%d %s", &size, &name); err == nil {
			cd.Entries = append(cd.Entries, &File{
				Name:   name,
				Size:   size,
				Parent: cd,
			})
			return
		}
		var path string
		if _, err := fmt.Sscanf(line, "$ cd %s", &path); err == nil {
			if path == ".." {
				cd = cd.Parent
			} else {
				for _, child := range cd.Entries {
					if child.Name == path {
						cd = child
						return
					}
				}
				panic(fmt.Sprint(line, path, cd.Entries))
			}
			return
		}
		panic(fmt.Sprintf("%q", line))
	})
	print([]*File{fileTree}, 0)

	freeAtLeast := 30000000
	totalDiskSpace := 70000000
	// 1
	var dirTotal int
	totalSize := size(fileTree, func(dirSize int) {
		if dirSize <= 100000 {
			dirTotal += dirSize
		}
	})
	fmt.Println(1)
	fmt.Println("dirTotal", dirTotal)
	fmt.Println()
	fmt.Println(2)
	fmt.Println("totalSize", totalSize)
	unusedSpace := totalDiskSpace - totalSize
	fmt.Println("unused", unusedSpace)

	smallest := totalSize
	// 2
	size(fileTree, func(dirSize int) {
		if unusedSpace+dirSize >= freeAtLeast {
			if dirSize < smallest {
				smallest = dirSize
			}
		}
	})
	fmt.Println(smallest)

}
