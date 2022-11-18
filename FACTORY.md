### CREATE MODULE 

```
ignite scaffold module factory --dep account,bank
```
### CREATE DENOM MAP
```
ignite scaffold map Denom description:string ticker:string precision:int siteUrl:string logoUrl:string maxSupply:int supply:int canChangeMaxSupply:bool --signer owner --index denom --module factory
```


## CREATE TOKEN

ummad tx factory create-denom usdt "Stable coin USD" USDT 6 "https://onemillion.uz" "https://onemillion.uz" 1000000000 true --from owner

## CHECK TOKEN LIST

ummad query factory list-denom

## MINT AND SEND TOKEN

ummad tx factory mint-and-send-tokens usdt 100 umma1gpsh2h2z828gu8kgt5w28fel2krkkgjmrpe9en --from owner

## BURN TOKEN
ummad tx factory burn-tokens usdt 10 --from owner


## ['RUN COMAND']
```sh
STAKE_TOKEN=aumma UNSAFE_CORS=true docker-compose up
```

### FAUCET ADD SEED PHRASE

want scorpion fly flame omit solid weird east surprise need invest asset series pepper blue around add slam country honey present accuse soldier laundry