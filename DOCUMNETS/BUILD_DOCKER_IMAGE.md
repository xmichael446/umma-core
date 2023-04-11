### docker build -t umma:0.1.0 . # 0.1.0 last builded image version
### docker run --name umma umma:0.1.0
### sudo apt install mc # total comander for linux terminal

### docker load -i /DATA/images/umma-0.1.0.tar

### docker-compose up -d umma


```
umma:
    image: umma:0.1.0
    command: ./setup_and_run.sh umma1gpsh2h2z828gu8kgt5w28fel2krkkgjmrpe9en
    volumes:
    - /home/owner/.umma:/root/.umma
    ports:
    - 1317:1317 # rest
    - 26656:26656 # p2p
    - 26657:26657 # rpc
    environment:
    - GAS_LIMIT=${GAS_LIMIT:-100000000}
    - STAKE_TOKEN=${STAKE_TOKEN:-aumma}
```


ACCESS TO GENESIS

chown owner /home/owner/.umma/config/genesis.json