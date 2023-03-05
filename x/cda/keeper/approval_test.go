package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Assert expected behavior when setting approving for the first time
func (suite *KeeperTestSuite) TestSetApproval() {
	signers := []*sdk.AccAddress{&suite.TestAccs[0]}
	cdaIds, signerIds := suite.PrepareCdas(signers, 1)
	k := suite.App.CdaKeeper

	err := k.SetApproval(suite.Ctx, cdaIds[0], signerIds[0])
	suite.NoError(err)
}

// Assert fails with error when attempting to approve twice
func (suite *KeeperTestSuite) TestSetApproval_ApproveTwice() {
	signers := []*sdk.AccAddress{&suite.TestAccs[0]}
	cdaIds, signerIds := suite.PrepareCdas(signers, 1)
	k := suite.App.CdaKeeper

	err := k.SetApproval(suite.Ctx, cdaIds[0], signerIds[0])
	suite.NoError(err)

	// Attempt to sign a second time
	err = k.SetApproval(suite.Ctx, cdaIds[0], signerIds[0])
	suite.EqualError(err, "The address has already given approval for this CDA.")
}

func (suite *KeeperTestSuite) TestHasApproval() {
	signers := []*sdk.AccAddress{&suite.TestAccs[0]}
	cdaIds, signerIds := suite.PrepareCdas(signers, 1)
	k := suite.App.CdaKeeper

	// Ensure false when approval has not been received
	suite.False(k.HasApproval(suite.Ctx, cdaIds[0], signerIds[0]))

	// Update state to set approval from signer
	k.SetApproval(suite.Ctx, cdaIds[0], signerIds[0])

	// Ensure true after update
	suite.True(k.HasApproval(suite.Ctx, cdaIds[0], signerIds[0]))
}
