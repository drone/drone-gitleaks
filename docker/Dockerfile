FROM alpine:3 as alpine
RUN apk add -U --no-cache ca-certificates

FROM alpine:3
ENV GODEBUG netdns=go

COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

RUN wget https://github.com/zricethezav/gitleaks/releases/download/v7.6.1/gitleaks-linux-amd64
RUN mv gitleaks-linux-amd64 /bin/gitleaks
RUN chmod +x /bin/gitleaks

RUN apk add --no-cache bash git openssh-client
ADD release/linux/amd64/plugin /bin/
ENTRYPOINT ["/bin/plugin"]