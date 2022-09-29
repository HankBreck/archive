package keeper_test

import (
	"archive/x/cda/keeper"
	"archive/x/cda/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func approveAll(k keeper.Keeper, ctx sdk.Context, cda *types.CDA, owners []*sdk.AccAddress) {
	// Send approval for each owner
	ownerships := (*cda).Ownership
	for _, owner := range owners {
		msg := types.MsgApproveCda{
			Creator:   owner.String(),
			CdaId:     cda.Id,
			Ownership: ownerships,
		}
		err := k.SetApproval(ctx, &msg)
		if err != nil {
			panic(err)
		}
	}
}

// Assert it works as normal
func (suite *KeeperTestSuite) TestFinalize() {
	k := suite.App.CdaKeeper
	owners := []*sdk.AccAddress{&suite.TestAccs[0], &suite.TestAccs[1]}
	ids := suite.PrepareCdasForOwner(owners, 1)
	cdas := suite.GetCdas(ids)
	cda := cdas[0]

	approveAll(k, suite.Ctx, cda, owners)

	// Send finalize message
	msg := types.MsgFinalizeCda{
		Creator: owners[0].String(),
		CdaId:   cda.Id,
	}
	err := k.Finalize(suite.Ctx, &msg)
	suite.NoError(err)

	queryMsg := types.QueryCdaRequest{
		Id: cda.Id,
	}
	res, err := suite.queryClient.Cda(suite.Ctx.Context(), &queryMsg)
	suite.NoError(err)
	suite.True(res.Cda.Approved)
}

// Assert that it fails on nonexisting CdaId
func (suite *KeeperTestSuite) TestFinalize__NonexistentCdaId() {
	k := suite.App.CdaKeeper
	owners := []*sdk.AccAddress{&suite.TestAccs[0], &suite.TestAccs[1]}
	ids := suite.PrepareCdasForOwner(owners, 1)
	cdas := suite.GetCdas(ids)
	cda := cdas[0]

	approveAll(k, suite.Ctx, cda, owners)

	// Send finalize message with a nonexistent CdaId
	msg := types.MsgFinalizeCda{
		Creator: owners[0].String(),
		CdaId:   cda.Id + 1,
	}
	err := k.Finalize(suite.Ctx, &msg)
	suite.EqualError(err, "Invalid CdaId. Please ensure the CDA exists for the given ID.")

	// Ensure the real CDA was not finalized
	queryMsg := types.QueryCdaRequest{
		Id: cda.Id,
	}
	res, err := suite.queryClient.Cda(suite.Ctx.Context(), &queryMsg)
	suite.NoError(err)
	suite.False(res.Cda.Approved)
}

// Assert that it fails when missing approvals
func (suite *KeeperTestSuite) TestFinalize__MissingApproval() {
	k := suite.App.CdaKeeper
	owners := []*sdk.AccAddress{&suite.TestAccs[0], &suite.TestAccs[1]}
	ids := suite.PrepareCdasForOwner(owners, 1)
	cdas := suite.GetCdas(ids)
	cda := cdas[0]

	// Send approval from owners[0] but not owners[1]
	approveAll(k, suite.Ctx, cda, owners[:1])

	// Send finalize message with a nonexistent CdaId
	msg := types.MsgFinalizeCda{
		Creator: owners[0].String(),
		CdaId:   cda.Id,
	}
	err := k.Finalize(suite.Ctx, &msg)
	expectedErr := types.ErrMissingApproval.Wrapf("The CDA with an ID of %d is missing approval from account %s", cda.Id, owners[1].String())
	suite.EqualError(err, expectedErr.Error())

	// Ensure the CDA was not finalized
	queryMsg := types.QueryCdaRequest{
		Id: cda.Id,
	}
	res, err := suite.queryClient.Cda(suite.Ctx.Context(), &queryMsg)
	suite.NoError(err)
	suite.False(res.Cda.Approved)
}

// Assert that it fails when already finalized
func (suite *KeeperTestSuite) TestFinalize_DoubleApproval() {
	k := suite.App.CdaKeeper
	owners := []*sdk.AccAddress{&suite.TestAccs[0], &suite.TestAccs[1]}
	ids := suite.PrepareCdasForOwner(owners, 1)
	cdas := suite.GetCdas(ids)
	cda := cdas[0]

	approveAll(k, suite.Ctx, cda, owners)

	// Send finalize message
	msg := types.MsgFinalizeCda{
		Creator: owners[0].String(),
		CdaId:   cda.Id,
	}
	err := k.Finalize(suite.Ctx, &msg)
	suite.NoError(err)

	// Assert failure on second finalize message
	err2 := k.Finalize(suite.Ctx, &msg)
	suite.EqualError(err2, "CDA has already been finalized.")
}
