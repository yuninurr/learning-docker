# learning-docker
Docker welcome - Yuni Nur Rohmatilah

## Golang Project the Serve http
#### go mod init exampledockerfile
#### go get -u github.com/labstack/echo/v4
#### go get -u github.com/labstack/echo/v4/middleware
#### go run main.go */

<img width="1440" alt="image" src="https://github.com/yuninurr/learning-docker/assets/89186114/d5c292ad-8e3e-4682-a649-96c0a483980b">


## Create Docker file to Build and Run Golang Project 
#### docker build --tag my-wise-word .
#### docker image ls 
#### docker run -d -p 8080:8080 --name my-dcoker-go-ping-container my-wise-word

<img width="1440" alt="image" src="https://github.com/yuninurr/learning-docker/assets/89186114/5c5d1815-0cfd-4dc0-9301-be7492380578">

<img width="1201" alt="image" src="https://github.com/yuninurr/learning-docker/assets/89186114/4b756caa-43a1-4388-b170-3d2d52d3f7ed">

## See Logs of "my-wise-word" 
#### docker logs -f --tail 10 my-dcoker-go-ping-container
<img width="1146" alt="image" src="https://github.com/yuninurr/learning-docker/assets/89186114/443c9ce1-ecae-45f0-b562-3e8b29c24f1f">



