package keeper_test

import (
	"archive/x/cda/types"
	goctx "context"
)

func (suite *KeeperTestSuite) TestQueryCdasOwned() {
	println("in testquerycdasowned")
	queryClient := suite.queryClient
	owner := suite.TestAccs[0]
	ids := suite.PrepareCdasForOwner(owner, 5)

	response, err := queryClient.CdasOwned(goctx.Background(), &types.QueryCdasOwnedRequest{Owner: owner.String()})

	// Ensure query ran successfully
	suite.Require().NoError(err)

	// Ensure same length
	suite.Require().Equal(len(ids), len(response.Ids))

	// Ensure every item in `ids` is in `response.Ids`
	for i := 0; i < len(ids); i++ {
		suite.Require().Contains(response.Ids, ids[i])
	}
}
