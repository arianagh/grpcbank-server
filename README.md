
## Overview

`grpcbank` is a backend that simulates the core operations of a bank using gRPC and Protocol Buffers. The project follows the Ports and Adapters (Hexagonal) architecture to ensure a clear separation of concerns and high maintainability. The service offers a set of APIs for handling bank operations, including balance inquiry, exchange rate fetching, transaction summarization, and multi-transfer processing.

## Features

### API Endpoints

The `grpcbank` service provides the following gRPC API methods:

1. **GetCurrentBalance**:
    - **Description**: Retrieves the current balance of an account.
    - **Request**: `CurrentBalanceRequest`
    - **Response**: `CurrentBalanceResponse`

2. **FetchExchangeRates**:
    - **Description**: Streams the current exchange rates for different currencies.
    - **Request**: `ExchangeRateRequest`
    - **Response**: Stream of `ExchangeRateResponse`

3. **SummarizeTransactions**:
    - **Description**: Streams transactions and receives a summary of all transactions.
    - **Request**: Stream of `Transaction`
    - **Response**: `TransactionSummary`

4. **TransferMultiple**:
    - **Description**: Processes multiple transfers and streams the results.
    - **Request**: Stream of `TransferRequest`
    - **Response**: Stream of `TransferResponse`

## Architecture

The project is structured based on the Ports and Adapters architecture, which includes:

- **Ports**: Interfaces that define the boundaries between the core application logic and the outside world.
- **Adapters**: Implementations of these interfaces, handling the specifics of how data flows in and out of the application.

This architecture ensures that the business logic is decoupled from external concerns like gRPC, making the system easier to maintain and extend.

## Running the Application

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.16+)
- [Protocol Buffers](https://developers.google.com/protocol-buffers) (for generating gRPC code)
- [gRPC](https://grpc.io/docs/languages/go/quickstart/)

### Installation

1. **Clone the Repository**:
   ```bash
   git clone <repository-url>
   cd grpcbank
   ```

2. **Install Dependencies**:
   ```
   go mod tidy
   ``` 
   
3. **Generate gRPC Code: If you have made changes to the .proto files, regenerate the Go code**:
    ```
    make bank
    make bank-grpc
    ```

## Running the Application

- **To start the grpcbank service, run the following command**:
    ```
    go run ./cmd
    ```

## Testing the APIs

You can test the APIs using Insomnia or Postman by importing the gRPC requests.

### Using Insomnia

- Install Insomnia:
    Download and install Insomnia.


- Create a New gRPC Request:
Open Insomnia and create a new gRPC request.


- Set the Proto File:
  Add your proto file to the request by clicking on the "Proto File" field and selecting the appropriate .proto file from the proto/ directory.


- Choose the appropriate method (GetCurrentBalance, FetchExchangeRates, etc.).
  Set the URL to localhost:9000.


- Set the Request Body:
  Enter the required parameters in JSON format in the request body.


- Send the Request:
  Click on the "Send" button to test the gRPC method.