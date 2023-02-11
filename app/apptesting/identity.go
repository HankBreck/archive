package apptesting

import (
	"github.com/HankBreck/archive/x/identity/types"

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

	k.SetInitialOperator(s.Ctx, id, *recipient)

	return id, nil
}

func (s *KeeperTestHelper) SetMembers(certificateId uint64, members []sdk.AccAddress) error {
	k := s.App.IdentityKeeper

	// Add members as pending
	err := k.UpdatePendingMembers(s.Ctx, certificateId, members, []sdk.AccAddress{})
	if err != nil {
		return err
	}

	// Add members as accepted
	err = k.UpdateAcceptedMembers(s.Ctx, certificateId, members, []sdk.AccAddress{})
	if err != nil {
		return err
	}

	// Removing members from pending
	err = k.UpdatePendingMembers(s.Ctx, certificateId, []sdk.AccAddress{}, members)
	if err != nil {
		return err
	}

	return nil
}

func (s *KeeperTestHelper) AddOperators(certificateId uint64, opers []sdk.AccAddress) error {
	k := s.App.IdentityKeeper

	err := s.SetMembers(certificateId, opers)
	if err != nil {
		return err
	}

	err = k.SetOperators(s.Ctx, certificateId, opers)
	if err != nil {
		return err
	}
	return nil
}

func (s *KeeperTestHelper) AcceptMembership(certificateId uint64, member sdk.AccAddress) error {
	k := s.App.IdentityKeeper

	err := k.UpdateMembershipStatus(s.Ctx, certificateId, member, true)
	return err
}
