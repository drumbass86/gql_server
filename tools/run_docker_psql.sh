#!/bin/bash
# Script for starting docker container PostgreSQL

container_name=gql_serv_db
docker_stat="$(docker container inspect -f '{{.State.Status}}' $container_name)"

if [[ $docker_stat != "running" ]]; then
   if [[ $docker_stat != "paused" || $docker_stat == "" ]]; then   
      echo "container $container_name dont paused or running. Kill and started new container!"
      docker container rm $container_name
      docker run -d \
      --name $container_name \
      -p 5432:4432 \
      -e POSTGRES_PASSWORD=123qwe \
      -e POSTGRES_USER=user \
      -v ${HOME}/postgres/data-gql-server:/var/lib/postgresql/data \
      postgres
   else
      docker container start $container_name
   fi
else
   echo "container $container_name already running!"
fi