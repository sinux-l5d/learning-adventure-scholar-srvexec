FROM srvexec:bin-python as runtime
FROM python:3.10-slim as executor

WORKDIR /app

COPY --from=runtime /srvexec-python-generic ./srvexec-python-generic

ENTRYPOINT [ "./srvexec-python-generic" ]
