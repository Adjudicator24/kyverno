package images

import (
	"context"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/sigstore/cosign/pkg/oci/remote"
)

type ImageVerifier interface {
	// VerifySignature verifies that the image has the expected signatures
	VerifySignature(ctx context.Context, opts Options) (*Response, error)
	// FetchAttestations retrieves signed attestations and decodes them into in-toto statements
	// https://github.com/in-toto/attestation/blob/main/spec/README.md#statement
	FetchAttestations(ctx context.Context, opts Options) (*Response, error)
}

type Client interface {
	Keychain() authn.Keychain
	RefreshKeychainPullSecrets(ctx context.Context) error
	BuildRemoteOption(context.Context) remote.Option
}

type Options struct {
	ImageRef             string
	Client               Client
	FetchAttestations    bool
	Key                  string
	Cert                 string
	CertChain            string
	Roots                string
	Subject              string
	Issuer               string
	AdditionalExtensions map[string]string
	Annotations          map[string]string
	Repository           string
	RekorURL             string
	SignatureAlgorithm   string
	PredicateType        string
	Type                 string
	Identities           string
}

type Response struct {
	Digest     string
	Statements []map[string]interface{}
}
