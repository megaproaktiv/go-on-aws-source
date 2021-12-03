## We specify the base image we need for our
## go application
FROM golang:bullseye
## We create an /app directory within our
## image that will hold our application source
## files
RUN mkdir /app /app/templates /app/assets /app/assets/css /app/assets/js
## We copy everything in the root directory
## into our /app directory
ADD dist/dosshow /app 
COPY "templates/*.tmpl" /app/templates
COPY "assets/js/*.js" /app/assets/js
COPY "assets/css/*.css" /app/assets/css
## We specify that we now wish to execute 
## any further commands inside our /app
## directory
WORKDIR /app
## we run go build to compile the binary
## executable of our Go program
## Our start command which kicks off
## our newly created binary executable
CMD ["/app/dosshow"]
