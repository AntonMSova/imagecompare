docker build -t kaosd/imagecompare:latest -t kaosd/imagecompare:$SHA .
docker push kaosd/imagecompare:latest
docker push kaosd/imagecompare:$SHA

kubectl apply -f k8s
kubectl set image deployments/server-deployment server=kaosd/imagecompare:$SHA
