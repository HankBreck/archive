package keeper_test

import (
	"testing"

	testkeeper "github.com/HankBreck/archive/testutil/keeper"
	"github.com/HankBreck/archive/x/identity/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
)

// Test Params

func TestParamsQuery(t *testing.T) {
	keeper, ctx := testkeeper.IdentityKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	keeper.SetParams(ctx, params)

	response, err := keeper.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}

// Test IdentityMembers

func (suite *KeeperTestSuite) TestQueryIdentityMembers_Pending() {
	suite.SetupTest()
	queryClient := suite.queryClient
	context := sdk.WrapSDKContext(suite.Ctx)

	// Setup certificate & membership for test
	issuer := suite.TestAccs[0]
	recipient := suite.TestAccs[1]
	id, _ := suite.PrepareCertificate(issuer, &recipient)
	req := types.QueryIdentityMembersRequest{
		Id:        id,
		IsPending: true,
	}

	// Test the query
	res, err := queryClient.IdentityMembers(context, &req)
	suite.NoError(err)
	suite.Len(res.Members, 1)
	suite.Equal(res.Members[0], recipient.String())
}

func (suite *KeeperTestSuite) TestQueryIdentityMembers_Accepted() {
	suite.SetupTest()
	queryClient := suite.queryClient
	context := sdk.WrapSDKContext(suite.Ctx)

	// Setup certificate & membership for test
	issuer := suite.TestAccs[0]
	recipient := suite.TestAccs[1]
	id, _ := suite.PrepareCertificate(issuer, &recipient)
	req := types.QueryIdentityMembersRequest{
		Id:        id,
		IsPending: false,
	}
	expectedMembers := suite.TestAccs[2:4]
	suite.SetMembers(id, expectedMembers)

	// Test the query
	res, err := queryClient.IdentityMembers(context, &req)
	suite.NoError(err)
	suite.Len(res.Members, len(expectedMembers))
	for i := range expectedMembers {
		suite.Equal(expectedMembers[i].String(), res.Members[i])
	}
}

func (suite *KeeperTestSuite) TestQueryIdentityMembers_Pagination() {
	suite.SetupTest()
	queryClient := suite.queryClient
	goCtx := sdk.WrapSDKContext(suite.Ctx)
	queryLimit := uint64(2)
	recoveredMembers := []string{}

	// Setup certificate & membership for test
	issuer := suite.TestAccs[0]
	recipient := suite.TestAccs[1]
	id, _ := suite.PrepareCertificate(issuer, &recipient)
	expectedMembers := suite.TestAccs[2:6]
	suite.SetMembers(id, expectedMembers)

	// Test the first query
	res, err := queryClient.IdentityMembers(goCtx, &types.QueryIdentityMembersRequest{
		Id:         id,
		IsPending:  false,
		Pagination: &query.PageRequest{Limit: queryLimit},
	})
	suite.NoError(err)
	suite.Len(res.Members, int(queryLimit))
	suite.NotNil(res.Pagination.NextKey)
	recoveredMembers = append(recoveredMembers, res.Members...)

	// Test the second query
	res, err = queryClient.IdentityMembers(goCtx, &types.QueryIdentityMembersRequest{
		Id:         id,
		IsPending:  false,
		Pagination: &query.PageRequest{Key: res.Pagination.NextKey, Limit: queryLimit},
	})
	suite.NoError(err)
	suite.Len(res.Members, int(queryLimit))
	suite.Nil(res.Pagination.NextKey)
	recoveredMembers = append(recoveredMembers, res.Members...)

	// Ensure recovered membership lists match (order not guaranteed to be the same)
	suite.Len(recoveredMembers, len(expectedMembers))
	for _, addr := range expectedMembers {
		suite.Contains(recoveredMembers, addr.String())
	}
}

func (suite *KeeperTestSuite) TestQueryIdentityMembers_InvalidCertificateId() {
	suite.SetupTest()
	queryClient := suite.queryClient
	context := sdk.WrapSDKContext(suite.Ctx)

	// Setup certificate & membership for test
	issuer := suite.TestAccs[0]
	recipient := suite.TestAccs[1]
	id, _ := suite.PrepareCertificate(issuer, &recipient)

	invalidId := id + 1

	// Test the query
	res, err := queryClient.IdentityMembers(context, &types.QueryIdentityMembersRequest{
		Id:        invalidId,
		IsPending: true,
	})
	suite.EqualError(err, sdkerrors.ErrNotFound.Wrapf("A certificate with an ID of %d was not found", invalidId).Error())
	suite.Nil(res)
}

// Test Issuers

func (suite *KeeperTestSuite) TestQueryIssuers() {
	suite.SetupTest()
	queryClient := suite.queryClient
	goCtx := sdk.WrapSDKContext(suite.Ctx)
	queryLimit := uint64(2)
	recoveredIssuers := []string{}

	// Mock issuers in our test
	expectedIssuers, _ := suite.MockIssuers(10)

	// Test first query
	res, err := queryClient.Issuers(goCtx, &types.QueryIssuersRequest{
		Pagination: &query.PageRequest{
			Limit: queryLimit,
		},
	})
	suite.NoError(err)
	suite.Len(res.Issuers, int(queryLimit))
	suite.NotNil(res.Pagination.NextKey)
	recoveredIssuers = append(recoveredIssuers, res.Issuers...)

	// Test second query
	res, err = queryClient.Issuers(goCtx, &types.QueryIssuersRequest{
		Pagination: &query.PageRequest{
			Key:   res.Pagination.NextKey,
			Limit: queryLimit,
		},
	})
	suite.NoError(err)
	suite.Len(res.Issuers, int(queryLimit))
	suite.NotNil(res.Pagination.NextKey)
	recoveredIssuers = append(recoveredIssuers, res.Issuers...)

	// Test final query
	newQueryLimit := uint64(6)
	res, err = queryClient.Issuers(goCtx, &types.QueryIssuersRequest{
		Pagination: &query.PageRequest{
			Key:   res.Pagination.NextKey,
			Limit: newQueryLimit,
		},
	})
	suite.NoError(err)
	suite.Len(res.Issuers, int(newQueryLimit))
	suite.Nil(res.Pagination.NextKey)
	recoveredIssuers = append(recoveredIssuers, res.Issuers...)

	// Ensure recovered issuer list matches expected (order not guaranteed to be the same)
	suite.Len(recoveredIssuers, len(expectedIssuers))
	for _, addr := range expectedIssuers {
		suite.Contains(recoveredIssuers, addr.String())
	}
}

// Test IssuerInfo
func (suite *KeeperTestSuite) TestQueryIssuerInfo() {
	suite.SetupTest()
	goCtx := sdk.WrapSDKContext(suite.Ctx)
	queryClient := suite.queryClient
	k := suite.App.IdentityKeeper

	// Mock issuer in storage
	issuerAddr := suite.TestAccs[0]
	expectedIssuer := types.Issuer{
		Creator:     issuerAddr.String(),
		Name:        "Issuer #1",
		MoreInfoUri: "https://google.com/issuers/1",
	}
	k.SetIssuer(suite.Ctx, expectedIssuer)

	// Ensure recovered info matches
	res, err := queryClient.IssuerInfo(goCtx, &types.QueryIssuerInfoRequest{Issuer: issuerAddr.String()})
	suite.NoError(err)
	suite.Equal(expectedIssuer, *res.IssuerInfo)
}

func (suite *KeeperTestSuite) TestQueryIssuerInfo_InvalidIssuerAddress() {
	suite.SetupTest()
	goCtx := sdk.WrapSDKContext(suite.Ctx)
	queryClient := suite.queryClient
	k := suite.App.IdentityKeeper

	// Mock issuer in storage
	k.SetIssuer(suite.Ctx, types.Issuer{
		Creator:     suite.TestAccs[0].String(),
		Name:        "Issuer #1",
		MoreInfoUri: "https://google.com/issuers/1",
	})

	// Ensure recovered info matches
	invalidAddress := "invalid address"
	res, err := queryClient.IssuerInfo(goCtx, &types.QueryIssuerInfoRequest{Issuer: invalidAddress})
	suite.EqualError(err, "decoding bech32 failed: invalid character in string: ' '")
	suite.Nil(res)
}

func (suite *KeeperTestSuite) TestQueryIssuerInfo_NonIssuerAddress() {
	suite.SetupTest()
	goCtx := sdk.WrapSDKContext(suite.Ctx)
	queryClient := suite.queryClient
	k := suite.App.IdentityKeeper

	// Mock issuer in storage
	k.SetIssuer(suite.Ctx, types.Issuer{
		Creator:     suite.TestAccs[0].String(),
		Name:        "Issuer #1",
		MoreInfoUri: "https://google.com/issuers/1",
	})

	// Ensure recovered info matches
	nonIssuerAddress := suite.TestAccs[1].String()
	res, err := queryClient.IssuerInfo(goCtx, &types.QueryIssuerInfoRequest{Issuer: nonIssuerAddress})
	suite.EqualError(err, sdkerrors.ErrNotFound.Wrapf("No Issuer found for address %s", nonIssuerAddress).Error())
	suite.Nil(res)
}

// Test Identity

// TODO:
//		Test Identity
//		Test Operators
// 		Test MemberRole
// 		Test IsFrozen
