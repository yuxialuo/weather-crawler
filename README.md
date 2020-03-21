# weather-crawler
获取中国部分地区当日天气。

1)拉取elasticsearch镜像
docker pull docker.elastic.co/elasticsearch/elasticsearch:7.6.1

2)创建并启动容器
docker run -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.6.1

3)go build

4)./weather-crawler
