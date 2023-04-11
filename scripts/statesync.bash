#!/bin/bash
# microtick and bitcanna contributed significantly here.
# Pebbledb state sync script.
set -uxe

# Set Golang environment variables.
export GOPATH=~/go
export PATH=$PATH:~/go/bin

# Install Umma with pebbledb 
go mod edit -replace github.com/tendermint/tm-db=github.com/notional-labs/tm-db@136c7b6
go mod tidy
go install -ldflags '-w -s -X github.com/cosmos/cosmos-sdk/types.DBBackend=pebbledb' -tags pebbledb ./...

# NOTE: ABOVE YOU CAN USE ALTERNATIVE DATABASES, HERE ARE THE EXACT COMMANDS
# go install -ldflags '-w -s -X github.com/cosmos/cosmos-sdk/types.DBBackend=rocksdb' -tags rocksdb ./...
# go install -ldflags '-w -s -X github.com/cosmos/cosmos-sdk/types.DBBackend=badgerdb' -tags badgerdb ./...
# go install -ldflags '-w -s -X github.com/cosmos/cosmos-sdk/types.DBBackend=boltdb' -tags boltdb ./...

# Initialize chain.
ummad init test

# Get Genesis OLD SOLUTION
#wget https://download.dimi.sh/juno-phoenix2-genesis.tar.gz # TODO:
#tar -xvf juno-phoenix2-genesis.tar.gz
#mv juno-phoenix2-genesis.json "$HOME/.umma/config/genesis.json"

# NEW SOLUTION
wget -O $HOME/.ummad/config/genesis.json https://raw.githubusercontent.com/umma-chain/mainnet/main/genesis.json



# Get "trust_hash" and "trust_height". TODO:
INTERVAL=1000
LATEST_HEIGHT="$(curl -s http://135.125.3.192:26657/block | jq -r .result.block.header.height)"
BLOCK_HEIGHT="$((LATEST_HEIGHT-INTERVAL))"
TRUST_HASH="$(curl -s "http://135.125.3.192:26657/block?height=$BLOCK_HEIGHT" | jq -r .result.block_id.hash)"

# Print out block and transaction hash from which to sync state.
echo "trust_height: $BLOCK_HEIGHT"
echo "trust_hash: $TRUST_HASH"

# Export state sync variables. TODO:
export UMMAD_STATESYNC_ENABLE=true
export UMMAD_P2P_MAX_NUM_OUTBOUND_PEERS=200
export UMMAD_STATESYNC_RPC_SERVERS="https://rpc-juno-ia.notional.ventures:443,https://juno-rpc.polkachu.com:443"
export UMMAD_STATESYNC_TRUST_HEIGHT=$BLOCK_HEIGHT
export UMMAD_STATESYNC_TRUST_HASH=$TRUST_HASH

# Fetch and set list of seeds from chain registry.
UMMAD_P2P_SEEDS="$(curl -s https://raw.githubusercontent.com/cosmos/chain-registry/master/juno/chain.json | jq -r '[foreach .peers.seeds[] as $item (""; "\($item.id)@\($item.address)")] | join(",")')"
export UMMAD_P2P_SEEDS

# Start chain.
ummad start --x-crisis-skip-assert-invariants --db_backend pebbledb
