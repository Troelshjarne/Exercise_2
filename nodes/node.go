package main

import (
	criticalpackage "github.com/Troelshjarne/Exercise_2/critical"
	//"google.golang.org/grpc"
)

var lamTime = 0

type Server struct {
	criticalpackage.UnimplementedCommunicationServer

	channel map[string][]chan *criticalpackage.Request
}

func (s *Server) joinCluster(ch *criticalpackage.Channel, requestStream criticalpackage.Communication_JoinClusterClient) error {
	requestChannel := make(chan *criticalpackage.Request)

	s.channel[ch.NodeId] = append(s.channel[ch.NodeId], requestChannel)

}

func (s *Server) sendRequest()
