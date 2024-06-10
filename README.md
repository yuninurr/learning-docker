# learning-docker
Docker welcome - Yuni Nur Rohmatilah

### Run a new Container
docker run -d -p 8080:80 --name welcome1 docker/welcome-to-docker

<img width="1440" alt="image" src="https://github.com/yuninurr/learning-docker/assets/89186114/b910841e-45b6-4228-96ca-7f3280527bc3">

### Try In Browser 
<img width="1440" alt="image" src="https://github.com/yuninurr/learning-docker/assets/89186114/20bfa45e-5057-473d-8f51-931a454ca2db">

### See Logs of running welcome1 container 
<img width="682" alt="image" src="https://github.com/yuninurr/learning-docker/assets/89186114/2a5c5fef-0327-4227-b9e6-253fe28ca309">

### Execute a Command inside running welcome1 container 
<img width="1440" alt="image" src="https://github.com/yuninurr/learning-docker/assets/89186114/68ebffa0-6294-4adf-9648-7ba0bc6bd595">

### Stop Container and see latest logs 
docker stop welcome1 
docker logs -f --tail 10 welcome1
<img width="682" alt="image" src="https://github.com/yuninurr/learning-docker/assets/89186114/250a9ade-38e9-41b9-b84c-23f11a25102c">



