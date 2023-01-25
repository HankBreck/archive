package simulation

import (
	"math/rand"

	"github.com/HankBreck/archive/x/cda/keeper"
	"github.com/HankBreck/archive/x/cda/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgFinalizeCda(
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgFinalizeCda{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the FinalizeCda simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "FinalizeCda simulation not implemented"), nil, nil
	}
}
