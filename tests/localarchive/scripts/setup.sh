#!/bin/sh

echo "starting setup.sh"

CHAIN_ID=localarchive
ARCHIVE_HOME=$HOME/.archived
CONFIG_FOLDER=$OSMOSIS_HOME/config
MONIKER=val
STATE='false'

MNEMONIC="bottom loan skill merry east cradle onion journey palm apology verb edit desert impose absurd oil bubble sweet glove shallow size build burst effort"

if [[ ! -d $CONFIG_FOLDER ]]
then
    echo $MNEMONIC | archived init -o --chain-id=$CHAIN_ID --home $ARCHIVE_HOME --recover $MONIKER
fi

archived start --home $ARCHIVE_HOME
