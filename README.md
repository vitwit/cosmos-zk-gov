# **Cosmos ZK Voting**

A secure and anonymous governance voting system leveraging Zero-Knowledge (ZK) proofs on the Cosmos blockchain.

## **Overview**

Cosmos ZK Voting is a governance voting system designed to ensure voter anonymity while maintaining the integrity and security of the voting process. Utilizing Zero-Knowledge proofs, this system allows voters to participate in governance without revealing their identities, fostering a more private and secure voting environment.

## **Features**

* **Anonymous Voting:** Protects voter identities using Zero-Knowledge proofs.  
* **Decentralized:** Built on the Cosmos blockchain for enhanced security and scalability.  
* **Secure Key Management:** Ensures the safety of prover and verifier keys.  
* **dispatcher Support:** Facilitates seamless vote transactions through a dedicated dispatcher service.  
* **Flexible Proposal Management:** Create and manage governance proposals with ease.

## **Getting Started**

Follow the steps below to set up and run Cosmos ZK Voting locally.

### **Prerequisites**

* [Go](https://golang.org/doc/install) (version 1.18 or higher)  
* Cosmos SDK installed  
* [Git](https://git-scm.com/downloads)

### **Clone the Repository**

bash

Copy code

`git clone https://github.com/vitwit/cosmos-zk-gov`

`git fetch`

`git checkout main`

### **Install zkappd  Binary**

Build and install the `zkappd ` binary required to run the Cosmos ZK Voting chain.

bash

Copy code

`make install`

### **Run Local Testnet**

Initialize and start a local testnet.

1. **Initialize and Start the Chain:**  
Modify the `scripts/init-simapp.sh` script according to your requirements.  
`make init-simapp` will the run chain.
bash  
Copy code  
`make init-simapp`


## **ZK Prover and Verifier Keys**

Before setting up the chain, generate the ZK prover and verifier keys.

**Create Keys Directory:**  
bash  
Copy code  
`mkdir keys`

1. 

**Generate ZK Keys:**  
bash  
Copy code  
`make generate-zk-keys`

2. This command runs the `groth16.setup()` function defined in `/x/zkgov/client/zk/main.go`.  
   **Note:** Multiple prover and verifier key pairs are generated (e.g., `verifier-2`, `verifier-3`, `verifier-4`, etc.) to accommodate varying Merkle proof sizes. Use the appropriate key pair dynamically as needed.

### **Future Enhancements**

* **Optimize Merkle Implementation:** Aim for a constant Merkle tree size of 2^31 and improve update operations from O(n) to O(log n).

## **Dispatcher**

The Dispatcher is an HTTP server that listens for vote transactions, signs them with a valid chain address, and broadcasts them on-chain.

### **Why Use a Dispatcher?**

Vote transactions must be sent from a different, unlinkable address than the registered address. Since most users may not have multiple unlinkable addresses, the dispatcher serves as an intermediary to handle this securely.

**Security Assurance:** The Dispatcher cannot manipulate transactions because the ZK proof guarantees transaction validity without exposing user secrets.

### **Running the Dispatcher**

**Navigate to the Dispatcher Directory:**  
bash  
Copy code  
`cd /x/zkgov/client/dispatcher`

1. 

**Start the dispatcher:**  
bash  
Copy code  
`zkappd  tx zk-gov run-dispatcher --from {key} --keyring-backend test --chain-id {chain-id} -y`

Alternatively, you can run a predefined dispatcher (e.g., Vishal):  
bash  
Copy code  
`make run-vishal-dispatcher`

2. This command starts an HTTP server on port `8080` by default. You can change the port using the `--dispatcherPort {port}` flag.

## **Transactions**

### **Create a Proposal**

Create a new governance proposal that can be voted on with a YES or NO.

**Note:** This implementation is basic and intended for testing. A more robust system is recommended for production use.

**Proposal ID:** Starts at 1 and increments sequentially.

bash

Copy code

`zkappd  tx zk-gov create-proposal [proposal-title] [proposal-description] --from [address] --keyring-backend test --chain-id [chain-id]`

Or use a make command for convenience:

bash

Copy code

`make create-proposal-a`

**Create Commitments Directory:**  
bash  
Copy code  
`mkdir commitments`


### **Register a Vote**

Register a vote commitment, which will be used later to anonymize the actual vote.

bash

Copy code

`zkappd  tx zk-gov register-vote [proposal-id] {"YES"/"NO"} --from [actual-voter-address] --keyring-backend test --chain-id [chain-id]`

Alternatively, use predefined make commands:

bash

Copy code

`make register-vishal-vote`

`make register-kanna-vote`

`make register-sai-vote`

`make register-teja-vote`

### **Cast a Vote**

After registering a vote commitment, generate a ZK proof to cast the actual vote. This proof verifies that you know the valid commitment without revealing your identity.

**Important:** The vote transaction must be signed by a different, unlinkable address from the one used to register the vote commitment. Use the dispatcher to facilitate this process.

#### **Without dispatcher**

bash

Copy code

`zkappd  tx zk-gov vote [proposal-id] [register-vote-address] --from [different-unlinkable-address] --keyring-backend test --chain-id [chain-id]`

Or use make commands:

bash

Copy code

`make broadcast-vishal-vote`

`make broadcast-kanna-vote`

#### **With Dispatcher**

bash

Copy code

`zkappd  tx zk-gov vote [proposal-id] [register-vote-address] --dispatcher [dispatcher-address]`

Or use make commands:

bash

Copy code

`make broadcast-sai-vote-via-dispatcher`

`make broadcast-teja-vote-via-dispatcher`

### **Query a Proposal**

View the state of a proposal, including commitments and votes.

bash

Copy code

`zkappd  q zk-gov get-proposal-info [proposal-id]`

## **Contributing**

Contributions are welcome\! Please fork the repository and submit a pull request for any enhancements or bug fixes.

## **License**

This project is licensed under the MIT License.

## **Contact**

For questions or support, please open an issue on the [GitHub repository](https://github.com/vitwit/cosmos-zk-voting) or reach out to contact@vitwit.com.
