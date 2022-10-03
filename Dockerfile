ARG image=registry.access.redhat.com/ubi8/ubi-minimal:8.6
FROM $image

ENV HOME=/tmp
WORKDIR /tmp

COPY customization.yaml /customization.yaml
COPY openmetadata-connector /openmetadata-connector

USER 10001
EXPOSE 8080

ENTRYPOINT ["/openmetadata-connector"]
CMD [ "run", "--customization", "/customization.yaml" ]
