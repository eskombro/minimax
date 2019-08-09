package minimax

import (
	"fmt"

	gr "github.com/eskombro/minimax/graph"
)

const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxValuesFromChildren(node *gr.Node) {
	selected := 0
	node.Value = MinInt
	for i, ch := range node.Children {
		if ch.UpdatedValue > node.Children[selected].UpdatedValue {
			selected = i
		}
	}
	node.UpdatedValue = node.Children[selected].UpdatedValue
	node.SelectedChild = node.Children[selected]
}

func minValuesFromChildren(node *gr.Node) {
	selected := 0
	node.Value = MaxInt
	for i, ch := range node.Children {
		if ch.UpdatedValue < node.Children[selected].UpdatedValue {
			selected = i
		}
	}
	node.UpdatedValue = node.Children[selected].UpdatedValue
	node.SelectedChild = node.Children[selected]
}

func MinimaxRecursive(graph *gr.Node, depth int, max bool, start bool) {
	for _, node := range graph.Children {
		if depth >= 0 {
			MinimaxRecursive(node, depth-1, !max, false)
			if len(node.Children) == 0 || depth == 0 {
				node.UpdatedValue = node.Value

			} else {
				if max {
					maxValuesFromChildren(node)
				} else {
					minValuesFromChildren(node)
				}
			}
		}
	}
	if start {
		maxValuesFromChildren(graph)
	}
}

func maxValuesFromChildrenPruning(node *gr.Node, alpha int, beta int) {
	selected := 0
	node.Value = MinInt
	for i, ch := range node.Children {
		if ch.UpdatedValue > node.Children[selected].UpdatedValue {
			selected = i
			alpha = maximum(alpha, ch.UpdatedValue)
			if alpha >= beta {
				break
			}
		}
	}
	node.UpdatedValue = node.Children[selected].UpdatedValue
	node.SelectedChild = node.Children[selected]
}

func minValuesFromChildrenPruning(node *gr.Node, alpha int, beta int) {
	selected := 0
	node.Value = MaxInt
	for i, ch := range node.Children {
		if ch.UpdatedValue < node.Children[selected].UpdatedValue {
			selected = i
			beta = min(beta, ch.UpdatedValue)
			if alpha <= beta {
				break
			}
		}
	}
	node.UpdatedValue = node.Children[selected].UpdatedValue
	node.SelectedChild = node.Children[selected]
}

func MinimaxRecursivePruning(graph *gr.Node, depth int, max bool, start bool, alpha int, beta int) {
	for _, node := range graph.Children {
		if depth >= 0 {
			MinimaxRecursivePruning(node, depth-1, !max, false, alpha, beta)
			if len(node.Children) == 0 || depth == 0 {
				node.UpdatedValue = node.Value

			} else {
				if max {
					maxValuesFromChildrenPruning(node, alpha, beta)

				} else {
					minValuesFromChildrenPruning(node, alpha, beta)
				}
			}
		}
	}
	if start {
		maxValuesFromChildren(graph)
	}
}

func LaunchMinimax(graph *gr.Node, depth int) {
	fmt.Println("Launching Minimax with depth ", depth)
	graph.UpdatedValue = graph.Value
	MinimaxRecursive(graph, depth, false, true)
	fmt.Println("==========================")
	fmt.Println("Final value:", graph.UpdatedValue)
}

func LaunchMinimaxPruning(graph *gr.Node, depth int) {
	alpha := MinInt
	beta := MaxInt
	fmt.Println("Launching Minimax Pruning with depth ", depth)
	graph.UpdatedValue = graph.Value
	MinimaxRecursivePruning(graph, depth, false, true, alpha, beta)
	fmt.Println("==========================")
	fmt.Println("Final value:", graph.UpdatedValue)
}
