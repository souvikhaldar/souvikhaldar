---
title: Distributed Hole Punching
date: 2022-12-16
draft: false
---
The web of interconnected devices is called the internet. But this interconnection currently heavily depends on servers and hence internet has become heavily monopolised as few tech giants hold the most computation power to serve. This centralisation is popularly termed web2. The utopia of the next generation of the internet is web3 where computers can actually talk to each other to serve as purposes that web2 currently does. 
But how do computers talk to each other directly given the fact that most of the computers are behind firewall or NAT (Network Address Translator)? Here comes the concept of **Hole Punching**, which helps in achieving that hence it can be thought of the basis of web3.
Here, I'll try to explain Hole Punching as very simple way in following steps:
1. Figure out if the computer can be connected directly: If a computer can be connected to directly then hole punching is not required. To figure this out two components of libp2p is required called [identify](https://github.com/libp2p/go-libp2p/tree/master/p2p/protocol/identify) and [autoNAT](https://github.com/libp2p/go-libp2p/tree/master/p2p/host/autonat). The node figures a range of its possible public addess using identify then asks neighbours to dial it. If one of the address succeeds, autoNAT declares that hole punching is not required. If hole punching is required we proceed with the following steps.
2. Advertising public address via relay nodes: Using Kademlia Distributed Hash Table, the private node finds the closest relay nodes (in [IPFS](IPFS) any public node can serve as relay). The node would advestise its address as `/<relay-addr>/p2p-circuit/<peer-id>` , which other nodes can use to connect to it via this relay node. 
3. Establishing direct connection: Once both nodes are connected via relay as mentioned in step 2, both send each other a `connect` message that contains each others address. Then both send each other `sync` message directly using the address obtained in others `connect` message. This concludes the process of **hole punching**.


## References
1. https://docs.libp2p.io/concepts/nat/hole-punching/