#begin 
FROM public.ecr.aws/lambda/provided:al2 AS build
ENV CGO_ENABLED=0
RUN mkdir -p /opt/extensions
RUN yum -y install go 
RUN go env -w GOPROXY=direct
ADD go.mod go.sum ./
RUN go mod download
COPY . ${LAMBDA_TASK_ROOT}
RUN env GOOS=linux GOARCH=arm64 go build -o=/main 
# copy artifacts to a clean image
FROM public.ecr.aws/lambda/provided:al2
COPY --from=build /main /main
ENTRYPOINT [ "/main" ]
#end