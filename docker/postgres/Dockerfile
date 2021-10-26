FROM postgres:12.2

COPY ./scripts/* /

ENTRYPOINT ["./entrypoint.sh"]

EXPOSE 5432
CMD ["postgres"]
