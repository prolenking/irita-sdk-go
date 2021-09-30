package codec

import (
	codectypes "github.com/bianjieai/irita-sdk-go/codec/types"
	"github.com/bianjieai/irita-sdk-go/crypto/keys/ed25519"
	"github.com/bianjieai/irita-sdk-go/crypto/keys/multisig"
	"github.com/bianjieai/irita-sdk-go/crypto/keys/secp256k1"
	"github.com/bianjieai/irita-sdk-go/crypto/keys/sm2"
	cryptotypes "github.com/bianjieai/irita-sdk-go/crypto/types"
)

// RegisterInterfaces registers the sdk.Tx interface.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	// TODO We now register both Tendermint's PubKey and our own PubKey. In the
	// long-term, we should move away from Tendermint's PubKey, and delete
	// these lines.
	registry.RegisterInterface("cosmos.crypto.Pubkey", (*cryptotypes.PubKey)(nil))
	registry.RegisterImplementations((*cryptotypes.PubKey)(nil), &sm2.PubKey{})
	registry.RegisterImplementations((*cryptotypes.PubKey)(nil), &ed25519.PubKey{})
	registry.RegisterImplementations((*cryptotypes.PubKey)(nil), &secp256k1.PubKey{})
	registry.RegisterImplementations((*cryptotypes.PubKey)(nil), &multisig.LegacyAminoPubKey{})
}
