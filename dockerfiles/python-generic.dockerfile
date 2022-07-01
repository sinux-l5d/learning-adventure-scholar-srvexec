FROM srvexec:bin-python-generic as runtime
FROM python:3.10-slim as executor

WORKDIR /app

COPY --from=runtime /srvexec-python-generic ./srvexec-python-generic

EXPOSE 8080

ENV SRVEXEC_LISTEN=0.0.0.0
ENV SRVEXEC_PORT=8080
ENV SRVEXEC_TIMEOUT="5s"
ENTRYPOINT [ "./srvexec-python-generic" ]
