package keeper_test

import (
	"archive/x/cda/types"
	goctx "context"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	CREATOR   = "creator address"
	VALID_CID = "QmSrnQXUtGqsVRcgY93CdWXf8GPE9Zjj7Tg3SZUgLKDN5W"
)

func (suite *KeeperTestSuite) TestCdaQuery() {
	queryClient := suite.queryClient
	owners := []*sdk.AccAddress{&suite.TestAccs[0]}
	ids := suite.PrepareCdasForOwner(owners, 1)
	owner := owners[0]

	response, err := queryClient.Cda(goctx.Background(), &types.QueryCdaRequest{Id: ids[0]})

	ownerships := make([]*types.Ownership, 1)
	ownerships[0] = &types.Ownership{
		Owner:     owner.String(),
		Ownership: uint64(100),
	}
	cda := types.CDA{
		Creator:    owner.String(),
		Id:         0,
		Cid:        "QmSrnQXUtGqsVRcgY93CdWXf8GPE9Zjj7Tg3SZUgLKDN5W",
		Ownership:  ownerships,
		Expiration: 4123503529000,
	}

	expected := types.QueryCdaResponse{
		Cda: &cda,
	}
	suite.Require().NoError(err)
	suite.Require().EqualValues(*response, expected)
}

func (suite *KeeperTestSuite) TestCdaQuery_UnsetId() {
	queryClient := suite.queryClient
	owners := []*sdk.AccAddress{&suite.TestAccs[0]}
	ids := suite.PrepareCdasForOwner(owners, 1)

	suite.Require().NotContains(ids, 1)

	// Attempt to query for unset ID
	_, err := queryClient.Cda(goctx.Background(), &types.QueryCdaRequest{Id: 1})

	suite.Require().EqualError(err, "not found")
}
