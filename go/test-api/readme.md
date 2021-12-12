#to deploy to docker
do docker-compose up in the project directory

#create an image using :

docker build -t dockerprofilename/imagename .

kubectl apply -f ./k8s

minikube service --url <service-name>

minikube service --url api-service

then use this ip to hit service apis
