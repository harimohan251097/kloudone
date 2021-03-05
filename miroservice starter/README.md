# web-microservice-starter
Starter application to build HTTP based web microservices

This starter application provides an example to build a simple articleging API using gin-gonic.

# Table of Contents
1. [Project Structure](#project-Structure)
2. [API Documentation](#api-documentation)
3. [Database in GKE](#database-in-gke)
4. [Configuration Management](#configuration-management)
5. [RBAC](#rbac)
6. [Build and deploy](#build-and-deploy)

# Project structure
This project follows MVC pattern, except the View. The source code related the web application is present inside `src` folder and follows the folder structure of a Java Spring based application. 
> Developers should not modify main.go or the utils folder.

# API Documentation
API Documentation is done using swag. See https://github.com/swaggo/swag#getting-started.
To generate documentation, run
```
swag init -g routes.go
```

The swagger documentation will be available at the docs/swagger.json.

# Database in GKE
If you don't have a database already running in GKE, you'll need to create one.

In a terminal, use port forward to connect to the MySQL instance running in accuknox-dev-mysql namespace.
`kubectl port-forward -n accuknox-dev-mysql svc/mysql 3306:3306`

In another terminal, login to MySQL using the following command. For password, please contact Saleem via slack or at s@accuknox.com
`mysql -u root -p -h 127.0.0.1`

Run the following in the mysql shell.
```
# Create the database
CREATE DATABASE '<db_name>';

# Create a new user
CREATE USER '<user>'@'%' IDENTIFIED BY 'password';

# Grant access to the newly created user on the database
GRANT ALL PRIVILEGES ON <db_name>.* TO '<user>'@'%';
```
Usually, &lt;user&gt; is the name of the microservice that's accessing the database.

# Configuration management
Use conf/app.yaml for local testing only. For deployment specific configurations, use the configmap in `deployment/<env>-config.yaml`. Use dev-config.yaml as an example and modify according to your needs.

# RBAC
Currently, only Istio RBAC policies are supported. OPA policies will be supported in the future.
Edit your policy at `authx-policies/<env>-istio-authx-policy.yaml`. Refer https://istio.io/latest/docs/reference/config/security/

Apply using: 
```
kubectl apply -f authx-policies/dev-istio-authx-policy.yaml -n saleem-demo
```

# Build and deploy

## Building the microservice binary
While developing locally or for dev environment
```
CGO_ENABLED=0 go build -o bin/server
```

For building to deploy to production
```
export GIN_MODE=release && CGO_ENABLED=0 go build -o bin/server
```

## Build docker image
```
docker build -t gcr.io/accuknox/web-microservice-starter .
```

## Tag image
```
docker tag accuknox/web-microservice-starter gcr.io/accuknox/web-microservice-starter:<version>
```

## Configure gcloud and docker
Install gcloud SDK and follow https://cloud.google.com/container-registry/docs/quickstart#auth to configure docker to use gcr as container repository.

## Push to gcr.io
```
docker push gcr.io/accuknox/web-microservice-starter:<version>
```
Update image in deployment/dev-deployment.yaml

## Istioctl installation
Sidecars are not automatically injected to the pods. It is required that istioctl is installed to perform a step below.
Install istioctl at https://istio.io/latest/docs/setup/getting-started/#download

## Deploy
Do not edit the deployment/deploy.yaml directly. Edit dev-deployment.yaml and get the final deploy.yaml to use only for deploying the service in GKE. To deploy to other envs, create &lt;env&gt;-deployment.yaml to make env specific changes and run the deploy.yaml generation command.

```
kubectl apply -f deployment/<env>-config.yaml -n <namespace>
istioctl kube-inject -f deployment/dev-deployment.yaml > deployment/deploy.yaml
kubectl apply -f deployment/deploy.yaml -n <namespace>
kubectl apply -f deployment/service.yaml -n <namespace>
```

> Note, if you are deploying the services for the first time, add your service basepath [example: /webmicroservicestarter (basepath present in app.yaml) ] at https://github.com/accuknox/platform-istio-gateway. Otherwise, your APIs will not be accessible via api-<env>.accuknox.com

