package rpc

import (
	"time"

	"context"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	types "github.com/Microsoft/presidio-genproto/golang"
)

func connect(addr string) (*grpc.ClientConn, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	callOpts := []grpc_retry.CallOption{
		grpc_retry.WithBackoff(grpc_retry.BackoffLinear(1 * time.Second)),
		grpc_retry.WithCodes(codes.NotFound, codes.Aborted),
		grpc_retry.WithMax(5),
	}

	conn, err := grpc.DialContext(ctx, addr,
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(callOpts...)),
		// TODO: We need to add TLS option as well
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithBackoffMaxDelay(5*time.Second),
		grpc.WithBalancerName("round_robin"))

	if err != nil {
		return nil, err
	}
	return conn, nil
}

//SetupAnonymizeService ...
func SetupAnonymizeService(address string) (types.AnonymizeServiceClient, error) {

	conn, err := connect(address)
	if err != nil {
		return nil, err
	}

	client := types.NewAnonymizeServiceClient(conn)
	return client, nil
}

//SetupAnalyzerService ...
func SetupAnalyzerService(address string) (types.AnalyzeServiceClient, error) {
	conn, err := connect(address)
	if err != nil {
		return nil, err
	}
	client := types.NewAnalyzeServiceClient(conn)
	return client, nil
}

//SetupDatasinkService ...
func SetupDatasinkService(address string) (types.DatasinkServiceClient, error) {
	conn, err := connect(address)
	if err != nil {
		return nil, err
	}
	client := types.NewDatasinkServiceClient(conn)
	return client, nil
}

//SetupSchedulerService ...
func SetupSchedulerService(address string) (types.SchedulerServiceClient, error) {
	conn, err := connect(address)
	if err != nil {
		return nil, err
	}
	client := types.NewSchedulerServiceClient(conn)
	return client, nil
}