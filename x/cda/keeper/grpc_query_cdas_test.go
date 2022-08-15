package keeper_test

import (
	testkeeper "archive/testutil/keeper"
	"archive/x/cda/types"
	"encoding/binary"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
)

func TestCdasQuery(t *testing.T) {
	keeper, ctx := testkeeper.CdaKeeper(t)

	testkeeper.PopulateCdas(CREATOR, 5, keeper, ctx)

	// Success on empty request
	req := types.QueryCdasRequest{}
	res, err := keeper.Cdas(sdk.WrapSDKContext(ctx), &req)
	require.Nil(t, err)
	require.Equal(t, len(res.CDAs), 5)
	require.Equal(t, res.CDAs[0].Id, uint64(0))
	require.Nil(t, res.Pagination.NextKey)

	// Reset res and err
	res = nil
	err = nil

	// Create the byte key for querying CDAs with ids from 2 to 4
	byteId := make([]byte, 8)
	binary.BigEndian.PutUint64(byteId, 2)

	// Success viewing CDAs with ids from 2 to 4
	req = types.QueryCdasRequest{
		Pagination: &query.PageRequest{
			Key: byteId,
		},
	}
	res, err = keeper.Cdas(sdk.WrapSDKContext(ctx), &req)
	require.Nil(t, err)
	require.Equal(t, len(res.CDAs), 3)
	require.Equal(t, res.CDAs[0].Id, uint64(2))
	require.Equal(t, res.CDAs[2].Id, uint64(4))
}
