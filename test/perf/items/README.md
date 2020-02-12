# API performance test example

## Choose a strategy and populate `private.json` first.

## Run performance test:

```bash
go run ./test/perf/items/ -config ./config/private.json -strategy saml -spo
```

## Random SPO:

![response sample](./resp-sample-1.png)

## Random On-Prem:

![response sample](./resp-sample-2.png)