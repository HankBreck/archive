package keeper

import (
	"archive/x/cda/types"
	"encoding/binary"
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Sets the CDA to the finalized state. This requires prior approval from every owner defined in the CDA's Ownership field
//
// Returns an error if:
//
// (1) The CDA does not exist
// (2) Not all owners have approved the CDA
// (3) The CDA has already been finalized
func (k Keeper) FinalizeCda(ctx sdk.Context, msg *types.MsgFinalizeCda) error {
	// Load CDA store
	cdaStore := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.CDAKey))

	bzCdaId := make([]byte, 8)
	binary.BigEndian.PutUint64(bzCdaId, msg.CdaId)
	bzCda := cdaStore.Get(bzCdaId)
	if bzCda == nil {
		return types.ErrNonExistentCdaId
	}
	var cda types.CDA
	if err := k.cdc.Unmarshal(bzCda, &cda); err != nil {
		return err
	}

	// Do not allow multiple final approvals
	if cda.Approved {
		return types.ErrAlreadyFinalized
	}

	// Load approvals store for the CDA
	keySuffix := strconv.FormatUint(cda.Id, 10)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.CDAApprovalKey+keySuffix))

	// Ensure each owner has approved this CDA
	for _, owner := range cda.Ownership {
		acc, err := sdk.AccAddressFromBech32(owner.Owner)
		if err != nil {
			// All addresses should be pre-verified, halt chain if this is not the case
			panic(err)
		}
		if !store.Has(acc.Bytes()) {
			return types.ErrMissingApproval.Wrapf("The CDA with an ID of %d is missing approval from account %s", cda.Id, acc.String())
		}
	}

	// If so, update CDA with Approved set to true and store
	cda.Approved = true
	newBzCda := k.cdc.MustMarshal(&cda)
	cdaStore.Set(bzCdaId, newBzCda)
	return nil
}
