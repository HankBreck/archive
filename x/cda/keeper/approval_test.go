package keeper_test

import (
	crtypes "github.com/HankBreck/archive/x/contractregistry/types"

	"github.com/HankBreck/archive/x/cda/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Assert expected behavior when setting approving for the first time
func (suite *KeeperTestSuite) TestSetApproval() {
	owners := []*sdk.AccAddress{&suite.TestAccs[0]}
	ids := suite.PrepareCdasForOwner(owners, 1)
	k := suite.App.CdaKeeper

	msg := types.MsgApproveCda{
		Creator:     owners[0].String(),
		CdaId:       ids[0],
		SigningData: suite.GetSigningData(),
	}
	err := k.SetApproval(suite.Ctx, &msg)
	suite.NoError(err)
}

// Assert fails with error when attempting to approve twice
func (suite *KeeperTestSuite) TestSetApproval_ApproveTwice() {
	owners := []*sdk.AccAddress{&suite.TestAccs[0]}
	ids := suite.PrepareCdasForOwner(owners, 1)
	k := suite.App.CdaKeeper
	msg := types.MsgApproveCda{
		Creator:     owners[0].String(),
		CdaId:       ids[0],
		SigningData: suite.GetSigningData(),
	}
	err := k.SetApproval(suite.Ctx, &msg)
	suite.NoError(err)

	// Attempt to sign a second time
	err2 := k.SetApproval(suite.Ctx, &msg)
	suite.EqualError(err2, "The address has already given approval for this CDA.")
}

// Assert fails with error on a CdaId that does not exist
func (suite *KeeperTestSuite) TestSetApproval_NonexistentCdaId() {
	owners := []*sdk.AccAddress{&suite.TestAccs[0]}
	_ = suite.PrepareCdasForOwner(owners, 1)
	k := suite.App.CdaKeeper
	msg := types.MsgApproveCda{
		Creator:     owners[0].String(),
		CdaId:       uint64(2), // The id 2 is not set in state
		SigningData: suite.GetSigningData(),
	}
	err := k.SetApproval(suite.Ctx, &msg)
	suite.EqualError(err, "Invalid CdaId. Please ensure the CDA exists for the given ID.")
}

// Assert fails with error on invalid signing data
func (suite *KeeperTestSuite) TestSetApproval_InvalidSigningData() {
	owners := []*sdk.AccAddress{&suite.TestAccs[0]}
	ids := suite.PrepareCdasForOwner(owners, 1)
	k := suite.App.CdaKeeper

	var invalidSigningData crtypes.RawSigningData
	invalidSigningData.UnmarshalJSON([]byte(`
	{
		"notOwnerships": [
			{ "owner": "address", "ownership_proportion": 1 },
			{ "owner": "address2", "ownership_proportion": 99 }
		]
	}`))
	msg := types.MsgApproveCda{
		Creator:     owners[0].String(),
		CdaId:       ids[0],
		SigningData: invalidSigningData,
	}
	err := k.SetApproval(suite.Ctx, &msg)
	suite.EqualError(err, "Signing data provided does not match the signing data stored in the CDA.")
}

// Assert fails with error on non-owner Creator
func (suite *KeeperTestSuite) TestSetApproval_UnauthorizedCreator() {
	owners := []*sdk.AccAddress{&suite.TestAccs[0]}
	ids := suite.PrepareCdasForOwner(owners, 1)
	k := suite.App.CdaKeeper
	msg := types.MsgApproveCda{
		Creator:     suite.TestAccs[1].String(),
		CdaId:       ids[0],
		SigningData: suite.GetSigningData(),
	}
	err := k.SetApproval(suite.Ctx, &msg)
	suite.EqualError(err, "Signer is not an owner of cda 0: unauthorized")
}

// Assert fails with error on invalid cda.Status
func (suite *KeeperTestSuite) TestSetApproval_WrongStatus() {
	signers := []*sdk.AccAddress{&suite.TestAccs[0]}
	id := suite.PrepareVoidedCdaForSigners(signers)
	k := suite.App.CdaKeeper

	msg := types.MsgApproveCda{
		Creator:     signers[0].String(),
		CdaId:       id,
		SigningData: suite.GetSigningData(),
	}
	err := k.SetApproval(suite.Ctx, &msg)
	suite.EqualError(err, "The CDA must have a status of pending to be approved: The CDA's status did not match the expected status.")
}
