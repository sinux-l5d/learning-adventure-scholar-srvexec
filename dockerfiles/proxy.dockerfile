FROM srvexec:bin-proxy as runtime
FROM scratch

WORKDIR /app

COPY --from=runtime /srvexec-proxy ./srvexec-proxy

EXPOSE 8080

ENV SRVEXEC_LISTEN=0.0.0.0
ENV SRVEXEC_PORT=8080
VOLUME [ "/app/.proxy.json" ]

ENTRYPOINT [ "./srvexec-proxy" ]