# learning-docker
# Task 3 : Run Multiple Containers - Yuni Nur Rohmatilah

## A. Single Database "yuni_access", single table "access_log" with single column named "timestamp" with data type timestamp or datetime
<img width="1440" alt="Screenshot 2024-06-11 at 16 17 00" src="https://github.com/yuninurr/learning-docker/assets/89186114/3410112d-c0b0-4ba1-8b4b-5e56937c7bd1">
<img width="1440" alt="Screenshot 2024-06-11 at 16 23 57" src="https://github.com/yuninurr/learning-docker/assets/89186114/b7956ade-36ef-409f-8c34-eafa48a8d31a">
<img width="1440" alt="Screenshot 2024-06-11 at 16 24 10" src="https://github.com/yuninurr/learning-docker/assets/89186114/347158b8-70fe-4d0e-af85-cb5155369007">

## B. Create or use existing golang project that serve http and connect to postgre (or any) databases
<img width="159" alt="image" src="https://github.com/yuninurr/learning-docker/assets/89186114/fc99bbf8-ac12-41e3-8842-c52ab234b982">
### - Path "/", print html text of your favorite wise word
#### -also when access to "/" print to console using fmt.Println with text "Ouch!"
<img width="451" alt="image" src="https://github.com/yuninurr/learning-docker/assets/89186114/94b8aa24-e162-408d-ac2d-d4dbd872e7b1">
### - Path "/ping" -> do ping to database and show the result (success connect or failed) in browser
<img width="459" alt="image" src="https://github.com/yuninurr/learning-docker/assets/89186114/c7801af6-eda1-46cb-a8c9-856621563877">
#### - new record current timestamp 2024-01-31 19:55
### - Expose to port 78 
<img width="658" alt="image" src="https://github.com/yuninurr/learning-docker/assets/89186114/1510198e-713d-4fc0-aa8c-2aeb0391d705">

## C. Create Dockerfile to build and run your golang project 
<img width="1440" alt="image" src="https://github.com/yuninurr/learning-docker/assets/89186114/72dbc867-4e5a-4fe1-90d2-a79613da4f33">
go run main.go 
<img width="451" alt="image" src="https://github.com/yuninurr/learning-docker/assets/89186114/94b8aa24-e162-408d-ac2d-d4dbd872e7b1">
<img width="459" alt="image" src="https://github.com/yuninurr/learning-docker/assets/89186114/c7801af6-eda1-46cb-a8c9-856621563877">

## D. Create docker-compose.yml and instruct to run 2 containers 
### - Golang Project (give name : "go_service"), expose container port 78 to host port 9911
<img width="1440" alt="image" src="https://github.com/yuninurr/learning-docker/assets/89186114/bc1169be-e4df-4a06-91ef-aa1e8eea24ea">
### - Postgres Image postgres:latest (give name : "Database")
<img width="1440" alt="image" src="https://github.com/yuninurr/learning-docker/assets/89186114/aaff9de7-42ed-4001-a15a-a246b78d8f62">

## E. Make sure the golang project can connect to postgresqll and postgresql data is persistent 
### - network name = "netwwc" and volume name = "volwwc"
<img width="1440" alt="image" src="https://github.com/yuninurr/learning-docker/assets/89186114/6ba34503-b5d1-4e91-9fac-9b93be09d2d3">

## F. Run Docker Compose with name "my-wise-word-compose"
docker compose up --build
<img width="1440" alt="image" src="https://github.com/yuninurr/learning-docker/assets/89186114/02e4efe2-8905-426f-a484-0cf0e823660a">
### - make sure access 9911 and show a wise word
<img width="623" alt="Screenshot 2024-06-12 at 22 41 44" src="https://github.com/yuninurr/learning-docker/assets/89186114/dba50881-23bc-45fd-b5b6-6e194e2e9eac">
### - make sure acces 9911/ping and show Success (with succes ping to database) and new data inserted in your database
<img width="624" alt="Screenshot 2024-06-12 at 22 41 56" src="https://github.com/yuninurr/learning-docker/assets/89186114/237cd671-dfa5-4a17-8513-8583e0715871">

## G. Check logs of your each container in "my-wise-word-compose"
<img width="1440" alt="Screenshot 2024-06-12 at 22 40 32" src="https://github.com/yuninurr/learning-docker/assets/89186114/286b7aad-74d4-4efa-96d7-0970d5e6024e">
<img width="1440" alt="Screenshot 2024-06-12 at 22 40 39" src="https://github.com/yuninurr/learning-docker/assets/89186114/becdb0e6-875e-4cec-a3e3-8218153a2af7">
