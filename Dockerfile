FROM scratch

ENV PORT 8000
EXPOSE $PORT

COPY go-k8s-microservice /
CMD ["/go-k8s-microservice"]