# LedgerFusion: Privacy-Preserving Decentralized DAG Consensus

LedgerFusion is a Go-based project that aims to revolutionize decentralized consensus mechanisms by integrating Directed Acyclic Graph (DAG) structures with homomorphic encryption (HE). This fusion enables privacy-preserving verifiable computation on permissionless blockchain substrates, addressing a critical need for secure and scalable decentralized applications in various domains, including finance, supply chain management, and healthcare. The project explores novel approaches to achieve consensus while maintaining data confidentiality, allowing nodes to process and validate transactions without directly accessing the underlying sensitive information.

The core concept revolves around leveraging the inherent parallelism of DAGs for increased throughput and the power of HE to perform computations on encrypted data. This combination offers a unique advantage over traditional blockchain consensus algorithms that often struggle with scalability and privacy concerns. LedgerFusion achieves verifiable computation by utilizing HE schemes that allow nodes to verify the correctness of computations performed on encrypted data without revealing the input data itself. This is accomplished by incorporating zero-knowledge proofs within the consensus protocol, ensuring that computations are not only private but also demonstrably accurate.

LedgerFusion focuses on creating a modular and extensible architecture, allowing developers to integrate the consensus mechanism into existing blockchain platforms or build entirely new decentralized applications on top of it. The project strives to balance security, performance, and privacy, recognizing the inherent trade-offs between these factors in decentralized systems. By carefully selecting and optimizing HE schemes and DAG structures, LedgerFusion aims to provide a practical and efficient solution for privacy-sensitive decentralized applications. The ultimate goal is to unlock new possibilities for decentralized systems by enabling secure and scalable computation on sensitive data without compromising privacy.

## Key Features

*   **Privacy-Preserving Consensus:** Employs homomorphic encryption to enable nodes to participate in the consensus process without revealing the underlying transaction data. The project currently uses the BFV scheme for practical performance, although future iterations will incorporate more advanced schemes such as CKKS and TFHE as hardware acceleration improves.
*   **DAG-Based Architecture:** Utilizes a Directed Acyclic Graph (DAG) to achieve high throughput and scalability by allowing concurrent transaction processing. The DAG implementation uses a modified version of the Tangle architecture, optimized for privacy-preserving operations. Conflict resolution is handled via a novel voting mechanism built on encrypted data.
*   **Verifiable Computation:** Integrates zero-knowledge proofs to verify the correctness of computations performed on encrypted data, ensuring that the results are accurate and trustworthy. Specifically, the project uses Bulletproofs for range proofs within the smart contract execution environment.
*   **Modular Design:** Features a modular architecture that allows developers to easily integrate LedgerFusion's consensus mechanism into existing blockchain platforms or build new decentralized applications. The core modules include the DAG engine, the HE engine, the ZKP engine, and the consensus engine. These modules are designed to be loosely coupled and easily replaceable.
*   **Go Implementation:** Developed in Go for performance, concurrency, and cross-platform compatibility. Go's built-in concurrency features are heavily utilized to maximize the parallelism of the DAG structure and the homomorphic encryption operations.
*   **Byzantine Fault Tolerance (BFT):** Designed to be resilient to Byzantine faults, ensuring that the consensus mechanism can function correctly even in the presence of malicious nodes. The BFT protocol is implemented on top of the DAG structure, ensuring that even if some nodes are malicious, the overall system can still reach consensus.
*   **Smart Contract Support:** Enables the execution of smart contracts on encrypted data, allowing developers to build privacy-preserving decentralized applications with complex logic. The smart contract execution environment is based on a modified version of the Ethereum Virtual Machine (EVM), adapted to support homomorphic encryption operations and zero-knowledge proofs.

## Technology Stack

*   **Go:** The primary programming language, chosen for its performance, concurrency support, and cross-platform compatibility.
*   **BFV (BrakerskiFanVercauteren) Scheme:** Used for homomorphic encryption, allowing computations on encrypted data. The BFV implementation uses the `go-bfv` library, with custom optimizations for the specific requirements of the LedgerFusion consensus mechanism.
*   **Bulletproofs:** Employed for generating zero-knowledge range proofs, ensuring that computations are performed within specified bounds without revealing the actual values. The `go-bulletproofs` library is used for this purpose.
*   **gRPC:** Utilized for inter-process communication between nodes in the network. This allows for efficient and scalable communication between different components of the system.
*   **LevelDB:** Used for persistent storage of the DAG structure and other system data. This provides a fast and reliable way to store and retrieve data.

## Installation

1.  Ensure you have Go installed (version 1.18 or higher). You can download it from [https://go.dev/dl/](https://go.dev/dl/).

2.  Clone the LedgerFusion repository:
    git clone [https://github.com/jjfhwang/LedgerFusion](https://github.com/jjfhwang/LedgerFusion)

3.  Navigate to the project directory:
    cd LedgerFusion

4.  Download dependencies:
    go mod download

5.  Build the LedgerFusion binary:
    go build -o ledgerfusion

## Configuration

LedgerFusion can be configured using environment variables. The following variables are available:

*   `NODE_ID`: A unique identifier for the node (default: 0).
*   `PORT`: The port on which the node will listen for connections (default: 8080).
*   `BOOTSTRAP_NODE`: The address of a bootstrap node in the network (optional). This is used to discover other nodes in the network.
*   `DATA_DIR`: The directory where the node will store its data (default: ./data).
*   `HE_PARAM_FILE`: The path to the file containing the homomorphic encryption parameters (optional). If not provided, default parameters will be used, but these are not suitable for production environments.

Example usage:

NODE_ID=1 PORT=8081 ./ledgerfusion

## Usage

Running the LedgerFusion binary will start a node in the network. The node will attempt to connect to other nodes in the network and participate in the consensus process.

The LedgerFusion API is exposed via gRPC. The API includes the following endpoints:

*   `SubmitTransaction`: Submits a new transaction to the network.
*   `GetDAG`: Retrieves the current state of the DAG.
*   `GetNodeInfo`: Retrieves information about the node.

Detailed API documentation is available in the `api` directory of the repository. Example usage for submitting a transaction:

First, define the protobuf structures for the request and response.
Then, use the gRPC client to call the SubmitTransaction endpoint.
Here's a simplified illustration, assuming you have a gRPC client set up:



## Contributing

We welcome contributions to LedgerFusion! Please follow these guidelines:

1.  Fork the repository.
2.  Create a new branch for your feature or bug fix.
3.  Write clear and concise commit messages.
4.  Submit a pull request.

All contributions must adhere to the project's coding style and be thoroughly tested.

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/jjfhwang/LedgerFusion/blob/main/LICENSE) file for details.

## Acknowledgements

We would like to thank the open-source community for providing the libraries and tools that made this project possible. We also acknowledge the research community for their work on homomorphic encryption and zero-knowledge proofs.