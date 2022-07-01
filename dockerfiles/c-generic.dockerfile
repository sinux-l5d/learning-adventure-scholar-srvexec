FROM srvexec:bin-c-generic as runtime
FROM gcc:12.1 as executor

WORKDIR /app

COPY --from=runtime /srvexec-c-generic ./srvexec-c-generic

EXPOSE 8080

ENV SRVEXEC_LISTEN=0.0.0.0
ENV SRVEXEC_PORT=8080
ENV SRVEXEC_TIMEOUT="5s"
ENTRYPOINT [ "./srvexec-c-generic" ]