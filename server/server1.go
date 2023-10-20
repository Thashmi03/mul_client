package main

import (
    "context"
    "fmt"
    "net"
    hw "grpcsample/proto"

    "google.golang.org/grpc"
)

type Customer struct {
    hw.UnimplementedCustomerServiceServer
}

func (c *Customer) CreateCustomer(ctx context.Context, req *hw.CustomerRequest) (*hw.CustomerResponse, error) {
    return &hw.CustomerResponse{
        Count: "1",
    }, nil
}
func main() {
    lis, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Printf("Failed to lIsten:%v", err)
        return
    }
    s := grpc.NewServer()
    hw.RegisterCustomerServiceServer(s, &Customer{})
    if err := s.Serve(lis); err != nil {
        fmt.Printf("failed to server:%v", err)
    }

}