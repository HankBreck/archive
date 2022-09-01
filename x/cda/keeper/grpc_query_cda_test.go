package keeper_test

import (
	testkeeper "archive/testutil/keeper"
	"archive/x/cda/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	CREATOR   = "creator address"
	VALID_CID = "QmSrnQXUtGqsVRcgY93CdWXf8GPE9Zjj7Tg3SZUgLKDN5W"
)

func TestCdaQuery(t *testing.T) {
	keeper, ctx := testkeeper.CdaKeeper(t)

	// Failure on empty request
	emptyRequest := types.QueryCdaRequest{}
	_, err := keeper.Cda(sdk.WrapSDKContext(ctx), &emptyRequest)
	require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))

	// Reset err variable
	err = nil

	// Failure on unset Id
	unsetIdRequest := types.QueryCdaRequest{Id: 0}
	_, err = keeper.Cda(sdk.WrapSDKContext(ctx), &unsetIdRequest)
	require.ErrorIs(t, err, sdkerrors.ErrNotFound)

	// Set CDA for a given id
	expected := types.CDA{
		Creator: CREATOR,
		Cid:     VALID_CID,
		Id:      0,
	}
	id := keeper.AppendCDA(ctx, expected)
	require.Equal(t, id, uint64(0))

	// Reset err variable
	err = nil

	// Call a query with the returned id and asset equality
	validRequest := types.QueryCdaRequest{Id: id}
	actual, err := keeper.Cda(sdk.WrapSDKContext(ctx), &validRequest)
	require.Nil(t, err)

	// Require that actual CDA's fields are the same as expected CDA's fields
	require.Equal(t, actual.Creator, expected.Creator)
	require.Equal(t, actual.Cid, expected.Cid)
	require.Equal(t, actual.Id, expected.Id)
}
