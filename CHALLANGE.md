# Результат

Сервис идентификации с Rest API, который предоставляет доступ к профилю пользователю и его балансу
в [FA1.2](https://assets.tqtezos.com/docs/token-contracts/fa12/1-fa12-intro/)
и [FA2](https://assets.tqtezos.com/docs/token-contracts/fa2/2-fa2-nft-tutorial/) токенах.

Доступ предоставляется по security-токену пользователя.

## Core-функционал

- Библиотека или другой подход безопасно генерирует security-токен пользователя, который основан на его
  blockchain-identity. Должен использоваться Temple или другой популярный кошелек Tezos. Решение может включать
  функцонал подписей на блокчейне или вспомогательные смарт-контракты.
- Сервис синхронизации для добавления или удаления определенных Telegram или Discord-каналов на основе баланса
  пользователя.
- Будет хорошо, если узел Tezos и RPC-прокси будут позволять только пользователям с определенным балансам целевых
  токенов получить доступ.
- Все должно быть связано
  в [Helm Chart](https://helm.sh/ru/docs/intro/using_helm/#:~:text=Chart%20%E2%80%93%20%D1%8D%D1%82%D0%BE%20%D0%BF%D0%B0%D0%BA%D0%B5%D1%82%20Helm.,charts%2D%D1%8B%20%D0%B8%20%D0%B4%D0%B5%D0%BB%D0%B8%D1%82%D1%8C%D1%81%D1%8F%20%D0%B8%D0%BC%D0%B8.&text=Helm%20%D1%83%D1%81%D1%82%D0%B0%D0%BD%D0%B0%D0%B2%D0%BB%D0%B8%D0%B2%D0%B0%D0%B5%D1%82%20charts%20%D0%B2%20Kubernetes%2C%20%D1%81%D0%BE%D0%B7%D0%B4%D0%B0%D0%B2%D0%B0%D1%8F%20%D0%BD%D0%BE%D0%B2%D1%8B%D0%B9%20release%20%D0%B4%D0%BB%D1%8F%20%D0%BA%D0%B0%D0%B6%D0%B4%D0%BE%D0%B9%20%D1%83%D1%81%D1%82%D0%B0%D0%BD%D0%BE%D0%B2%D0%BA%D0%B8.)
  и настраиваться по таким параметрам:
    - целевой [FA1.2](https://assets.tqtezos.com/docs/token-contracts/fa12/1-fa12-intro/)
      / [FA2](https://assets.tqtezos.com/docs/token-contracts/fa2/2-fa2-nft-tutorial/) токен, по которому пользователь
      получает доступ (адрес контракта и/или id токена);
    - Tezos RPC и URL узла;
    - id Telegram и Discord-каналов;
    - ключ администратора для этих каналов;
    - минимальный баланс в целевых токенах для синхронизации с каналами.

## Подсказки и рекомендации

* Использовать [TypeScript](https://www.typescriptlang.org/)
* По возможности — [k8s](https://kubernetes.io/)
* [Dappetizer](https://docs.dappetizer.dev/)
* [Taquito](https://tezostaquito.io/) для доступа к блокчейну [Tezos](https://tezos.com/).

## Линки

* [Критерии оценивания челленджей и детальное описание](https://colossal-hexagon-b6c.notion.site/3e1f9bd04dac4d9bab833db0b0c0e8ce)
* [Temple - Tezos Wallet](https://templewallet.com/)
* [Tezos Dappetizer](https://docs.dappetizer.dev/)
* [kubernetes (k8s)](https://kubernetes.io/)
* [Helm Chart](https://helm.sh/ru/docs/intro/using_helm/#:~:text=Chart%20%E2%80%93%20%D1%8D%D1%82%D0%BE%20%D0%BF%D0%B0%D0%BA%D0%B5%D1%82%20Helm.,charts%2D%D1%8B%20%D0%B8%20%D0%B4%D0%B5%D0%BB%D0%B8%D1%82%D1%8C%D1%81%D1%8F%20%D0%B8%D0%BC%D0%B8.&text=Helm%20%D1%83%D1%81%D1%82%D0%B0%D0%BD%D0%B0%D0%B2%D0%BB%D0%B8%D0%B2%D0%B0%D0%B5%D1%82%20charts%20%D0%B2%20Kubernetes%2C%20%D1%81%D0%BE%D0%B7%D0%B4%D0%B0%D0%B2%D0%B0%D1%8F%20%D0%BD%D0%BE%D0%B2%D1%8B%D0%B9%20release%20%D0%B4%D0%BB%D1%8F%20%D0%BA%D0%B0%D0%B6%D0%B4%D0%BE%D0%B9%20%D1%83%D1%81%D1%82%D0%B0%D0%BD%D0%BE%D0%B2%D0%BA%D0%B8.)
