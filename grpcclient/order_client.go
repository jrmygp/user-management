package grpcclient

import (
	"context"
	"time"

	transactionpb "github.com/jrmygp/contracts/proto/transactionpb"
	"google.golang.org/grpc"
)

type OrderClient struct {
	client transactionpb.OrderServiceClient
}

func NewOrderClient() (*OrderClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}

	client := transactionpb.NewOrderServiceClient(conn)

	return &OrderClient{client: client}, conn, nil
}

func (u *OrderClient) GetOrderByMidtransID(id string) (*transactionpb.GetOrderResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return u.client.FindByMidtransOrderID(ctx, &transactionpb.GetOrderByMidtransRequest{
		Id: id,
	})
}
