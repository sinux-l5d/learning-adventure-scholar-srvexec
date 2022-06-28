FROM srvexec:bin-python-generic as runtime
FROM python:3.10-slim as executor

WORKDIR /app

COPY --from=runtime /srvexec-python-generic ./srvexec-python-generic

EXPOSE 8080

ENTRYPOINT [ "./srvexec-python-generic" ]
