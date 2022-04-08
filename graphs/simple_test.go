package graphs

import (
	"fmt"
	"math/rand"
	"testing"
)

func randomMapValue(src Nodes) *Node {
	for _, n := range src {
		return n
	}
	return nil
}

const maxPower = 9001

// randomGraph generates random graph and
func randomGraph(t *testing.T, size int) (Graph, Nodes) {
	t.Helper()

	nodes := make(Nodes, size)
	for i := 0; i < size; i++ {
		n := &Node{
			Name:  fmt.Sprintf("node-%d", i),
			Power: rand.Intn(maxPower),
		}
		nodes[n.Name] = n
	}
	allNodes := make(Graph, size)
	for _, n := range nodes {
		numNeighbour := rand.Intn(size-1) + 1
		neighbours := make(Neighbours, numNeighbour)
		for i := 0; i < numNeighbour; {
			v := randomMapValue(nodes)
			if _, ok := neighbours[v.Name]; !ok {
				neighbours[v.Name] = nodes[v.Name]
				i++
			}
		}
		allNodes[n.Name] = neighbours
	}

	if size <= 10 {
		t.Logf("Created graph: %+v", allNodes)
	}
	return allNodes, nodes
}

func TestRandomGraph(t *testing.T) {
	randomGraph(t, 5)
}

func TestFindNode(t *testing.T) {
	testGraph, nodeCache := randomGraph(t, 10)

	t.Run("very powerful", func(t *testing.T) {
		midPower := maxPower / 10
		minPower := rand.Intn(midPower) + midPower
		t.Logf("Will search for any node with power >= %d", minPower)
		startNode := randomMapValue(nodeCache) // cheeeeat!
		t.Logf("Related to node %s", startNode.Name)
		node, err := FindNode(testGraph, startNode, func(n *Node) bool {
			return n.Power >= minPower
		})
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("Local powerful node: %s(%+v)", node.Name, node)
	})
	t.Run("second one", func(t *testing.T) {
		startNode := randomMapValue(nodeCache) // cheeeeat!
		targetName := "node-1"

		t.Logf("Related to node %s", startNode.Name)
		node, err := FindNode(testGraph, startNode, func(n *Node) bool {
			return n.Name == targetName
		})
		if err != nil {
			t.Fatal(err)
		}
		if node != nil {
			t.Logf("There is path from %s to %s", startNode.Name, targetName)
		}
	})
}
