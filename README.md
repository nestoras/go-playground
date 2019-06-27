### Create PostgreSQL database in Docker

```
docker run -d -p 5432:5432 --name my-postgres -e POSTGRES_PASSWORD=123456 postgres
docker exec -it my-postgres bash
psql -U postgres
CREATE DATABASE golang_course;

docker start my-postgres


# Mysql
docker run -d -p 3306:3306  --name mysql -e MYSQL_ROOT_PASSWORD=my-secret-pw -d mysql:latest
docker exec -it mysql bash 
mysql mysql -u root -p
CREATE DATABASE golang_course;
docker start mysql
```