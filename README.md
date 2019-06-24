# Chat app by Golang

## Targets

### Performance
- Using preview image to replace the original image, it's useful expecially for group chat
- Increasing the resourse services, e.g. using cloud service (qos/alioss)
- Compressed message body, using url for the resource (content => url)

### Concurancy
- Increasing single machine performance
- Distibuted deployment
- Dynamic expand capacity

### Requirements and strategy
- Message data structure

```go

type Message struct {
  Id      int64  `json:"id,omitempty" form :"id"`         // message id
  Userid  int64  `json:"userid,omitempty" form:"userid"`   // sender id
  Cmd     int    `json:"cmd,omitempty" form:"cmd"`         // group or private chat
  Dstid   int64  `json:"dstid,omitempty" form:"dstid"`     // receiver(person or group) id
  Media   int64  `json:"media,omitempty" form:"media"`     // meida type
  Content string `json:"content,omitempty" form:"content"` // content
  Pic     string `json:"pic,omitempty" form:"pic"`         // preview pic
  Url     string `json:"url,omitempty" form:"url"`         // service url
  Memo    string `json:"memo,omitempty" form:"memo"`       // simple description
  Amount int    `json:"amount,omitempty" form:"amount"`    // amount, meta info of the content
}
```


## Common IM System Structure

### Front-end
- iOS, Android, Webapp
- SDK, API, websocket

### Interface Layer
- TCP, HTTPS, HTTP2, websocket

### Logic layer
- auth, login, group chat, signle chat, notification...

### Storage layer
- Mysql, Redis, Mongodb, Hbase, Hive, file server ( cloud based )



## Optimize single machine performance
- Optimize Map
  - Using read and write map
    - In the app, read will be a lot of greater the the write
    - map should not be to large, not bigger than 100,000 users
  - Linux
    - Adjust maximum files "最大文件数"
  - Opimize CPU
    - Decrease JSON encode/decode, it's comsuming performance
  - I/O
    - Combine multiple DB operation
    - Optimize read operation
    - Using cache as much as possible
  - Application and resource servers separate
    - Using cloud service for resources


# Websocket

## Steps of sending message
- User A opens websocket, send /chat?id=xxx&token=yyy
- Server auth and create the map of userid => websocket(channel)
- Using coroutine, conn.ReadMessage to wait and read message
- A sends json string with dstid
- If the id is a group, send to group
- ClientMap[userId] gets conn
- conn.WriteMessage

## Lib
- github.com/gorilla/websocket
- golang.org/x/net/websocket

## auth
- ws://192.168.0.100/chat?id=uid&token=token
- auth by checking id and token
- define a ClientNode to store the connection information

## websocket heart beat
- every 30s send
- next 30s after the latest message
- heart beat will affect the server's performance

## Front-end sendding message
- using queue
- 
```js
var dataqueue = []
function push (m) {
  if (!dataqueue) { dataqueue = [] }
  dataqueue.push(m)
}

function pop () {
  if (dataqueue) return dataqueue.shift()
  else return null
}
```

# DB

## Database CRUD Steps
```golang
// 1. init db
xorm.NewSession(drivername, datasourcename)

// 2. model layer entity - define model or entity
type User struct {
    Id      int64    `xorm: "pk autoincr bigint(20)"`
    Mobile  string   `xorm: "varchar(20)"`
    Name    string   `...`
}

// 3. service layer
- find one
DbEngine.ID(userId).Get(&User)

- search, return multiple records
result := make([]User, 0)
DbEngine.where("a=? and b=?", a, b).Find(&result)

- create single record
DbEngine.InsertOne(&User)

- update
DbEngine.ID(userId).Update(&User)
DbEngine.Where("a=? and b=?", a, b).Update(&User)
DbEngine.ID(userId).Cols("field1, field2, field3").Update(&User)

- delete
DbEngine.ID(userId).Delete(&User)

- md5
import (
    "crypto/md5"
    "encoding/hex"
    "strings"
)

func Md5Encode(data string) string {
    h := md5.New()
    h.Write([]byte(data))
    cipherStr := h.Sum(hil)
    
    return hex.EncodeToString(cipherStr)
}

func MD5Encode(data string) string {
    return strings.ToUpper(Md5Encode(data))
} 

```


















