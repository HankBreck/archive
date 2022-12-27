package apptesting

import (
	"archive/x/identity/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// PrepareCertificate registers the issuer, issues a certificate, and adds recipient as a member if the pointer is not nil.
//
// Returns the id of the certificate. Returns an error if one arises.
func (s *KeeperTestHelper) PrepareCertificate(issuer sdk.AccAddress, recipient *sdk.AccAddress) (uint64, error) {
	k := s.App.IdentityKeeper

	// Create the issuer
	err := k.SetIssuer(s.Ctx, types.Issuer{
		Creator:     issuer.String(),
		Name:        "Test Issuer",
		MoreInfoUri: "google.com/more-info",
		Cost:        0,
	})
	if err != nil {
		return 0, err
	}

	// Issue the certificate
	cert := types.Certificate{
		IssuerAddress:     issuer.String(),
		Salt:              "salt",
		MetadataSchemaUri: "google.com/metadata-schema",
		Hashes: []*types.HashEntry{
			{Field: "field1", Hash: "hash1"},
			{Field: "field2", Hash: "hash2"},
		},
	}
	id := k.AppendCertificate(s.Ctx, cert)

	if recipient != nil {
		k.CreateMembership(s.Ctx, id, *recipient)
	}

	return id, nil
}
