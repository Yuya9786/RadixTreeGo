package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const ValueBits int = 32

type RadixNode struct {
	valid uint8
	value uint32
	depth uint8
	count uint64
	children [2]*RadixNode
}

func NewRadixNode(depth uint8) *RadixNode {
	return &RadixNode{
		valid: 0,
		value: 0,
		depth: depth,
		count: 0,
		children: [2]*RadixNode{nil, nil},
	}
}

func InsertRadixNode(radixnode *RadixNode, value uint32, depth uint8) {
	if depth == radixnode.depth {
		radixnode.valid = 1
		radixnode.value = value
	} else {
		child := (value << radixnode.depth) >> (ValueBits - 1)
		if radixnode.children[child] == nil {
			radixnode.children[child] = NewRadixNode(radixnode.depth + 1)
		}
		InsertRadixNode(radixnode.children[child], value, depth)
	}
	radixnode.count++
}

func SearchRadixNode(radixnode *RadixNode, value uint32) (uint64, bool) {
	var lastValidNode *RadixNode = nil
	node := radixnode
	indexPlusList := []uint64{0}

	for depth := 0; depth < ValueBits; depth++ {
		child := (value << depth) >> (ValueBits - 1)
		if node.children[child] == nil {
			break
		}

		v := uint64(node.valid)
		if child == 1 && node.children[0] != nil {
			v += node.children[0].count
		}

		indexPlusList = append(indexPlusList, v + indexPlusList[len(indexPlusList)-1])

		node = node.children[child]
		if node.valid == 1 {
			lastValidNode = node
		}
	}
	if lastValidNode != nil {
		return indexPlusList[lastValidNode.depth], true
	} else {
		return 0, false
	}
}

func addrStr2uint32(addrStr string) uint32 {
	slice := strings.Split(addrStr, ".")
	var addrUint uint32 = 0
	for i, v := range slice {
		x, _ := strconv.Atoi(v)
		addrUint |= uint32(x) << ((3-i) * 8)
	}
	return addrUint
}

func main() {
	radixTree := NewRadixNode(0)

	// read
	data, _ := os.Open("route-02.txt")
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := scanner.Text()
		addr := strings.Split(line, "/")

		prefixLength, _ := strconv.Atoi(addr[1])
		a := addrStr2uint32(addr[0])
		// fmt.Printf("%d\n", a)
		InsertRadixNode(radixTree, a, uint8(prefixLength))
	}

	answer, _ := SearchRadixNode(radixTree, addrStr2uint32("41.74.1.1"))
	fmt.Printf("%d\n", answer)
	answer, _ = SearchRadixNode(radixTree, addrStr2uint32("66.31.10.3"))
	fmt.Printf("%d\n", answer)
	answer, _ = SearchRadixNode(radixTree, addrStr2uint32("133.5.1.1"))
	fmt.Printf("%d\n", answer)
	answer, _ = SearchRadixNode(radixTree, addrStr2uint32("209.143.75.1"))
	fmt.Printf("%d\n", answer)
	answer, _ = SearchRadixNode(radixTree, addrStr2uint32("221.121.128.1"))
	fmt.Printf("%d\n", answer)
}