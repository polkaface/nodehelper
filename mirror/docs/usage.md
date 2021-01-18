# Useage
Polkadot的节点搭建，可以参考官方文档 
[Fast Install Instructions](https://wiki.polkadot.network/docs/en/maintain-sync#fast-install-instructions-linux)

这里我们一Linux（Unbuntu 18.04)为例：

S1: 安装Rust：

    curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh

S2: Clone Polkadot代码：

    git clone https://github.com/paritytech/polkadot

S3: 初始化编译环境(这里指定v0.8.27版本)：

    cd polkadot
    git checkout tags/v0.8.27
    ./scripts/init.sh

S4：编译polkadot程序：

    cargo build --release

S5: 启动polkadot节点：

    ./target/release/polkadot --name "My node's name"

得益于cargo的构建环境，整个过程还是非常简单的。同时主网和测试的相关配置也是直接在程序资源中，
不需要像其他公链一样，要自己手动去配置创世节点信息什么的。直接一键启动。

启动后，可以在[远程监控](https://telemetry.polkadot.io/#list/Polkadot)搜索到你的节点。

其中的Block列可以看到你的节点当前同步到哪个区块高度了。也可以在日志中看到：

    2021-01-18 11:12:40  🔍 Discovered new external address for our node: /ip4/172.18.0.2/tcp/30333/p2p/12D3KooWEdhVmtPqyK4PrCxRZHnVT4mBgSUQwY2sroxfAkH7f1Db
    2021-01-18 11:12:43  ⚙️  Syncing 28.2 bps, target=#3390373 (7 peers), best: #242788 (0x3729…c8e2), finalized #242688 (0x9948…6585), ⬇ 13.2kiB/s ⬆ 16.7kiB/s
    2021-01-18 11:12:48  ⚙️  Syncing 28.5 bps, target=#3390374 (7 peers), best: #242931 (0x1420…6882), finalized #242688 (0x9948…6585), ⬇ 16.5kiB/s ⬆ 21.8kiB/s
    2021-01-18 11:12:53  ⚙️  Syncing 33.7 bps, target=#3390374 (9 peers), best: #243100 (0xfeab…5b00), finalized #242688 (0x9948…6585), ⬇ 11.2kiB/s ⬆ 17.1kiB/s
    2021-01-18 11:12:58  ⚙️  Syncing 30.1 bps, target=#3390375 (9 peers), best: #243251 (0x32d2…cbdc), finalized #243200 (0xe19d…2f2a), ⬇ 11.0kiB/s ⬆ 17.9kiB/s
    2021-01-18 11:13:03  ⚙️  Syncing 33.3 bps, target=#3390376 (9 peers), best: #243418 (0x0e49…b0a3), finalized #243200 (0xe19d…2f2a), ⬇ 9.7kiB/s ⬆ 15.7kiB/s

这里的target高度。


上面会发现，如果要等待区块同步到当前高度：#3,390,401 （2020-01-18）是要等待非常长时间的。
并且上面的节点默认是运行一个同步节点，还不是验证节点，同步节点仅保留了过去256个区块的状态信息。
因此要运行一个全数据的验证节点，这里同步的速度还会更慢。我们测试一台AWS(4C16G)的机器，需要
使用48H才能同步到3M的高度。

因此我们提供了区块数据库镜像的功能，这里，我们用新命令来启动一个archive的全数据节点：

    ./target/release/polkadot --name "NodeName" --base-path ~/polkadot/data --pruning=archive

然后我们进入到节点数据目录：

    cd ~/polkadot/data/chains/polkadot

这里会看到:

    db  keystore  network

几个目录，然后从polkaface查询最新的镜像文件，假设是[polcadot-2021-01-17-11-46-15.UTC.tgz](https://polkaface-node-db-mirror.s3.ap-southeast-1.amazonaws.com/polcadot-2021-01-17-11-46-15.UTC.tgz?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAYXGKVB6E7YW4MX5T%2F20210118%2Fap-southeast-1%2Fs3%2Faws4_request&X-Amz-Date=20210118T011635Z&X-Amz-Expires=259200&X-Amz-SignedHeaders=host&X-Amz-Signature=1a9d70e3e110facc6b40d749b29018a9b778e92ac670ba3c65f49289ba086278)

下载镜像文件：

    wget -c https://polkaface-node-db-mirror.s3.ap-southeast-1.amazonaws.com/polcadot-2021-01-17-11-46-15.UTC.tgz?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAYXGKVB6E7YW4MX5T%2F20210118%2Fap-southeast-1%2Fs3%2Faws4_request&X-Amz-Date=20210118T011635Z&X-Amz-Expires=259200&X-Amz-SignedHeaders=host&X-Amz-Signature=1a9d70e3e110facc6b40d749b29018a9b778e92ac670ba3c65f49289ba086278

    unzip polcadot-2021-01-17-11-46-15.UTC.tgz

然后将里面的 "chains/polkadot/db"目录copy出来，替换这里的db目录。
之后再执行上面的启动命令，然后观察日志就可以看到当前的同步区块已经很接近最新高度了。

这里我们假设你的环境下载这个文件的速度是1OMB/s。那么3M的高度大概30G，也就是大概1H就同步到3M高度
了，比让节点自己同步的48H，节省了两天时间。
