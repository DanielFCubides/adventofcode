package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Parent   *Node
	Children map[string]*Node
	Size     int
	Name     string
	Type     string
}

func ReadCommand(command string, node *Node) *Node {
	args := strings.Split(command, " ")
	if args[0] == "$" {
		if args[1] == "cd" {
			if args[2] == ".." {
				return node.Parent
			} else {
				return node.Children[args[2]]
			}
		}
	} else {
		if args[0] == "dir" {
			node.Children[args[1]] = &Node{
				Parent:   node,
				Children: map[string]*Node{},
				Size:     0,
				Name:     args[1],
				Type:     "dir",
			}
		} else {
			value, _ := strconv.Atoi(args[0])
			//fmt.Printf("creating file child %s value %d\n", args[1], value)
			child := &Node{
				Parent:   node,
				Children: map[string]*Node{},
				Size:     value,
				Name:     args[1],
				Type:     "file",
			}
			node.Children[args[1]] = child
			addValue(node, value)
		}

	}

	return node
}

func addValue(node *Node, value int) {
	parent := node.Parent
	for parent != nil {
		//fmt.Printf("node %s = %d + %d\n", parent.Name, parent.Size, value)
		parent.Size += value
		parent = parent.Parent
	}
	//fmt.Printf("node %s = %d + %d\n", node.Name, node.Size, value)
	node.Size += value
}
func readFile(filename string) Node {
	root := Node{
		Parent:   nil,
		Children: map[string]*Node{},
		Size:     0,
		Name:     "/",
		Type:     "dir",
	}
	//fmt.Printf("Address of node=\t%p\n", &root)
	readFile, err := os.Open(filename)
	if err != nil {
		panic("no file")
	}
	fileScanner := bufio.NewScanner(readFile)
	line := 0
	node := &root
	//fmt.Printf("Address of node=\t%p\n", node)
	for fileScanner.Scan() {
		if line != 0 {
			node = ReadCommand(fileScanner.Text(), node)
		}
		line++

	}
	return root
}

func printNode(node *Node, output *int) {
	fmt.Printf("node %s : ", node.Name)
	if node.Size < 100000 {
		fmt.Printf("value %d added to output %d, output now is : %d", node.Size, *output, *output+node.Size)
		*output = *output + node.Size
	}
	for _, child := range node.Children {
		if child.Type == "dir" {
			printNode(child, output)
		}
	}
	fmt.Println(" - ")
}

func main() {
	puzzle := readFile("input.txt")
	fmt.Println("finish reading")
	result := 0
	printNode(&puzzle, &result)
	fmt.Println(result)

}
