# Webflix

A simple website written in HTML, CSS, Javascript and Golang.

## App Preview

![website image](https://github.com/user-attachments/assets/9143272f-f1c3-4d92-b3e6-61f6d30d59a2)

# CI/CD

 - Continous building and testing, and further containerization of the app is handled by Git Actions (cicd.yaml). 
 - The app itself is run on a WSL 2 ubuntu container inside a minikube cluster. 
 - The app is exposed using ingress addons(uses nginx ingress controller) provided by minikube. 
 - Continous deployment is facilitated by Helm charts + ArgoCD (GitOps)

## App Architecture

![Architecture image](https://github.com/user-attachments/assets/d841a524-e3c1-4ef8-b94e-811279310b2d)

## App directory structure

![Directory structure image](https://github.com/user-attachments/assets/52c99795-5559-4b77-9e9d-b05456786f2d)

## Running the server

To run the server, execute the following command:

```bash
go run main.go
```

The server will start on port 8080. You can access it by navigating to `http://localhost:8080/home` in your web browser.



