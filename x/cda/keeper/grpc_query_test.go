package keeper_test

import (
	"context"
	"encoding/binary"
	"testing"

	"github.com/HankBreck/archive/x/cda/types"

	testkeeper "github.com/HankBreck/archive/testutil/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
)

const (
	CREATOR   = "creator address"
	VALID_CID = "QmSrnQXUtGqsVRcgY93CdWXf8GPE9Zjj7Tg3SZUgLKDN5W"
)

/*
 * Query CDA Tests
 *
 */

func (suite *KeeperTestSuite) TestCdaQuery() {
	queryClient := suite.queryClient
	signers := []*sdk.AccAddress{&suite.TestAccs[0]}
	cdaIds, signerIds := suite.PrepareCdas(signers, 1)
	cda := suite.GetTemplateCda(*signers[0], signerIds)
	cda.Id = 1
	expected := types.QueryCdaResponse{
		Cda: &cda,
	}

	response, err := queryClient.Cda(context.Background(), &types.QueryCdaRequest{Id: cdaIds[0]})
	suite.NoError(err)
	suite.EqualValues(*response, expected)
}

func (suite *KeeperTestSuite) TestCdaQuery_UnsetId() {
	queryClient := suite.queryClient
	owners := []*sdk.AccAddress{&suite.TestAccs[0]}
	ids, _ := suite.PrepareCdas(owners, 1)

	suite.NotContains(ids, 2)

	// Attempt to query for unset ID
	_, err := queryClient.Cda(context.Background(), &types.QueryCdaRequest{Id: 2})

	suite.EqualError(err, "Invalid CdaId. Please ensure the CDA exists for the given ID.")
}

/*
 * Query CDAs Tests
 *
 */

func (suite *KeeperTestSuite) TestQueryCdas() {
	queryClient := suite.queryClient
	owners := []*sdk.AccAddress{&suite.TestAccs[0]}
	ids, _ := suite.PrepareCdas(owners, 5)

	response, err := queryClient.Cdas(context.Background(), &types.QueryCdasRequest{})

	suite.Require().NoError(err)
	suite.Require().Nil(response.Pagination.NextKey)
	suite.Require().EqualValues(5, len(response.CDAs))

	for i := 0; i < len(response.CDAs); i++ {
		suite.Require().Contains(ids, response.CDAs[i].Id)
		suite.Require().EqualValues(owners[0].String(), response.CDAs[i].Creator)
	}
}

func (suite *KeeperTestSuite) TestQueryCdas_Paginate() {
	queryClient := suite.queryClient
	firstOwners := []*sdk.AccAddress{&suite.TestAccs[0]}
	secondOwners := []*sdk.AccAddress{&suite.TestAccs[0]}
	firstIds, _ := suite.PrepareCdas(firstOwners, 5)
	secondIds, _ := suite.PrepareCdas(secondOwners, 5)

	// Get the first 5 elements, starting with 0
	bzKey := make([]byte, 8)
	binary.BigEndian.PutUint64(bzKey, 0)
	pagination := &query.PageRequest{
		Key:   bzKey,
		Limit: 5,
	}
	response, err := queryClient.Cdas(context.Background(), &types.QueryCdasRequest{Pagination: pagination})

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
	finalResponse, finalErr := queryClient.Cdas(context.Background(), &types.QueryCdasRequest{Pagination: finalPagination})

	suite.Require().NoError(finalErr)
	suite.Require().Len(finalResponse.CDAs, 5)
	for i := 0; i < 5; i++ {
		suite.Require().Equal(secondIds[i], finalResponse.CDAs[i].Id)
		suite.Require().Equal(secondOwners[0].String(), finalResponse.CDAs[i].Creator)
	}
}

func (suite *KeeperTestSuite) TestQueryCdas_PaginateReversed() {
	queryClient := suite.queryClient
	signers := []*sdk.AccAddress{&suite.TestAccs[0]}
	cdaIds, _ := suite.PrepareCdas(signers, 10)

	// Get the last 5 elements stored
	pagination := &query.PageRequest{
		// For some reason we can't pass the max key...
		Limit:      5,
		Reverse:    true,
		CountTotal: true,
	}
	response, err := queryClient.Cdas(context.Background(), &types.QueryCdasRequest{Pagination: pagination})

	suite.Require().EqualValues(uint64(10), response.Pagination.Total)
	suite.Require().NoError(err)
	suite.Require().Len(response.CDAs, 5)
	for i := 0; i < 5; i++ {
		// Order is reversed so need to index ids with len(ids)-i
		suite.Require().Equal(cdaIds[9-i], response.CDAs[i].Id)
		suite.Require().Equal(signers[0].String(), response.CDAs[i].Creator)
	}

	// Fetch the next five using response.Pagination.NextKey
	finalPagination := &query.PageRequest{
		Key:        response.Pagination.NextKey,
		Limit:      5,
		Reverse:    true,
		CountTotal: true,
	}
	finalResponse, finalErr := queryClient.Cdas(context.Background(), &types.QueryCdasRequest{Pagination: finalPagination})

	suite.Require().NoError(finalErr)
	suite.Require().Len(finalResponse.CDAs, 5)
	for i := 0; i < 5; i++ {
		// Order is reversed and this is the second batch,
		// so need to index ids with len(ids)-5-i
		suite.Require().Equal(cdaIds[4-i], finalResponse.CDAs[i].Id)
		suite.Require().Equal(signers[0].String(), finalResponse.CDAs[i].Creator)
	}
}

func (suite *KeeperTestSuite) TestQueryCdas_InvalidKey() {
	queryClient := suite.queryClient
	signers := []*sdk.AccAddress{&suite.TestAccs[0]}
	suite.PrepareCdas(signers, 5)

	invalidKey := make([]byte, 8)
	binary.BigEndian.PutUint64(invalidKey, uint64(6))
	pagination := &query.PageRequest{
		Key: invalidKey,
	}
	response, err := queryClient.Cdas(context.Background(), &types.QueryCdasRequest{Pagination: pagination})

	suite.Require().NoError(err)
	suite.Require().Len(response.CDAs, 0)
}

/*
 * Query CDAs Owned Tests
 *
 */

// func (suite *KeeperTestSuite) TestQueryCdasOwned() {
// 	queryClient := suite.queryClient
// 	owners := []*sdk.AccAddress{&suite.TestAccs[0]}
// 	ids := suite.PrepareCdasForOwner(owners, 5)

// 	response, err := queryClient.CdasOwned(context.Background(), &types.QueryCdasOwnedRequest{Owner: owners[0].String()})

// 	suite.Require().NoError(err)
// 	suite.Require().Equal(len(ids), len(response.Ids))

// 	for i := 0; i < len(ids); i++ {
// 		suite.Require().Contains(response.Ids, ids[i])
// 	}
// }

// func (suite *KeeperTestSuite) TestQueryCdasOwned_InvalidOwner() {
// 	queryClient := suite.queryClient
// 	owner := "invalid address"

// 	_, err := queryClient.CdasOwned(context.Background(), &types.QueryCdasOwnedRequest{Owner: owner})
// 	suite.Require().EqualError(err, "decoding bech32 failed: invalid character in string: ' '")
// }

// func (suite *KeeperTestSuite) TestQueryCdasOwned_OwnerNotFound() {
// 	queryClient := suite.queryClient
// 	owner := suite.TestAccs[0]

// 	res, err := queryClient.CdasOwned(context.Background(), &types.QueryCdasOwnedRequest{Owner: owner.String()})

// 	suite.Require().Nil(res.Ids)
// 	suite.Require().NoError(err)
// }

// func (suite *KeeperTestSuite) TestQueryCdasOwned_EmptyStringRequest() {
// 	queryClient := suite.queryClient

// 	_, err := queryClient.CdasOwned(context.Background(), &types.QueryCdasOwnedRequest{Owner: ""})
// 	suite.Require().EqualError(err, "empty address string is not allowed")
// }

// func (suite *KeeperTestSuite) TestQueryCdasOwned_EmptyRequest() {
// 	queryClient := suite.queryClient

// 	_, err := queryClient.CdasOwned(context.Background(), &types.QueryCdasOwnedRequest{})
// 	suite.Require().EqualError(err, "empty address string is not allowed")
// }

/*
 * Query Approval Tests
 *
 * TODO: Add nonexistent owner check
 */

func (suite *KeeperTestSuite) TestApprovalQuery() {
	queryClient := suite.queryClient
	signers := []*sdk.AccAddress{&suite.TestAccs[0]}

	// Setup test env
	cdaIds, signerIds := suite.PrepareCdas(signers, 1)
	suite.ApproveCda(*signers[0], cdaIds[0], signerIds[0])

	// Perform the test
	req := types.QueryApprovalRequest{
		CdaId:    cdaIds[0],
		SignerId: signerIds[0],
	}
	res, err := queryClient.Approval(context.Background(), &req)
	suite.NoError(err)
	suite.True(res.Approved)
}

func (suite *KeeperTestSuite) TestApprovalQuery_NonexistentCda() {
	queryClient := suite.queryClient
	req := types.QueryApprovalRequest{}
	res, err := queryClient.Approval(context.Background(), &req)
	suite.Nil(res)
	suite.EqualError(err, "CDA with ID 0 not found: Invalid CdaId. Please ensure the CDA exists for the given ID.")
}

// Test SigningData
func (suite *KeeperTestSuite) TestSigningData() {
	k := suite.App.CdaKeeper
	goCtx := sdk.WrapSDKContext(suite.Ctx)
	cdaIds, _ := suite.PrepareCdas([]*sdk.AccAddress{&suite.TestAccs[0]}, 1)

	// Mock signing data in storage
	var expectedSigningData types.RawSigningData
	expectedSigningData.UnmarshalJSON([]byte("hello world"))
	k.SetSigningData(suite.Ctx, cdaIds[0], expectedSigningData)

	// Test SigningData query
	res, err := suite.queryClient.SigningData(goCtx, &types.QuerySigningDataRequest{Id: cdaIds[0]})
	suite.NoError(err)
	suite.Equal(expectedSigningData.Bytes(), res.SigningData)
}

func (suite *KeeperTestSuite) TestSigningData_InvalidId() {
	k := suite.App.CdaKeeper
	goCtx := sdk.WrapSDKContext(suite.Ctx)
	contract := suite.PrepareContract()

	// Mock signing data in storage
	var expectedSchema types.RawSigningData
	expectedSchema.UnmarshalJSON([]byte("hello world"))
	k.SetSigningData(suite.Ctx, contract.Id, expectedSchema)

	// Test SigningDataSchema query
	invalidId := contract.Id + 1
	res, err := suite.queryClient.SigningData(goCtx, &types.QuerySigningDataRequest{Id: invalidId})
	suite.EqualError(err, types.ErrNonExistentSigningData.Wrapf("id %d not found", invalidId).Error())
	suite.Nil(res)
}

// Test SigningDataSchema
func (suite *KeeperTestSuite) TestSigningDataSchema() {
	k := suite.App.CdaKeeper
	goCtx := sdk.WrapSDKContext(suite.Ctx)
	contract := suite.PrepareContract()

	// Mock signing data in storage
	var expectedSchema types.RawSigningData
	expectedSchema.UnmarshalJSON([]byte("hello world"))
	k.SetSigningDataSchema(suite.Ctx, contract.Id, expectedSchema)

	// Test SigningDataSchema query
	res, err := suite.queryClient.SigningDataSchema(goCtx, &types.QuerySigningDataSchemaRequest{Id: contract.Id})
	suite.NoError(err)
	suite.Equal(expectedSchema.Bytes(), res.Schema)
}

func (suite *KeeperTestSuite) TestSigningDataSchema_InvalidId() {
	k := suite.App.CdaKeeper
	goCtx := sdk.WrapSDKContext(suite.Ctx)
	contract := suite.PrepareContract()

	// Mock signing data in storage
	var expectedSchema types.RawSigningData
	expectedSchema.UnmarshalJSON([]byte("hello world"))
	k.SetSigningDataSchema(suite.Ctx, contract.Id, expectedSchema)

	// Test SigningDataSchema query
	invalidId := contract.Id + 1
	res, err := suite.queryClient.SigningDataSchema(goCtx, &types.QuerySigningDataSchemaRequest{Id: invalidId})
	suite.EqualError(err, types.ErrNonExistentSigningData.Wrapf("id %d not found", invalidId).Error())
	suite.Nil(res)
}

/*
 * Query Params Tests
 *
 */

func TestParamsQuery(t *testing.T) {
	keeper, ctx := testkeeper.CdaKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	keeper.SetParams(ctx, params)

	response, err := keeper.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
