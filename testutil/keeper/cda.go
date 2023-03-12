package keeper

import (
	"testing"

	"github.com/HankBreck/archive/x/cda/keeper"
	"github.com/HankBreck/archive/x/cda/types"

	identitykeeper "github.com/HankBreck/archive/x/identity/keeper"
	identitytypes "github.com/HankBreck/archive/x/identity/types"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"
)

func CdaKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	idStoreKey := sdk.NewKVStoreKey(identitytypes.StoreKey)
	idMemStoreKey := storetypes.NewMemoryStoreKey(identitytypes.MemStoreKey)
	idParamsSubspace := typesparams.NewSubspace(
		cdc,
		identitytypes.Amino,
		idStoreKey,
		idMemStoreKey,
		"IdentityParams",
	)
	idKeeper := identitykeeper.NewKeeper(
		cdc,
		idStoreKey,
		idMemStoreKey,
		idParamsSubspace,
	)

	paramsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		storeKey,
		memStoreKey,
		"CdaParams",
	)
	k := keeper.NewKeeper(
		cdc,
		storeKey,
		memStoreKey,
		paramsSubspace,
		idKeeper,
		nil, // TODO: need wasm keeper
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, ctx
}
