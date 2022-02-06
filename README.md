# Flame DeFi telegram and discord channel service

Identity Awareness service with HTTP Rest API, which provides access to the user's profile and balance
in [FA1.2](https://assets.tqtezos.com/docs/token-contracts/fa12/1-fa12-intro/)
and [FA2](https://assets.tqtezos.com/docs/token-contracts/fa2/2-fa2-nft-tutorial/) tokens.

> Access is provided by the user's security token.

## Helpful commands

Test tezos smart-contract:

```shell
bin/ligo test --format='human-readable' smart1.ligo
```

## Links

* [Flame Tezos DeFi Hackathon 2022 Challenge (rus)](CHALLANGE.md)
* [Guidelines for participants on Tezos DeFi Hackathon 2022 (rus)](GUIDLINES.md)
* [How to run Tezos](https://tezos.gitlab.io/introduction/howtoget.html)

### Wallet

* [Create Tezos Wallet](https://tezos.com/create-wallet/) - How to create wallet guide.
* [Temple](https://templewallet.com/) - Tezos Wallet.

### Smart contracts

* [What Are Smart Contracts and How They Work](https://forklog.com/sp/dev-on-tezos/en/tezos-introduction)
* [LIGO Smart-Contract IDE](https://ide.ligolang.org/)
* [SmartPy IDE](https://smartpy.io/ide)

### Framework and libraries

* [PyTezos](https://github.com/murbard/pytezos/) - Python utils for Tezos.
* [Tezos Dappetizer](https://docs.dappetizer.dev/) - Framework for building  indexer apps using TypeScript (or JavaScript). Its main strength is its versatility - it can be used to rapidly develop a simple smart contract indexer or collect particular block data, but also to index data from the entire blockchain or compose multiple indexers.
* [Tezos RPC API  indexers](https://tezosguides.com/infrastructure/indexer/) - The Tezos RPC API exposes all the information you could possibly need to build applications using blockchain data, however, the RPC API alone might not be efficient for your particular use-case, this is where indexers come into play, here we are going to look at two different indexers for tezos.

### CI/CD

* [Kubernetes](https://kubernetes.io/) - Kubernetes, also known as K8s, is an open-source system for automating deployment, scaling, and management of containerized applications.
* [Helm Chart](https://helm.sh/ru/docs/intro/using_helm/#:~:text=Chart%20%E2%80%93%20%D1%8D%D1%82%D0%BE%20%D0%BF%D0%B0%D0%BA%D0%B5%D1%82%20Helm.,charts%2D%D1%8B%20%D0%B8%20%D0%B4%D0%B5%D0%BB%D0%B8%D1%82%D1%8C%D1%81%D1%8F%20%D0%B8%D0%BC%D0%B8.&text=Helm%20%D1%83%D1%81%D1%82%D0%B0%D0%BD%D0%B0%D0%B2%D0%BB%D0%B8%D0%B2%D0%B0%D0%B5%D1%82%20charts%20%D0%B2%20Kubernetes%2C%20%D1%81%D0%BE%D0%B7%D0%B4%D0%B0%D0%B2%D0%B0%D1%8F%20%D0%BD%D0%BE%D0%B2%D1%8B%D0%B9%20release%20%D0%B4%D0%BB%D1%8F%20%D0%BA%D0%B0%D0%B6%D0%B4%D0%BE%D0%B9%20%D1%83%D1%81%D1%82%D0%B0%D0%BD%D0%BE%D0%B2%D0%BA%D0%B8.) - Package manager for Kubernetes
