FROM debian:stable-slim AS build

WORKDIR /usr/local/src

COPY . .

RUN dpkg-deb --build . battery-notify.deb


FROM scratch

COPY --from=build /usr/local/src/battery-notify.deb .
