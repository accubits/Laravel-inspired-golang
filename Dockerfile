FROM golang

ARG app_env
ENV APP_ENV $app_env
ENV GO111MODULE=on

COPY ./app /go/src/myproject/app
WORKDIR /go/src/myproject/app


COPY ./app/go.mod . 
COPY ./app/go.sum .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .

CMD if [ ${APP_ENV} = production ]; \
	then \
	app; \
	else \
	go get github.com/pilu/fresh && \
	fresh; \
	fi
	
EXPOSE 8080

