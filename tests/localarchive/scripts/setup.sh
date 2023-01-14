#!/bin/sh

echo "starting setup.sh"

CHAIN_ID=localarchive
ARCHIVE_HOME=$HOME/.archive
CONFIG_FOLDER=$ARCHIVE_HOME/config
MONIKER=val
STATE='false'

MNEMONIC="bottom loan skill merry east cradle onion journey palm apology verb edit desert impose absurd oil bubble sweet glove shallow size build burst effort"

edit_genesis() {
	GENESIS=$CONFIG_FOLDER/genesis.json

	# Fund initial validator
    archived add-genesis-account archive12smx2wdlyttvyzvzg54y2vnqwq2qjatekl5jhc 5000000000token,1000000000stake --home $ARCHIVE_HOME

    # Create initial validator
    echo $MNEMONIC | archived keys add $MONIKER --recover --keyring-backend=test --home $ARCHIVE_HOME
    archived gentx $MONIKER 100000000stake --keyring-backend=test --chain-id=$CHAIN_ID --home $ARCHIVE_HOME

    # Create initial genesis file
    archived collect-gentxs --home $ARCHIVE_HOME
}

if [[ ! -d $CONFIG_FOLDER ]]
then
    echo "Initializing validator!"	
    echo $MNEMONIC | archived init -o --chain-id=$CHAIN_ID --home $ARCHIVE_HOME --recover $MONIKER
    edit_genesis
fi

echo "Starting blockchain!"

# TODO: Move this to config like here: https://github.com/osmosis-labs/osmosis/blob/main/tests/localosmosis/scripts/setup.sh#L90
archived start --home $ARCHIVE_HOME --rpc.laddr tcp://0.0.0.0:26657 & 

wait
