package app

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/ignite/cli/ignite/pkg/cosmoscmd"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

var defaultGenesisBz []byte
var defaultEncodingConfig = cosmoscmd.MakeEncodingConfig(ModuleBasics)

func getDefaultGenesisStateBytes() []byte {
	if len(defaultGenesisBz) == 0 {
		genesisState := NewDefaultGenesisState(defaultEncodingConfig.Marshaler)
		stateBytes, err := json.MarshalIndent(genesisState, "", " ")
		if err != nil {
			panic(err)
		}
		defaultGenesisBz = stateBytes
	}
	return defaultGenesisBz
}

// Setup initializes a new app
func Setup(isCheckTx bool) *App {
	db := dbm.NewMemDB()
	app := New(
		log.NewNopLogger(),
		db,
		nil,
		true,
		map[int64]bool{},
		DefaultNodeHome,
		5,
		defaultEncodingConfig,
		simapp.EmptyAppOptions{},
	)
	if !isCheckTx {
		stateBytes := getDefaultGenesisStateBytes()

		app.InitChain(
			abci.RequestInitChain{
				Validators:      []abci.ValidatorUpdate{},
				ConsensusParams: simapp.DefaultConsensusParams,
				AppStateBytes:   stateBytes,
			},
		)
	}
	return app
}