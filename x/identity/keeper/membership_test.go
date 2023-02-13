package keeper_test

import (
	"github.com/HankBreck/archive/x/identity/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
)

func (suite *KeeperTestSuite) TestCreateMembership() {
	k := suite.App.IdentityKeeper

	// Setup default values
	defaultIssuer := types.Issuer{
		Creator:     suite.TestAccs[0].String(),
		Name:        "Test Issuer",
		MoreInfoUri: "google.com/more-info",
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
		"no_certificate_set": {
			inputIssuer:      &defaultIssuer,
			inputCertificate: nil,
			inputRecipient:   &defaultRecipient,
			expPanic:         true,
		},
		"simple_set": {
			inputIssuer:      &defaultIssuer,
			inputCertificate: &defaultCert,
			inputRecipient:   &defaultRecipient,
			expPanic:         false,
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

func (suite *KeeperTestSuite) TestHasPendingMember() {
	k := suite.App.IdentityKeeper
	issuer := suite.TestAccs[0]
	recipient := suite.TestAccs[1]
	id, _ := suite.PrepareCertificate(issuer, &recipient)
	suite.True(k.HasPendingMember(suite.Ctx, id, recipient))
	suite.False(k.HasPendingMember(suite.Ctx, id, issuer))
}

func (suite *KeeperTestSuite) TestHasMember() {
	k := suite.App.IdentityKeeper
	issuer := suite.TestAccs[0]
	recipient := suite.TestAccs[1]
	id, _ := suite.PrepareCertificate(issuer, &recipient)
	suite.SetMembers(id, []sdk.AccAddress{recipient})

	suite.True(k.HasMember(suite.Ctx, id, recipient))
	suite.False(k.HasPendingMember(suite.Ctx, id, recipient))
}

func (suite *KeeperTestSuite) TestGetMembers() {
	k := suite.App.IdentityKeeper
	issuer := suite.TestAccs[0]
	recipient := suite.TestAccs[1]

	id, _ := suite.PrepareCertificate(issuer, &recipient)

	secondMember := suite.TestAccs[2]
	toAdd := []sdk.AccAddress{secondMember}
	toRemove := []sdk.AccAddress{}
	k.UpdatePendingMembers(suite.Ctx, id, toAdd, toRemove)
	expected := append(toAdd, recipient)

	members, pageRes, err := k.GetMembers(suite.Ctx, id, true, &query.PageRequest{Limit: 1})
	suite.NoError(err)
	suite.NotNil(secondMember.Bytes(), pageRes.NextKey)
	memberAddr, _ := sdk.AccAddressFromBech32(members[0])
	suite.Contains(expected, memberAddr)

	members, pageRes, err = k.GetMembers(suite.Ctx, id, true, &query.PageRequest{Limit: 1, Key: pageRes.NextKey})
	suite.NoError(err)
	suite.Nil(pageRes.NextKey)
	memberAddr, _ = sdk.AccAddressFromBech32(members[0])
	suite.Contains(expected, memberAddr)
}

func (suite *KeeperTestSuite) TestUpdatePendingMembers() {
	// Setup initial certificate
	k := suite.App.IdentityKeeper
	issuer := suite.TestAccs[0]
	recipient := suite.TestAccs[1]
	id, _ := suite.PrepareCertificate(issuer, &recipient)

	// Assert initial pending member is present
	suite.True(k.HasPendingMember(suite.Ctx, id, recipient))

	// Test add & remove pending members
	toAdd := []sdk.AccAddress{suite.TestAccs[2]}
	toRemove := []sdk.AccAddress{recipient}
	err := k.UpdatePendingMembers(suite.Ctx, id, toAdd, toRemove)
	suite.NoError(err)
	suite.False(k.HasPendingMember(suite.Ctx, id, recipient))
	suite.True(k.HasPendingMember(suite.Ctx, id, toAdd[0]))
}

func (suite *KeeperTestSuite) TestUpdatePendingMembers_NonexistentCert() {
	k := suite.App.IdentityKeeper
	recipient := suite.TestAccs[1]
	toAdd := []sdk.AccAddress{suite.TestAccs[2]}
	toRemove := []sdk.AccAddress{recipient}
	err := k.UpdatePendingMembers(suite.Ctx, uint64(10), toAdd, toRemove)
	suite.Error(err)
}

func (suite *KeeperTestSuite) TestUpdateAcceptedMembers() {
	// Setup initial certificate
	k := suite.App.IdentityKeeper
	issuer := suite.TestAccs[0]
	recipient := suite.TestAccs[1]
	id, _ := suite.PrepareCertificate(issuer, &recipient)
	suite.AcceptMembership(id, recipient)
	// remove operator status so recipient can be removed
	k.RemoveOperators(suite.Ctx, id, []sdk.AccAddress{recipient})

	// Assert initial state
	hasInitialMember, _ := k.HasMember(suite.Ctx, id, recipient)
	hasInitialOperator, _ := k.HasOperator(suite.Ctx, id, recipient)
	suite.True(hasInitialMember)
	suite.False(hasInitialOperator)

	// Test add & remove members
	toAdd := []sdk.AccAddress{suite.TestAccs[2]}
	toRemove := []sdk.AccAddress{recipient}
	err := k.UpdateAcceptedMembers(suite.Ctx, id, toAdd, toRemove)
	suite.NoError(err)
	hasRecipient, _ := k.HasMember(suite.Ctx, id, recipient)
	hasNew, _ := k.HasMember(suite.Ctx, id, toAdd[0])
	suite.False(hasRecipient)
	suite.True(hasNew)
}

func (suite *KeeperTestSuite) TestUpdateAcceptedMembers_NonexistentCert() {
	k := suite.App.IdentityKeeper
	recipient := suite.TestAccs[1]
	toAdd := []sdk.AccAddress{suite.TestAccs[2]}
	toRemove := []sdk.AccAddress{recipient}
	err := k.UpdateAcceptedMembers(suite.Ctx, uint64(10), toAdd, toRemove)
	suite.Error(err)
}
