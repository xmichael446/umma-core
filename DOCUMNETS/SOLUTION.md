#### BURN qilishni to'xtatib qo'ysak bo'ladi

```
cat genesis.json | jq .app_state.auth.accounts

cat /home/owner/.umma/config/genesis.json | jq .app_state.auth.accounts

```

shu query request yordamida biz burn qilish qo'shilganmi yoki yo'qmi bilib olamiz

#### Tx fee ni change qilsak bo'ladi buning uchun biz genesis.json ga o'zgartirish kiritamiz

```
cat genesis.json | jq .app_state.staking

cat /home/owner/.umma/config/genesis.json | jq .app_state.staking
```

shu query request yordamida qayerdan o'zgartirishimiz kerakligi to'g'risida ma'lumot olamiz


#### UPDATE COSMOS SDK

```
https://docs.cosmos.network/v0.45/migrations/chain-upgrade-guide-044.html
```

#### MINT CONTROLLER CUSTOM MODULE & CLAIMS MODULE

```
URL: https://github.com/ignite/modules
DISCORD: https://discord.com/channels/893126937067802685/894182472567382038/1030096984025071616
```

#### FeesToken
```
URL: https://github.com/osmosis-labs/osmosis/blob/main/x/txfees/keeper/feetokens.go
```

white list e'lon qilib shuni ichiga fee o'rnida to'lansa bo'ladigon coinlar or tokenlar ro'yxati qilish uchun osmosisda bu narsa qiligan ekan shuni o'zimizga qo'shsak bo'ladi

### OSMOSIZ EPOCH CONTROLLER

```
URL : https://github.com/osmosis-labs/osmosis/tree/main/x/epochs
```

### OSMOSIS MAKEFILE

```
URL : https://github.com/osmosis-labs/osmosis/blob/main/Makefile
```

### OSMOSIS MINT MODULE FOR EXAMPLE

```
URL: https://github.com/osmosis-labs/osmosis/tree/main/x/mint
```

### GAS & Fees

Validatorlar start qilish jarayonida qaysi coindan minimalniy commison olishlarini belgilab qo'ya olishadi va narx oshib borish xisobiga bu larni o'zgartira olishadi
```
ummad start ... --minimum-gas-prices=0.00001usdt;0.05atom
```


```javascript
The initial supply for your coins for your chain is simply the sum of all account balances in your genesis.
The bank module genesis state can also contains a supply value. If set, this enforces the initial supply of the chain and the genesis initialization fails if the sum of balances is not equal to this value.
The total supply is dynamic with token inflation. Inflation is configurable with the mint module.
-----
Первоначальный запас ваших монет для вашей цепочки — это просто сумма всех балансов счетов в вашем генезисе.
Состояние генезиса банковского модуля также может содержать значение поставки. Если установлено, это обеспечивает начальное снабжение цепочки, и инициализация генезиса завершается сбоем, если сумма балансов не равна этому значению.
Общее предложение динамично зависит от инфляции токенов. Инфляция настраивается с помощью модуля монетного двора.
```


### PRODUCTIONDA AIRDOR RUN QILISH UCHUN

```
URL: https://github.com/tendermint/faucet
```
Cosmos sdk da ko'tarilgan wasm yoki ethermint bo'ladimi xammasi uchun universal qilingan faucet bor ekan buni run qilishga urinib ko'rdim lekin xozircha to'liq run qila olmadim shuni fix qilish kerak bo'ladi shuni topsak faucet qilish muammosi xal bo'lgan bo'ladi

### ACCESS TO DIRECTORY
chmod -R 755 ~/home/owner/.umma/

which docker

### SNAPSHOT ENABLED FROM APP.TOML FILE

```
    sudo nano ~/.umma/config/app.toml
    snapshot-interval = 100 or 1000 # default snapshot interval = 0
    SAVE FILE CTRL+X AFTER ENTER

```