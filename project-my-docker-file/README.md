# learning-docker
## Project 1 - Build Docker File to Docker Image

## 1. Create a new Golang project that serve http
- Path "/", do print string/html text of your favorite wise word
- When you access path "" , also print to log/console using fm. Println with text "Ouch!"
- Start HTTP on port 77 (or 78 if your chrome blocked the port)
- Modification go.mod line 3 "go 1...."change into "go 1.21"
<img width="1440" alt="Screenshot 2024-06-14 at 00 00 30" src="https://github.com/yuninurr/learning-docker/assets/89186114/64e38a8c-e339-428f-84f0-180723e543eb">
<img width="501" alt="Screenshot 2024-06-14 at 00 01 29" src="https://github.com/yuninurr/learning-docker/assets/89186114/1c958f12-32e9-489a-919b-e9e8a312600e">
 
## 2. Create file "AUTHORS.md" and "LINKS.md" within same folder with "Dockerfile"
- Edit "AUTHORS.md" file and fill it with your first name or your fake name
- Edit "LINKS.md" file and fill it with your github profile link

## 3. Create Dockerfile to build and run your golang project
- Use Base Image "golang:1.21"
- Set WORKDIR with /myapp
- RUN some command "go version" after WORKDIR
- COPY "AUTHORS.md" to image BEFORE run build golang (go build)
- COPY "LINKS.md" to image BEFORE run build golang. (go build)
- Make golang build output with name "my-go-app" and make sure it will run correctly
- <img width="1440" alt="Screenshot 2024-06-14 at 00 06 18" src="https://github.com/yuninurr/learning-docker/assets/89186114/d5122116-8277-431a-b8d4-fdfd41cf8038">

## 4. Build Dockerfile image with name "my-go-app:v2"
<img width="1440" alt="Screenshot 2024-06-14 at 00 06 18" src="https://github.com/yuninurr/learning-docker/assets/89186114/0205595f-153a-4411-afef-8a2f01b3b5e2">

## 5. Run the image into container with name "go-app" and expose to port host 5555
<img width="1440" alt="Screenshot 2024-06-14 at 00 06 29" src="https://github.com/yuninurr/learning-docker/assets/89186114/1d177dab-d039-48c9-8f9e-8f5fa9afe989">

## 6. Access your golang inside docker via localhost:5555 and see logs of your container
"go-app"
<img width="1440" alt="Screenshot 2024-06-14 at 00 07 20" src="https://github.com/yuninurr/learning-docker/assets/89186114/abd79887-7375-4736-9ca2-81de0c6e9988">
<img width="1440" alt="Screenshot 2024-06-14 at 00 07 12" src="https://github.com/yuninurr/learning-docker/assets/89186114/6eb680d2-f129-4313-9fd7-ad3aed8a2106">



