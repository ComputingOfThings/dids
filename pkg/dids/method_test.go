package dids

import (
	"crypto/sha256"
	"encoding/hex"
	"testing"

	"github.com/ComputingOfThings/dids/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/suutaku/go-bbs/pkg/bbs"
)

func TestCreateDID(t *testing.T) {
	pub, priv, err := bbs.GenerateKeyPair(sha256.New, nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, priv)
	assert.NotEmpty(t, pub)

	keyStr, err := test.GetTestResource("private-key-2.txt")
	assert.NoError(t, err)
	bs, err := hex.DecodeString(string(keyStr))
	assert.NoError(t, err)
	priv2, err := bbs.UnmarshalPrivateKey(bs)
	assert.NoError(t, err)
	pub2 := priv2.PublicKey()

	did := NewDID(pub2)
	assert.NotEmpty(t, did)
	t.Log(did.String())
	t.Logf("id length %d\n", len(did.String()))

	doc, err := did.CreateDoc(false)
	assert.NoError(t, err)
	assert.NotEmpty(t, doc)
	require.Len(t, doc.VerificationMethod, 1)
	require.NotEmpty(t, doc.VerificationMethod[0].PublicKeyBase58)
	require.Empty(t, doc.VerificationMethod[0].PublicKeyJWK)
	require.Equal(t, typeG2, doc.VerificationMethod[0].Type)

	str, err := doc.MarshalIndent()
	assert.NoError(t, err)
	t.Logf("%s\n", str)

	doc2, err := did.CreateDoc(true)
	assert.NoError(t, err)
	assert.NotEmpty(t, doc2)
	require.Len(t, doc2.VerificationMethod, 1)
	require.NotEmpty(t, doc2.VerificationMethod[0].PublicKeyJWK)
	require.Empty(t, doc2.VerificationMethod[0].PublicKeyBase58)
	require.Equal(t, ldKeyType, doc2.VerificationMethod[0].Type)

	str2, err := doc.MarshalIndent()
	assert.NoError(t, err)
	t.Logf("%s\n", str2)

}
