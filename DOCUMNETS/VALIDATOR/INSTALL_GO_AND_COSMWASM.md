## Установите Go и Cosmovisor

#### Не стесняйтесь пропустить этот шаг, если у вас уже есть Go и Cosmovisor.

##### **Установка Go**

Мы будем использовать Go v1.19.3в качестве примера здесь.
Приведенный ниже код также полностью удаляет все предыдущие установки Go.

```bash
sudo rm -rvf /usr/local/go/
wget https://golang.org/dl/go1.19.3.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.19.3.linux-amd64.tar.gz
rm go1.19.3.linux-amd64.tar.gz
```
#### **Настройка Go**
Если вы не хотите настраивать нестандартным способом, задайте их в ~/.profileфайле.

```bash
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export GO111MODULE=on
export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin
```
#### **Установка Cosmovisor**
Мы будем использовать Cosmovisor в v1.0.0качестве примера здесь.
```bash
    go install github.com/cosmos/cosmos-sdk/cosmovisor/cmd/cosmovisor@v1.0.0
```

### **Установка узла**
Установите текущую версию двоичного файла узла.
```bash
git clone https://github.com/umma-chain/umma-core umma-core
cd umma-core
git checkout v1.0.0
make install
```
#### **Настройка узла**
##### **Инициализировать узел**
Пожалуйста, замените YOUR_MONIKER на свой собственный псевдоним.

```bash
ummad init YOUR_MONIKER --chain-id umma-1
```
#### **Скачать Genesis**
Ссылка на файл genesis ниже - это зеркальная загрузка Polkachu.
Лучше всего найти официальную ссылку для скачивания genesis.

```bash
wget -O genesis.json https://github.com/umma-chain/mainnet/genesis.json --inet4-only
mv genesis.json ~/.umma/config
```

#### **Настройка Seed**

На наш взгляд, наилучшей практикой является использование начального узла для начальной загрузки.
В качестве альтернативы вы можете использовать addrbook или persistent_peers.

```bash
sed -i 's/seeds = ""/seeds = "ade4d8bc8cbe014af6ebdf3cb7b1e9ad36f412c0@seeds.ummma-core.com:12656"/' ~/.umma/config/config.toml
```
### **Запуск узла**
#### **Настройка папки Cosmovisor**
Создайте папки Cosmovisor и загрузите двоичный файл узла.

```bash
mkdir -p ~/.umma/cosmovisor/genesis/bin
mkdir -p ~/.umma/cosmovisor/upgrades
cp ~/go/bin/ummad ~/.umma/cosmovisor/genesis/bin
```

### **Создать служебный файл**
Создайте umma.service файл в /etc/systemd/system папке со следующим фрагментом кода.
Обязательно USERзамените его своим именем пользователя Linux.
Для выполнения этого шага вам понадобится sudo previlege.
```bash
sudo nano /etc/systemd/system/umma.service
```

```bash
[Unit]
Description="ummad node"
After=network-online.target

[Service]
User=YOUR_USER_NAME
ExecStart=/home/USER/go/bin/cosmovisor start
Restart=always
RestartSec=3
LimitNOFILE=4096
Environment="DAEMON_NAME=ummad"
Environment="DAEMON_HOME=/home/USER/.umma"
Environment="DAEMON_ALLOW_DOWNLOAD_BINARIES=false"
Environment="DAEMON_RESTART_AFTER_UPGRADE=true"
Environment="UNSAFE_SKIP_BACKUP=true"

[Install]
WantedBy=multi-user.target
```
#### **Загрузить Моментальный снимок**
Пожалуйста, используйте наш популярный сервис загрузки [моментальных](https://google.com) снимков,
чтобы загрузить и извлечь Umma snapshot.
#### **Запустить службу узлов**
```bash
# Enable service
sudo systemctl enable umma.service

# Start service
sudo service umma start

# Check logs
sudo journalctl -fu umma
```

#### **Другие соображения**
Это руководство по установке - это минимум для запуска узла. По мере того,
как вы станете более опытным оператором узла, вам следует учитывать следующее.

+ Используйте скрипт [Ansible](https://github.com/polkachu/cosmos-validators) для автоматизации процесса установки узла

+ Настройте брандмауэр так, чтобы он закрывал большинство портов, оставляя открытым только порт p2p (обычно 26656)
+ Используйте пользовательские порты для каждого узла, чтобы вы могли запускать несколько узлов на одном сервере

*Если вы обнаружите ошибку в этом руководстве по установке, пожалуйста, свяжитесь с нашим сервером [**Telegram**](https://google.com) и сообщите нам об этом.*