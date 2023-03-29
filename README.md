# archive
**archive** is a blockchain built using Cosmos SDK and Tendermint. 

To learn more about how archive works, check out [the docs](https://arc-h1ve.gitbook.io/arc-h1ve-documentation/).

## Get started

### Run a local node

Build the local archive Docker container:
```bash
make localnet-init
```

Start the local archive Docker container:
```bash
make localnet-start
```

### Run the tests
```bash
make test
```

### Build reproducible binary
```bash
make build-reproducible
```
