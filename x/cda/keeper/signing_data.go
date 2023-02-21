package keeper

import (
	"encoding/binary"
	"encoding/json"
	"fmt"

	"github.com/HankBreck/archive/x/cda/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/santhosh-tekuri/jsonschema/v5"
)

// SetSigningDataSchema stores signingData with a key of id. The most recent contract stored must have a contract.Id == id.
// This function should only be called from RegisterContract.
//
// Panics on any id that is not equal to the current contract count - 1.
//
// Returns an error if the signingData is nil.
func (k Keeper) SetSigningDataSchema(ctx sdk.Context, signingData types.RawSigningData, id uint64) error {
	// Ensure the signing data references a valid contract
	if !k.HasContract(ctx, id) {
		return types.ErrNonExistentContract.Wrapf("Could not find a contract with an id of %d", id)
	}

	// Ensure there is not a repeat ID
	if k.HasSigningDataSchema(ctx, id) {
		return types.ErrExistingEntry.Wrapf("Signing data already stored for id %d", id)
	}

	err := k.uncheckedSetSigningDataSchema(ctx, signingData, id)
	if err != nil {
		return err
	}
	return nil
}

// GetSigningDataSchema fetches the contract stored under the key of id. If no contract is found, an error is returned
func (k Keeper) GetSigningDataSchema(ctx sdk.Context, id uint64) (types.RawSigningData, error) {
	store := k.getSigningDataSchemaStore(ctx)
	bzKey := make([]byte, 8)
	binary.BigEndian.PutUint64(bzKey, id)

	bzData := store.Get(bzKey)
	if len(bzData) == 0 {
		return nil, types.ErrNonExistentSigningData.Wrapf("Signing data not found for id %d", id)
	}
	var result types.RawSigningData
	result.UnmarshalJSON(bzData)

	return result, nil
}

func (k Keeper) HasSigningDataSchema(ctx sdk.Context, id uint64) bool {
	store := k.getSigningDataSchemaStore(ctx)
	bzKey := make([]byte, 8)
	binary.BigEndian.PutUint64(bzKey, id)

	return store.Has(bzKey)
}

// MatchesSigningDataSchema takes rawInputData and targetContractId as inputs.
//
// Returns true if it matches the schema specified by the contract with an id
// of targetContractId. Returns false if not.
func (k Keeper) MatchesSigningDataSchema(ctx sdk.Context, targetContractId uint64, rawInputData types.RawSigningData) (bool, error) {
	// Fetch schema if it exists
	rawSchema, err := k.GetSigningDataSchema(ctx, targetContractId)
	if err != nil {
		return false, err
	}

	// Compile the json schema
	schema, err := jsonschema.CompileString(fmt.Sprintf("com.archive/%d", targetContractId), string(rawSchema.Bytes()))
	if err != nil {
		return false, err
	}

	// Unmarshal input data
	var inputData map[string]interface{}
	err = json.Unmarshal(rawInputData.Bytes(), &inputData)
	if err != nil {
		return false, err
	}

	// Validate inputData against the schema
	err = schema.Validate(inputData)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Stores the contract's signing data with the contract's ID as the key. The contract.Id field must be set by a calling function.
// The signing data and ID passed as arguments are assumed to be valid, so calling functions must assure this.
func (k Keeper) uncheckedSetSigningDataSchema(ctx sdk.Context, signingData types.RawSigningData, id uint64) error {
	store := k.getSigningDataSchemaStore(ctx)

	bzKey := make([]byte, 8)
	binary.BigEndian.PutUint64(bzKey, id)

	// NOTE: this implementation of MarshalJSON does not seem to return any errors.
	// Might want to add an error if bzData == []byte("null")
	bzData, err := signingData.MarshalJSON()
	if err != nil {
		return err
	}

	store.Set(bzKey, bzData)
	return nil
}

func (k Keeper) getSigningDataSchemaStore(ctx sdk.Context) prefix.Store {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SigningDataKey))
	return store
}
