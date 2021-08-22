FROM scratch

ENV SERVICE_HOME=/opt/health-sync-server/services/{{Product}}

RUN mkdir -p /opt/health-sync-server/binaries && mkdir -p $SERVICE_HOME/bin/linux-amd64

# Copy gödel build output for linux-amd64 os/arch
COPY {{InputBuildArtifact Product "linux-amd64"}} /opt/health-sync-server/binaries

RUN ln -s /opt/health-sync-server/binaries/{{Product}} $SERVICE_HOME/bin/linux-amd64/{{Product}}

EXPOSE 8443
EXPOSE 8444

WORKDIR $SERVICE_HOME

ENTRYPOINT ["bin/linux-amd64/{{Product}}"]
CMD ["server"]
