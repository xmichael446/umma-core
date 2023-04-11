### **Umma Node Snapshot**

#### **Настройка нашего сервера моментальных снимков Umma**

##### **Мы делаем снимок узла каждый день. Затем мы удаляем все предыдущие снимки, чтобы освободить место на файловом сервере.**

Снимок предназначен для операторов узлов для запуска эффективной службы проверки подлинности в цепочке Umma.
Чтобы сделать снимок как можно меньше, сохраняя при этом жизнеспособность в качестве средства проверки,
мы используем следующую настройку для экономии места на диске. Мы предлагаем вам выполнить
ту же настройку и на вашем узле. Пожалуйста, обратите внимание, что ваш узел будет иметь
очень ограниченную функциональность, помимо подписывания блоков с эффективным использованием
дискового пространства. Например, ваш узел не сможет служить конечной точкой RPC
(которую в любом случае не рекомендуется запускать на узле проверки подлинности).

Поскольку мы периодически синхронизируем состояние наших узлов моментальных снимков,
вы можете заметить, что иногда размер нашего снимка на удивление мал.

#### **app.toml**
```bash
# Prune Type
pruning = "custom"

# Prune Strategy
pruning-keep-recent = "100"
pruning-keep-every = "0"
pruning-interval = "10"
```

#### **config.toml**
```bash
indexer = "null"
```
#### **Как обработать снимок Umma**
##### *При необходимости установите lz4*
```bash
sudo apt update
sudo apt install snapd -y
sudo snap install lz4
```
##### Загрузите снимок
```bash
wget -O umma_6228819.tar.lz4 https://snapshots.validator.uz/snapshots/umma_6228819.tar.lz4 --inet4-only
```
##### Остановите свой узел
```bash
sudo service umma stop
```
Сбросьте свой узел. **ПРЕДУПРЕЖДЕНИЕ**: это приведет к удалению базы данных вашего узла.
Если вы уже используете validator,
убедитесь, что вы создали резервную копию **`priv_validator_key.json`** перед запуском команды.
Команда не удаляет файл. Однако у вас уже должна быть его резервная копия в безопасном месте.

```bash
# On some tendermint chains
ummad unsafe-reset-all

# On other tendermint chains
ummad tendermint unsafe-reset-all --home $HOME/.umma --keep-addr-book
```
##### Распакуйте снимок в папку вашей базы данных. Ваше местоположение в базе данных будет зависеть от реализации вашего узла.

```bash
lz4 -c -d umma_6228819.tar.lz4  | tar -x -C $HOME/.umma
```
##### Если все хорошо, теперь перезапустите свой узел
```bash
sudo service umma start
```
##### Удалите загруженный снимок, чтобы освободить место
```bash
rm -v umma_6228819.tar.lz4
```
##### Убедитесь, что ваш узел запущен
```bash
sudo service umma status
sudo journalctl -u umma -f
```
**РАСШИРЕННЫЙ МАРШРУТ**: вышеупомянутое решение требует, чтобы вы загрузили сжатый файл,
распаковали его, а затем удалили исходный файл. Для этого требуется дополнительное
место для хранения на вашем сервере. Вы можете выполнить следующую комбинированную команду,
чтобы передать снимок в папку вашей базы данных. Только для опытных пользователей:

```bash
curl -o - -L https://snapshots.validator.uz/snapshots/umma_6228819.tar.lz4 | lz4 -c -d - | tar -x -C $HOME/.umma
```
##### **АЛЬТЕРНАТИВНЫЙ МАРШРУТ:** У нас также есть служба [синхронизации состояния Umma](https://validator.uz/state_sync), которая поможет вам загрузить узел.