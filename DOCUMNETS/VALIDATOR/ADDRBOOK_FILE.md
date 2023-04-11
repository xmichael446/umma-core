### Addrbook for Umma
#### Дополнительная книга для Umma
Иногда операторы узлов сталкиваются с проблемами пиринга с остальной частью сети.
Часто это можно решить с помощью хорошего начального узла или списка стабильных постоянных узлов.

Однако, когда все остальное терпит неудачу, вы можете доверить нашему
*`addrbook.json`* файлу начальную загрузку соединения вашего узла с сетью.

Остановите свой узел, загрузите и замените *`addrbook.json`* его на описанные ниже действия,
а затем перезапустите свой узел.

```bash
service umma stop

wget -O addrbook.json https://snapshots.validator.uz/addrbook/addrbook.json --inet4-only

mv addrbook.json ~/.umma/config

service umma start

service umma status

```
