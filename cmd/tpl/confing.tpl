server:
  name: %package%
  port: 9932
  env: local
  version: 1.0
  weight: 1
mysql:
  enable: true
  oaauser: root:root@tcp(127.0.0.1:3306)/oaauser
nacos:
  enable: true
  ipaddr: nacos.com
  logdir: ./nacos
  cachedir: ./nacos
  dataid: %package%
  group: %package%
kafka:
  consumer:
    enable: true
    nodes:
      - 0.0.0.0:9093
    topic:
      - TestComposeTopic
    groupid: qqqq
  producer:
    enable: true
    nodes:
      - 0.0.0.0:9093
    topic: TestComposeTopic
redis:
  enable: true
  addr: 127.0.0.1:6937
  password: 123456
  db: 1
logger:
  enablekafka: true
  path: ./logs
  name:  %package%
docker:
  harbor:
    url: 192.168.3.85:8088