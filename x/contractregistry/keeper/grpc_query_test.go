package keeper_test

import (
	"archive/x/contractregistry/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestParamsQuery() {
	k := suite.App.ContractregistryKeeper
	params := types.DefaultParams()
	k.SetParams(suite.Ctx, params)

	response, err := k.Params(sdk.WrapSDKContext(suite.Ctx), &types.QueryParamsRequest{})
	suite.NoError(err)
	suite.Equal(&types.QueryParamsResponse{Params: params}, response)
}