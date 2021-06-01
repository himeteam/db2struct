# db2struct

mysql schema to golang struct tool
support Gorm, gorm, json tags

Install

```
    git clone https://github.com/himeteam/db2struct.git
    cd db2struct
    go mod download
    mkdir build
    go build -o build/db2struct ./cmd
```

Usage:

```
build/db2struct -H 127.0.0.1 -u root -p xxxxpassword -d databaseName -t tableName --json --gorm
```
Generate the structure of the following style:

```
package model

type OauthAccessToken struct {
        ID        string     `json:"id" gorm:"primaryKey;not null"`
        UserId    *int64     `json:"user_id"`
        ClientId  int        `json:"client_id" gorm:"not null"`
        Name      *string    `json:"name"`
        Scopes    *string    `json:"scopes"`
        Revoked   int        `json:"revoked" gorm:"not null"`
        CreatedAt *time.Time `json:"created_at" gorm:"autoCreateTime"`
        UpdatedAt *time.Time `json:"updated_at" gorm:"autoUpdateTime"`
        ExpiresAt *time.Time `json:"expires_at"`
}

```