package main

import (
	"fmt"
	"log"
	"net"

	criticalpackage "github.com/Troelshjarne/Exercise_2/critical"
	"google.golang.org/grpc"
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

func main() {

	fmt.Println("=== Node starting up ===")
	list, err := net.Listen("tcp", ":9080")

	if err != nil {
		log.Fatalf("Failed to listen on port 9080: %v", err)
	}

	var options []grpc.ServerOption
	grpcServer := grpc.NewServer(options...)

	criticalpackage.RegisterCommunicationServer(grpcServer, &Server{
		channel: make(map[string][]chan *criticalpackage.Request),
	})
	grpcServer.Serve(list)
}
