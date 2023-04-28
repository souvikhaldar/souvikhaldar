---
title: Elements of Bitcoin
date: 2023-04-28
draft: false
---
Bitcoin is the cryptocurrency that set out a revolution in the world of finance as well as technology. It was created by someone named Satoshi Nakamoto, who is still a mystery. Let's discuss some important concepts of Bitcoin. The average size of each block in [Bitcoin](Bitcoin.md) is 1 MB.
The process of mining is generating a new block to be added to the blockchain so that, the hash of 
1. Timestamp
2. Previous block Hash
3. Merkle root
4. Nonce
5. Version
6. Difficulty bits
comes out having a certain number of leading zeros called difficulty "bits". All the above items are known except the nonce, hence figuring out the nonce is the challenge. 

The target number of leading zeros is computed by:
target = coefficient * 2 ^ (8*(index-3))
where,
index = first two bytes
coefficient = rest of the bytes
Eg bits - 0x170e2632
index= 17 in hex, so 23 in decimal
coefficient= 0e2632 in hex 
The difficulty is set according to the computing power available throughout the network and it changes typically over 2016 blocks. 

For transferring bitcoins,ß scripts are sent, `scriptPubKey` is the script for transferring/sending bitcoin and `scriptSig` is used for verification or claiming the received bitcoin. Transactions are lucidly explained here https://developer.bitcoin.org/devguide/transactions.html.

Bitcoin network has a random topology and runs over TCP.

Transactions submitted to a node eventually get sent to most of the nodes on the network by a process called flooding in which the given node shares its transactions with its neighbour which then shares with its neighbours and so on. 
Once a block is successfully mined by any miner, it floods the new block to all nodes. Each node then verifies that it has the correct nonce then adds it to its local copy of the blockchain. Now, other nodes will remove the transactions that are part of the new block and continue to mine the next block. 
Since multiple miners are mining simultaneously it is possible that different nodes accept different block as the next node and hence forks might get created. This forking is resolved by accepting the **longest chain** as the blockchain. 

## References
1. https://drive.google.com/file/d/1o8xBkDpGgx58Z8DFezOs0xEG2GLd4Ld1/view 
2. https://drive.google.com/file/d/1xlLDYve7Pfx3-e09MaYZsv91aRUN3Qcq/view
