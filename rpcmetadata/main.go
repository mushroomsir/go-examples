package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

const (
	headerAuthorize = "authorization"
	scheme          = "Bearer"
)

func main() {
	ctx := context.Background()
	ctx = SetToken(ctx, "xxx")

	testtoken, testerr := AuthFromMD(ctx, scheme)
	fmt.Println(testtoken, testerr)
}

// SetToken set token into context
func SetToken(ctx context.Context, token string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, headerAuthorize, fmt.Sprintf("%s %s", scheme, token))
}

// AuthFromMD is a helper function for extracting the :authorization header from the gRPC metadata of the request.
//
// It expects the `:authorization` header to be of a certain scheme (e.g. `basic`, `bearer`), in a
// case-insensitive format (see rfc2617, sec 1.2). If no such authorization is found, or the token
// is of wrong scheme, an error with gRPC status `Unauthenticated` is returned.
func AuthFromMD(ctx context.Context, expectedScheme string) (string, error) {
	val := metautils.ExtractOutgoing(ctx).Get(headerAuthorize)
	if val == "" {
		return "", grpc.Errorf(codes.Unauthenticated, "Request1 unauthenticated with "+expectedScheme)

	}
	splits := strings.SplitN(val, " ", 2)
	if len(splits) < 2 {
		return "", grpc.Errorf(codes.Unauthenticated, "Bad authorization string")
	}
	if strings.ToLower(splits[0]) != strings.ToLower(expectedScheme) {
		return "", grpc.Errorf(codes.Unauthenticated, "Request unauthenticated with "+expectedScheme)
	}
	return splits[1], nil
}
