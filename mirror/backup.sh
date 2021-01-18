#!/usr/bin/env bash
datetime=`date "+%Y-%m-%d-%H-%M-%S"`
echo $datetime
rm ../backup/*
tgz_file_name="polcadot-"$datetime".UTC.tgz"
tar --use-compress-program=pigz -p 12 -cvf ../backup/$tgz_file_name ../chains/polkadot/db
aws s3 cp ../backup/$tgz_file_name  s3://polkaface-node-db-mirror/ >> ./s3_cp.log
