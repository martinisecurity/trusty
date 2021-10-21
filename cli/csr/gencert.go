package csr

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/go-phorce/dolly/algorithms/guid"
	"github.com/go-phorce/dolly/ctl"
	"github.com/martinisecurity/trusty/authority"
	"github.com/martinisecurity/trusty/cli"
	"github.com/martinisecurity/trusty/pkg/csr"
	"github.com/martinisecurity/trusty/pkg/print"
	"github.com/pkg/errors"
)

// GenCertFlags specifies flags for GenCert command
type GenCertFlags struct {
	SelfSign *bool
	// CACert specifies file name with CA cert
	CACert *string
	// CAKey specifies file name with CA key
	CAKey *string
	// CAConfig specifies file name with ca-config
	CAConfig *string
	// CsrProfile specifies file name with CSR profile
	CsrProfile *string
	// Label specifies name for generated key
	KeyLabel *string
	// SAN specifies coma separated alt names for generated cert
	SAN *string
	// Profile specifies the profile name from ca-config
	Profile *string
	// Output specifies the optional prefix for output files,
	// if not set, the output will be printed to STDOUT only
	Output *string
}

// GenCert generates a cert
func GenCert(c ctl.Control, p interface{}) error {
	flags := p.(*GenCertFlags)

	cryptoprov, defaultCrypto := c.(*cli.Cli).CryptoProv()
	if cryptoprov == nil {
		return errors.Errorf("unsupported command for this crypto provider")
	}

	isscfg := &authority.IssuerConfig{}

	if flags.SelfSign != nil && *flags.SelfSign {
		if *flags.CACert != "" || *flags.CAKey != "" {
			return errors.Errorf("--self-sign can not be used with --ca-key")
		}
	} else {
		if *flags.CACert == "" || *flags.CAKey == "" {
			return errors.Errorf("CA certificate and key are required")
		}
		isscfg.CertFile = *flags.CACert
		isscfg.KeyFile = *flags.CAKey
	}

	// Load CSR
	csrf, err := c.(*cli.Cli).ReadFileOrStdin(*flags.CsrProfile)
	if err != nil {
		return errors.WithMessage(err, "read CSR profile")
	}

	prov := csr.NewProvider(defaultCrypto)
	req := csr.CertificateRequest{
		KeyRequest: prov.NewKeyRequest(prefixKeyLabel(*flags.KeyLabel), "ECDSA", 256, csr.SigningKey),
	}

	err = json.Unmarshal(csrf, &req)
	if err != nil {
		return errors.WithMessage(err, "invalid CSR profile")
	}
	if flags.SAN != nil && len(*flags.SAN) > 0 {
		req.SAN = strings.Split(*flags.SAN, ",")
	}

	// Load ca-config
	cacfg, err := authority.LoadConfig(*flags.CAConfig)
	if err != nil {
		return errors.WithMessage(err, "ca-config")
	}
	err = cacfg.Validate()
	if err != nil {
		return errors.WithMessage(err, "invalid ca-config")
	}

	isscfg.Profiles = cacfg.Profiles

	var key, csrPEM, certPEM []byte

	if flags.SelfSign != nil && *flags.SelfSign {
		certPEM, csrPEM, key, err = authority.NewRoot(*flags.Profile,
			cacfg,
			defaultCrypto, &req)
		if err != nil {
			return errors.WithStack(err)
		}
	} else {
		issuer, err := authority.NewIssuer(isscfg, cryptoprov)
		if err != nil {
			return errors.WithMessage(err, "create issuer")
		}

		csrPEM, key, _, _, err = prov.CreateRequestAndExportKey(&req)
		if err != nil {
			key = nil
			return errors.WithMessage(err, "process CSR")
		}

		signReq := csr.SignRequest{
			Request: string(csrPEM),
			Profile: *flags.Profile,
		}

		_, certPEM, err = issuer.Sign(signReq)
		if err != nil {
			return errors.WithMessage(err, "sign request")
		}
	}

	if *flags.Output == "" {
		print.CSRandCert(c.Writer(), key, csrPEM, certPEM)
	} else {
		err = SaveCert(*flags.Output, key, csrPEM, certPEM)
		if err != nil {
			return errors.WithMessagef(err, "unable to save generated files")
		}
	}

	return nil
}

// prefixKeyLabel adds a date prefix to label for a key
func prefixKeyLabel(label string) string {
	if strings.HasSuffix(label, "*") {
		g := guid.MustCreate()
		t := time.Now().UTC()
		label = strings.TrimSuffix(label, "*") +
			fmt.Sprintf("_%04d%02d%02d%02d%02d%02d_%s", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), g[:4])
	}

	return label
}
