#### Install postgres and pgadmin in docker


#### Install Docker:
```bash
$ apt-get update
$ apt-get install docker.io
```

#### Install postgres
```bash
$ docker volume create postgresqldata
$ docker run -d -v postgresqldata:/data/db -e POSTGRES_PASSWORD=123456 --name postgres -p 5432:5432 postgres:13
$ docker volume ls
$ docker ps -a
$ docker exec -it postgres /bin/bash
```

#### Install pgadmin
```bash
$ docker run --name pgadmin -e "PGADMIN_DEFAULT_EMAIL=myemail@gmail.com" -e "PGADMIN_DEFAULT_PASSWORD=123456" -p 8000:80 -d dpage/pgadmin4 
$ docker ps -a
```

#### Link them with Docker network and set up Pgadmin
```bash
$ docker network create --driver bridge pgnetwork
$ docker network ls
$ docker network connect pgnetwork pgadmin
$ docker network connect pgnetwork postgres
$ docker network inspect pgnetwork
```

#### Set up Pgadmin and confirm connection postgres
```bash
http://localhost:8000/
```

Thanks,
