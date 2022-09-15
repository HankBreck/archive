package keeper_test

import (
	"archive/x/cda/types"
	goctx "context"
)

const (
	CREATOR   = "creator address"
	VALID_CID = "QmSrnQXUtGqsVRcgY93CdWXf8GPE9Zjj7Tg3SZUgLKDN5W"
)

func (suite *KeeperTestSuite) TestCdaQuery() {
	queryClient := suite.queryClient
	owner := suite.TestAccs[0]
	ids := suite.PrepareCdasForOwner(owner, 1)

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
	owner := suite.TestAccs[0]
	ids := suite.PrepareCdasForOwner(owner, 1)

	suite.Require().NotContains(ids, 1)

	// Attempt to query for unset ID
	_, err := queryClient.Cda(goctx.Background(), &types.QueryCdaRequest{Id: 1})

	suite.Require().EqualError(err, "not found")
}
