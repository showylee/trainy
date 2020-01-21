FROM node:12.14.1-alpine

# locale & timezone
ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

RUN apk update && \
    npm install -g npm && \
    yarn global add @vue/cli @vue/cli-init


ADD src /src
WORKDIR /src/app

ENV HOST 0.0.0.0
EXPOSE 3000
#CMD ["/bin/ash"]

RUN yarn install
CMD ["yarn", "run", "dev"]
