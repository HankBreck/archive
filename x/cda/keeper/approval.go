package keeper

import (
	"bytes"
	"strconv"

	"github.com/HankBreck/archive/x/cda/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Adds the approval for the (CDA, owner) pair
//
// Returns an error if:
//
// (1) The CDA does not exist
// (2) The message's SigningData does not match the stored SigningData
// (3) The Creator has already approved the CDA
func (k Keeper) SetApproval(ctx sdk.Context, msg *types.MsgApproveCda) error {
	// Validate Creator address
	msgSigner := sdk.MustAccAddressFromBech32(msg.Creator)

	// Ensure the CDA exists
	cda, err := k.GetCDA(ctx, msg.CdaId)
	if err != nil {
		return err
	}

	// Only allow approvals when in the pending state
	if cda.Status != types.CDA_Pending {
		return types.ErrInvalidCdaStatus.Wrap("The CDA must have a status of pending to be approved")
	}

	// Ensure signing data matches
	metadata, err := k.GetSigningData(ctx, msg.CdaId)
	if err != nil {
		return err
	}
	if !bytes.Equal(metadata.Bytes(), msg.SigningData.Bytes()) {
		return types.ErrInvalidSigningData
	}

	// Ensure the sender is a valid signer
	includesSender := false
	for _, party := range cda.SigningParties {
		if msgSigner.String() == party {
			includesSender = true
		}
	}
	if !includesSender {
		return sdkerrors.ErrUnauthorized.Wrapf("Signer is not an owner of cda %d", msg.CdaId)
	}

	// Check if msgSigner has already approved the CDA
	if k.HasApproval(ctx, msg.CdaId, msgSigner) {
		return types.ErrExistingApproval
	}

	// If not, update the store to include their address
	k.uncheckedSetApproval(ctx, msg.CdaId, msgSigner)
	return nil
}

// Checks if the store contains an entry for signer.
// Returns true if an entry is found
func (k Keeper) HasApproval(ctx sdk.Context, cdaId uint64, signer sdk.AccAddress) bool {
	store := k.getApprovalStore(ctx, cdaId)
	bzApproval := store.Get(signer.Bytes())
	return bzApproval != nil
}

func (k Keeper) uncheckedSetApproval(ctx sdk.Context, cdaId uint64, signer sdk.AccAddress) {
	store := k.getApprovalStore(ctx, cdaId)
	store.Set(signer.Bytes(), []byte("x"))
}

func (k Keeper) getApprovalStore(ctx sdk.Context, cdaId uint64) prefix.Store {
	keySuffix := strconv.FormatUint(cdaId, 10)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.CDAApprovalKey+keySuffix))
	return store
}
