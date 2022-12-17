---
title: Foundations of Blockchain
date: 2022-11-17
draft: false
---

# Introduction
Blockchain unlocks a whole new world of use-cases of Internet that we know but powers we still don't. It is like a distributed computing environment that is not owned by no central authority, a true democracy in the digital world. 
To understand the blockchain stack, it would be convinient to break it down into multiple layers but there is no strict distinction between them and this layering is just for purpose of reference. 

# Layers of Blockchain
1. Layer 0: The internet is going to be the foundation of blockchain stack. It is the medium by which peers can interact amoung each other.
2. Layer 1: The consensus layer is layer where [[Bitcoin]] and [[Ethereum]] lie. It is effectively a mechanism to keep chain in sync accross the distributed computers.
3. Layer 2: The scaling layer is the layer where we try to optimize the use of layer 1 and scale it well to be used by large number of users. 
4. Layer 3: Application is the third layer where we build applications on top of blockchain stack, like Distributed Finance i.e [[DeFi]], Decentralized Autonomous Organization i.e [[DAO]], etc.

# Digital Signature Scheme
The ability to communicate reliably and securely in a peer-to-peer network is very important because any of the peers can have mal intentions. The digital signature scheme is used for this purpose and its primary constituents are the following algorithms:
1. Key-generation algorithm: We need to generate a pair of public and private keys using some algorithm. We can use `ssh-keygen` for it.  $f(random\space seed)=(private key, public key)$
2. Signing algorithm: The message to be delivered and private key of the sender is used to create the signature, and then the message in clear text along with this signature is sent. So, under this scheme, the signature is not only dependent on the sender but also on the content of the message being sent. $f(message, private key) = mesage, signature$ 
3. Signature verification: In order to verify that the signature actually belongs to the sender this algorithm is used, where if we provide the input as signature, message and the public key of the sender we can get a boolen response indicating verified or not. $f(mesage,signature,public key of sender)= true/false$ 

# State Machine Replication problem

A [State Machine](https://en.wikipedia.org/wiki/Finite-state_machine) is an abstract machine that be in exactly one state at any given time. For e.g. A database is a state machine that should have a certain set of data (its state) at any given point of time, even after performing any number of operations on it like insertion, updation, deletion, etc.  
But in case of blockchain, the state machine is referred to the history of all the transactions ordered by timestamp or "Ledger".
Since in blockchain each participating node is supposed to have a local copy of the history (or ledger) hence it is replicated. This replication of the ledger is termed as State machine replication.
The problem that arises out of this replication is the probability of inconsistency of the ledger across nodes since they can be located anywhere on the planet. The ideal solution to SMR problem should satisfy the following conditions:
	1. Consistency: Each node should have the same copy of the ledger or each copy can be a prefix of the same history.
	2. Liveness: Each valid transaction should get added to the ledger eventually.


# Assumptions for the solution of SMR problem
The solution for SMR problem needs a few assumptions to begin with, which can be relaxed aferwards. There are the assumptions:
1. Permissioned- The number of nodes that going to run the protocol are known before the running starts. The names of the nodes are distinct, like IP addresses. Nodes need to have permission before running the protocol, unlike permissionless protocols like [Bitcoin](Bitcoin.md), [Ethereum](Ethereum), etc which any node can join.  
2. Public-key infrastructure - All the participating nodes have the public-private key pair and every node knows each other's public key.  
3. Synchronous - All nodes share a global clock and messages sent to each other always arrives by next time iterval, where time is divided into intervals of some duration, say 10 seconds.
4. All nodes are honest- All nodes are running the protocol in exactly the same way as expected. This is kinda ridiculous and will be relaxed really quick.  

A **Dis-honest** node is a node that does not follow the protocol. They are also known as faulty nodes. There are roughly three kinds of faults:  
1. Crash fault: When a node crashes.  
2. Omission fault: When a node does not pass the message (or transaction) that it was supposed to pass.  
3. Byzantine fault: A Byzantine faulty node is a node whose behavior can not be determined. It can do almost anything and we can make no assumptions.  

Now the goal can shift from all honest nodes to having some byzantine nodes and the network still works. 

## References
1. https://youtu.be/KNJGPI0fuFA
