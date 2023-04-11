## Инициализировать цепочку

Перед фактическим запуском узла нам необходимо инициализировать цепочку и, что наиболее важно, ее файл genesis. Это делается с initпомощью подкоманды:

```
# The argument <moniker> is the custom username of your node, it should be human-readable.
ummad init <moniker> --chain-id my-test-chain
```

Приведенная выше команда создает все файлы конфигурации, необходимые для запуска вашего узла, а также файл genesis по умолчанию, который определяет начальное состояние сети. Все эти файлы конфигурации находятся в ~/.ummad по умолчанию, но вы можете перезаписать расположение этой папки, передав --home флаг.

**~/.umma Папка имеет следующую структуру:**

```
.                                   # ~/.umma
|- data                           # Contains the databases used by the node.
|- config/
|- app.toml                   # Application-related configuration file.
|- config.toml                # Tendermint-related configuration file.
|- genesis.json               # The genesis file.
|- node_key.json              # Private key to use for node authentication in the p2p protocol.
|- priv_validator_key.json    # Private key to use as a validator in the consensus protocol.
```

### Обновление некоторых настроек по умолчанию

Если вы хотите изменить какие-либо значения полей в файлах конфигурации (например, genesis.json), вы можете использовать jq(installation & docs) & sedcommands для этого. Здесь перечислены несколько примеров.

```
# to change the chain-id
jq '.chain_id = "testing"' genesis.json > temp.json && mv temp.json genesis.json

# to enable the api server
sed -i '/\[api\]/,+3 s/enable = false/enable = true/' app.toml

# to change the voting_period
jq '.app_state.gov.voting_params.voting_period = "600s"' genesis.json > temp.json && mv temp.json genesis.json

# to change the inflation
jq '.app_state.mint.minter.inflation = "0.300000000000000000"' genesis.json > temp.json && mv temp.json genesis.json
```

## Добавление учетных записей Genesis

Перед запуском цепочки необходимо заполнить состояние хотя бы одной учетной записью. Для этого сначала создайте новую учетную запись в связке ключей с именем my_validatorпод бэкэндом testсвязки ключей (не стесняйтесь выбирать другое имя и другой бэкэнд).

Теперь, когда вы создали локальную учетную запись, предоставьте ей несколько stakeтокенов в файле genesis вашей цепочки. Это также позволит убедиться, что ваша сеть знает о существовании этой учетной записи:

```
ummad add-genesis-account $MY_VALIDATOR_ADDRESS 100000000000stake
```

Напомним, что $MY_VALIDATOR_ADDRESSэто переменная, которая содержит адрес my_validator ключа в связке ключей. Также обратите внимание, что токены в Cosmos SDK имеют {amount}{denom}формат: amountis - это десятичное число с точностью до 18 цифр, и denomэто уникальный идентификатор токена с его ключом обозначения (напримерatom, или uatom). Здесь мы предоставляем stakeтокены, как stakeи идентификатор токена, используемый для размещения ставокsimapp. Вместо этого следует использовать этот идентификатор токена для вашей собственной цепочки с собственным значением размещения.


Теперь, когда в вашей учетной записи есть несколько токенов, вам нужно добавить валидатор в свою цепочку. Валидаторы - это специальные полные узлы, которые участвуют в процессе согласования (реализованном в базовом механизме согласования), чтобы добавлять новые блоки в цепочку. Любая учетная запись может заявить о своем намерении стать оператором валидатора, но только те, у кого достаточно делегирования, могут войти в активный набор (например, только 125 лучших кандидатов в валидаторы с наибольшим делегированием становятся валидаторами в Cosmos Hub). В этом руководстве вы добавите свой локальный узел (созданный с помощью initприведенной выше команды) в качестве средства проверки вашей цепочки. Валидаторы могут быть объявлены до первого запуска цепочки с помощью специальной транзакции, включенной в файл genesis, который называется gentx:


```
# Create a gentx.
ummad gentx my_validator 100000000stake --chain-id my-test-chain --keyring-backend test

# Add the gentx to the genesis file.
ummad collect-gentxs
```

#### A gentx выполняет три действия:

> 01 Регистрирует validatorсозданную вами учетную запись как учетную запись оператора валидатора  (то есть учетную запись, которая управляет валидатором).


> 02 Самостоятельно делегирует предоставленные amountтокены для размещения.

> 03 Свяжите учетную запись оператора с публичным ключом узла Tendermint, который будет использоваться для подписи блоков. Если --pubkeyфлаг не указан, по умолчанию используется общедоступный ключ локального узла, созданный с помощью ummad initприведенной выше команды.

Для получения дополнительной информации gentxиспользуйте следующую команду:

```
ummad gentx --help
```

### Настройка узла с использованием app.toml и config.toml

Cosmos SDK автоматически генерирует два файла конфигурации внутри~/.ummad/config:

> config.toml: используется для настройки Tendermint, подробнее о документации Tendermint,

> **app.toml**: генерируется Cosmos SDK и используется для настройки вашего приложения, например, стратегий сокращения состояния, телеметрии, конфигурации gRPC и серверов REST, синхронизации состояния...


##### Оба файла сильно прокомментированы, пожалуйста, обратитесь к ним напрямую, чтобы настроить свой узел.

Одним из примеров конфигурации для настройки является minimum-gas-pricesполе внутриapp.toml, которое определяет минимальные цены на газ, которые узел проверки готов принять для обработки транзакции. В зависимости от цепочки, это может быть пустая строка или нет. Если оно пустое, обязательно отредактируйте поле с некоторым значением, например10token, иначе узел остановится при запуске. Для целей этого руководства давайте установим минимальную цену на газ равной 0:


```
# The minimum gas prices a validator is willing to accept for processing a
# transaction. A transaction's fees must meet the minimum of any denomination
# specified in this config (e.g. 0.25token1;0.0001token2).
minimum-gas-prices = "0stake"
```

### Запуск локальной сети

Теперь, когда все настроено, вы можете, наконец, запустить свой узел:
```
ummad start
```
Вы должны увидеть, как появляются блоки.

Предыдущая команда позволяет запускать один узел. Этого достаточно для следующего раздела, посвященного взаимодействию с этим узлом, но вы можете захотеть запустить несколько узлов одновременно и посмотреть, как происходит консенсус между ними.

Наивным способом было бы снова запустить те же команды в отдельных окнах терминала. Это возможно, однако в Cosmos SDK мы используем возможности Docker Compose для запуска локальной сети. Если вам нужна информация о том, как настроить собственную локальную сеть с помощью Docker Compose, вы можете ознакомиться с пакетом SDK Cosmos docker-compose.yml.

## Ведение журнала

Ведение журнала позволяет увидеть, что происходит с узлом. По умолчанию установлен уровень информации. Это глобальный уровень, и все информационные журналы будут выводиться на терминал. Если вы хотите фильтровать определенные журналы в терминале, а не все, то настройка module:log_levelзаключается в том, как это может работать.

Пример:
В config.toml:

```
log_level: "state:info,p2p:info,consensus:info,x/staking:info,x/ibc:info,*error"
```
--------------------------------

### Взаимодействие с узлом

>  **КРАТКОЕ ОПИСАНИЕ**
Существует несколько способов взаимодействия с узлом: использование CLI, использование gRPC или использование конечных точек REST.

> **ПРИМЕЧАНИЕ**
Необходимые показания

- Конечные точки gRPC, REST и Tendermint
- Запуск узла

### Использование командной строки

Теперь, когда ваша цепочка запущена, пришло время попробовать отправить токены с первой созданной вами учетной записи на вторую учетную запись. В новом окне терминала начните с выполнения следующей команды запроса:

```
ummad query bank balances $MY_VALIDATOR_ADDRESS --chain-id umma
```
Вы должны увидеть текущий баланс созданной вами учетной записи, равный исходному балансу stakeпредоставленной вами учетной записи за вычетом суммы, которую вы делегировали через gentx. Теперь создайте вторую учетную запись:

```
ummad keys add recipient --keyring-backend test

# Put the generated address in a variable for later use.
RECIPIENT=$(ummad keys show recipient -a --keyring-backend test)
```
Приведенная выше команда создает локальную пару ключей, которая еще не зарегистрирована в цепочке. Учетная запись создается при первом получении токенов от другой учетной записи. Теперь выполните следующую команду, чтобы отправить токены на recipientучетную запись:

```
ummad tx bank send $MY_VALIDATOR_ADDRESS $RECIPIENT 1000000stake --chain-id umma --keyring-backend test

# Check that the recipient account did receive the tokens.
ummad query bank balances $RECIPIENT --chain-id umma
```

Наконец, делегируйте некоторые токены ставки, отправленные на recipientучетную запись, валидатору:

```
ummad tx staking delegate $(ummad keys show my_validator --bech val -a --keyring-backend test) 500stake --from recipient --chain-id umma --keyring-backend test

# Query the total delegations to `validator`.
ummad query staking delegations-to $(ummad keys show my_validator --bech val -a --keyring-backend test) --chain-id umma
```

Вы должны увидеть два делегирования, первое из которых сделано из gentx, а второе, которое вы только что выполнили из recipientучетной записи

## Использование gRPC

Экосистема Protobuf разработала инструменты для различных вариантов использования, включая генерацию кода из *.protoфайлов на разных языках. Эти инструменты позволяют легко создавать клиентов. Часто клиентское соединение (то есть транспорт) можно подключить и заменить очень легко. Давайте рассмотрим один из самых популярных видов транспорта: gRPC.

Поскольку библиотека генерации кода во многом зависит от вашего собственного стека технологий, мы представим только три альтернативы:

- grpcurl для общей отладки и тестирования,
- программно через Go,
- CosmJS для разработчиков JavaScript / TypeScript.

## grpcurl

grpcurl похожcurl, но для gRPC. Он также доступен как библиотека Go, но мы будем использовать его только как команду CLI для целей отладки и тестирования. Следуйте инструкциям в предыдущей ссылке, чтобы установить его.

Предполагая, что у вас запущен локальный узел (либо локальная сеть, либо подключена живая сеть), вы должны иметь возможность выполнить следующую команду, чтобы перечислить доступные службы Protobuf (вы можете заменить localhost:9000конечной точкой сервера gRPC другого узла, которая настроена в grpc.addressполе внутриapp.toml):

```
grpcurl -plaintext localhost:9090 list
```
Вы должны увидеть список сервисов gRPC, например cosmos.bank.v1beta1.Query. Это называется отражением, которое представляет собой конечную точку Protobuf, возвращающую описание всех доступных конечных точек. Каждый из них представляет отдельную службу Protobuf, и каждая служба предоставляет несколько методов RPC, к которым вы можете запросить.

Чтобы получить описание сервиса, вы можете выполнить следующую команду:

```
grpcurl \
localhost:9090 \
describe cosmos.bank.v1beta1.Query                  # Service we want to inspect
```

Также возможно выполнить вызов RPC для запроса информации к узлу:

```
grpcurl \
-plaintext
-d '{"address":"$MY_VALIDATOR"}' \
localhost:9090 \
cosmos.bank.v1beta1.Query/AllBalances
```

Список всех доступных конечных точек запросов gRPC скоро появится.

Запрос исторического состояния с использованием grpcurl

Вы также можете запросить исторические данные, передав некоторые метаданные gRPC в запрос: x-cosmos-block-heightметаданные должны содержать блок для запроса. Используя grpcurl, как указано выше, команда выглядит следующим образом:

```
grpcurl \
-plaintext \
-H "x-cosmos-block-height: 279256" \
-d '{"address":"$MY_VALIDATOR"}' \
localhost:9090 \
cosmos.bank.v1beta1.Query/AllBalances
```
### CosmJS

Документацию CosmJS можно найти по адресу https://cosmos.github.io/cosmjs . По состоянию на январь 2021 года документация CosmJS все еще находится в стадии разработки.

#### Использование конечных точек REST

Как описано в руководстве по gRPC, все службы gRPC в Cosmos SDK становятся доступными для более удобных запросов на основе REST через gRPC-gateway. Формат URL-пути основан на полном имени сервисного метода Protobuf, но может содержать небольшие настройки, чтобы конечные URL-адреса выглядели более идиоматично. Например, конечной точкой REST для cosmos.bank.v1beta1.Query/AllBalancesметода является GET /cosmos/bank/v1beta1/balances/{address}. Аргументы запроса передаются в качестве параметров запроса.

В качестве конкретного примера curlкоманда для запроса баланса:

```
curl \
-X GET \
-H "Content-Type: application/json" \
http://localhost:1317/cosmos/bank/v1beta1/balances/$MY_VALIDATOR
```

Обязательно замените localhost:1317 на конечную точку REST вашего узла, настроенную в api.addressполе.


Список всех доступных конечных точек REST доступен в виде файла спецификации Swagger, его можно просмотреть по адресу localhost:1317/swagger. Убедитесь, что api.swaggerв вашем файле для поля установлено значение trueapp.toml

## Запрос исторического состояния с использованием REST

Запрос исторического состояния выполняется с использованием заголовка HTTP x-cosmos-block-height. Например, команда curl будет выглядеть следующим образом:

```
curl \
-X GET \
-H "Content-Type: application/json" \
-H "x-cosmos-block-height: 279256"
http://localhost:1317/cosmos/bank/v1beta1/balances/$MY_VALIDATOR
```

Предполагая, что состояние в этом блоке еще не было удалено узлом, этот запрос должен вернуть непустой ответ.

## Совместное использование ресурсов из разных источников (CORS)

Политики CORS по умолчанию не включены для обеспечения безопасности. Если вы хотите использовать rest-сервер в общедоступной среде, мы рекомендуем вам предоставить обратный прокси-сервер, это можно сделать с помощью nginx. Для целей тестирования и разработки внутри есть **enabled-unsafe-cors** поле **app.toml**.


```
ignite network chain publish github.com/fadeev/example
```

genrated chain id  ?

```
ignite network chain init 37
```

QUESTION 3-4

```
ignite network chain join 37 --amount 95000000stake
```

```
ignite n request list 37
```

SHOW GENESIS ACCOUNTS


```
ignite n request qpprove $CHAIN_ID 1,2
```
APPROVE VALIDATORS TO CHAIN


```
ignite nework chain launch $CHAIN_ID
```

```
```
ignite network chain prepare $CHAIN_ID
```
AFTER RUN
example ummad --home ./home/genesis.json

VALIDATORLIK ga kirmoqchoqchi bo'lgan user

```
ignite network chain join $CHAIN_ID --amount 95000000stake
```


### [GENESIS URL](https://raw.githubusercontent.com/umma-chain/mainnet/main/genesis.json)
### [STATUS NODE](http://135.125.3.192:26657/status?)
### [REST API](http://135.125.3.192:1317)
#### GET NODE ID -> docker exec -it umma_umma_1 ummad tendermint show-node-id
#### [JITIPOLA](https://github.com/NibiruChain/installer)
### ONE LINE CMD
```bash
 wget -q -O ummad.sh https://gist.githubusercontent.com/FounderDAO/56f4285dded55dc55a976d0f812a7f28/raw/ff853493aaa7a2814c5036f78d14a1af2b2cdbbd/ummad.sh && chmod +x ummad.sh && sudo /bin/bash ummad.sh
 wget -q -O ummad.sh https://bit.ly/3BZv1gc && chmod +x ummad.sh && sudo /bin/bash ummad.sh
```
