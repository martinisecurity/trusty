package certpublisher

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-phorce/dolly/xlog"
	"github.com/juju/errors"
	"github.com/martinisecurity/trusty/api/v1/pb"
	"github.com/martinisecurity/trusty/pkg/storage"
)

var logger = xlog.NewPackageLogger("github.com/martinisecurity/trusty/pkg", "certpublisher")

// Publisher interface
type Publisher interface {
	// PublishCertificate publishes issued cert
	PublishCertificate(context.Context, *pb.Certificate, string) (string, error)
	// PublishCRL publishes issued CRL
	PublishCRL(context.Context, *pb.Crl) (string, error)
}

type publisher struct {
	cfg *Config
}

// NewPublisher returns new Publisher
func NewPublisher(cfg *Config) (Publisher, error) {
	logger.KV(xlog.INFO, "cert_bucket", cfg.CertsBucket, "crl_bucket", cfg.CRLBucket)
	return &publisher{cfg}, nil
}

// PublishCertificate publishes issued cert
func (p *publisher) PublishCertificate(ctx context.Context, cert *pb.Certificate, filename string) (string, error) {
	location := fmt.Sprintf("%s/%s", p.cfg.CertsBucket, filename)

	logger.KV(xlog.INFO, "location", location)

	pem := strings.TrimSpace(cert.Pem)
	if len(cert.IssuersPem) > 0 {
		pem = pem + "\n" + strings.TrimSpace(cert.IssuersPem)
	}

	_, err := storage.WriteFile(ctx, location, []byte(pem))
	if err != nil {
		return "", errors.Annotatef(err, "unable to write file to: "+location)
	}
	return location, nil
}

// PublishCRL publishes issued CRL
func (p *publisher) PublishCRL(ctx context.Context, crl *pb.Crl) (string, error) {
	fileName := fmt.Sprintf("%s/%s", p.cfg.CertsBucket, string(crl.Ikid))

	logger.KV(xlog.INFO, "location", fileName)

	_, err := storage.WriteFile(ctx, fileName, []byte(crl.Pem))
	if err != nil {
		return "", errors.Annotatef(err, "unable to write file to: "+fileName)
	}
	return fileName, nil
}
