### **STATE SYNC**

#### **Настройка нашего сервера синхронизации состояния**

Наши app.toml настройки, связанные с синхронизацией состояния, следующие.
Это информация только для вас. Вам не нужно следовать той же настройке на вашем узле.

```bash
# Prune Type
pruning = "custom"

# Prune Strategy
pruning-keep-every = 2000

# State-Sync Snapshot Strategy
snapshot-interval = 2000
snapshot-keep-recent = 5

```
#### **Наш RPC-сервер с синхронизацией состояния для Umma:**

```bash

https://umma-rpc.validator.uz:443 # -> http://135.125.3.192:26657
```
#### **Инструкция**

Мы предполагаем, что вы используете Cosmovisor для управления своим узлом.
Если вы не используете Cosmovisor, вам нужно будет немного настроить следующую инструкцию.

Создайте повторно используемый сценарий оболочки, например, *`state_sync.sh`* с
помощью следующего кода. Код получит важную информацию о синхронизации
состояния (например, высоту блока и хэш доверия) с нашего сервера
и соответствующим образом обновит ваш *`config.toml`* файл.

```bash
#!/bin/bash

SNAP_RPC="https://umma-rpc.validator.uz:443" # http://135.125.3.192:26657

LATEST_HEIGHT=$(curl -s $SNAP_RPC/block | jq -r .result.block.header.height); \
BLOCK_HEIGHT=$((LATEST_HEIGHT - 2000)); \
TRUST_HASH=$(curl -s "$SNAP_RPC/block?height=$BLOCK_HEIGHT" | jq -r .result.block_id.hash)

sed -i.bak -E "s|^(enable[[:space:]]+=[[:space:]]+).*$|\1true| ; \
s|^(rpc_servers[[:space:]]+=[[:space:]]+).*$|\1\"$SNAP_RPC,$SNAP_RPC\"| ; \
s|^(trust_height[[:space:]]+=[[:space:]]+).*$|\1$BLOCK_HEIGHT| ; \
s|^(trust_hash[[:space:]]+=[[:space:]]+).*$|\1\"$TRUST_HASH\"|" $HOME/.umma/config/config.toml
```

##### Предоставьте пользователю право на выполнение сценария, а затем запустите сценарий:

```bash
chmod 700 state_sync.sh 
./state_sync.sh
```

#### **Остановить узел:**

```bash
sudo service umma.service stop
```

#### **Сброс узла:**

```bash
# On some tendermint chains
ummad unsafe-reset-all

# On other tendermint chains
ummad tendermint unsafe-reset-all --home $HOME/.umma --keep-addr-book
```

#### **Перезапустите узел:**

```bash
sudo service umma.service start
```
Если все идет хорошо, ваш узел должен начать синхронизацию в течение 10 минут.

##### АЛЬТЕРНАТИВНЫЙ МАРШРУТ: у нас также есть служба [моментальных снимков Umma](https://google.com), которая поможет вам загрузить узел.
---
#### **Устранение неполадок**
##### **Я не могу подключиться к серверу синхронизации состояния**
Во-первых, посетите нашу целевую [страницу](https://google.com) сервера синхронизации состояния, чтобы убедиться, что он не отключен.

Во-вторых, добавьте сервер синхронизации состояния Validator в качестве своего партнера. Вы можете найти информацию о узле Validator [здесь](https://google.com).

##### **Я могу подключиться, но сразу получаю ошибки AppHash**
Убедитесь, что вы используете последнюю версию узла цепочки при синхронизации состояния.

#### **Другие вопросы?**

Временами синхронизация состояния может быть неустойчивой. Присоединяйтесь к нашему [серверу Discord](https://google.com), если у вас возникнут какие-либо проблемы. Счастливая синхронизация состояния!