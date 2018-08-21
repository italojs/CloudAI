FROM golang:1.10.2-alpine3.7
RUN apk add -U git
COPY . /go/src/github.com/italojs/AI-Cloud/projects/backend/micro-services/crop-image
WORKDIR /go/src/github.com/italojs/AI-Cloud/projects/backend/micro-services/crop-image/app
RUN go get -u github.com/Masterminds/glide/...
RUN glide up
RUN go build -o /usr/bin/app .
RUN chmod +x .
EXPOSE 3001
CMD ["app"]

# FROM alpine:3.6
# RUN mkdir crop-image
# COPY --from=builder /go/src/github.com/italojs/AI-Cloud/projects/backend/micro-services/crop-image/app/app ./crop-image
# WORKDIR /crop-image
