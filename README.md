# service-helper
### service helpers is a package center for micro services

### Installation

```shell
go get github.com/Decem-Technology/service-helper

```

### Service Connections
import file

```go
import "github.com/Decem-Technology/service-helper/bootstrap"
```

#### MySQL
Environment

```
MYSQL_HOST=your redis host
MYSQL_PORT=3306
MYSQL_USERNAME=username
MYSQL_PASSWORD=secret
MYSQL_DBNAME=database name
```

```go
mysql := bootstrap.CreateMySQLConnection()
defer mysql.Close()
```

#### MongoDB
Environment

```
MONGODB_CONNECTION=mongodb://[username:password@]host1[:port1][,...hostN[:portN]][/[defaultauthdb][?options]]
```

```go
bootstrap.CreateMongoConnection()
```

#### Redis
Environment

```
REDIS_HOST=your redis host:6379
REDIS_PASSWORD=
```

```go
bootstrap.CreateRedisConnection()
```

#### Redis Cluster
Environment

```
REDIS_CLUSTER_HOST=your redis cluster ex. 127.0.0.1,127.0.0.2
REDIS_PASSWORD=
```

```go
bootstrap.CreateRedisClusterConnection()
```

#### SMTP Mail
Environment

```
MAIL_SMTP=smtp.google.com
MAIL_PORT=543
MAIL_USERNAME=username
MAIL_PASSWORD=password
```

```
bootstrap.CreateMailerConnection()
```

---

#### Start Project With load private key using 
Environment

```
LOAD_PRIVATE_KEY=true
```

*** make sure you have private.key in directory storage/ 

<BR />

### Use case example
Example Find data in mysql
#repositories/test-repository.go

```go
package repositories

import (
    "github.com/Decem-Technology/service-helper/bootstrap"
    "github.com/test-service/{service name}/entities"
)

type TestRepository struct {
    mysql bootstrap.MySQL
    // redis bootstrap.RedisDB
    // redis bootstrap.RedisClusterDB
    // mongo bootstrap.MongoDB
    // micro bootstrap.MicroModule
    // mail  bootstrap.Mailer
}

func (r *TestRepository) FindTest(id uint) (*entities.Test, error) {
    model := entities.Test{}
    if err := r.mysql.DB().Where("id = ?", id).First(&model).Error; err != nil {
        return nil, err
    }

    return model, nil
}

```

---

Example publish message to queue message
<br/>

#main.go > main()

```go
// New Service
service := micro.NewService(
    micro.Name("{service name}"),
    micro.Version("latest"),
)
service.Init()
// register client
bootstrap.RegisterClient(service.Client())

consumerGroup := os.Getenv("KAFKA_CONSUMERGROUP")
topics := map[string]string{
    "topic-key": os.Getenv("KAFKA_TOPIC_TEST"),
}

bootstrap.RegisterPublishers(topics)

micro.RegisterSubscriber(topics["topic-key"], service.Server(), subscribers.Test, server.SubscriberQueue(consumerGroup))

```
*** example at https://github.com/project/{service name}

```go
package repositories

import (
    "context"
    "github.com/Decem-Technology/service-helper/bootstrap"
    eventPB "github.com/Decem-Technology/service-helper/proto/event"
	somePB "github.com/project/{service name}/proto/test"

)

type SomeHandler struct {
    micro bootstrap.MicroModule
}

func (ctl *TestHandler) Find(ct context.Context, req *somePB.FindTestRequest, res *somePB.TestResponse) error {
    ctl.micro.Publisher("topic-key").Publish(context.TODO(), &eventPB.Event{
        "name": "some data"
    })
    res.Message = "success"
    return nil
}
```

Example subscriber function
```go
package subscribers

import (
	"context"
	"github.com/Decem-Technology/service-helper/helpers/dump"
	eventPB "github.com/Decem-Technology/service-helper/proto/event"
)

func Test(ct context.Context, msg *eventPB.Event) error {
	dump.DD(msg)
	return nil
}
```
---
Example verify request token

```go
package repositories

import (
    "context"
    "github.com/Decem-Technology/service-helper/bootstrap"
    "github.com/Decem-Technology/service-helper/contracts"
    eventPB "github.com/Decem-Technology/service-helper/proto/event"
	somePB "github.com/project/{service name}/proto/test"

)

type SomeHandler struct {
    micro bootstrap.MicroModule
}

func (ctl *TestHandler) Find(ct context.Context, req *somePB.FindTestRequest, res *somePB.TestResponse) error {
	ctx := contracts.AppContext{Context: ct}
    uID, claims, err := ctx.VerifyToken("*")
	if err != nil {
		return microError.Unauthorized("401", "%s", "Unauthorized")
	}
	if allow, err := ctx.VerifyPermission("view-stock"); (err != nil || allow == false) && claims.Audience != "machine" {
		return microError.Forbidden("403", "%s", "Forbidden")
    }
    
    _ = uID // user_id
    
    ctl.micro.Publisher("topic-key").Publish(context.TODO(), &eventPB.Event{
        "name": "some data"
    })
    res.Message = "success"
    return nil
}
```

---

generate protobuf file using

```bash
❯ protoc -I. --go_out=plugins=grpc,paths=source_relative:. --micro_out=paths=source_relative:. --proto_path=$GOPATH/src --proto_path=../  proto/response/meta.proto

❯ protoc -I. --go_out=plugins=grpc,paths=source_relative:. --micro_out=paths=source_relative:. --proto_path=$GOPATH/src --proto_path=../  proto/event/event.proto
```