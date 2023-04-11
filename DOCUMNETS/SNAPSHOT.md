### CREATE SNAPSHOT

#### SET ACCESS TO CHANGE umma directory

```bash
    sudo -R 755 $HOME/.umma/
```
#### CREATE LOG FILE

```bash
    sudo nano $HOME/.umma/umma_log.txt
    SPACE CTRL+X Y ENTER
```
#### SET ACCESS TO READ&WRITE LOG FILE
```javascript
    chmod +x $HOME/.ummma/umma_log.txt
```
#### GENERATE SNAPSHOT tar archive
```bash
    bash scripts/umma_snapshot.sh
```
#### CHECK CREATED NEW SNAPSHOT

```bash
    ls -la $HOME/.umma/data/snapshots/
```

#### RESPONSE EXAMPLE
```bash
    umma-1_2022-12-12.tar
```