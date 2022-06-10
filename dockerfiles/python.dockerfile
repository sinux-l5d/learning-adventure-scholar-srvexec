FROM srvexec:bin-python as runtime
FROM python:3.10-slim as executor

WORKDIR /app

COPY --from=runtime /srvexec-python ./srvexec-python

ENTRYPOINT [ "./srvexec-python" ]
