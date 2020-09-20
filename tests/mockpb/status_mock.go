package mockpb

import (
	"context"

	"github.com/go-phorce/trusty/api/v1/serverpb"
	"github.com/gogo/protobuf/proto"
)

// MockStatusServer for testing
type MockStatusServer struct {
	serverpb.StatusServer

	Reqs []proto.Message

	// If set, all calls return this error.
	Err error

	// responses to return if err == nil
	Resps []proto.Message
}

// Version returns the server version.
func (m *MockStatusServer) Version(ctx context.Context, req *serverpb.EmptyRequest) (*serverpb.VersionResponse, error) {
	//m.reqs = append(m.reqs, req)
	if m.Err != nil {
		return nil, m.Err
	}
	return m.Resps[0].(*serverpb.VersionResponse), nil
}

// Server returns the server statum.
func (m *MockStatusServer) Server(ctx context.Context, req *serverpb.EmptyRequest) (*serverpb.ServerStatusResponse, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return m.Resps[0].(*serverpb.ServerStatusResponse), nil
}

// Caller returns the caller statum.
func (m *MockStatusServer) Caller(ctx context.Context, req *serverpb.EmptyRequest) (*serverpb.CallerStatusResponse, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return m.Resps[0].(*serverpb.CallerStatusResponse), nil
}