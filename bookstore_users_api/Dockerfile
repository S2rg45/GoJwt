FROM golang:1.12.13

ENV REPO_URL=SiigoApiGo/bookstore_users_api

ENV GOPATH=/app

ENV APP_PATH=${GOPATH}/${REPO_URL}

ENV WORKPATH=${APP_PATH}/src
COPY  src ${WORKPATH}
WORKDIR ${WORKPATH}

RUN go get -d -v ./...
# RUN go install -v ./...
RUN go build -o dev-billing .
#Expose 8080 to the world
EXPOSE 8080

CMD [ "./dev-billing" ]
