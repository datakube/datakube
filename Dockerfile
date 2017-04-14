FROM scratch
MAINTAINER Manuel Laufenberg <hello@manuel-laufenberg.de>

COPY dist/datahamster-worker /datahamster-worker
ENTRYPOINT ["/datahamster-worker"]