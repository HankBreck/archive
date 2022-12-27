package keeper_test

import (
	"archive/x/identity/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// TODO:
// 	GetMembers
// 	UpdatePendingMembers
//		Fails when cert doesnt exist
//  UpdateMembershipStatus
//		Fails when cert doesnt exist
//		Fails when member is not in pending state
//		Base case
// 			Removes member from pending list
//			Adds member to accepted list
//		Fails when trying to accept twice
// 	HasMember

func (suite *KeeperTestSuite) TestCreateMembership() {
	k := suite.App.IdentityKeeper

	// Setup default values
	defaultIssuer := types.Issuer{
		Creator:     suite.TestAccs[0].String(),
		Name:        "Test Issuer",
		MoreInfoUri: "google.com/more-info",
		Cost:        0,
	}
	defaultCert := types.Certificate{
		IssuerAddress:     defaultIssuer.Creator,
		Salt:              "salt",
		MetadataSchemaUri: "google.com/metadata-schema",
		Hashes: []*types.HashEntry{
			{Field: "field1", Hash: "hash1"},
			{Field: "field2", Hash: "hash2"},
		},
	}
	defaultRecipient := suite.TestAccs[1]

	tests := map[string]struct {
		inputIssuer      *types.Issuer
		inputCertificate *types.Certificate
		inputRecipient   *sdk.AccAddress
		expPanic         bool
	}{
		"simple_set": {
			inputIssuer:      &defaultIssuer,
			inputCertificate: &defaultCert,
			inputRecipient:   &defaultRecipient,
			expPanic:         false,
		},
		"no_certificate_set": {
			inputIssuer:      &defaultIssuer,
			inputCertificate: nil,
			inputRecipient:   &defaultRecipient,
			expPanic:         true,
		},
	}

	for name, test := range tests {
		suite.Run(name, func() {
			id := uint64(0)

			// Mock the certificate creation step
			k.SetIssuer(suite.Ctx, *test.inputIssuer)
			if test.inputCertificate != nil {
				id = k.AppendCertificate(suite.Ctx, *test.inputCertificate)
			}

			// Test creation of membership
			if !test.expPanic {
				suite.NotPanics(func() { k.CreateMembership(suite.Ctx, id, *test.inputRecipient) })
			} else {
				suite.Panics(func() { k.CreateMembership(suite.Ctx, id, *test.inputRecipient) })
			}
		})
	}
}

func (suite *KeeperTestSuite) TestDoubleAddMember() {
	k := suite.App.IdentityKeeper
	issuer := suite.TestAccs[0]
	recipient := suite.TestAccs[1]
	id, _ := suite.PrepareCertificate(issuer, nil)
	suite.NotPanics(func() { k.CreateMembership(suite.Ctx, id, recipient) })
	suite.Panics(func() { k.CreateMembership(suite.Ctx, id, recipient) })
}

func (suite *KeeperTestSuite) TestHasPendingMember() {
	k := suite.App.IdentityKeeper
	issuer := suite.TestAccs[0]
	recipient := suite.TestAccs[1]
	id, _ := suite.PrepareCertificate(issuer, &recipient)
	suite.True(k.HasPendingMember(suite.Ctx, id, recipient))
	suite.False(k.HasPendingMember(suite.Ctx, id, issuer))
}
