FROM alpine:3
RUN apk --no-cache add nginx
COPY default.conf /etc/nginx/http.d/default.conf
CMD ["nginx", "-g", "daemon off;"]