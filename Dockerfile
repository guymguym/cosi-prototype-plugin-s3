FROM fedora

COPY ./plugin /usr/local/bin/plugin

ENTRYPOINT ["plugin"]
