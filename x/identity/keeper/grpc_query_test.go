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
	for _, addr := range expectedMembers {
		suite.Contains(res.Members, addr.String())
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
func (suite *KeeperTestSuite) TestQueryIdentity() {
	suite.SetupTest()
	goCtx := sdk.WrapSDKContext(suite.Ctx)
	queryClient := suite.queryClient
	id, _ := suite.PrepareCertificate(suite.TestAccs[0], &suite.TestAccs[1])

	// Pulled from /app/apptesting/identity.go:28
	expectedCertificate := types.Certificate{
		Id:                id,
		IssuerAddress:     suite.TestAccs[0].String(),
		Salt:              "salt",
		MetadataSchemaUri: "google.com/metadata-schema",
		Hashes: []*types.HashEntry{
			{Field: "field1", Hash: "hash1"},
			{Field: "field2", Hash: "hash2"},
		},
	}

	// Test the query
	res, err := queryClient.Identity(goCtx, &types.QueryIdentityRequest{Id: id})
	suite.NoError(err)
	suite.Equal(expectedCertificate, *res.Certificate)
}

func (suite *KeeperTestSuite) TestQueryIdentity_InvalidCertificateId() {
	suite.SetupTest()
	goCtx := sdk.WrapSDKContext(suite.Ctx)
	queryClient := suite.queryClient
	id, _ := suite.PrepareCertificate(suite.TestAccs[0], &suite.TestAccs[1])

	// Test the query
	invalidId := id + 1
	res, err := queryClient.Identity(goCtx, &types.QueryIdentityRequest{Id: invalidId})
	suite.EqualError(err, types.ErrNonexistentCertificate.Wrapf("no certificate found for ID: %d", invalidId).Error())
	suite.Nil(res)
}

// Test Operators
func (suite *KeeperTestSuite) TestQueryOperators() {
	suite.SetupTest()
	goCtx := sdk.WrapSDKContext(suite.Ctx)
	queryClient := suite.queryClient
	id, _ := suite.PrepareCertificate(suite.TestAccs[0], &suite.TestAccs[1])
	recoveredOperators := []string{}

	// Mock operators in storage
	expectedOperators := suite.TestAccs[2:6]
	expectedOperators = append(expectedOperators, suite.TestAccs[1]) // initial recipient is also an operator
	suite.AddOperators(id, expectedOperators)

	// Test the first query
	res, err := queryClient.Operators(goCtx, &types.QueryOperatorsRequest{
		Id:         id,
		Pagination: &query.PageRequest{Limit: 2},
	})
	suite.NoError(err)
	suite.Len(res.Operators, 2)
	suite.NotNil(res.Pagination.NextKey)
	recoveredOperators = append(recoveredOperators, res.Operators...)

	// Test the seconds query
	res, err = queryClient.Operators(goCtx, &types.QueryOperatorsRequest{
		Id: id,
		Pagination: &query.PageRequest{
			Key:   res.Pagination.NextKey,
			Limit: 3,
		},
	})
	suite.NoError(err)
	suite.Len(res.Operators, 3)
	suite.Nil(res.Pagination.NextKey)
	recoveredOperators = append(recoveredOperators, res.Operators...)

	// Test that all operators were succesfully recovered
	suite.Len(recoveredOperators, len(expectedOperators))
	for _, addr := range expectedOperators {
		suite.Contains(recoveredOperators, addr.String())
	}
}

func (suite *KeeperTestSuite) TestQueryOperators_InvalidCertificateId() {
	suite.SetupTest()
	goCtx := sdk.WrapSDKContext(suite.Ctx)
	queryClient := suite.queryClient
	id, _ := suite.PrepareCertificate(suite.TestAccs[0], &suite.TestAccs[1])

	// Test the query
	invalidId := id + 1
	res, err := queryClient.Operators(goCtx, &types.QueryOperatorsRequest{
		Id: invalidId,
	})
	suite.EqualError(err, types.ErrNonexistentCertificate.Error())
	suite.Nil(res)
}

// Test MemberRole
func (suite *KeeperTestSuite) TestQueryMemberRole() {
	suite.SetupTest()
	goCtx := sdk.WrapSDKContext(suite.Ctx)
	queryClient := suite.queryClient

	// Mock storage for test
	id, _ := suite.PrepareCertificate(suite.TestAccs[0], &suite.TestAccs[1])
	expectedOperator := suite.TestAccs[2]
	suite.AddOperators(id, []sdk.AccAddress{expectedOperator})
	expectedAcceptedMember := suite.TestAccs[3]
	suite.SetMembers(id, []sdk.AccAddress{expectedAcceptedMember})
	expectedPendingMember := suite.TestAccs[4]
	suite.SetPendingMembers(id, []sdk.AccAddress{expectedPendingMember})

	// Test operator query
	res, err := queryClient.MemberRole(goCtx, &types.QueryMemberRoleRequest{
		Id:     id,
		Member: expectedOperator.String(),
	})
	suite.NoError(err)
	suite.Equal("operator", res.Role)

	// Test accepted member query
	res, err = queryClient.MemberRole(goCtx, &types.QueryMemberRoleRequest{
		Id:     id,
		Member: expectedAcceptedMember.String(),
	})
	suite.NoError(err)
	suite.Equal("accepted-member", res.Role)

	// Test pending member query
	res, err = queryClient.MemberRole(goCtx, &types.QueryMemberRoleRequest{
		Id:     id,
		Member: expectedPendingMember.String(),
	})
	suite.NoError(err)
	suite.Equal("pending-member", res.Role)
}

func (suite *KeeperTestSuite) TestQueryMemberRole_InvalidAddress() {
	suite.SetupTest()
	goCtx := sdk.WrapSDKContext(suite.Ctx)
	queryClient := suite.queryClient

	// Mock storage for test
	id, _ := suite.PrepareCertificate(suite.TestAccs[0], &suite.TestAccs[1])

	// Test operator query
	res, err := queryClient.MemberRole(goCtx, &types.QueryMemberRoleRequest{
		Id:     id,
		Member: "invalid address",
	})
	suite.EqualError(err, "decoding bech32 failed: invalid character in string: ' '")
	suite.Nil(res)
}

func (suite *KeeperTestSuite) TestQueryMemberRole_InvalidCertificateId() {
	suite.SetupTest()
	goCtx := sdk.WrapSDKContext(suite.Ctx)
	queryClient := suite.queryClient

	// Mock storage for test
	id, _ := suite.PrepareCertificate(suite.TestAccs[0], &suite.TestAccs[1])
	expectedOperator := suite.TestAccs[2]
	suite.AddOperators(id, []sdk.AccAddress{expectedOperator})
	expectedAcceptedMember := suite.TestAccs[3]
	suite.SetMembers(id, []sdk.AccAddress{expectedAcceptedMember})
	expectedPendingMember := suite.TestAccs[4]
	suite.SetPendingMembers(id, []sdk.AccAddress{expectedPendingMember})

	// Test operator query
	invalidId := id + 1
	res, err := queryClient.MemberRole(goCtx, &types.QueryMemberRoleRequest{
		Id:     invalidId,
		Member: expectedOperator.String(),
	})
	suite.EqualError(err, types.ErrNonexistentCertificate.Wrapf("no certificate found for ID: %d", invalidId).Error())
	suite.Nil(res)
}

func (suite *KeeperTestSuite) TestQueryMemberRole_NotAMember() {
	suite.SetupTest()
	goCtx := sdk.WrapSDKContext(suite.Ctx)
	queryClient := suite.queryClient

	// Mock storage for test
	id, _ := suite.PrepareCertificate(suite.TestAccs[0], &suite.TestAccs[1])
	expectedOperator := suite.TestAccs[2]
	suite.AddOperators(id, []sdk.AccAddress{expectedOperator})
	expectedAcceptedMember := suite.TestAccs[3]
	suite.SetMembers(id, []sdk.AccAddress{expectedAcceptedMember})
	expectedPendingMember := suite.TestAccs[4]
	suite.SetPendingMembers(id, []sdk.AccAddress{expectedPendingMember})

	// Test operator query
	nonMemberAddr := suite.TestAccs[5].String()
	res, err := queryClient.MemberRole(goCtx, &types.QueryMemberRoleRequest{
		Id:     id,
		Member: nonMemberAddr,
	})
	expectedError := sdkerrors.ErrNotFound.Wrapf("account (%s) is not a member of identity %d", nonMemberAddr, id).Error()
	suite.EqualError(err, expectedError)
	suite.Nil(res)
}

// Test IsFrozen

func (suite *KeeperTestSuite) TestQueryIsFrozen() {
	suite.SetupTest()
	goCtx := sdk.WrapSDKContext(suite.Ctx)
	queryClient := suite.queryClient
	k := suite.App.IdentityKeeper

	// Mock storage for test
	id, _ := suite.PrepareCertificate(suite.TestAccs[0], &suite.TestAccs[1])

	// Ensure identities are initially frozen
	res, err := queryClient.IsFrozen(goCtx, &types.QueryIsFrozenRequest{Id: id})
	suite.NoError(err)
	suite.False(res.IsFrozen)

	// Freeze the identity
	k.Freeze(suite.Ctx, id)

	// Ensure query reads that the identity is now frozen
	res, err = queryClient.IsFrozen(goCtx, &types.QueryIsFrozenRequest{Id: id})
	suite.NoError(err)
	suite.True(res.IsFrozen)
}

func (suite *KeeperTestSuite) TestQueryIsFrozen_InvalidCertificateId() {
	suite.SetupTest()
	goCtx := sdk.WrapSDKContext(suite.Ctx)
	queryClient := suite.queryClient

	// Mock storage for test
	id, _ := suite.PrepareCertificate(suite.TestAccs[0], &suite.TestAccs[1])

	// Test the query
	invalidId := id + 1
	res, err := queryClient.IsFrozen(goCtx, &types.QueryIsFrozenRequest{Id: invalidId})
	suite.EqualError(err, types.ErrNonexistentCertificate.Error())
	suite.Nil(res)
}
