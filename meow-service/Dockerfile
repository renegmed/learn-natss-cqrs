FROM golang:1.13.1-alpine as builder

WORKDIR /app

COPY . .
 
RUN go mod vendor
RUN go build -o meow .

# RUN pwd
# RUN ls -als

#ENTRYPOINT ["./meow"]




# Using golang:latest creates an image that consumes a lot 
# of memory and hence not production friendly. 
#
# So we used golang:alpine which is very light.
#
# old image size 561MB
# new image size 19.2MB

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/meow .
#COPY --from=builder /app/.env .       

# Expose port 8080 to the outside world
EXPOSE 8080

#Command to run the executable
#CMD ["./main"]

ENTRYPOINT ["./meow"]
