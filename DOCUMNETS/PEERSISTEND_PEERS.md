### GET NODE PEERS ID

```javascript
    docker exec -it umma_umma_1 ummad  tendermint show-node-id [OR] ummad tendermint show-node-id
    Response example : bf2058aa0502ae39759095d4dd7de6e9ab068e33
```

```json
    ummad init <'moniker'> --chain-id umma-1
```
 qilingan vaqtda
    node_key.json
    priv_validator_key.json
    genesis.json
    client.toml
    config.toml
    addrbook.json
    app.toml
    getnx/gentx-bf2058aa0502ae39759095d4dd7de6e9ab068e33.json
yaraladi va shu node uchun unique id generate bo'ladi biz shu UNIQUE ID ni olib boshqa node
larda persistent_peersga shuni bf2058aa0502ae39759095d4dd7de6e9ab068e33@NODE_IP:26656 config.toml ga
yozib qo'yishimiz kerak bo'ladi ketma-ket qilib orasiga , qo'yib yozib qo'yaveriladi


