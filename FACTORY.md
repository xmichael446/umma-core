### CREATE MODULE 

```
ignite scaffold module factory --dep account,bank
```
### CREATE DENOM MAP
```
ignite scaffold map Denom description:string ticker:string precision:int siteUrl:string logoUrl:string maxSupply:int supply:int canChangeMaxSupply:bool --signer owner --index denom --module factory
```