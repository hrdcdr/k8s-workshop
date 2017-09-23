FROM scratch

ENV K8SAPP_LOCAL_HOST 0.0.0.0
ENV K8SAPP_LOCAL_PORT 8080
ENV K8SAPP_LOG_LEVEL 0

EXPOSE $K8SAPP_LOCAL_PORT

COPY certs /etc/ssl/certs/
COPY bin/linux-amd64/k8s-workshop /

CMD ["/k8s-workshop"]
