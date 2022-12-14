package D07

import (
	"bufio"
	"fmt"
	"io"
	"sort"
)

type Node struct {
	prev  *Node
	leafs []*Node
	size  int
	name  string
}

func (n Node) isDir() bool {
	return !(n.size > 0)
}

func (n Node) getAllDirSizes() []int {
	var dirSizes []int

	sum := 0
	for _, leaf := range n.leafs {
		if leaf.isDir() {
			leafDirSizes := leaf.getAllDirSizes()
			sum += leafDirSizes[len(leafDirSizes)-1]
			dirSizes = append(dirSizes, leafDirSizes...)
		} else {
			sum += leaf.size
		}
	}

	return append(dirSizes, sum)
}

func NoSpaceLeftOnDevicePart1(input io.Reader) int {
	scanner := bufio.NewScanner(input)

	root := buildTree(scanner)
	dirSizes := root.getAllDirSizes()
	sum := 0
	for _, size := range dirSizes {
		if size < 100000 {
			sum += size
		}
	}

	return sum
}

func NoSpaceLeftOnDevicePart2(input io.Reader) int {
	scanner := bufio.NewScanner(input)

	root := buildTree(scanner)
	dirSizes := root.getAllDirSizes()

	sort.Ints(dirSizes)
	spaceHasToBeDeleted := 30000000 - (70000000 - dirSizes[len(dirSizes)-1])
	for _, size := range dirSizes {
		if size > spaceHasToBeDeleted {
			return size
		}
	}

	return 0
}

func buildTree(scanner *bufio.Scanner) Node {
	scanner.Split(bufio.ScanLines)

	var root Node
	var cur *Node
	scanner.Scan()
	if scanner.Text() == "$ cd /" {
		root = Node{name: "/"}
		cur = &root
	}

	for scanner.Scan() {
		commandString := scanner.Text()
		if commandString[0] == '$' {
			if commandString[2:4] == "cd" {
				arg := commandString[5:]
				if arg == ".." {
					cur = cur.prev
				} else {
					newNode := Node{name: arg, prev: cur}
					cur.leafs = append(cur.leafs, &newNode)
					cur = &newNode
				}
			}
			// ignore ls command
		} else {
			if commandString[:3] == "dir" {
				continue
			}
			var size int
			var name string
			_, err := fmt.Sscanf(commandString, "%d %s", &size, &name)
			if err != nil {
				panic(err)
			}
			fileNode := Node{name: name, size: size, prev: cur}
			cur.leafs = append(cur.leafs, &fileNode)
		}
	}
	return root
}
