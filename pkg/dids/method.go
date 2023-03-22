package dids

import (
	"crypto/sha256"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	"github.com/suutaku/go-bbs/pkg/bbs"
)

type DID struct {
	did string
	pub *bbs.PublicKey
}

const (
	DIDPrefix  = "did"
	MethodName = "cot"
)

func NewDID(pub *bbs.PublicKey) *DID {
	bs, err := pub.Marshal()
	if err != nil {
		return nil
	}
	h := sha256.New()
	h.Write(bs)
	hashBs := h.Sum(nil)
	enc := base58.Encode(hashBs)
	str := fmt.Sprintf("%s:%s:%s", DIDPrefix, MethodName, enc)
	ret := DID{
		did: str,
		pub: pub,
	}
	return &ret
}

func (d *DID) String() string {
	return d.did
}

func (d *DID) CreateDoc(jwk bool) (*DIDDocument, error) {
	doc := &DIDDocument{
		Context:    defaultContext,
		ID:         d.String(),
		Controller: d.String(),
	}
	keyType := typeG2
	if jwk {
		keyType = ldKeyType
	}
	verificationMethod := CreateVerificationMethod("owner", keyType, d.String(), d.pub)
	doc.VerificationMethod = []VerificationMethod{*verificationMethod}
	return doc, nil
}

func CreateVerificationMethod(keyReference, ldType, did string, pub *bbs.PublicKey) *VerificationMethod {

	fixed := did
	if keyReference != "" {
		fixed = did + "#" + keyReference
	}
	ret := &VerificationMethod{
		ID:         fixed,
		Type:       ldType,
		Controller: did,
	}
	switch ldType {
	case ldKeyType:
		ret.PublicKeyJWK = pub.ToJWK()
	case typeG2:
		bs, err := pub.Marshal()
		if err != nil {
			return nil
		}
		enc := base58.Encode(bs)
		ret.PublicKeyBase58 = enc
	case typeG1:

	}
	return ret
}
