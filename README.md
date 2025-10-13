# Photon Core Starter

Go 應用程式啟動框架，提供配置管理、依賴注入、生命週期管理等核心功能。



## 安裝

```bash
go get github.com/Phofuture/photon-core-starter
```

## 快速開始

```go
package main

import (
    photon "github.com/Phofuture/photon-core-starter"
)

func main() {
    photon.Run()
}
```

## Configuration


#### 檔案命名規則

1. **基礎配置檔案**：`app.yml` 或 `app.yaml`
2. **環境特定配置**：`{env}.yml` 或 `{env}.yaml`（例如：`dev.yml`、`prod.yml`）

#### 配置檔案搜尋路徑

系統會依序在以下目錄中搜尋配置檔案：
1. `./src/resources/`
2. `./src/`
3. `./`（專案根目錄）

### 配置載入順序

1. 載入 `app.yml` 作為基礎配置
2. 根據 `env.name` 的值載入對應的環境配置（如 `dev.yml`）
3. 環境變數覆蓋配置項（使用 `_` 替代 `.`）

### 配置檔案範例

#### app.yml（基礎配置）
```yaml
env:
  name: dev  # 環境名稱，決定要載入哪個環境配置檔案

app:
  name: my-application
  port: 8080

database:
  host: localhost
  port: 5432
```

#### dev.yml（開發環境配置）
```yaml
database:
  host: localhost
  port: 5432
  name: myapp_dev
  username: dev_user
  password: dev_password

log:
  level: debug
```

#### prod.yml（生產環境配置）
```yaml
database:
  host: prod-db.example.com
  port: 5432
  name: myapp_prod
  # 生產環境的敏感資訊建議使用環境變數

log:
  level: info
```

### 使用配置

#### 方法 1：註冊配置結構

```go
package main

import (
    "github.com/Phofuture/photon-core-starter/configuration"
)

type AppConfig struct {
    App struct {
        Name string `mapstructure:"name"`
        Port int    `mapstructure:"port"`
    } `mapstructure:"app"`

    Database struct {
        Host     string `mapstructure:"host"`
        Port     int    `mapstructure:"port"`
        Name     string `mapstructure:"name"`
        Username string `mapstructure:"username"`
        Password string `mapstructure:"password"`
    } `mapstructure:"database"`
}

var Config = &AppConfig{}

func init() {
    // 註冊配置結構，會在 InitConfiguration() 時自動解析
    configuration.Register(Config)
}
```

#### 方法 2：動態獲取配置

```go
package main

import (
    "context"
    "github.com/Phofuture/photon-core-starter/configuration"
)

func GetConfig(ctx context.Context) (*AppConfig, error) {
    return configuration.Get[AppConfig](ctx)
}
```

### 環境變數覆蓋

配置項可以透過環境變數覆蓋，使用 `_` 替代 `.`：

```bash
# 覆蓋 database.host
export DATABASE_HOST=custom-host

# 覆蓋 app.port
export APP_PORT=9090

# 覆蓋 env.name 以切換環境
export ENV_NAME=prod
```
