# openmetadata-connector
Fybrik OpenMetadata Connector.
Implements the Fybrik Data-Catalog specification.

## Compiling and Running
The first time you wish to compile the OpenMetadata Connector, run:
```bash
make
```
`make` automatically generates prerequisite code and compile the connector.

The next times you want to compile, run:
```bash
make compile
```

To run the connector locally, run:
```bash
make run
```

`make run` runs the connector with the configuration file in `conf/conf.yaml`. This configuration assumes that OM is running on localhost, listening on port 8585, and the Vault is running on localhost, listening on port 8200. If that is not the case, change the configuration file or employ port-forwarding.
The configuration file also contains the path of a JWT file which is used to identify against Vault.

## Directory Structure
- [database-types](database-types): Currently, the OM connector supports two types of databases: mysql and s3. The direcory contains `database_type.go` which defines the DatabaseType interface. It also contains `mysql.go` and `s3.go` which describe the `mysql` and `s3` structures (respectively), that implement the DatabaseType interface
- [utils](utils): Includes utility methods used in the connector code
- [vault](vault): Includes methods to obtain a token and secrets from Vault
- [conf](conf): Contains a sample configuration file
- [client-swagger](client-swagger): The OpenMetadata API Specification, used to OpenMetadata client code in the golang language. This directory contains the original YAML file and a modified YAML file to fix the auto-generated code
