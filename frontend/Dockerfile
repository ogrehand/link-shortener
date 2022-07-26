FROM node:lts-alpine3.15

COPY ./shortenerFE /usr/bin/

WORKDIR /usr/bin/shortenerFE

RUN npm install

ENTRYPOINT ["npm", "run", "serve"]