FROM gcr.io/distroless/base-debian10
WORKDIR /usr/src/good-ponds
COPY server .
CMD [ "/usr/src/good-ponds/server" ]