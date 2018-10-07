# build stage
FROM golang:alpine AS build-stage
COPY . /src
RUN cd /src && go build -o htmlserver 

# final stage
FROM alpine
COPY --from=build-stage /src/htmlserver /usr/local/bin/
COPY ./html /html
ENTRYPOINT ["htmlserver"]
CMD ["run"]
