### Чит-лист Umma CLI

В этой контрольной таблице собраны часто используемые команды CLI,
которые операторы узлов могут легко копировать и вставлять. Несколько условностей, которым мы следуем:

+ Слова , написанные заглавными буквами , обозначают заполнители
+ Всегда используйте наши собственные конечные точки RPC Validator
+ Всегда указывайте *`--chain-id`* и *`--node`* флаги, даже если они не нужны
+ Команда командной строки запроса всегда использует *`--output json`* флаг и передает результат через *`jq`*

#### **Банк: Отправить**
```bash
ummad tx bank send KEY RECEIVER_ADDRESS 1000000aumma \
--chain-id umma-1 \
--node https://umma-rpc.validator.uz:443 --gas-prices 0.025aumma --gas-adjustment 1.8 --gas 250000 \
--from KEY
```
#### **Распределение: Вывод вознаграждений, включая комиссию**
```bash
ummad tx distribution withdraw-rewards VALIDATOR_OPERATOR \
--commission \
--chain-id umma-1 \
--node https://umma-rpc.validator.uz:443 --gas-prices 0.025aumma --gas-adjustment 1.8 --gas 250000 \
--from KEY
```
#### **Правительство: Запрос предложения**
```bash
ummad query gov proposal PROPOSAL_NUMBER \
--chain-id umma-1 \
--node https://umma-rpc.validator.uz:443 \
--output json | jq
```
#### **Правительство: Голосуйте**
##### VOTE_OTION: *`yes`*, *`no`*, *`no_with_veto`* и *`abstain`*.
```bash
ummad tx gov vote PROPOSAL_NUMBER VOTE_OPTION \
--chain-id umma-1 \
--node https://umma-rpc.validator.uz:443 --gas-prices 0.025aumma --gas-adjustment 1.8 --gas 250000 \
--from KEY
```
#### **Рубящий удар: Unjail**
```bash
ummad tx slashing unjail \
--chain-id umma-1 \
--node https://umma-rpc.validator.umma:443 --gas-prices 0.025aumma --gas-adjustment 1.8 --gas 250000 \
--from KEY
```
#### **Размещение ставок: Создание валидатора**
Примечание: В этой команде мы используем примерные значения вместо фиктивных слов с заглавной буквы для
демонстрационных целей. Пожалуйста, не забудьте соответствующим образом настроить его для вашего использования.

```bash
ummad tx staking create-validator \
--amount 1000000aumma \
--commission-max-change-rate "0.05" \
--commission-max-rate "0.10" \
--commission-rate "0.05" \
--min-self-delegation "1" \
--pubkey=$(ummad tendermint show-validator) \
--moniker 'validator.uz' \
--website "https://validator.uz" \
--identity "0A6AF02D1557E5B4" \
--details "Validator is the trusted staking service provider for blockchain projects. 100% refund for downtime slash. Contact us at hello@valitator.uz" \
--security-contact="hello@validator.uz" \
--chain-id umma-1 \
--node https://umma-rpc.validator.uz:443 --gas-prices 0.025aumma --gas-adjustment 1.8 --gas 250000 \
--from KEY
```

##### Если у вас есть отзывы или вы нашли ошибки в этом чите, пожалуйста, сообщите нам об этом на нашем [сервере Discord.](https://google.com) Спасибо!