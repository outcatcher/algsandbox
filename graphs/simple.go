package graphs

import (
	"container/list"
	"errors"
	"fmt"
)

type Node struct {
	Name  string
	Power int
}

func (n *Node) String() string {
	return fmt.Sprintf("%d", n.Power)
}

type Neighbours map[string]*Node

type Graph map[string]Neighbours

type Nodes map[string]*Node

var (
	ErrInvalidQueryElem = errors.New("query element is not a string")
)

type ErrNodeNotFound struct {
	NodeName string
}

func (e ErrNodeNotFound) Error() string {
	return fmt.Sprintf("node %s is not found in the graph", e.NodeName)
}

type CheckFunc func(*Node) bool

func FindNode(g Graph, start *Node, checkFunc CheckFunc) (*Node, error) {
	nodeQueue := list.New() // queue of names
	nodeQueue.PushBack(start)
	currNode := nodeQueue.Front()
	for nodeQueue.Len() != 0 {
		if currNode == nil {
			break // empty queue
		}
		nextNode, ok := currNode.Value.(*Node)
		if !ok {
			return nil, ErrInvalidQueryElem
		}
		if checkFunc(nextNode) {
			return nextNode, nil
		}
		neighbours, ok := g[nextNode.Name]
		if !ok {
			return nil, ErrNodeNotFound{NodeName: nextNode.Name}
		}
		for _, node := range neighbours {
			if checkFunc(node) {
				return node, nil
			}
			nodeQueue.PushBack(node)
		}
		currNode = currNode.Next()
	}
	return nil, nil
}
