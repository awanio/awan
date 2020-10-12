# Awan
An open source Platform as Services base on Kubernetes

## Code Gudeline

Each of domain consist of:

* model This file will store our all model struct.
* repository The repository is responsible for database related job such as querying, inserting/storing or deleting. No business logic is implemented here.
* controller/handler This file accepts the request, call the repository and satisfy the business process and send the response.