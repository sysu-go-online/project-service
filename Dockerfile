FROM ubuntu
ADD main /
RUN apt update && apt install -y git

ENTRYPOINT ["/main"]
EXPOSE 8080
