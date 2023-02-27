FROM golang

RUN mkdir /usr/local/go/src/plDockerKubernetes
COPY main.go /usr/local/go/src/plDockerKubernetes
COPY server/ /usr/local/go/src/plDockerKubernetes/server
RUN cd /usr/local/go/src/plDockerKubernetes && go mod init plDockerKubernetes
RUN cd /usr/local/go/src/plDockerKubernetes && ls
RUN cd /usr/local/go/src/plDockerKubernetes && go build -o server.exe ./main.go
RUN ls /usr/local/go/src/plDockerKubernetes
ENTRYPOINT ["/usr/local/go/src/plDockerKubernetes/server.exe"]

