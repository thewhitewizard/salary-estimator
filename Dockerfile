
# same version as Scone
FROM golang:1.19.7-alpine3.17 AS build
RUN apk add --no-cache gcc build-base gcc-go
COPY . /app
WORKDIR /app

RUN go build -compiler=gccgo -buildmode=exe 


FROM alpine:3.17.2
RUN apk add --no-cache libgo
RUN mkdir /iexec_in /iexec_out
COPY --from=build /app/salary-estimator /app/salary-estimator
ENV IEXEC_IN=/iexec_in
ENV IEXEC_OUT=/iexec_out
ENTRYPOINT ["/app/salary-estimator"]
