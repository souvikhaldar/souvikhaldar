---
title: Basic Cryptographic Primitives in Blockchain
date: 2023-01-31
draft: false
---
Hash functions are used to connect blocks in a tamper proof way and also in digitally signing documents. [Bitcoin](Bitcoin.md) uses SHA256 hashing algorithm for mining. The ideal properties of a hash function:
1. Collition free - The time taken to generate same hash is proportional to number of bits in the hash.  
2. Puzzle friendly- It is difficult to find k such that Z = H(M || k), where Z and M are known and M is concatenated with k.
3. Deterministic- Always yields the same value on same input.  
4. Hiding- Nearly impossible to guess the input value looking at the output hash.  
In Public Key Encryption or Public Key Infrastructure (PKI), the encryption of messages is done with the public key and decryption is done using the private key.

A cryptographic hash pointer is the location to which some information is stored and is also the hash of the information stored. If m is the information and h(m) is its hash, then h(m) is the cryptographic hash pointer to location where m is stored. 
![](/images/Pasted%20image%2020230424132821.png)

This way we can link two blocks of data in a tamper proof way, because if data in previous block is changed then the hash will change and as a result the link will be broken. 
We can make this tampering computationally difficult by concatenating a nonce to the data in the block and then putting certain restriction to the output hash, like the output hash should has x number of leading zeroes.

![](/images/Pasted%20image%2020230424133117.png)

A merkle tree works in a similar fashion where each parent node is hash of its child nodes. In the world of blockchain, each such nodes is a transaction and each block of the blockchain contains the root hash, nonce, previous hash and the hash of the entire current block. See the image below:-

![](/images/Pasted%20image%2020230424131453.png)

Blockchain is a hashchain which is connected by hash pointers.  

## References
1. https://youtu.be/uTsAiyZ_cZ4 
2. Slides- https://drive.google.com/file/d/1uNevL5jVBU538JbzjpMIgI-ErhBiJHrk/view
