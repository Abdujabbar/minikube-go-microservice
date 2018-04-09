FROM scratch

COPY go-k8s-microservice /

ENV PORT 8000
EXPOSE $PORT
CMD ["/go-k8s-microservice"]