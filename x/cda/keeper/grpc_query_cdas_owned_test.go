package keeper_test

import (
	"archive/x/cda/types"
	goctx "context"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestQueryCdasOwned() {
	queryClient := suite.queryClient
	owners := []*sdk.AccAddress{&suite.TestAccs[0]}
	ids := suite.PrepareCdasForOwner(owners, 5)

	response, err := queryClient.CdasOwned(goctx.Background(), &types.QueryCdasOwnedRequest{Owner: owners[0].String()})

	suite.Require().NoError(err)
	suite.Require().Equal(len(ids), len(response.Ids))

	for i := 0; i < len(ids); i++ {
		suite.Require().Contains(response.Ids, ids[i])
	}
}

func (suite *KeeperTestSuite) TestQueryCdasOwned_InvalidOwner() {
	queryClient := suite.queryClient
	owner := "invalid address"

	_, err := queryClient.CdasOwned(goctx.Background(), &types.QueryCdasOwnedRequest{Owner: owner})
	suite.Require().EqualError(err, "decoding bech32 failed: invalid character in string: ' '")
}

func (suite *KeeperTestSuite) TestQueryCdasOwned_OwnerNotFound() {
	queryClient := suite.queryClient
	owner := suite.TestAccs[0]

	res, err := queryClient.CdasOwned(goctx.Background(), &types.QueryCdasOwnedRequest{Owner: owner.String()})

	suite.Require().Nil(res.Ids)
	suite.Require().NoError(err)
}

func (suite *KeeperTestSuite) TestQueryCdasOwned_EmptyStringRequest() {
	queryClient := suite.queryClient

	_, err := queryClient.CdasOwned(goctx.Background(), &types.QueryCdasOwnedRequest{Owner: ""})
	suite.Require().EqualError(err, "empty address string is not allowed")
}

func (suite *KeeperTestSuite) TestQueryCdasOwned_EmptyRequest() {
	queryClient := suite.queryClient

	_, err := queryClient.CdasOwned(goctx.Background(), &types.QueryCdasOwnedRequest{})
	suite.Require().EqualError(err, "empty address string is not allowed")
}
