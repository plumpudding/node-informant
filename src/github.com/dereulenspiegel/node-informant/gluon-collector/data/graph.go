package data

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
)

type GraphGenerator struct {
	Store *SimpleInMemoryStore
}

func FindInLinks(links []*GraphLink, sourceIndex, targetIndex int) (link *GraphLink, err error) {
	for _, item := range links {
		if item.Source == sourceIndex && item.Target == targetIndex {
			link = item
			err = nil
			return
		}
	}
	err = fmt.Errorf("Link not found")
	return
}

func (g *GraphGenerator) GenerateGraphJson() GraphJson {
	nodeTable := make(map[string]*GraphNode)
	log.Printf("We have %d neighbour discovery packets", len(g.Store.neighbourInfos))
	y := 0
	for nodeId, neighbourInfo := range g.Store.neighbourInfos {
		for mac, _ := range neighbourInfo.Batdv {
			nodeTable[mac] = &GraphNode{
				Id:     mac,
				NodeId: nodeId,
			}
			y = y + 1
		}
	}
	log.Printf("iterated over %d entries", y)
	log.Printf("Node table has %d entries", len(nodeTable))

	nodeList := make([]*GraphNode, 0, len(nodeTable))
	i := 0
	for _, item := range nodeTable {
		item.tableId = i
		nodeList = append(nodeList, item)
		i = i + 1
	}
	log.Printf("Node list has %d entries", len(nodeList))

	allLinks := make([]*GraphLink, 0, len(g.Store.neighbourInfos)*5)

	for _, neighbours := range g.Store.neighbourInfos {
		for ownMac, neighbour := range neighbours.Batdv {
			for peerMac, linkInfo := range neighbour.Neighbours {
				source, sourceExists := nodeTable[ownMac]
				target, targetExists := nodeTable[peerMac]
				if !sourceExists || !targetExists {
					log.WithFields(log.Fields{
						"source-mac": ownMac,
						"target-mac": peerMac,
					}).Warning("Tried to build link to unknown peer")
					continue
				}
				link := &GraphLink{
					Source: source.tableId,
					Target: target.tableId,
					Tq:     float64(linkInfo.Tq),
				}
				allLinks = append(allLinks, link)
			}
		}
	}

	bidirectionalLinks := make([]*GraphLink, 0, len(g.Store.neighbourInfos)*5)
	unidirectionalLinks := make([]*GraphLink, 0, len(g.Store.neighbourInfos))
	for _, link := range allLinks {
		_, err := FindInLinks(allLinks, link.Target, link.Source)
		if err != nil {
			link.Bidirect = false
			unidirectionalLinks = append(unidirectionalLinks, link)
		} else {
			link.Bidirect = true
			_, err := FindInLinks(allLinks, link.Source, link.Target)
			if err != nil {
				bidirectionalLinks = append(bidirectionalLinks, link)
			}
		}
	}

	allLinks = make([]*GraphLink, 0, len(bidirectionalLinks)+len(unidirectionalLinks))
	allLinks = append(allLinks, bidirectionalLinks...)
	allLinks = append(allLinks, unidirectionalLinks...)
	log.Printf("All links %v", allLinks)
	for _, link := range allLinks {
		if link == nil {
			log.Warnf("Link is nil!")
			continue
		}
		source := nodeList[link.Source]
		target := nodeList[link.Target]
		_, sourceGW := g.Store.gatewayList[source.Id]
		_, targetGW := g.Store.gatewayList[target.Id]
		if sourceGW || targetGW {
			link.Vpn = true
		}
	}

	batGraph := BatadvGraph{
		Multigraph: false,
		Directed:   false,
		Links:      allLinks,
		Nodes:      nodeList,
	}

	graphJson := GraphJson{
		Batadv:  batGraph,
		Version: 1,
	}
	return graphJson
}
