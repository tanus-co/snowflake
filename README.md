# snowflake
分布式唯一主键生成服务，snowflake算法实现

## server  配置
app.yml
```yml
snowflake:
  worker-id: 1
  data-center-id: 1
  server-port: 8080
```

## client 配置
app.yml
```yml
snowflake:
  server-port: 8080
  server-host: 127.0.0.1
```