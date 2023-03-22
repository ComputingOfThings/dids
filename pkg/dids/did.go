package dids

import (
	"encoding/json"

	"github.com/suutaku/go-bbs/pkg/bbs"
)

const (
	ldKeyType = "JsonWebKey2020"
	typeG1    = "Bls12381G1Key2020"
	typeG2    = "Bls12381G2Key2020"
)

var defaultContext = []string{
	"https://www.w3.org/ns/did/v1",
	"https://w3id.org/security/suites/bls12381-2020/v1",
}

type DIDDocument struct {
	Context interface{} `json:"@context,omitempty"`
	// As per https://www.w3.org/TR/did-core/#did-subject intermediate representations of DID Documents do not
	// require an ID property. The provided test vectors demonstrate IRs. As such, the property is optional.
	ID                   string                  `json:"id,omitempty"`
	Controller           string                  `json:"controller,omitempty"`
	AlsoKnownAs          string                  `json:"alsoKnownAs,omitempty"`
	VerificationMethod   []VerificationMethod    `json:"verificationMethod,omitempty" validate:"dive"`
	Authentication       []VerificationMethodSet `json:"authentication,omitempty" validate:"dive"`
	AssertionMethod      []VerificationMethodSet `json:"assertionMethod,omitempty" validate:"dive"`
	KeyAgreement         []VerificationMethodSet `json:"keyAgreement,omitempty" validate:"dive"`
	CapabilityInvocation []VerificationMethodSet `json:"capabilityInvocation,omitempty" validate:"dive"`
	CapabilityDelegation []VerificationMethodSet `json:"capabilityDelegation,omitempty" validate:"dive"`
	Services             []Service               `json:"service,omitempty" validate:"dive"`
}

func (dc *DIDDocument) Marshal() ([]byte, error) {
	return json.Marshal(dc)
}

func (dc *DIDDocument) MarshalIndent() ([]byte, error) {
	return json.MarshalIndent(dc, "", "  ")
}

func (dc *DIDDocument) Unmarshal(data []byte) error {
	return json.Unmarshal(data, dc)
}

type VerificationMethod struct {
	ID              string `json:"id" validate:"required"`
	Type            string `json:"type" validate:"required"`
	Controller      string `json:"controller" validate:"required"`
	PublicKeyBase58 string `json:"publicKeyBase58,omitempty"`
	// must conform to https://datatracker.ietf.org/doc/html/rfc7517
	PublicKeyJWK *bbs.PublicKeyJWK `json:"publicKeyJwk,omitempty" validate:"omitempty,dive"`
	// https://datatracker.ietf.org/doc/html/draft-multiformats-multibase-03
	PublicKeyMultibase string `json:"publicKeyMultibase,omitempty"`
	// for PKH DIDs - https://github.com/w3c-ccg/did-pkh/blob/90b28ad3c18d63822a8aab3c752302aa64fc9382/did-pkh-method-draft.md
	BlockchainAccountID string `json:"blockchainAccountId,omitempty"`
}

type VerificationMethodSet interface{}

// Service is a property compliant with the did-core spec https://www.w3.org/TR/did-core/#services
type Service struct {
	ID   string `json:"id" validate:"required"`
	Type string `json:"type" validate:"required"`
	// A string, map, or set composed of one or more strings and/or maps
	// All string values must be valid URIs
	ServiceEndpoint interface{} `json:"serviceEndpoint" validate:"required"`
	RoutingKeys     []string    `json:"routingKeys,omitempty"`
	Accept          []string    `json:"accept,omitempty"`
}
