package keeper_test

import (
	"archive/x/cda/types"
	goctx "context"
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/types/query"
)

func (suite *KeeperTestSuite) TestQueryCdas() {
	queryClient := suite.queryClient
	owner := suite.TestAccs[0]
	ids := suite.PrepareCdasForOwner(owner, 5)

	response, err := queryClient.Cdas(goctx.Background(), &types.QueryCdasRequest{})

	suite.Require().NoError(err)
	suite.Require().Nil(response.Pagination.NextKey)
	suite.Require().EqualValues(5, len(response.CDAs))

	for i := 0; i < len(response.CDAs); i++ {
		suite.Require().Contains(ids, response.CDAs[i].Id)
		suite.Require().EqualValues(owner.String(), response.CDAs[i].Creator)
	}
}

func (suite *KeeperTestSuite) TestQueryCdas_Paginate() {
	queryClient := suite.queryClient
	firstOwner := suite.TestAccs[0]
	secondOwner := suite.TestAccs[1]
	firstIds := suite.PrepareCdasForOwner(firstOwner, 5)
	secondIds := suite.PrepareCdasForOwner(secondOwner, 5)

	// Get the first 5 elements, starting with 0
	bzKey := make([]byte, 8)
	binary.BigEndian.PutUint64(bzKey, 0)
	pagination := &query.PageRequest{
		Key:   bzKey,
		Limit: 5,
	}
	response, err := queryClient.Cdas(goctx.Background(), &types.QueryCdasRequest{Pagination: pagination})

	// firstIds references the first 5 elements in storage
	suite.Require().NoError(err)
	suite.Require().Len(response.CDAs, 5)
	for i := 0; i < 5; i++ {
		suite.Require().Equal(firstIds[i], response.CDAs[i].Id)
	}

	// Fetch the next five using response.Pagination.NextKey
	finalPagination := &query.PageRequest{
		Key:        response.Pagination.NextKey,
		Limit:      5,
		CountTotal: true,
	}
	finalResponse, finalErr := queryClient.Cdas(goctx.Background(), &types.QueryCdasRequest{Pagination: finalPagination})

	suite.Require().NoError(finalErr)
	suite.Require().Len(finalResponse.CDAs, 5)
	for i := 0; i < 5; i++ {
		suite.Require().Equal(secondIds[i], finalResponse.CDAs[i].Id)
		suite.Require().Equal(secondOwner.String(), finalResponse.CDAs[i].Creator)
	}
}

func (suite *KeeperTestSuite) TestQueryCdas_PaginateReversed() {
	queryClient := suite.queryClient
	owner := suite.TestAccs[0]
	ids := suite.PrepareCdasForOwner(owner, 10)

	// Get the last 5 elements stored
	pagination := &query.PageRequest{
		// For some reason we can't pass the max key...
		Limit:      5,
		Reverse:    true,
		CountTotal: true,
	}
	response, err := queryClient.Cdas(goctx.Background(), &types.QueryCdasRequest{Pagination: pagination})

	suite.Require().EqualValues(uint64(10), response.Pagination.Total)
	suite.Require().NoError(err)
	suite.Require().Len(response.CDAs, 5)
	for i := 0; i < 5; i++ {
		// Order is reversed so need to index ids with len(ids)-i
		suite.Require().Equal(ids[9-i], response.CDAs[i].Id)
		suite.Require().Equal(owner.String(), response.CDAs[i].Creator)
	}

	// Fetch the next five using response.Pagination.NextKey
	finalPagination := &query.PageRequest{
		Key:        response.Pagination.NextKey,
		Limit:      5,
		Reverse:    true,
		CountTotal: true,
	}
	finalResponse, finalErr := queryClient.Cdas(goctx.Background(), &types.QueryCdasRequest{Pagination: finalPagination})

	suite.Require().NoError(finalErr)
	suite.Require().Len(finalResponse.CDAs, 5)
	for i := 0; i < 5; i++ {
		// Order is reversed and this is the second batch,
		// so need to index ids with len(ids)-5-i
		suite.Require().Equal(ids[4-i], finalResponse.CDAs[i].Id)
		suite.Require().Equal(owner.String(), finalResponse.CDAs[i].Creator)
	}
}

func (suite *KeeperTestSuite) TestQueryCdas_InvalidKey() {
	queryClient := suite.queryClient
	owner := suite.TestAccs[0]
	suite.PrepareCdasForOwner(owner, 5)

	invalidKey := make([]byte, 8)
	binary.BigEndian.PutUint64(invalidKey, uint64(6))
	pagination := &query.PageRequest{
		Key: invalidKey,
	}
	response, err := queryClient.Cdas(goctx.Background(), &types.QueryCdasRequest{Pagination: pagination})

	suite.Require().NoError(err)
	suite.Require().Len(response.CDAs, 0)
}
