

scope, user_id,act




request: 
```json
{
    "act":"xty",
    "user_id":"12333", // 用户id
}
```

system:
```json
{
    "weight":2, // 权重 默认为1
    "timestamp": 1723321312, // 当前时间发送时间
}
```

rules:
```json
{
    "scope":[
        "ident":["xyt","ff"],
        "valiter"["xyt"=>"act","ff"=>"act2"]
    ],
    "type": "limit", // 限制类型
    "value": "8*3600", // 限制值
}
```

用户ID合法性校验
身份的合法性校验


- 活动限制每人参加次数；
  - key为: act:xyt:user_id:12333
  - weight:2

- 活动总限制次数；
  - key为: act:xyt
  - weight:2

- 活动限制8小时内参与人数；
  - key为： act:xyt
  - weight:2
  - 



request: 

```json
{
    "act":"xty",
    "user_id":"12333", // 用户id
    "ident":"经纪人"
}
``` 



```
{
    "path": "(a,c)|c,d,e",
    "rules":[
        "a": []
    ]
}
``````










