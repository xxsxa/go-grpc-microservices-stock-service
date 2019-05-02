package main

import (
	"context"
	"github.com/micro/go-micro"
	"log"
	pb "stock-service/proto/stock"
)

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

type Repository struct {
	consignments []*pb.Consignment
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	update := append(repo.consignments, consignment)
	repo.consignments = update
	return consignment, nil
}

// GET ALL STOCK
func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

type service struct {
	repo repository
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}
	res.Created = true
	res.Consignment = consignment
	return nil
}
func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments := s.repo.GetAll()
	res.Consignments = consignments
	return nil
}

func main() {
	repo := &Repository{}
	Server := micro.NewService(micro.Name("shippy.service.stock"))
	Server.Init()
	pb.RegisterShippingServiceHandler(Server.Server(), &service{repo})
	if err := Server.Run(); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
