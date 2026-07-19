FROM golang:1.25-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux \
  go build -trimpath -ldflags="-s -w" \
  -o /out/stagepass .

FROM scratch

COPY --from=builder /out/stagepass /stagepass

USER 65532:65532

EXPOSE 8080

ENTRYPOINT ["/stagepass"]