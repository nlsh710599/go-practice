# A practice in Go

## Modules

1. Blockchain API server
2. Blockchain indexer

## Environment variables

| Key                    | Default                                                             | Description                                         |
| ---------------------- | ------------------------------------------------------------------- | --------------------------------------------------- |
| Port                   | 8080                                                                | port that API server binds                          |
| RpcUrl                 | https://nd-391-648-435.p2pify.com/899fe2afc3bc6426d419f89800c2d871  | rpc for the chain                                   |
| WsRpcUrl               | wss://ws-nd-391-648-435.p2pify.com/899fe2afc3bc6426d419f89800c2d871 | websocket for the chain                             |
| PostgresHost           | localhost                                                           | postgres host                                       |
| PostgresPort           | 5432                                                                | postgres port                                       |
| PostgresDatabase       | postgres                                                            | postgres database                                   |
| PostgresUser           | postgres                                                            | postgres user                                       |
| PostgresPassword       | docker                                                              | postgres password                                   |
| ConfirmationBlockCount | 20                                                                  | a block seems to be confirmed after how many blocks |

## Terms could be improved

1. handle websocket reconnection
2. query blocks/transactions/logs parallelly
3. host a full node to avoid rate limit issue
4. refine retry machenism when catch err
