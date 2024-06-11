# learning-docker
Container Volume - Yuni Nur Rohmatilah

### 1. Menghapus container postgre lama 

### 2. Run Docker SQL "my-postgres-yuni" with "my-pg-volume-yuni"
#### docker run -d \
#### > --name my-postgres-yuni \
#### > -e POSTGRES_USER=postgres \
#### > -e POSTGRES_PASSWORD=password \
#### > -v my-pg-volume-yuni:/var/lib/postgresql/data \
#### > -p 5430:5432 postgres 

<img width="995" alt="image" src="https://github.com/yuninurr/learning-docker/assets/89186114/9eb38ef5-ace1-46ad-a7e0-74aa6337a14b">

### 3. Connection Succes
#### Start Container
![image](https://github.com/yuninurr/learning-docker/assets/89186114/a49a276b-521e-41cc-a1a0-26a29efaad9d)
#### Connect Database 
<img width="1440" alt="image" src="https://github.com/yuninurr/learning-docker/assets/89186114/28079abe-2957-45cc-95c5-293d5711be93">

#### 4. Create New Table In PgAdmin
<img width="1440" alt="image" src="https://github.com/yuninurr/learning-docker/assets/89186114/a3b168d4-d0a8-4951-aca6-fd09f135409e">

### 5. Stop and Delete First Container
![image](https://github.com/yuninurr/learning-docker/assets/89186114/c97b980c-8acf-4eba-a0d1-12bf44d0c8b5)
![image](https://github.com/yuninurr/learning-docker/assets/89186114/031ab8e4-263e-4d53-83a6-49897ae52bf7)

### 6. New name Postgres
#### docker run -d \
#### --name my-postgres-v2-yuni \
#### -e POSTGRES_USER=postgres \
#### -e POSTGRES_PASSWORD=password \
#### -v my-pg-volume-yuni:/var/lib/postgresql/data \
#### -p 5430:5432 postgres
<img width="1001" alt="image" src="https://github.com/yuninurr/learning-docker/assets/89186114/da2edbe2-89dc-4531-a844-d7edaa67d1b1">
<img width="1440" alt="image" src="https://github.com/yuninurr/learning-docker/assets/89186114/c1df3770-8398-4ffa-8d0e-d2494ae1b4c7">
<img width="1440" alt="image" src="https://github.com/yuninurr/learning-docker/assets/89186114/8312a896-1066-4b3c-8abf-435a2e9adc39">

### 7. Check Table in PgAdmin
<img width="1440" alt="image" src="https://github.com/yuninurr/learning-docker/assets/89186114/f92a0ad1-a7b5-416f-a7ed-9a5f164a74e9">




