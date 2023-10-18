package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

type Directory struct {
	parent   *Directory
	children map[string]*Directory
	files    map[string]int
	name     string
	filesize int
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("could not get cwd")
		os.Exit(2)
	}

	data, err := os.ReadFile(filepath.Join(cwd, "inputs", "day7.txt"))
	if err != nil {
		fmt.Println("could not read file")
		os.Exit(2)
	}

	lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
	root := Directory{nil, make(map[string]*Directory), make(map[string]int), "/", 0}
	currentDir := &root

	for _, val := range lines[1:] {
		if len(val) == 0 {
			break
		}
		fields := strings.Fields(val)

		if fields[1] == "cd" {
			if fields[2] == ".." {
				currentDir = currentDir.parent
			} else {
				currentDir = currentDir.children[fields[2]]
			}
		}

		if fields[0] != "$" {
			size, err := strconv.Atoi(fields[0])
			if err != nil { // dir
				newDir := Directory{
					currentDir,
					make(map[string]*Directory),
					make(map[string]int),
					fields[1],
					0,
				}
				currentDir.children[fields[1]] = &newDir
			} else { // file
				currentDir.files[fields[1]] = size
			}
		}
	}
	calculateFilesizes(&root)

	// part 1
	maxSize := 100000
	fmt.Println(sumOfSizes(&root, maxSize))

	totalDiskSpace := 70000000
	sizeNeededForUpdate := 30000000
	totalSize := root.filesize
	// i need to subtract
	freeSpace := totalDiskSpace - totalSize
	needToFree := sizeNeededForUpdate - freeSpace

	dirSizes := make([]int, 0)
	dirsToSizes(&root, &dirSizes, needToFree)
	fmt.Println(slices.Min(dirSizes))
}

// appends directory file size to array if it is larger than needToFree
func dirsToSizes(root *Directory, arr *[]int, needToFree int) {
	if root.filesize > needToFree {
		*arr = append(*arr, root.filesize)
	}
	for _, child := range root.children {
		dirsToSizes(child, arr, needToFree)
	}
}

func calculateFilesizes(root *Directory) {
	var calculate func(node *Directory) int
	calculate = func(node *Directory) int {
		totalSize := 0
		for _, size := range node.files {
			totalSize += size
		}
		for _, child := range node.children {
			totalSize += calculate(child)
		}
		node.filesize = totalSize
		return totalSize
	}
	calculate(root)
}

// Helper function to print a tree recursively
func printTree(root *Directory) {
	var print func(node *Directory, indent string)
	print = func(node *Directory, indent string) {
		fmt.Printf("%s- %s (dir, size=%d)\n", indent, node.name, node.filesize)
		for _, child := range node.children {
			print(child, indent+" ")
		}

		for filename, size := range node.files {
			fmt.Printf("%s> %s (file, size = %d)\n", indent, filename, size)
		}
	}

	print(root, "")
}

func sumOfSizes(root *Directory, maxSize int) int {
	totalSize := 0
	if root.filesize < maxSize {
		totalSize += root.filesize
	}
	for _, child := range root.children {
		totalSize += sumOfSizes(child, maxSize)
	}
	return totalSize
}
