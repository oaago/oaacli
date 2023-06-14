server:
  name: oaatpl
  port: 9901
  env: local
  version: 1.0
  weight: 1
mysql:
  enable: false
  oaauser: oaa:123456@tcp(127.0.0.1:3306)/oaauser
nacos:
  enable: false
  ipaddr: nacos.com
  logdir: ./nacos
  cachedir: ./nacos
  dataid: oaatpl
  group: oaatpl
kafka:
  consumer:
    enable: false
    nodes:
      - 0.0.0.0:9093
    topic:
      - TestComposeTopic
    groupid: qqqq
  producer:
    enable: false
    nodes:
      - 0.0.0.0:9093
    topic: TestComposeTopic
redis:
  enable: false
  addr: 127.0.0.1:6937
  password: 123456
  db: 1
logger:
  enablekafka: true
  path: ./logs
  name:  oaatpl
docker:
  harbor:
    url: 192.168.3.85:8088