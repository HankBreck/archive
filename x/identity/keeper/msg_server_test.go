package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/HankBreck/archive/testutil/keeper"
	"github.com/HankBreck/archive/x/identity/keeper"
	"github.com/HankBreck/archive/x/identity/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.IdentityKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

// Register Issuer
//		nil message
//		Invalid creator address
//		Duplicate issuer
// Issuer Certificate
//		nil message
//		Invalid creator address
//		Invalid issuer (not registered)
//		Invalid recipient address
//		Overwrites preset cert ID
//		Incremental IDs
//		Recipient is pending member
//		Recipient is operator
// Accept Identity
//		nil message
//		Invalid creator address
//		Invalid sender (not pending member)
//		Invalid cert ID
//		Success adds sender to accepted members
//		Success removes sender from pending members
//		Success keeps sender as an operator
// Reject Identity
//		nil message
//		Invalid creator address
//		Invalid sender (not pending member)
//		Invalid cert ID
//		Success removes sender from pending members
//		Success removes sender from operators
// Renounce Identity
//		nil message
//		Invalid creator address
//		Invalid cert ID
//		Not an accepted member
//		Sender still an operator
//		Success removes sender from accepted members
// Update Members
//		nil message
//		Invalid creator address
//		Invalid cert ID
//		Fail if sender not an operator or issuer
//		Fail if one of toAdd addrs invalid
//		Fail if one of toRemove addrs invalid
//		Fail if one of toRemove addrs is an operator
//		Success adds and removes members correctly
//		Think about add / remove order
// Update Operators
//		nil message
//		Invalid creator address
//		Invalid cert ID
//		Fail if sender not an operator or issuer
//		Fail if one of toAdd addrs invalid
//		Fail if one of toRemove addrs invalid
//		Fail when adding duplicate operators
//		Fail when adding an operator that is not an accepted member
//		Success adds and removed members correctly
//		Think about add / remove order
