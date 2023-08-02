package main

import (
	"bufio"
	"fmt"
	"log"
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
	fmt.Printf("node=\t%p\tparent=%p\n", node, node.Parent)
	args := strings.Split(command, " ")
	if args[0] == "$" {
		if args[1] == "cd" {
			if args[2] == ".." {
				log.Printf("moving to parent %s \n", node.Parent.Name)
				return node.Parent
			} else {
				log.Printf("moving to child %s \n", node.Children[args[2]].Name)
				return node.Children[args[2]]
			}
		}
	} else {
		if args[0] == "dir" {
			log.Printf("creating dir child %s \n", args[1])
			node.Children[args[1]] = &Node{
				Parent:   node,
				Children: map[string]*Node{},
				Size:     0,
				Name:     args[1],
				Type:     "dir",
			}
		} else {
			value, _ := strconv.Atoi(args[0])
			log.Printf("creating file child %s \n", args[1])
			child := &Node{
				Parent:   node,
				Children: map[string]*Node{},
				Size:     value,
				Name:     args[1],
				Type:     "file",
			}
			node.Children[args[1]] = child
			node.Size += value
		}

	}

	return node
}
func readFile(filename string) Node {
	root := Node{
		Parent:   nil,
		Children: map[string]*Node{},
		Size:     0,
		Name:     "/",
		Type:     "dir",
	}
	fmt.Printf("Address of node=\t%p\n", &root)
	readFile, err := os.Open(filename)
	if err != nil {
		panic("no file")
	}
	fileScanner := bufio.NewScanner(readFile)
	line := 0
	node := &root
	fmt.Printf("Address of node=\t%p\n", node)
	for fileScanner.Scan() {
		if line != 0 {
			node = ReadCommand(fileScanner.Text(), node)
		}
		line++

	}
	return *node
}

func main() {
	puzzle := readFile("example.txt")
	log.Printf("%s - %d", puzzle.Name, puzzle.Size)

}
