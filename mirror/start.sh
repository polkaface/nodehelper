#!/usr/bin/env bash
nohup ./polkadot --name "PolkadotDBMirror" --base-path /data/polkadot --pruning=archive  --rpc-cors=all  &
