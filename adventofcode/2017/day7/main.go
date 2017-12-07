package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
	"sort"
)

func main() {
	tree := make(map[string]string)
	reversed := make(map[string][]string)
	weights := make(map[string]int)
	entrypoint := ""
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := strings.Fields(scanner.Text())
		weight, _ := strconv.Atoi(strings.Trim(strings.Trim(text[1], "("), ")"))
		weights[text[0]] = weight
		if len(text) > 3 {
			for i := 3; i < len(text); i++ {
				replace := strings.Trim(text[i], ",")
				tree[replace] = text[0]
				reversed[text[0]] = append(reversed[text[0]], replace)
				entrypoint = replace
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	root := parent(tree, entrypoint)
	node := subtree(reversed, weights, root)
	fmt.Println(newWeight(node))
}

func newWeight(node node) int {
	current := node
	difference := 0
	for {
		balance := current.difference()
		if balance == 0 {
			return current.original_weight - difference
		} else {
			difference = balance
			current = current.heavier()
		}
	}
}

type node struct {
	name            string
	original_weight int
	weight          int

	children []node
}

func (n *node) isBalanced() bool {
	first := n.children[0]
	last := n.children[len(n.children)-1]
	return first.weight == last.weight
}

func (n *node) difference() int {
	first := n.children[0]
	last := n.children[len(n.children)-1]
	return last.weight - first.weight
}

func (n *node) heavier() node {
	return n.children[len(n.children)-1]
}

func subtree(reversed map[string][]string, weights map[string]int, current string) node {
	next, ok := reversed[current]
	if !ok {
		return node{current, weights[current], weights[current], []node{}}
	}
	var children []node
	sum := weights[current]
	for _, n := range next {
		child := subtree(reversed, weights, n)
		children = append(children, child)
		sum += child.weight
	}
	sort.Slice(children, func(i, j int) bool {
		return children[i].weight < children[j].weight
	})
	return node{current, weights[current], sum, children}
}

func parent(tree map[string]string, last string) string {
	current := last
	for {
		parent, ok := tree[current]
		if !ok {
			break
		}
		current = parent
	}
	return current
}
