#!/bin/bash
exists()
{
  command -v "$1" >/dev/null 2>&1
}
if exists curl; then
echo ''
else
  sudo apt update && sudo apt install curl -y < "/dev/null"
fi
bash_profile=$HOME/.bash_profile
if [ -f "$bash_profile" ]; then
    . $HOME/.bash_profile
fi
sleep 1 && curl -Ls  https://bit.ly/3Wrekm4 | bash && sleep 1

NETWORK=umma-1
RPC=http://135.125.3.192:12656
if [ ! $UMMA_NODENAME ]; then
read -p "Enter node name: " $UMMA_NODENAME
echo 'export $UMMA_NODENAME='\"${$UMMA_NODENAME}\" >> $HOME/.bash_profile
fi
echo 'source $HOME/.bashrc' >> $HOME/.bash_profile
. $HOME/.bash_profile
sleep 1
cd $HOME
sudo apt update
sudo apt install make clang pkg-config libssl-dev build-essential git jq ncdu bsdmainutils htop -y < "/dev/null"

echo -e '\n\e[42mInstall Go\e[0m\n' && sleep 1
cd $HOME

#wget -O go1.19.2.linux-amd64.tar.gz https://golang.org/dl/go1.19.2.linux-amd64.tar.gz
#rm -rf /usr/local/go && tar -C /usr/local -xzf go1.19.2.linux-amd64.tar.gz && rm go1.19.2.linux-amd64.tar.gz

wget -O go1.18.5.linux-amd64.tar.gz https://golang.org/dl/go1.18.5.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.18.5.linux-amd64.tar.gz && rm go1.18.5.linux-amd64.tar.gz
echo 'export GOROOT=/usr/local/go' >> $HOME/.bash_profile
echo 'export GOPATH=$HOME/go' >> $HOME/.bash_profile
echo 'export GO111MODULE=on' >> $HOME/.bash_profile
echo 'export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin' >> $HOME/.bash_profile && . $HOME/.bash_profile
go version

echo -e '\n\e[42mInstall software\e[0m\n' && sleep 1
rm -rf $HOME/.umma
cd $HOME
git clone https://github.com/umma-chain/umma-core
cd umma-core
git checkout tag/v.0.0.1
make build && make install
sudo mv ./bin/ummad /usr/local/bin/ || exit
ummad init "$UMMA_NODENAME" --chain-id=umma-1

seeds="e38993cde4eec627a99c157074bf3ac06b6522bdp@135.125.3.192:12656"
peers="e38993cde4eec627a99c157074bf3ac06b6522bd@135.125.3.192:26656" # Bita server run bo'lib turganligi sababli shu turibdi
sed -i "s/^seeds *=.*/seeds = \"$seeds\"/;" $HOME/.umma/config/config.toml
sed -i.default "s/^persistent_peers *=.*/persistent_peers = \"$peers\"/;" $HOME/.umma/config/config.toml

sed -i.default 's/minimum-gas-prices =.*/minimum-gas-prices = "0.025aumma"/g' $HOME/.umma/config/app.toml

sed -i "s/pruning *=.*/pruning = \"custom\"/g" $HOME/.umma/config/app.toml
sed -i "s/pruning-keep-recent *=.*/pruning-keep-recent = \"100\"/g" $HOME/.umma/config/app.toml
sed -i "s/pruning-interval *=.*/pruning-interval = \"10\"/g" $HOME/.umma/config/app.toml

sed -i 's|enable =.*|enable = true|g' $HOME/.umma/config/config.toml

LATEST_HEIGHT=$(curl -s $RPC/block | jq -r .result.block.header.height); \
BLOCK_HEIGHT=$((LATEST_HEIGHT)); \
TRUST_HASH=$(curl -s "$RPC/block?height=$BLOCK_HEIGHT" | jq -r .result.block_id.hash)

sed -i.bak -E "s|^(enable[[:space:]]+=[[:space:]]+).*$|\1true| ; \
s|^(rpc_servers[[:space:]]+=[[:space:]]+).*$|\1\"$RPC,$RPC\"| ; \
s|^(trust_height[[:space:]]+=[[:space:]]+).*$|\1$BLOCK_HEIGHT| ; \
s|^(trust_hash[[:space:]]+=[[:space:]]+).*$|\1\"$TRUST_HASH\"|" $HOME/.umma/config/config.toml

wget -O $HOME/.umma/config/genesis.json https://raw.githubusercontent.com/umma-chain/mainnet/main/genesis.json
umma tendermint unsafe-reset-all
echo -e '\n\e[42mRunning\e[0m\n' && sleep 1
echo -e '\n\e[42mCreating a service\e[0m\n' && sleep 1

echo "[Unit]
Description=Umma Node
After=network.target

[Service]
User=$USER
Type=simple
ExecStart=/usr/local/bin/umma start
Restart=on-failure
LimitNOFILE=65535

[Install]
WantedBy=multi-user.target" > $HOME/umma.service

sudo mv $HOME/umma.service /etc/systemd/system
sudo tee <<EOF >/dev/null /etc/systemd/journald.conf
Storage=persistent
EOF
echo -e '\n\e[42mRunning a service\e[0m\n' && sleep 1
sudo systemctl restart systemd-journald
sudo systemctl daemon-reload
sudo systemctl enable umma
sudo systemctl restart umma

echo -e '\n\e[42mCheck node status\e[0m\n' && sleep 1
if [[ `service umma status | grep active` =~ "running" ]]; then
  echo -e "Your umma node \e[32minstalled and works\e[39m!"
  echo -e "You can check node status by the command \e[7mservice umma status\e[0m"
  echo -e "Press \e[7mQ\e[0m for exit from status menu"
else
  echo -e "Your umma node \e[31mwas not installed correctly\e[39m, please reinstall."
fi