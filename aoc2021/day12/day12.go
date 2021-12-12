package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

type Edge struct {
	start, end string
}

func parseInput() []Edge {
	data, err := ioutil.ReadFile("./input1.txt")
	if err != nil {
		fmt.Println("Failed to read file..")
	}
	inputArr := strings.Split(string(data), "\n")

	var edges []Edge
	for _, line := range inputArr {
		if line == "" {
			continue
		}
		parts := strings.Split(line, "-")
		edges = append(edges, Edge{start: parts[0], end: parts[1]})
		if parts[0] != "start" {
			edges = append(edges, Edge{start: parts[1], end: parts[0]})
		}
	}
	return edges
}

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func getConnectingEdges(edges []Edge, pos string) []Edge {
	var connectingEdges []Edge
	for _, edge := range edges {
		if edge.start == pos {
			connectingEdges = append(connectingEdges, edge)
		}
	}
	return connectingEdges
}

func hasVisited(visitedSmallNodes []string, pos string) bool {
	for _, visitedSmallNode := range visitedSmallNodes {
		if visitedSmallNode == pos {
			return true
		}
	}
	return false
}

func hasVisitedMostTwice(smallCaveTwice string, visitedSmallNodes []string, pos string) bool {
	sum := 0
	for _, v1 := range visitedSmallNodes {
		if v1 == pos {
			if v1 != smallCaveTwice {
				sum += 2
			} else {
				sum += 1
			}
		}
	}
	return sum > 1
}

func findAllSmallCaves(edges []Edge) []string {
	var smallCaves []string
	for _, edge := range edges {
		alreadyInserted := false
		for _, e2 := range smallCaves {
			if edge.start == e2 {
				alreadyInserted = true
				break
			}
		}
		if IsLower(edge.start) && edge.start != "start" && edge.start != "end" && !alreadyInserted {
			smallCaves = append(smallCaves, edge.start)
		}
	}
	return smallCaves
}

func traversePart2(edges []Edge, pos string, currentPath string, visitedSmallNodes []string, allPaths *[]string, smallCaveTwice string) string {
	if pos == "end" {
		newPath := fmt.Sprintf("%s,%s", currentPath, "end")
		*allPaths = append(*allPaths, newPath)
		return newPath
	}

	if IsLower(pos) {
		visitedSmallNodes = append(visitedSmallNodes, pos)
	}
	connectingEdges := getConnectingEdges(edges, pos)
	var paths string
	for _, connectedEdge := range connectingEdges {
		if hasVisitedMostTwice(smallCaveTwice, visitedSmallNodes, connectedEdge.end) {
			continue
		}
		paths = traversePart2(edges, connectedEdge.end, fmt.Sprintf("%s,%s", currentPath, connectedEdge.start), visitedSmallNodes, allPaths, smallCaveTwice)
	}
	return paths
}

func traverse(edges []Edge, pos string, currentPath string, visitedSmallNodes []string, allPaths *[]string) string {
	if pos == "end" {
		newPath := fmt.Sprintf("%s,%s", currentPath, "end")
		*allPaths = append(*allPaths, newPath)
		return newPath
	}
	if IsLower(pos) {
		visitedSmallNodes = append(visitedSmallNodes, pos)
	}
	connectingEdges := getConnectingEdges(edges, pos)
	var paths string
	for _, connectedEdge := range connectingEdges {
		if hasVisited(visitedSmallNodes, connectedEdge.end) {
			continue
		}
		paths = traverse(edges, connectedEdge.end, fmt.Sprintf("%s,%s", currentPath, connectedEdge.start), visitedSmallNodes, allPaths)
	}
	return paths
}

func part1() {
	edges := parseInput()

	var allPaths []string
	traverse(edges, "start", "", []string{}, &allPaths)

	fmt.Printf("Solution part 1: %v\n", len(allPaths))
}

func part2() {
	edges := parseInput()

	var allPaths []string

	// Try all different possibility of a small cave that can be visited
	// once
	for _, smallCave := range findAllSmallCaves(edges) {
		traversePart2(edges, "start", "", []string{}, &allPaths, smallCave)
	}

	// Remove duplicates
	var uniquePaths []string
	for _, p1 := range allPaths {
		hasPath := false
		for _, p2 := range uniquePaths {
			if p1 == p2 {
				hasPath = true
				break
			}
		}
		if !hasPath {
			uniquePaths = append(uniquePaths, p1)
		}
	}

	fmt.Printf("Solution part 2: %v\n", len(uniquePaths))
}

func main() {
	part1()
	part2()
}
