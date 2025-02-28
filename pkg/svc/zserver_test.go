package svc

import (
	"context"
	"testing"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/infobloxopen/acme/pkg/pb"
)

func TestGetVersion(t *testing.T) {
	var tests = []struct {
		name     string
		expected *pb.VersionResponse
		err      error
	}{
		{
			name:     "verify service version",
			expected: &pb.VersionResponse{Version: version},
			err:      nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server, err := NewBasicServer(nil)
			if err != test.err {
				t.Errorf("Unexpected error when creating server: %v - expected: %v",
					err, test.err,
				)
			}
			res, err := server.GetVersion(context.Background(), &empty.Empty{})
			if res.Version != test.expected.Version {
				t.Errorf("Unexpected version in response: %v - expected: %v",
					res.Version, test.expected.Version,
				)
			}
		})
	}
}
