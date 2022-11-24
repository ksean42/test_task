FROM golang:latest

COPY ./ ./
ENV GOPATH=/
RUN go build cmd/userGradeService.go
CMD ["./userGradeService"]