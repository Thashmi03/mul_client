package main

import (
	"context"
	"fmt"
	"log"

	q "grpcsample/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("192.168.1.10:8080", grpc.WithInsecure()) //swetha pc -ip address(192.168.1.10)
	//conn2,err:= grpc.Dial("localhost:8081",grpc.WithInsecure())
	if err != nil {
		log.Fatal("failed1", err)
	}
	defer conn.Close()
	//defer conn2.Close()

	res := q.NewCustomerServiceClient(conn)

	response, err := res.CreateCustomer(context.Background(), &q.CustomerRequest{
		CustomerId: 1,
	})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)

}

// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"time"

// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials/insecure"
// 	ecpb "grpcsample/proto"
// 	"google.golang.org/grpc/resolver"
// )

// const (
// 	exampleScheme      = "example"
// 	exampleServiceName = "lb.example.grpc.io"
// )

// var addrs = []string{"192.168.1.10:50051", "192.168.1.10:50052"}

// func callUnaryEcho(c ecpb.CustomerServiceClient, message string) {
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()
// 	r, err := c.CreateCustomer(ctx, &ecpb.CustomerRequest{
// 		CustomerId: 0,
// 	})
// 	if err != nil {
// 		log.Fatalf("could not greet: %v", err)
// 	}
// 	fmt.Println(r.Count)
// }

// func makeRPCs(cc *grpc.ClientConn, n int) {
// 	hwc := ecpb.NewCustomerServiceClient(cc)
// 	for i := 0; i < n; i++ {
// 		callUnaryEcho(hwc, "this is examples/load_balancing")
// 	}
// }

// func main() {

// 	// Make another ClientConn with round_robin policy.
// 	roundrobinConn, err := grpc.Dial(
// 		fmt.Sprintf("%s:///%s", exampleScheme, exampleServiceName),
// 		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`), // This sets the initial balancing policy.
// 		grpc.WithTransportCredentials(insecure.NewCredentials()),
// 	)
// 	if err != nil {
// 		log.Fatalf("did not connect: %v", err)
// 	}
// 	defer roundrobinConn.Close()

// 	fmt.Println("--- calling helloworld.Greeter/SayHello with round_robin ---")
// 	makeRPCs(roundrobinConn, 10)
// }

// // Following is an example name resolver implementation. Read the name
// // resolution example to learn more about it.

// type exampleResolverBuilder struct{}

// func (*exampleResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
// 	r := &exampleResolver{
// 		target: target,
// 		cc:     cc,
// 		addrsStore: map[string][]string{
// 			exampleServiceName: addrs,
// 		},
// 	}
// 	fmt.Println("build")
// 	r.start()
// 	return r, nil
// }
// func (*exampleResolverBuilder) Scheme() string {
// 	fmt.Println("scheme")
// 	return exampleScheme
// }

// type exampleResolver struct {
// 	target     resolver.Target
// 	cc         resolver.ClientConn
// 	addrsStore map[string][]string
// }

// func (r *exampleResolver) start() {
// 	addrStrs := r.addrsStore[r.target.Endpoint()]
// 	addrs := make([]resolver.Address, len(addrStrs))
// 	for i, s := range addrStrs {
// 		addrs[i] = resolver.Address{Addr: s}
// 	}
// 	r.cc.UpdateState(resolver.State{Addresses: addrs})
// }
// func (*exampleResolver) ResolveNow(o resolver.ResolveNowOptions) {}
// func (*exampleResolver) Close()                                  {}

// func init() {
// 	fmt.Println("init")
// 	resolver.Register(&exampleResolverBuilder{})
// }