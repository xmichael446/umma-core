#### CHECK NAME API

```javaScript
    URL: http://localhost:1317/umma/nameservice/nameservice/whois/SULTAN
```
### CHECK USERNAME API
```javaScript
    curl -s http://localhost:1317/umma/nameservice/nameservice/whois/SULTAN
```
### RESPONSE DATA FROM API

```javaScript
{
    "whois": {
        "index": "SULTAN",
        "name": "SULTAN",
        "value": "WORLD",
        "price": "1000000aumma",
        "owner": "umma1urf7xtpj3ey7t5pttahrlmxma0mnsgkur3ty7p"
    }
}

```

```json
    curl -X GET "http://135.125.3.192:1317/umma/nameservice/nameservice/whois/SULTAN" -H  "accept: application/json"
```

### BUY NAME
```json
    docker exec -it umma_umma_1 ummad tx nameservice buy-name SULTAN 1000000aumma --from founder --chain-id umma-1
```

### SET VALUE TO BUYED NAME

```json
    docker exec -it umma_umma_1 ummad tx nameservice set-name SULTAN WORLD --from founder --chain-id umma-1
```

### CHECK NAME OWNER
```json
    docker exec -it umma_umma_1 ummad q nameservice show-whois SULTAN
```