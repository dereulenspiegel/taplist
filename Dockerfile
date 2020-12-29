FROM scratch
COPY taplist /taplist
EXPOSE 9099
VOLUME ["/data"]
ENTRYPOINT ["/taplist"]
