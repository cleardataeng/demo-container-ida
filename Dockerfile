FROM scratch
ADD ida /
EXPOSE 8081
USER 1000
CMD ["/ida"]
