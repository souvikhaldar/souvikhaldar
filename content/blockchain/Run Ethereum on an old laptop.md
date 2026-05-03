---
title: How To Run Ethereum On An Old Laptop
date: 2023-01-06
draft: false
description: "A practical walkthrough for running an Ethereum node on an old laptop with geth, Prysm, SSD storage, and home-network configuration."
---
*A decentralized immutable append-only public ledger.*  
By running a node that connects to the blockchain network, you become part of it and help decentralize it further, propagating the principle further, pat on back!  
My other usecase is, I want to connect to the network directly so that I can deploy my smart contracts directly and not use other node as service services for deployment like Infura.io  
So here is how I approached in joining the ethereum main network (aka mainnet) using my old personal laptop.  
I'm using my old [HP x360](HP%20x360) laptop to run Ethereum on it. It has 16 GB RAM with 512 GB SSD secondary memory, along with it I've inserted [[Sandisk Extreme 2 TB SSD]]. 
I'm following [this](https://ethereum.org/en/developers/docs/nodes-and-clients/run-a-node/) article to follow the process of setting it up. 
These are the following steps I'm undertaking: 
1. Format the SSD it to `ext4` using the command `sudo mkfs -t ext4 /dev/sda2`. 
2. Static mount the external SSD to `/mnt/eth`.  
3. Then do `sudo chown -R user:user /mnt/eth` to own it. Now you can read/write to disk with user privilege.  
4. Enabled the **UPnP** option on the router so that the node can be connected with other peers.  
5. Run `geth` as the Execution Client using the command `geth --config ~/.eth/config.toml --http --http.api eth,net,engine,admin` with the following config in the `config.toml` file (some details have been removed for safety purpose):  


	```
	[Eth]
	NetworkId = 1
	SyncMode = "snap"
	EthDiscoveryURLs = []
	SnapDiscoveryURLs = []
	NoPruning = false
	NoPrefetch = false
	TxLookupLimit = 2350000
	LightPeers = 100
	UltraLightFraction = 75
	DatabaseCache = 512
	DatabaseFreezer = ""
	TrieCleanCache = 154
	TrieCleanCacheJournal = "triecache"
	TrieCleanCacheRejournal = 3600000000000
	TrieDirtyCache = 256
	TrieTimeout = 3600000000000
	SnapshotCache = 102
	Preimages = false
	FilterLogCacheSize = 32
	EnablePreimageRecording = false
	RPCGasCap = 50000000
	RPCEVMTimeout = 5000000000
	RPCTxFeeCap = 1e+00

	[Eth.Miner]
	GasFloor = 0
	GasCeil = 30000000
	GasPrice = 1000000000
	Recommit = 2000000000
	Noverify = false

	[Eth.Ethash]
	CacheDir = "ethash"
	CachesInMem = 2
	CachesOnDisk = 3
	CachesLockMmap = false
	DatasetDir = "/mnt/eth/ethereum/.ethash"
	DatasetsInMem = 1
	DatasetsOnDisk = 2
	DatasetsLockMmap = false
	PowMode = 0
	NotifyFull = false

	[Eth.TxPool]
	Locals = []
	NoLocals = false
	Journal = "transactions.rlp"
	Rejournal = 3600000000000
	PriceLimit = 1
	PriceBump = 10
	AccountSlots = 16
	GlobalSlots = 5120
	AccountQueue = 64
	GlobalQueue = 1024
	Lifetime = 10800000000000

	[Eth.GPO]
	Blocks = 20
	Percentile = 60
	MaxHeaderHistory = 1024
	MaxBlockHistory = 1024
	MaxPrice = 500000000000
	IgnorePrice = 2

	[Node]
	DataDir = "/mnt/eth/ethereum"
	IPCPath = "geth.ipc"
	HTTPHost = ""
	HTTPPort = 8545
	HTTPVirtualHosts = ["localhost"]
	HTTPModules = ["net", "web3", "eth"]
	AuthAddr = "localhost"
	AuthPort = 8551
	AuthVirtualHosts = ["localhost"]
	WSHost = ""
	WSPort = 8546
	WSModules = ["net", "web3", "eth"]
	GraphQLVirtualHosts = ["localhost"]

	[Node.P2P]
	MaxPeers = 30
	NoDiscovery = false
	BootstrapNodes = []
	BootstrapNodesV5 = []
	StaticNodes = []
	TrustedNodes = []
	ListenAddr = ":30303"
	DiscAddr = ""
	EnableMsgEvents = false

	[Node.HTTPTimeouts]
	ReadTimeout = 30000000000
	ReadHeaderTimeout = 30000000000
	WriteTimeout = 30000000000
	IdleTimeout = 120000000000

	[Metrics]
	HTTP = "127.0.0.1"
	Port = 6060
	InfluxDBEndpoint = "http://localhost:8086"
	InfluxDBDatabase = "geth"
	InfluxDBUsername = "test"
	InfluxDBPassword = "test"
	InfluxDBTags = "host=localhost"
	InfluxDBToken = "test"
	InfluxDBBucket = "geth"
	InfluxDBOrganization = "geth"  

	```


6. Run the consensus client:  
	   1. `mkdir prysm && cd prysm; curl https://raw.githubusercontent.com/prysmaticlabs/prysm/master/prysm.sh --output prysm.sh && chmod +x prysm.sh`   
           2. `./prysm.sh beacon-chain --execution-endpoint=/mnt/ssd/ethereum/geth.ipc --checkpoint-sync-url=https://beaconstate.ethstaker.cc --datadir=/mnt/ssd/ethereum/consensus/prysm`  



This is how the final results looks to me and I <3 it! Yay! This is actually my macbook where I do all my work, but here I've sshed into my HP laptop connected to Tailscale, which is another piece of software I love!

{{< x user="s0uvikhaldar" id="1613389779969937409" >}}

**Happy Hacking!**

## References
1. https://ethereum.org/en/developers/docs/nodes-and-clients/
