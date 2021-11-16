package main

import (
	criticalpackage "github.com/Troelshjarne/Exercise_2/critical"
	//"google.golang.org/grpc"
	//"github.com/hashicorp/go-msgpack/codec"
	//"github.com/hashicorp/logutils"
	//"github.com/hashicorp/serf/coordinate"
)

var lamTime = 0

// set by flag ?
var nodeID = 0

// increments each time a client joins the cluster
var nodes = 0

type Server struct {
	criticalpackage.UnimplementedCommunicationServer

	channel map[string][]chan *criticalpackage.Request
}

func (s *Server) joinCluster(ch *criticalpackage.Channel, requestStream criticalpackage.Communication_JoinClusterClient) error {
	requestChannel := make(chan *criticalpackage.Request)

	s.channel[ch.NodeId] = append(s.channel[ch.NodeId], requestChannel)

}

func (s *Server) sendRequest()
