FROM node:lts-alpine3.15

COPY ./shortener-fe /usr/bin/

WORKDIR /usr/bin/shortener-fe

RUN npm install

ENTRYPOINT ["npm", "run", "serve"]