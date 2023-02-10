# Docker Filesystems

## Camunda with MariaDB


> start processes

```sh
cd docker
docker-compose -f camunda-mariadb.yml up
```

> deploy war into container

```sh
docker cp C:\Users\d.lupo\test\spring-camunda-template\target\spring-camunda-template-1.3-SNAPSHOT.war docker_camunda-mariadb_1:/camunda/webapps
```


### References


https://github.com/camunda/docker-camunda-bpm-platform

https://hub.docker.com/_/mariadb

https://docs.docker.com/engine/reference/commandline/cp/

https://docs.docker.com/compose/environment-variables/
