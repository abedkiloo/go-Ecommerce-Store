package account

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	 "github.com/abedkiloo/go-Ecommerce-Similutaion/pb" // Import generated protobuf package

)

type grpcServer struct {
	service Service
}

func ListenGRPC(service Service, port int) error {
	listen, err := net.Listen("tcp", fmt.Sprintf("%d", port))
	if err != nil {
		return nil
	}
	serv := grpc.NewServer()
	pb.
	{
		serv,
	}
	reflection.Register(serv)
	return serv.Serve(listen)
}

func (s *grpcServer) PostAccount(ctx context.Context, r *pb.PostAccountRequest) (*pb.PostAccountResponse, error) {
	account, err := s.service.PostAccounts(ctx, r.Name)
	if err != nil {
		return nil, err
	}
	return &pb.PostAccountResponse{Account: &pb.Account{
		Id:   account.ID,
		Name: account.Name,
	}}, nil
}
func (s *grpcServer) GetAccount(ctx context.Context, r *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
	account, err := s.service.GetAccounts(ctx, r.ID)
	if err != nil {
		return nil, err
	}
	return &pb.GetAccountResponse{Account: &pb.Account{
		Id:   account.ID,
		Name: account.Name,
	}}, nil
}
func (s *grpcServer) GetAccounts(ctx context.Context, r *pb.GetAccountsRequest) (*pb.GetAccountsResponse, error) {
	res, err := s.service.GetAccounts(ctx, r.ID)
	if err != nil {
		return nil, err
	}
	accounts := []*pb.ACcount{}
	for _, p := range res {
		accounts = append(accounts, &pb.Account{
			Id:   p.ID,
			Name: p.Name,
		})
	}
	return &pb.GetAccountsResponse{Account: accounts}, nil
}
