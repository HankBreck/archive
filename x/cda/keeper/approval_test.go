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

// Assert fails with error on a CdaId that does not exist
func (suite *KeeperTestSuite) TestSetApproval_NonexistentCdaId() {
	k := suite.App.CdaKeeper
	signers := []*sdk.AccAddress{&suite.TestAccs[0]}
	cdaIds, signerIds := suite.PrepareCdas(signers, 1)

	err := k.SetApproval(suite.Ctx, cdaIds[0], signerIds[0])
	suite.EqualError(err, "Invalid CdaId. Please ensure the CDA exists for the given ID.")
}
