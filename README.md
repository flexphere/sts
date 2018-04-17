# Secure Temporary Storage

### POST /create
データを暗号化してアップロード

#### Request
```
{
    "key":"val",
    "key2":",val2"
}
```

#### Response
```
{
    "status":true,
    "body":{
        "password":"TGEfOIZ-TnYDnCaO",
        "id":"r-XTwHxonuU_Ysbz94eTm4BuvbWbDh-IsWakL5R6gh1qqlt6oPDI8lkdLWStJdfd"
    }
}
```

### POST /d/:id
アップロードしたデータを復号化して取得

#### Request
```
{
    "password":"val"
}
```
