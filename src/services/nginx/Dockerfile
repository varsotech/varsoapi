FROM node:21-alpine3.18 as builder

WORKDIR /app

COPY ./src/services/www/package.json ./
RUN npm i

COPY ./src/services/www/ .
RUN npm run build


FROM nginx:latest

RUN mkdir -p /etc/letsencrypt
RUN curl -s https://raw.githubusercontent.com/certbot/certbot/master/certbot-nginx/certbot_nginx/_internal/tls_configs/options-ssl-nginx.conf > "/etc/nginx/options-ssl-nginx.conf"
RUN curl -s https://raw.githubusercontent.com/certbot/certbot/master/certbot/certbot/ssl-dhparams.pem > "/etc/nginx/ssl-dhparams.pem"

# Copy the init.sh script to the container
COPY ./src/services/nginx/init.sh /docker-entrypoint.d/

# Set execute permissions for the script
RUN chmod +x /docker-entrypoint.d/init.sh

# Copy the nginx configuration
COPY ./src/services/nginx/nginx.conf /etc/nginx/nginx.conf
COPY ./src/services/nginx/cors.conf /etc/nginx/cors.conf

# Copy the built frontend code
COPY --from=builder /app/build /var/www/varso.org/html

EXPOSE 443
EXPOSE 80


CMD ["nginx", "-u", "-g", "daemon off;"]