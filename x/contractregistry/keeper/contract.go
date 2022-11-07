package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"archive/x/contractregistry/types"
)

// AppendContract stores the contract in prefixed storage using contract.Id as the key.
// Expects contract.Id to be unset, so an existing value will be overwritten by this function.
//
// Returns the uint64 id of the contract
func (k Keeper) AppendContract(ctx sdk.Context, contract types.Contract) uint64 {
	count := k.getContractCount(ctx)
	if k.HasContract(ctx, count) {
		panic("Duplicate Contract ID found... this should never happen")
	}
	contract.Id = count

	// TODO: do we need any checks on contract?
	k.uncheckedSetContract(ctx, contract)
	k.setContractCount(ctx, count+1)
	return count
}

// GetContract fetches the contract stored under the key of id. If no contract is found, an error is returned
func (k Keeper) GetContract(ctx sdk.Context, id uint64) (*types.Contract, error) {
	store := k.getContractStore(ctx)
	bzKey := make([]byte, 8)
	binary.BigEndian.PutUint64(bzKey, id)

	var contract types.Contract
	bzContract := store.Get(bzKey)
	if len(bzContract) == 0 {
		return nil, types.ErrNonExistentContract.Wrapf("The ID of %d was not found", id)
	}

	k.cdc.MustUnmarshal(bzContract, &contract)
	return &contract, nil
}

// HasContract returns true if a contract is stored under id, else false
func (k Keeper) HasContract(ctx sdk.Context, id uint64) bool {
	store := k.getContractStore(ctx)
	bzKey := make([]byte, 8)
	binary.BigEndian.PutUint64(bzKey, id)
	return store.Has(bzKey)
}

// GetContracts uses pagination to find the next N contracts
//
// Returns a tuple of: the contracts found, the page response, and an error.
func (k Keeper) GetContracts(ctx sdk.Context, pageReq *query.PageRequest) ([]types.Contract, *query.PageResponse, error) {
	store := k.getContractStore(ctx)
	contracts := []types.Contract{}

	pageRes, err := query.Paginate(store, pageReq, func(_, value []byte) error {
		var contract *types.Contract
		err := k.cdc.Unmarshal(value, contract)
		if err != nil {
			return err
		}
		contracts = append(contracts, *contract)
		return nil
	})
	if err != nil {
		return nil, nil, err
	}

	return contracts, pageRes, nil
}

// Stores the contract with a key of contract.Id. The contract.Id field must be set by a calling function.
// The contract passed as an argument is assumed to be valid, so calling functions must assure this.
func (k Keeper) uncheckedSetContract(ctx sdk.Context, contract types.Contract) {
	store := k.getContractStore(ctx)
	bzKey := make([]byte, 8)
	binary.BigEndian.PutUint64(bzKey, contract.Id)
	bzContract := k.cdc.MustMarshal(&contract)
	store.Set(bzKey, bzContract)
}

func (k Keeper) getContractStore(ctx sdk.Context) prefix.Store {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.ContractKey))
	return store
}

func (k Keeper) getContractCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.ContractCountKey))
	bzCount := store.Get([]byte{0})
	if bzCount == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bzCount)
}

func (k Keeper) setContractCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.ContractCountKey))
	bzCount := make([]byte, 8)
	binary.BigEndian.PutUint64(bzCount, count)
	store.Set([]byte{0}, bzCount)
}
