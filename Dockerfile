FROM fedora
RUN dnf install -y golang
COPY hello_openshift.go /hello-openshift.go
RUN go build /hello-openshift.go
EXPOSE 8080 8888
ENTRYPOINT ["/hello-openshift"]
