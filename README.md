# YYPush Configuration Tool

`ZOOKEEPER` environmental variable should be defined before you can use this tool.

```
export ZOOKEEPER=10.13.2.43:2181,10.13.2.44:2181
```

## Get Configuration

```
yypushc get /logpush/access_log
```

## Add IP(s) to configuration

```
yypushc ipadd /logpush/access_log 10.85.91.167
```

## Delete IP(s) from configuration

```
yypushc ipdel /logpush/access_log 10.85.91.167,10.85.91.168
```
