# Setup WASP Node

## Download and Install Wasp v0.2.5 and wasp-cli

Install RockDB <https://github.com/facebook/rocksdb/blob/master/INSTALL.md>

<https://wiki.iota.org/smart-contracts/guide/chains_and_nodes/running-a-node#download-wasp>

```SHELL
git clone -b v0.2.5 https://github.com/iotaledger/wasp
cd wasp
make install # or make build if you don't want the path
```

## Run A Wasp node

Run A Wasp node will connecting to goshimmer using TXStream.

you can find testnet endpoint here <https://wiki.iota.org/smart-contracts/guide/chains_and_nodes/testnet#endpoints>

for now we using `goshimmer.sc.iota.org:5000`

```shell
make run-wasp
```

Run local Go Shimmer and Wasp node by Docker (Optional)

```shell
cd tools/devnet
HOST=0.0.0.0 docker-compose up -d
```

## Configuring wasp-cli

Set up Go Shimmer API

```shell
wasp-cli init

# Set go shimmer api
wasp-cli set goshimmer.api https://api.goshimmer.sc.iota.org
```

For local Go Shimmer api (Optional)

```shell
# or for local go shimmer
wasp-cli set goshimmer.api 127.0.0.1:8080

# set wasp address for a local node
wasp-cli set wasp.0.api 127.0.0.1:9090
wasp-cli set wasp.0.nanomsg 127.0.0.1:5550
wasp-cli set wasp.0.peering 127.0.0.1:4000
```

Requesting fund

```shell
# Request fund
wasp-cli request-funds
wasp-cli balance
```

## Deploy a Chain

Trust node

```Shell
# Trust node
wasp-cli peering info
wasp-cli peering trust {PubKey} 127.0.0.1:4000
# or
wasp-cli peering trust $(wasp-cli peering info | grep PubKey | sed -e "s/^PubKey\:\s//") 127.0.0.1:4000

# checking trusted nodes
wasp-cli peering list-trusted
```

Deploy Chain

```shell
# Deploy The Chain
# `committee` will correspond to wasp.0, wasp.1 in `wasp-cli.json`
# `quorum` is the minimum amount of nodes required to form a consensus (recommend floor(N*2/3)+1 whre N = number of nodes)
wasp-cli chain deploy --committee=0 --quorum=1 --chain=iscp-lottery-chain --description="ISCP Lottery Chain"

# Deposit money to the chain
wasp-cli chain deposit IOTA:10000

# Set wasp-cli chain ID you can find chain id from `http://127.0.0.1:7000/chains` usr/pass: wasp/wasp
wasp-cli set chains.test-chain {chain-id}
wasp-cli set chain test-chain
```

## Deactivate chain

```SHELL
wasp-cli chain deactivate
```

## Resources

<https://wiki.iota.org/wasp/overview>
