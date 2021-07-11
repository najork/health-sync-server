FROM scratch

ENV SERVICE_HOME=/opt/apollo/services/{{Product}}

RUN mkdir -p /opt/apollo/binaries && mkdir -p $SERVICE_HOME/bin/linux-amd64

# Copy gödel build output for linux-amd64 os/arch
COPY {{InputBuildArtifact Product "linux-amd64"}} /opt/apollo/binaries

RUN ln -s /opt/apollo/binaries/{{Product}} $SERVICE_HOME/bin/linux-amd64/{{Product}}

EXPOSE 8443
EXPOSE 8444

WORKDIR $SERVICE_HOME

ENTRYPOINT ["bin/linux-amd64/{{Product}}"]
CMD ["server"]
