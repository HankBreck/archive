package keeper

import (
	"archive/x/identity/types"
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// AppendCertificate stores the certificate in prefixed storage using certificate.Id as the key.
// Expects certificate.Id to be unset, so an existing value will be overwritten by this function.
//
// Returns the uint64 id of the certificate
func (k Keeper) AppendCertificate(ctx sdk.Context, certificate types.Certificate) uint64 {
	count := k.getCertificateCount(ctx)
	if k.HasCertificate(ctx, count) {
		panic("Duplicate Certificate ID found... this should never happen")
	}
	certificate.Id = count

	// TODO: Any more checks necessary? (maybe hashes != nil)
	k.uncheckedSetCertificate(ctx, certificate)
	k.setCertificateCount(ctx, count+1)
	return count
}

// GetCertificate fetches the certificate stored under the key of id.
// If no contract is found, an error is returned
func (k Keeper) GetCertificate(ctx sdk.Context, id uint64) (*types.Certificate, error) {
	store := k.getCertificateStore(ctx)
	bzKey := make([]byte, 8)
	binary.BigEndian.PutUint64(bzKey, id)

	var certificate types.Certificate
	bzCert := store.Get(bzKey)
	if len(bzCert) == 0 {
		return nil, sdkerrors.ErrNotFound.Wrapf("A certificate with an ID of %d was not found", id)
	}

	k.cdc.MustUnmarshal(bzCert, &certificate)
	return &certificate, nil
}

// HasCertificate returns true if a certificate is stored under id, else false
func (k Keeper) HasCertificate(ctx sdk.Context, id uint64) bool {
	store := k.getCertificateStore(ctx)
	bzKey := make([]byte, 8)
	binary.BigEndian.PutUint64(bzKey, id)
	return store.Has(bzKey)
}

// HasCertificate returns true if issuer matches the ID pointing to the certificate's issuer field, else false
func (k Keeper) HasIssuerForId(ctx sdk.Context, id uint64, issuer sdk.AccAddress) (bool, error) {
	cert, err := k.GetCertificate(ctx, id)
	if err != nil {
		return false, err
	}
	addr, err := sdk.AccAddressFromBech32(cert.IssuerAddress)
	if err != nil {
		return false, err
	}
	return !issuer.Equals(addr), nil
}

// Stores the certificate with a key of certificate.Id. The certificate.Id field must be set by a calling function.
// The certificate passed as an argument is assumed to be valid, so calling functions must assure this.
func (k Keeper) uncheckedSetCertificate(ctx sdk.Context, certificate types.Certificate) {
	store := k.getCertificateStore(ctx)
	bzKey := make([]byte, 8)
	binary.BigEndian.PutUint64(bzKey, certificate.Id)
	bzCert := k.cdc.MustMarshal(&certificate)
	store.Set(bzKey, bzCert)
}

// getCertificateCount returns the next available id for a Certificate to use
func (k Keeper) getCertificateCount(ctx sdk.Context) uint64 {
	// Fetch value from KV store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.CertificateCountKey))
	bzCount := store.Get([]byte{0})

	// Return 0 if the key is nil (first time accessing)
	if bzCount == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bzCount)
}

// setCertificateCount sets the value of CertificateCountKey to count in the prefixed storage
func (k Keeper) setCertificateCount(ctx sdk.Context, count uint64) {
	// Load prefix storage
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.CertificateCountKey))

	// Convert count to bytes
	bzCount := make([]byte, 8)
	binary.BigEndian.PutUint64(bzCount, count)

	// Set the count in storage
	store.Set([]byte{0}, bzCount)
}

func (k Keeper) getCertificateStore(ctx sdk.Context) prefix.Store {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.CertificateKey))
	return store
}
