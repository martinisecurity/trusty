package authority_test

import (
	"strings"
	"testing"

	"github.com/go-phorce/dolly/algorithms/guid"
	"github.com/go-phorce/dolly/xpki/cryptoprov"
	"github.com/martinisecurity/trusty/authority"
	"github.com/martinisecurity/trusty/backend/config"
	"github.com/martinisecurity/trusty/pkg/awskmscrypto"
	"github.com/martinisecurity/trusty/pkg/csr"
	"github.com/martinisecurity/trusty/pkg/gcpkmscrypto"
	"github.com/martinisecurity/trusty/tests/testutils"
	"github.com/stretchr/testify/suite"
)

const (
	ca1CertFile    = "/tmp/trusty/certs/trusty_l1_ca.pem"
	ca1KeyFile     = "/tmp/trusty/certs/trusty_l1_ca.key"
	ca2CertFile    = "/tmp/trusty/certs/trusty_l2_ca.pem"
	ca2KeyFile     = "/tmp/trusty/certs/trusty_l2_ca.key"
	caBundleFile   = "/tmp/trusty/certs/trusty_cabundle.pem"
	rootBundleFile = "/tmp/trusty/certs/trusty_root_ca.pem"
)

var (
	falseVal = false
	trueVal  = true
)

type testSuite struct {
	suite.Suite

	cfg    *config.Configuration
	crypto *cryptoprov.Crypto
}

func (s *testSuite) SetupSuite() {
	var err error

	s.cfg, err = testutils.LoadConfig(projFolder, "UNIT_TEST")
	s.Require().NoError(err)

	cryptoprov.Register("AWSKMS", awskmscrypto.KmsLoader)
	cryptoprov.Register("GCPKMS", gcpkmscrypto.KmsLoader)
	cryptoprov.Register("SoftHSM", cryptoprov.Crypto11Loader)
	s.crypto, err = cryptoprov.Load(s.cfg.CryptoProv.Default, s.cfg.CryptoProv.Providers)
	s.Require().NoError(err)
}

func (s *testSuite) TearDownSuite() {
}

func TestAuthority(t *testing.T) {
	suite.Run(t, new(testSuite))
}

func (s *testSuite) TestNewAuthority() {
	//
	// Test empty config
	//
	cfg := &authority.Config{}
	_, err := authority.NewAuthority(cfg, s.crypto)
	s.Require().Error(err)
	s.Equal("missing Authority configuration", err.Error())

	cfg, err = authority.LoadConfig(projFolder + "/etc/dev/ca-config.dev.yaml")
	s.Require().NoError(err)

	/*
		//
		// Test 0 default durations
		//
		cfg2 := cfg.Copy()
		s.Require().Equal(*cfg, *cfg2)

		cfg2.Authority.DefaultAIA = &authority.AIAConfig{
			CRLExpiry:  0,
			OCSPExpiry: 0,
			CRLRenewal: 0,
		}

		_, err = authority.NewAuthority(cfg2, s.crypto)
		s.Require().NoError(err)
	*/

	//
	// Test invalid Issuer files
	//
	cfg3 := cfg.Copy()
	cfg3.Authority.Issuers = []authority.IssuerConfig{
		{
			Label:    "disabled",
			Disabled: &trueVal,
		},
		{
			Label:   "badkey",
			KeyFile: "not_found",
		},
	}

	_, err = authority.NewAuthority(cfg3, s.crypto)
	s.Require().Error(err)
	s.Equal("unable to create issuer: \"badkey\": unable to create signer: load key file: open not_found: no such file or directory", err.Error())

	//
	// test default Expiry and Renewal from Authority config
	//
	cfg4 := cfg.Copy()
	for i := range cfg4.Authority.Issuers {
		cfg4.Authority.Issuers[i].AIA = &authority.AIAConfig{}
	}

	a, err := authority.NewAuthority(cfg4, s.crypto)
	s.Require().NoError(err)
	issuers := a.Issuers()
	s.Equal(len(cfg4.Authority.Issuers), len(issuers))

	for _, issuer := range issuers {
		s.NotContains(issuer.AiaURL(), "${ISSUER_ID}")
		s.NotContains(issuer.CrlURL(), "${ISSUER_ID}")
		s.NotContains(issuer.OcspURL(), "${ISSUER_ID}")

		i, err := a.GetIssuerByLabel(issuer.Label())
		s.NoError(err)
		s.NotNil(i)

		i, err = a.GetIssuerByKeyID(issuer.SubjectKID())
		s.NoError(err)
		s.NotNil(i)

		for name := range cfg.Profiles {
			_, err = a.GetIssuerByProfile(name)
			if strings.Contains(name, "${LABEL}") {
				s.Error(err)
			} else {
				s.NoError(err)
			}
		}
	}

	_, err = a.GetIssuerByKeyID("xxxx")
	s.Error(err)
	s.Equal("issuer not found: xxxx", err.Error())

	_, err = a.GetIssuerByLabel("wrong")
	s.Error(err)
	s.Equal("issuer not found: wrong", err.Error())

	_, err = a.GetIssuerByProfile("wrong_profile")
	s.Error(err)
	s.Equal("issuer not found for profile: wrong_profile", err.Error())
}

func (s *testSuite) TestIssuerSign() {
	crypto := s.crypto.Default()
	kr := csr.NewKeyRequest(crypto, "TestNewRoot"+guid.MustCreate(), "ECDSA", 256, csr.SigningKey)
	rootReq := csr.CertificateRequest{
		CommonName: "[TEST] Trusty Root CA",
		KeyRequest: kr,
	}
	rootPEM, _, rootKey, err := authority.NewRoot("ROOT", rootCfg, crypto, &rootReq)
	s.Require().NoError(err)

	rootSigner, err := authority.NewSignerFromPEM(s.crypto, rootKey)
	s.Require().NoError(err)

	cfg := &authority.IssuerConfig{
		AIA: &authority.AIAConfig{
			AiaURL:  "https://localhost/v1/cert/${ISSUER_ID}",
			OcspURL: "https://localhost/v1/ocsp",
			CrlURL:  "https://localhost/v1/crl/${ISSUER_ID}",
		},
		Label: "TrustyRoot",
		Profiles: map[string]*authority.CertProfile{
			"L1": {
				Usage:       []string{"cert sign", "crl sign"},
				Expiry:      1 * csr.OneYear,
				OCSPNoCheck: true,
				CAConstraint: authority.CAConstraint{
					IsCA:       true,
					MaxPathLen: 1,
				},
				Policies: []csr.CertificatePolicy{
					{
						ID: csr.OID{1, 2, 1000, 1},
						Qualifiers: []csr.CertificatePolicyQualifier{
							{Type: csr.CpsQualifierType, Value: "CPS"},
							{Type: csr.UserNoticeQualifierType, Value: "notice"},
						},
					},
				},
				AllowedExtensions: []csr.OID{
					{1, 3, 6, 1, 5, 5, 7, 1, 1},
				},
			},
			"RestrictedCA": {
				Usage:       []string{"cert sign", "crl sign"},
				Expiry:      1 * csr.OneYear,
				Backdate:    0,
				OCSPNoCheck: true,
				CAConstraint: authority.CAConstraint{
					IsCA:       true,
					MaxPathLen: 0,
				},
				AllowedNames: "^[Tt]rusty CA$",
				AllowedDNS:   "^trusty\\.com$",
				AllowedEmail: "^ca@trusty\\.com$",
				AllowedURI:   "^spiffe://trusty/.*$",
				AllowedCSRFields: &csr.AllowedFields{
					Subject:        true,
					DNSNames:       true,
					IPAddresses:    true,
					EmailAddresses: true,
					URIs:           true,
				},
				AllowedExtensions: []csr.OID{
					{1, 3, 6, 1, 5, 5, 7, 1, 1},
					{2, 5, 29, 17},
				},
			},
			"RestrictedServer": {
				Usage:        []string{"server auth", "signing", "key encipherment"},
				Expiry:       1 * csr.OneYear,
				Backdate:     0,
				AllowedNames: "trusty.com",
				AllowedDNS:   "^(www\\.)?trusty\\.com$",
				AllowedEmail: "^ca@trusty\\.com$",
				AllowedURI:   "^spiffe://trusty/.*$",
				AllowedCSRFields: &csr.AllowedFields{
					Subject:        true,
					DNSNames:       true,
					IPAddresses:    true,
					EmailAddresses: true,
					URIs:           true,
				},
				AllowedExtensions: []csr.OID{
					{1, 3, 6, 1, 5, 5, 7, 1, 1},
					{2, 5, 29, 17},
				},
			},
			"default": {
				Usage:        []string{"server auth", "signing", "key encipherment"},
				Expiry:       1 * csr.OneYear,
				Backdate:     0,
				AllowedNames: "trusty.com",
				AllowedURI:   "^spiffe://trusty/.*$",
				AllowedCSRFields: &csr.AllowedFields{
					Subject:  true,
					DNSNames: true,
					URIs:     true,
				},
				AllowedExtensions: []csr.OID{
					{1, 2, 3},
					{2, 5, 29, 17},
				},
			},
		},
	}

	for name, profile := range cfg.Profiles {
		s.NoError(profile.Validate(), "failed to validate %s profile", name)
	}

	rootCA, err := authority.CreateIssuer(cfg, rootPEM, nil, nil, rootSigner)
	s.Require().NoError(err)

	s.Run("default", func() {
		req := csr.CertificateRequest{
			CommonName: "trusty.com",
			SAN:        []string{"www.trusty.com", "127.0.0.1", "server@trusty.com", "spiffe://trusty/test"},
			KeyRequest: kr,
		}

		csrPEM, _, _, _, err := csr.NewProvider(crypto).CreateRequestAndExportKey(&req)
		s.Require().NoError(err)

		sreq := csr.SignRequest{
			Request: string(csrPEM),
			SAN:     req.SAN,
			Extensions: []csr.X509Extension{
				{
					ID:    csr.OID{1, 2, 3},
					Value: "0500",
				},
			},
		}

		crt, _, err := rootCA.Sign(sreq)
		s.Require().NoError(err)
		s.Equal(req.CommonName, crt.Subject.CommonName)
		s.Equal(rootReq.CommonName, crt.Issuer.CommonName)
		s.False(crt.IsCA)
		s.Equal(-1, crt.MaxPathLen)
		s.NotEmpty(crt.IPAddresses)
		s.NotEmpty(crt.EmailAddresses)
		s.NotEmpty(crt.DNSNames)
		s.NotEmpty(crt.URIs)

		// test unknown profile
		sreq.Profile = "unknown"
		_, _, err = rootCA.Sign(sreq)
		s.Require().Error(err)
		s.Equal("unsupported profile: unknown", err.Error())
	})

	s.Run("Valid L1", func() {
		caReq := csr.CertificateRequest{
			CommonName: "[TEST] Trusty Level 1 CA",
			KeyRequest: kr,
		}

		caCsrPEM, _, _, _, err := csr.NewProvider(crypto).CreateRequestAndExportKey(&caReq)
		s.Require().NoError(err)

		sreq := csr.SignRequest{
			SAN:     []string{"trusty@ekspand.com", "trusty.com", "127.0.0.1"},
			Request: string(caCsrPEM),
			Profile: "L1",
			Subject: &csr.X509Subject{
				CommonName: "Test L1 CA",
			},
		}

		caCrt, _, err := rootCA.Sign(sreq)
		s.Require().NoError(err)
		s.Equal(sreq.Subject.CommonName, caCrt.Subject.CommonName)
		s.Equal(rootReq.CommonName, caCrt.Issuer.CommonName)
		s.True(caCrt.IsCA)
		s.Equal(1, caCrt.MaxPathLen)
	})

	s.Run("RestrictedCA/NotAllowedCN", func() {
		caReq := csr.CertificateRequest{
			CommonName: "[TEST] Trusty Level 2 CA",
			KeyRequest: kr,
			SAN:        []string{"ca@trusty.com", "trusty.com", "127.0.0.1"},
			Names: []csr.X509Name{
				{
					O: "trusty",
					C: "US",
				},
			},
		}

		caCsrPEM, _, _, _, err := csr.NewProvider(crypto).CreateRequestAndExportKey(&caReq)
		s.Require().NoError(err)

		sreq := csr.SignRequest{
			Request: string(caCsrPEM),
			Profile: "RestrictedCA",
		}

		_, _, err = rootCA.Sign(sreq)
		s.Require().Error(err)
		s.Equal("CommonName does not match allowed list: [TEST] Trusty Level 2 CA", err.Error())
	})

	s.Run("RestrictedCA/NotAllowedDNS", func() {
		caReq := csr.CertificateRequest{
			CommonName: "trusty CA",
			KeyRequest: kr,
			SAN:        []string{"ca@trusty.com", "trustyca.com", "127.0.0.1"},
			Names: []csr.X509Name{
				{
					O: "trusty",
					C: "US",
				},
			},
		}

		caCsrPEM, _, _, _, err := csr.NewProvider(crypto).CreateRequestAndExportKey(&caReq)
		s.Require().NoError(err)

		sreq := csr.SignRequest{
			Request: string(caCsrPEM),
			Profile: "RestrictedCA",
		}

		_, _, err = rootCA.Sign(sreq)
		s.Require().Error(err)
		s.Equal("DNS Name does not match allowed list: trustyca.com", err.Error())
	})

	s.Run("RestrictedCA/NotAllowedURI", func() {
		caReq := csr.CertificateRequest{
			CommonName: "trusty CA",
			KeyRequest: kr,
			SAN:        []string{"ca@trusty.com", "127.0.0.1", "spiffe://google.com/ca"},
			Names: []csr.X509Name{
				{
					O: "trusty",
					C: "US",
				},
			},
		}

		caCsrPEM, _, _, _, err := csr.NewProvider(crypto).CreateRequestAndExportKey(&caReq)
		s.Require().NoError(err)

		sreq := csr.SignRequest{
			SAN:     caReq.SAN,
			Request: string(caCsrPEM),
			Profile: "RestrictedCA",
		}

		_, _, err = rootCA.Sign(sreq)
		s.Require().Error(err)
		s.Equal("URI does not match allowed list: spiffe://google.com/ca", err.Error())
	})

	s.Run("RestrictedCA/NotAllowedEmail", func() {
		caReq := csr.CertificateRequest{
			CommonName: "trusty CA",
			KeyRequest: kr,
			SAN:        []string{"rootca@trusty.com", "trusty.com", "127.0.0.1"},
			Names: []csr.X509Name{
				{
					O: "trusty",
					C: "US",
				},
			},
		}

		caCsrPEM, _, _, _, err := csr.NewProvider(crypto).CreateRequestAndExportKey(&caReq)
		s.Require().NoError(err)

		sreq := csr.SignRequest{
			Request: string(caCsrPEM),
			Profile: "RestrictedCA",
		}

		_, _, err = rootCA.Sign(sreq)
		s.Require().Error(err)
		s.Equal("Email does not match allowed list: rootca@trusty.com", err.Error())
	})

	s.Run("RestrictedCA/Valid", func() {
		caReq := csr.CertificateRequest{
			CommonName: "trusty CA",
			KeyRequest: kr,
			//SAN:        []string{"ca@trusty.com", "trusty.com", "127.0.0.1"},
			Names: []csr.X509Name{
				{
					O: "trusty",
					C: "US",
				},
			},
		}

		caCsrPEM, _, _, _, err := csr.NewProvider(crypto).CreateRequestAndExportKey(&caReq)
		s.Require().NoError(err)

		sreq := csr.SignRequest{
			Request: string(caCsrPEM),
			Profile: "RestrictedCA",
		}

		caCrt, _, err := rootCA.Sign(sreq)
		s.Require().NoError(err)
		s.Equal(caReq.CommonName, caCrt.Subject.CommonName)
		s.Equal(rootReq.CommonName, caCrt.Issuer.CommonName)
		s.True(caCrt.IsCA)
		s.Equal(0, caCrt.MaxPathLen)
		s.True(caCrt.MaxPathLenZero)
		// for CA, these are not set:
		s.Empty(caCrt.DNSNames)
		s.Empty(caCrt.EmailAddresses)
		s.Empty(caCrt.IPAddresses)
	})

	s.Run("RestrictedServer/Valid", func() {
		req := csr.CertificateRequest{
			CommonName: "trusty.com",
			KeyRequest: kr,
			SAN:        []string{"ca@trusty.com", "www.trusty.com", "127.0.0.1"},
			Names: []csr.X509Name{
				{
					O: "trusty",
					C: "US",
				},
			},
		}

		csrPEM, _, _, _, err := csr.NewProvider(crypto).CreateRequestAndExportKey(&req)
		s.Require().NoError(err)

		sreq := csr.SignRequest{
			Request: string(csrPEM),
			Profile: "RestrictedServer",
		}

		crt, _, err := rootCA.Sign(sreq)
		s.Require().NoError(err)
		s.Equal(req.CommonName, crt.Subject.CommonName)
		s.Equal(rootReq.CommonName, crt.Issuer.CommonName)
		s.False(crt.IsCA)
		s.Contains(crt.DNSNames, "www.trusty.com")
		s.Contains(crt.EmailAddresses, "ca@trusty.com")
		s.NotEmpty(crt.IPAddresses)
	})
}
