package mockpb

import (
	"context"

	"github.com/gogo/protobuf/proto"
	"github.com/martinisecurity/trusty/api/v1/pb"
)

// MockCAServer for testing
type MockCAServer struct {
	pb.CAServiceServer

	Reqs []proto.Message

	// If set, all calls return this error.
	Err error

	// responses to return if err == nil
	Resps []proto.Message
}

// SetResponse sets a single response without errors
func (m *MockCAServer) SetResponse(r proto.Message) {
	m.Err = nil
	m.Resps = []proto.Message{r}
}

// ProfileInfo returns the certificate profile info
func (m *MockCAServer) ProfileInfo(context.Context, *pb.CertProfileInfoRequest) (*pb.CertProfile, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return m.Resps[0].(*pb.CertProfile), nil
}

// GetIssuer returns the issuing CA
func (m *MockCAServer) GetIssuer(ctx context.Context, req *pb.IssuerInfoRequest) (*pb.IssuerInfo, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return m.Resps[0].(*pb.IssuerInfo), nil
}

// SignCertificate returns the certificate
func (m *MockCAServer) SignCertificate(context.Context, *pb.SignCertificateRequest) (*pb.CertificateResponse, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return m.Resps[0].(*pb.CertificateResponse), nil
}

// ListIssuers returns the issuing CAs
func (m *MockCAServer) ListIssuers(context.Context, *pb.ListIssuersRequest) (*pb.IssuersInfoResponse, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return m.Resps[0].(*pb.IssuersInfoResponse), nil
}

// PublishCrls returns published CRLs
func (m *MockCAServer) PublishCrls(context.Context, *pb.PublishCrlsRequest) (*pb.CrlsResponse, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return m.Resps[0].(*pb.CrlsResponse), nil
}

// GetCertificate returns the certificate
func (m *MockCAServer) GetCertificate(ctx context.Context, in *pb.GetCertificateRequest) (*pb.CertificateResponse, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return m.Resps[0].(*pb.CertificateResponse), nil
}

// RevokeCertificate returns the revoked certificate
func (m *MockCAServer) RevokeCertificate(ctx context.Context, in *pb.RevokeCertificateRequest) (*pb.RevokedCertificateResponse, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return m.Resps[0].(*pb.RevokedCertificateResponse), nil
}

// ListCertificates returns stream of Certificates
func (m *MockCAServer) ListCertificates(ctx context.Context, in *pb.ListByIssuerRequest) (*pb.CertificatesResponse, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return m.Resps[0].(*pb.CertificatesResponse), nil
}

// ListRevokedCertificates returns stream of Revoked Certificates
func (m *MockCAServer) ListRevokedCertificates(ctx context.Context, in *pb.ListByIssuerRequest) (*pb.RevokedCertificatesResponse, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return m.Resps[0].(*pb.RevokedCertificatesResponse), nil
}

// GetCRL returns the CRL
func (m *MockCAServer) GetCRL(ctx context.Context, in *pb.GetCrlRequest) (*pb.CrlResponse, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return m.Resps[0].(*pb.CrlResponse), nil
}

// SignOCSP returns OCSP response
func (m *MockCAServer) SignOCSP(ctx context.Context, in *pb.OCSPRequest) (*pb.OCSPResponse, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return m.Resps[0].(*pb.OCSPResponse), nil
}

// UpdateCertificateLabel returns the updated certificate
func (m *MockCAServer) UpdateCertificateLabel(ctx context.Context, in *pb.UpdateCertificateLabelRequest) (*pb.CertificateResponse, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return m.Resps[0].(*pb.CertificateResponse), nil
}

// ListOrgCertificates returns the Org certificates
func (m *MockCAServer) ListOrgCertificates(ctx context.Context, in *pb.ListOrgCertificatesRequest) (*pb.CertificatesResponse, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return m.Resps[0].(*pb.CertificatesResponse), nil
}

// RegisterIssuer creates Issuer
func (m *MockCAServer) RegisterIssuer(ctx context.Context, in *pb.RegisterIssuerRequest) (*pb.IssuerInfo, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return m.Resps[0].(*pb.IssuerInfo), nil
}

// RegisterProfile registers the certificate profile
func (m *MockCAServer) RegisterProfile(ctx context.Context, req *pb.RegisterProfileRequest) (*pb.CertProfile, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return m.Resps[0].(*pb.CertProfile), nil
}
