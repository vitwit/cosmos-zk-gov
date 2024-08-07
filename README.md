# Anonymous Voting System using Zero Knowledge (ZK)

A governance voting system but the voter's identity should remain anonymous.

## Run Locally
Follow the below steps to run the chain locally.

### Clone
```
git clone https://github.com/vishal-kanna/zk-gov
git fetch
git checkout teja/dev-final
```

### Install Simd Binary
```
make install
```

### Run local testnet
- Modify the scripts/init-simapp script according to your requirements
```
make init-simapp
```

## ZK prover and verifier keys
- Zk prover and verifier keys needs to be setup beforehand. Ideally pk and vk should be generated before chain setup and the vk should set in the chain state in the genesis.
For now, we will be storing those keys in a keys folder

```
mkdir keys
make generate-zk-keys
```

- This will run groth16.setup() defined in /x/zkgov/client/zk/main.go

    Note: we need multiple pk & vk pairs as merkle proof size varies (2 to 32?) depending on the number of commitments. Ideally the commitments size should be constant (some 2^31) and updates should be optimized from O(n) to O(log n).

    For now, we will be generating multiple pk & vk pairs (i.e verifier-2, verifier-3, verifier-4...), use the one needed dynamically

##### Future scope: optimize merkle implementation (constant size of 2^31, update from O(n) to O(log n))

## Relayer
Relayer is a simple http sever which will listen to vote transactions, signs them with a valid chain address and broadcast them to on chain.

##### Why do we need Relayer?
Vote transaction needs to be sent by a different unlinkable address from the registered address (register & vote are explained below). Since most users won't have 2 completely unlinkable address, they can rather use relayer to do the job.

It's not possible for the relayer to manipulate the transaction as ZK proof acts like a guarantee and it is improbable to modify ZK proof without knowing user-known Secret codes.

For more information, spec will be released soon. You may refer that.

### Run Relayer
- Relayer is implemented in /x/zkgov/client/relayer
```
    simd tx zk-gov run-relayer --from {key} --keyring-backend test --chain-id {chain id} -y

    // (or simply run the below command to run alice as relayer)

    make run-alice-relayer
```

- This will start an http server on port 8080. You can change the port using the flag --relayerPort {port}


## Transactions

### Create A proposal
- A simple proposal which can be later voted with YES / NO.

    Note: proposal implementation here is very simple and for testing purpose. More robust implementation is required to be used.

    Note: proposal id is a counter with start from 1 and so on..

```
simd tx zk-gov create-proposal [proposal-title] [proposal-description] --from [address] --keyring-backend test --chain-id [chain id]

(or simple run a make command)

make create-proposal-a
```

### Register Vote
- Register vote commitment, which will be used for anonymity later
```
simd tx zk-gov register-vote [proposal-id] {"YES"/"NO} --from [address you want actually vote] --keyring-backend test --chain-id [chain id]

(or)

make register-alice-vote
make register-bob-vote
make register-sai-vote
make register-teja-vote
```

### Vote
The previous register transaction stores a commitment for a vote but not actual vote itself. Now we will generate a zk proof claiming that we know the values for a valid register commitment. The zk proof will be verified and the vote will be processed now.

##### Note: it's important to remember that whoever signs this transaction doesn't have anything to do the actual vote itself. The actual voter is the previous address who registered the commitment. So this vote should be signed by a different address for anonymity. We can also used relayer for this.

#### Without Relayer
```
simd tx zk-gov vote [proposal id] [address with which the register-vote tx was done] --from [different unlinkable address] --keyring-backend test --chain-id [chain id]

    (or)

make broadcast-alice-vote
make broadcast-bob-vote
```

#### With Relayer

```
simd tx zk-gov vote [proposal id] [address with which the register-vote tx was done] --relayer [relayer address]

or 

make broadcast-sai-vote-via-relayer
make broadcast-teja-vote-via-relayer
```

#### Query Proposal
You can view all the state regarding a proposal including commitments and votes using below command
```
simd q zk-gov get-proposal-info [proposal id]
```
