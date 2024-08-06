# Logger

I created this pkg for write log to clichouse via https://vector.dev/


### vector config

```yaml
sources:
  http_input:
    type: http_server
    address: 0.0.0.0:80
    encoding: binary
    headers:
      - User-Agent
    host_key: hostname
    method: POST
    path: /
    path_key: path
    query_parameters:
      - application
    response_code: 200
    strict_path: true
    auth:
      username: "vector"
      password: "vector_password"
  
sinks:
  clickhouse:
    type: http
    inputs:
      - http_input
    uri: http://localhost:8123/?query=INSERT+INTO+logs.logs2+FORMAT+JSONEachRow
    auth:
      strategy: basic
      user: default
      password: password
    encoding: 
      codec: "text"
```

clickshouse table
```sql
CREATE TABLE logs.logs2
(

    `time` DateTime64(3),

    `app` String,

    `build` String,

    `level` String,

    `msg` String,

    `source` String,

    `attrs` String
)
ENGINE = MergeTree
ORDER BY time
SETTINGS index_granularity = 8192;
```

### clickhouse docker compose
```yaml
version: '3.8'

services:
  clickhouse:
    image: clickhouse/clickhouse-server:24.6.2.17
    container_name: clickhouse-server
    ports:
      - "8123:8123"   # HTTP interface
      - "9000:9000"   # Native interface
      - "9009:9009"   # TCP interserver communications
    volumes:
      - ./docker-volume/lib/clickhouse:/var/lib/clickhouse
      - ./docker-volume/logs:/var/log/clickhouse-server/
      
    environment:
      - CLICKHOUSE_DB=default
      - CLICKHOUSE_USER=default
      - CLICKHOUSE_PASSWORD=password
      - CLICKHOUSE_DEFAULT_ACCESS_MANAGEMENT=1
```

## Example usage

```go
	v := &Vector{
		Client:   &http.Client{},
		Url:      "http://localhost",
		User:     "vector",
		Password: "vector_password",
	}

	logger := NewLogger(&Options{
		Level:     slog.LevelDebug,
		AddSource: true,
		Writer:    v,
		App:       "app",
		Build:     "v1.1",
	})
	logger.Info("test logger", slog.String("test_attr", "test_attr_value"))
```


#### systemctl config
```
[Unit]
Description=Vector Logger
After=network.target

[Service]
Type=simple
WorkingDirectory=/home/user/projects/logs/vector
ExecStart=vector --config /home/user/projects/logs/vector/http_log.yaml
Restart=always
RestartSec=3
StandardOutput=journal
StandardError=journal

[Install]
WantedBy=multi-user.target
```