#!/bin/sh

certdir="/etc/letsencrypt/live"
realcert="${certdir}/.realcert"
fakecert="${certdir}/.fakecert"

# Only create dummy certificate if no cert exists
if [ ! -f "$path/fullchain.pem" ]; then
    for domain in ${CERTBOT_DOMAINS}; do
        path="/etc/letsencrypt/live/${domain}"
        mkdir -p "$path"
        openssl req -x509 -nodes -newkey rsa:4096 -days 1 -keyout "$path/privkey.pem" -out "$path/fullchain.pem" -subj "/CN=localhost"
    done

    mkdir -p "${certdir}"
    touch "$fakecert"
fi

# If we are not in DEV mode
if [ "$APP_ENV" != "DEV" ]; then
    # Wait for certbot to create a real certificate
    (
        while [ ! -f "$realcert" ]; do
            echo "Waiting for certbot to create .realcert.."
            sleep 5  # Adjust the sleep duration if needed
        done

        nginx -s reload
    ) &
fi
