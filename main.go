package main

import (
	"fmt"

	"github.com/khelben/fnimport/dto"
	"github.com/khelben/fnimport/importers"
)

func main() {
	// pois, err := importers.ReadFietsNetPois()
	// if err != nil {
	// 	panic(err)
	// }
	// importers.ImportFietsNetNodes(pois.Pois)

	rtrees, err := importers.ReadFietsNetRtrees()
	if err != nil {
		panic(err)
	}
	readTracks(rtrees.RootNode.Children)
	// fmt.Printf("%#v\n", rtrees)
}

func readTracks(nodes []dto.Node) error {
	for _, node := range nodes {
		if node.Text != "" {
			fmt.Printf("Track: %s\n", node.Text)
		}
		if len(node.Children) > 0 {
			readTracks(node.Children)
		}
	}
	return nil
}
