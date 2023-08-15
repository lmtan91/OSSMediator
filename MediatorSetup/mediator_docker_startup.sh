#!/bin/bash
set -e

version=$(cat ../VERSION)
heap_size=${heap_size:-2g}
while [ $# -gt 0 ]; do

   if [[ $1 == *"--"* ]]; then
        param="${1/--/}"
        declare $param="$2"
   fi

  shift
done

if ! [ -x "$(command -v docker)" ]; then
  echo "Docker not installed, please install docker and re-run the script."
	exit 1
fi

echo "Installing OSSMediator $version version"

mkdir -p es_data
chmod g+rwx es_data
chgrp 0 es_data
chown 1000:1000 es_data/
docker start ndac_oss_opensearch 2>/dev/null || docker run --name "ndac_oss_opensearch" --restart=always -t -d -p 9200:9200 -p 9600:9600 --ulimit nofile=65535:65535 -e "discovery.type=single-node" -e 'DISABLE_SECURITY_PLUGIN=true' -e OPENSEARCH_JAVA_OPTS="-Xms$heap_size -Xmx$heap_size" -v $(pwd)/es_data:/usr/share/opensearch/data opensearchproject/opensearch:2.8.0

mkdir -p grafana_storage
chown 472:472 grafana_storage/
docker start ndac_grafana 2>/dev/null || docker run -d --name "ndac_grafana" --network host -e "GF_INSTALL_PLUGINS=grafana-opensearch-datasource" -e "GF_USERS_DEFAULT_THEME=light" -v $(pwd)/grafana_data/provisioning/dashboards:/etc/grafana/provisioning/dashboards -v $(pwd)/grafana_data/provisioning/datasources:/etc/grafana/provisioning/datasources -v $(pwd)/grafana_data/dashboards:/etc/grafana/dashboards grafana/grafana-oss:9.3.6

## Cutomizing login page
# Replace Logo
docker cp ./customizied_login_page/Logo_Celanese.svg ndac_grafana:/usr/share/grafana/public/img/grafana_icon.svg
# Update Title
# Update Login Title
docker cp ./customizied_login_page/build/2362.5e93872490cc5d80351e.js ndac_grafana:/usr/share/grafana/public/build/ 
docker cp ./customizied_login_page/build/2362.5e93872490cc5d80351e.js.map ndac_grafana:/usr/share/grafana/public/build/
# Update background
docker cp ./customizied_login_page/light-background.svg ndac_grafana:/usr/share/grafana/public/img/g8_login_light.svg

docker start ndac_oss_collector 2>/dev/null || docker run -d --name "ndac_oss_collector" --network host -v $(pwd)/reports:/reports  -v $(pwd)/collector/log:/collector/log -v $(pwd)/collector/checkpoints:/collector/bin/checkpoints -v $(pwd)/.secret:/collector/bin/.secret -v $(pwd)/collector_conf.json:/collector/resources/conf.json ossmediatorcollector:$version
sleep 30

docker start ndac_oss_elasticsearchplugin 2>/dev/null || docker run -d --name "ndac_oss_elasticsearchplugin" --network host -v $(pwd)/reports:/reports -v $(pwd)/plugin/log:/plugin/log -v $(pwd)/plugin_conf.json:/plugin/resources/conf.json elasticsearchplugin:$version

echo "OSSMediator is started, open http://<IP_Address>:3000/dashboards to view the dashboards."
