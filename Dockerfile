FROM scratch

COPY app /

EXPOSE 8000
CMD ["/app"]