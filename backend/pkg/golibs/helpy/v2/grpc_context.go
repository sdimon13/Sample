package helpy

import (
	"context"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func ExtractHeaderFromContext(ctx context.Context, key string) (value string, isOk bool) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return
	}

	val := md[key]
	if len(val) == 0 || val[0] == "" {
		return
	}

	return val[0], true
}

func ExtractHeaderFromContextInt32(ctx context.Context, key string) (value int32, isOk bool) {
	strValue, isOk := ExtractHeaderFromContext(ctx, key)
	if !isOk {
		return 0, false
	}
	val, err := strconv.ParseInt(strValue, 10, 32)
	if err != nil {
		return 0, false
	}
	return int32(val), true
}

func ExtractHeaderFromContextInt64(ctx context.Context, key string) (value int64, isOk bool) {
	strValue, isOk := ExtractHeaderFromContext(ctx, key)
	if !isOk {
		return 0, false
	}
	val, err := strconv.ParseInt(strValue, 10, 64)
	if err != nil {
		return 0, false
	}
	return val, true
}

func GrpcDialContext(ctx context.Context, host string, debug ...any) (*grpc.ClientConn, error) {
	if len(debug) > 0 {
		// For local services without TLS
		return grpc.DialContext(ctx, host, grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	return grpc.DialContext(ctx, host, grpc.WithBlock(), grpc.WithTransportCredentials(TLSInsecure()))
}
