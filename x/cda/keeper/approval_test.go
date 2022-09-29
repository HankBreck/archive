package keeper_test

import (
	"archive/x/cda/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Assert expected behavior when setting approving for the first time
func (suite *KeeperTestSuite) TestSetApproval() {
	owners := []*sdk.AccAddress{&suite.TestAccs[0]}
	ids := suite.PrepareCdasForOwner(owners, 1)
	k := suite.App.CdaKeeper
	cdas := suite.GetCdas(ids)

	ownerships := (*cdas[0]).Ownership
	msg := types.MsgApproveCda{
		Creator:   owners[0].String(),
		CdaId:     ids[0],
		Ownership: ownerships,
	}
	err := k.SetApproval(suite.Ctx, &msg)
	suite.NoError(err)
}

// Assert fails with error when attempting to approve twice
func (suite *KeeperTestSuite) TestSetApproval_ApproveTwice() {
	owners := []*sdk.AccAddress{&suite.TestAccs[0]}
	ids := suite.PrepareCdasForOwner(owners, 1)
	k := suite.App.CdaKeeper
	cdas := suite.GetCdas(ids)

	ownerships := (*cdas[0]).Ownership
	msg := types.MsgApproveCda{
		Creator:   owners[0].String(),
		CdaId:     ids[0],
		Ownership: ownerships,
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
	ids := suite.PrepareCdasForOwner(owners, 1)
	k := suite.App.CdaKeeper
	cdas := suite.GetCdas(ids)

	ownerships := (*cdas[0]).Ownership
	msg := types.MsgApproveCda{
		Creator:   owners[0].String(),
		CdaId:     1, // id of 1 does not exist
		Ownership: ownerships,
	}
	err := k.SetApproval(suite.Ctx, &msg)
	suite.EqualError(err, "Invalid CdaId. Please ensure the CDA exists for the given ID.")
}

// Assert fails with error on invalid ownership length
func (suite *KeeperTestSuite) TestSetApproval_InvalidOwnershipLength() {
	owners := []*sdk.AccAddress{&suite.TestAccs[0]}
	ids := suite.PrepareCdasForOwner(owners, 1)
	k := suite.App.CdaKeeper
	cdas := suite.GetCdas(ids)

	ownerships := (*cdas[0]).Ownership
	extraOwnership := types.Ownership{
		Owner:     suite.TestAccs[1].String(),
		Ownership: 10,
	}
	ownerships = append(ownerships, &extraOwnership)
	msg := types.MsgApproveCda{
		Creator:   owners[0].String(),
		CdaId:     ids[0],
		Ownership: ownerships,
	}
	err := k.SetApproval(suite.Ctx, &msg)
	suite.EqualError(err, "Invalid ownership map")
}

// Assert fails with error on mismatched ownerships
func (suite *KeeperTestSuite) TestSetApproval_WrongOwnership() {
	owners := []*sdk.AccAddress{&suite.TestAccs[0]}
	ids := suite.PrepareCdasForOwner(owners, 1)
	k := suite.App.CdaKeeper
	cdas := suite.GetCdas(ids)

	ownerships := (*cdas[0]).Ownership
	ownerships[0].Ownership += 1 // Edit the real Ownership struct
	msg := types.MsgApproveCda{
		Creator:   owners[0].String(),
		CdaId:     ids[0],
		Ownership: ownerships,
	}
	err := k.SetApproval(suite.Ctx, &msg)
	suite.EqualError(err, "Invalid ownership map")
}

// Assert fails with error on non-owner Creator
func (suite *KeeperTestSuite) TestSetApproval_UnauthorizedCreator() {
	owners := []*sdk.AccAddress{&suite.TestAccs[0]}
	ids := suite.PrepareCdasForOwner(owners, 1)
	k := suite.App.CdaKeeper
	cdas := suite.GetCdas(ids)

	ownerships := (*cdas[0]).Ownership
	msg := types.MsgApproveCda{
		Creator:   suite.TestAccs[1].String(),
		CdaId:     ids[0],
		Ownership: ownerships,
	}
	err := k.SetApproval(suite.Ctx, &msg)
	suite.EqualError(err, "Signer is not an owner of cda 0: unauthorized")
}
