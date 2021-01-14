#!/usr/bin/env bash
ps aux |grep polkadot |grep -v grep | awk '{print $2}' | xargs kill -2
