package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val    int
	Left   *TreeNode
	Right  *TreeNode
	Parent *TreeNode
}

func kthSmallest(numNodes int, root *TreeNode, k int) int {

	var smallestNode *TreeNode
	ks := 1

	if numNodes > root.Right.Val {
		smallestNode = root.Right
	} else {
		smallestNode = root.Left
	}

	for smallestNode.Left != nil && smallestNode.Left.Val != -1 {
		smallestNode = smallestNode.Left
	}

	for ks < k {
		fmt.Printf("Smallest Node: %+v\n", smallestNode)
		smallestNode = smallestNode.Parent
		ks++
	}

	return smallestNode.Val

}

func parseInput(input string) ([]int, int) {
	parts := strings.Split(input, ", k = ")
	treePart := strings.TrimPrefix(parts[0], "root = [")
	treePart = strings.TrimSuffix(treePart, "]")

	var treeArr []int
	for _, val := range strings.Split(treePart, ",") {
		val = strings.TrimSpace(val)
		if val == "null" {
			treeArr = append(treeArr, -1)
		} else {
			num, _ := strconv.Atoi(val)
			treeArr = append(treeArr, num)
		}
	}

	k, _ := strconv.Atoi(strings.TrimSpace(parts[1]))

	return treeArr, k
}

func buildTree(arr []int) (*TreeNode, int) {
	if len(arr) == 0 || arr[0] == -1 {
		return nil, -1
	}
	length := 0
	root := &TreeNode{Val: arr[0], Parent: nil}
	queue := []*TreeNode{root}
	index := 1

	for index < len(arr) {
		node := queue[0]
		queue = queue[1:]

		if index < len(arr) && arr[index] != -1 {
			node.Left = &TreeNode{Val: arr[index], Parent: node}
			length++
			queue = append(queue, node.Left)
		}
		index++

		if index < len(arr) && arr[index] != -1 {
			node.Right = &TreeNode{Val: arr[index], Parent: node}
			queue = append(queue, node.Right)
			length++
		}
		index++

	}

	return root, length
}

func inorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	inorderTraversal(root.Left)
	fmt.Print(root.Val, " ")
	inorderTraversal(root.Right)
}

func main() {

	input := "root = [3,1,4,null,2], k = 1"

	treeArr, k := parseInput(input)

	root, size := buildTree(treeArr)

	fmt.Print("Inorder Traversal: ")
	inorderTraversal(root)
	fmt.Println()

	fmt.Printf("Value of k: %d\n\n", k)

	fmt.Println(kthSmallest(size, root, k))
}
