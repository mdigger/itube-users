############################
# STEP 1 build executable binary
############################
# golang alpine
FROM golang:alpine AS builder

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

# Create appuser
ENV USER=appuser
ENV UID=10001

# See https://stackoverflow.com/a/55757473/12429735
RUN adduser \
--disabled-password \
--gecos "" \
--home "/nonexistent" \
--shell "/sbin/nologin" \
--no-create-home \
--uid "${UID}" \
"${USER}"
WORKDIR $GOPATH/src/itube/user/
COPY . .

# Fetch dependencies.
# RUN go get -d -v
RUN go mod download
RUN go mod verify

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
-ldflags='-w -s -extldflags "-static"' -a \
-o /go/bin/server ./cmd/itube-users-server/main.go


############################
# STEP 2 build a small image
############################
FROM scratch

# Import from builder.
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Copy our static executable
COPY --from=builder /go/bin/server /server

# Use an unprivileged user.
USER appuser:appuser

# Run the server.
ENV PORT=50051 DSN= GOOGLE_CLIENT_ID= GOOGLE_SECRET= SMTP= TEMPLATES= LOG_LEVEL=
EXPOSE ${PORT}
VOLUME ["/templates"]
ENTRYPOINT ["/server"]