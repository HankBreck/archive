package keeper_test

import (
	"context"

	"archive/x/cda/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Cases
// 	Nonexistent owner
//

func (suite *KeeperTestSuite) TestApprovalQuery() {
	queryClient := suite.queryClient
	owners := []*sdk.AccAddress{&suite.TestAccs[0]}
	ids := suite.PrepareCdasForOwner(owners, 1)
	err := suite.ApproveCda(ids[0], owners[0])
	suite.NoError(err)

	req := types.QueryApprovalRequest{
		CdaId: ids[0],
		Owner: owners[0].String(),
	}
	res, err := queryClient.Approval(context.Background(), &req)
	suite.NoError(err)
	suite.True(res.Approved)
}

func (suite *KeeperTestSuite) TestApprovalQuery_NonexistentCda() {
	queryClient := suite.queryClient
	req := types.QueryApprovalRequest{
		CdaId: 0, // unset
		Owner: suite.TestAccs[0].String(),
	}
	res, err := queryClient.Approval(context.Background(), &req)
	suite.Nil(res)
	suite.EqualError(err, "Could not find the cda with an id of 0: key not found")
}

func (suite *KeeperTestSuite) TestApprovalQuery_EmptyRequest() {
	queryClient := suite.queryClient
	req := types.QueryApprovalRequest{}
	res, err := queryClient.Approval(context.Background(), &req)
	suite.EqualError(err, "empty address string is not allowed")
	suite.Nil(res)
}
