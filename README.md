# **Cosmos ZK Voting**

A secure and anonymous governance voting system leveraging Zero-Knowledge (ZK) proofs on the Cosmos blockchain.

## **Overview**

Cosmos ZK Voting is a governance voting system designed to ensure voter anonymity while maintaining the integrity and security of the voting process. Utilizing Zero-Knowledge proofs, this system allows voters to participate in governance without revealing their identities, fostering a more private and secure voting environment.

## **Features**

* **Anonymous Voting:** Protects voter identities using Zero-Knowledge proofs.  
* **Decentralized:** Built on the Cosmos blockchain for enhanced security and scalability.  
* **Secure Key Management:** Ensures the safety of prover and verifier keys.  
* **Relayer Support:** Facilitates seamless vote transactions through a dedicated relayer service.  
* **Flexible Proposal Management:** Create and manage governance proposals with ease.

## **Getting Started**

Follow the steps below to set up and run Cosmos ZK Voting locally.

### **Prerequisites**

* [Go](https://golang.org/doc/install) (version 1.18 or higher)  
* Cosmos SDK installed  
* [Git](https://git-scm.com/downloads)

### **Clone the Repository**
```bash
git clone https://github.com/your-username/cosmos-zk-voting
cd cosmos-zk-voting
git fetch
git checkout main
```

### **Install zkappd  Binary**

Build and install the `zkappd ` binary required to run the Cosmos ZK Voting chain.

```bash
make install
```

### **Run Local Testnet**

Initialize and start a local testnet.

**Initialize the Chain:**  
Modify the `scripts/init-simapp.sh` script according to your requirements.  
```bash
make init-simapp
```

1. **Start the Testnet:**  
```bash
zkappd  start
```

2. ## **ZK Prover and Verifier Keys**

Before setting up the chain, generate the ZK prover and verifier keys.

**Create Keys Directory:**  
```bash
mkdir keys
```

1. **Generate ZK Keys:**  
```bash
make generate-zk-keys
```

2. This command runs the `groth16.setup()` function defined in `/x/zkgov/client/zk/main.go`.  
   **Note:** Multiple prover and verifier key pairs are generated (e.g., `verifier-2`, `verifier-3`, `verifier-4`, etc.) to accommodate varying Merkle proof sizes. Use the appropriate key pair dynamically as needed.

### **Future Enhancements**

* **Optimize Merkle Implementation:** Aim for a constant Merkle tree size of 2^31 and improve update operations from O(n) to O(log n).

## **Relayer**

The Relayer is an HTTP server that listens for vote transactions, signs them with a valid chain address, and broadcasts them on-chain.

### **Why Use a Relayer?**

Vote transactions must be sent from a different, unlinkable address than the registered address. Since most users may not have multiple unlinkable addresses, the Relayer serves as an intermediary to handle this securely.

**Security Assurance:** The Relayer cannot manipulate transactions because the ZK proof guarantees transaction validity without exposing user secrets.

### **Running the Relayer**

**Navigate to the Relayer Directory:**  
```bash
cd /x/zkgov/client/relayer
```

1. **Start the Relayer:**  
```bash
zkappd  tx zk-gov run-relayer --from {key} --keyring-backend test --chain-id {chain-id} -y
```

Alternatively, you can run a predefined relayer (e.g., Alice):  
```bash
make run-alice-relayer
```

2. This command starts an HTTP server on port `8080` by default. You can change the port using the `--relayerPort {port}` flag.

## **Transactions**

### **Create a Proposal**

Create a new governance proposal that can be voted on with a YES or NO.

**Note:** This implementation is basic and intended for testing. A more robust system is recommended for production use.

**Proposal ID:** Starts at 1 and increments sequentially.

```bash
zkappd  tx zk-gov create-proposal [proposal-title] [proposal-description] --from [address] --keyring-backend test --chain-id [chain-id]
```

Or use a make command for convenience:

```bash
make create-proposal-a
```

### **Register a Vote**

Register a vote commitment, which will be used later to anonymize the actual vote.

```bash
zkappd  tx zk-gov register-vote [proposal-id] {"YES"/"NO"} --from [actual-voter-address] --keyring-backend test --chain-id [chain-id]
```

Alternatively, use predefined make commands:

```bash
mkdir commitments
make register-alice-vote
make register-bob-vote
make register-sai-vote
make register-teja-vote
```

### **Cast a Vote**

After registering a vote commitment, generate a ZK proof to cast the actual vote. This proof verifies that you know the valid commitment without revealing your identity.

**Important:** The vote transaction must be signed by a different, unlinkable address from the one used to register the vote commitment. Use the Relayer to facilitate this process.

#### **Without Relayer**

```bash
zkappd  tx zk-gov vote [proposal-id] [register-vote-address] --from [different-unlinkable-address] --keyring-backend test --chain-id [chain-id]
```

Or use make commands:

```bash
make broadcast-alice-vote
make broadcast-bob-vote
```

#### **With Relayer**

```bash
zkappd  tx zk-gov vote [proposal-id] [register-vote-address] --relayer [relayer-address]
```

Or use make commands:

```bash
make broadcast-sai-vote-via-relayer
make broadcast-teja-vote-via-relayer
```

### **Query a Proposal**

View the state of a proposal, including commitments and votes.

```bash
zkappd  q zk-gov get-proposal-info [proposal-id]
```

## **Contributing**

Contributions are welcome\! Please fork the repository and submit a pull request for any enhancements or bug fixes.

## **License**

This project is licensed under the MIT License.

## **Contact**

For questions or support, please open an issue on the [GitHub repository](https://github.com/vitwit/cosmos-zk-voting) or reach out to contact@vitwit.com.
