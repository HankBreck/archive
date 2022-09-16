package keeper_test

import (
	"archive/x/cda/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Assert that it fails on nonexisting CdaId
// Assert that it fails when missing approvals
// Assert that it fails when already approved

// Assert it works as normal
func (suite *KeeperTestSuite) TestFinalizeCda() {
	k := suite.App.CdaKeeper
	owners := []*sdk.AccAddress{&suite.TestAccs[0], &suite.TestAccs[1]}
	ids := suite.PrepareCdasForOwner(owners, 1)
	cdas := suite.GetCdas(ids)

	cda := cdas[0]

	// Send approval for each owner
	ownerships := (*cda).Ownership
	for _, owner := range owners {
		msg := types.MsgApproveCda{
			Creator:   owner.String(),
			CdaId:     ids[0],
			Ownership: ownerships,
		}
		err := k.SetApproval(suite.Ctx, &msg)
		suite.NoError(err)
	}

	// Send finalize message
	msg := types.MsgFinalizeCda{
		Creator: owners[0].String(),
		CdaId:   cda.Id,
	}
	err := k.FinalizeCda(suite.Ctx, &msg)
	suite.NoError(err)

	queryMsg := types.QueryCdaRequest{
		Id: cda.Id,
	}
	res, err := suite.queryClient.Cda(suite.Ctx.Context(), &queryMsg)
	suite.NoError(err)
	suite.True(res.Cda.Approved)
}
