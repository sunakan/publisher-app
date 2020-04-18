ARG DOCKER_NODE_TAG
FROM node:${DOCKER_NODE_TAG} as base
ARG TINI_VERSION
ENV TINI_VERSION ${TINI_VERSION}
RUN apt update \
  && apt install --yes --quiet \
    git
ADD https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini \
  /usr/local/bin/tini
RUN chmod +x /usr/local/bin/tini
EXPOSE 3000
WORKDIR /home/node/app
RUN chown node:node --recursive /home/node/app
USER node
#COPY --chown=node:node package.json package-lock.json ./
#RUN npm ci \
#  && npm cache clean --force

################################################################################
# Development
################################################################################
FROM base as development
ENV NODE_ENV development
ENV PATH=/home/node/app/node_modules/.bin:$PATH
CMD ["npm", "run", "dev"]

################################################################################
# Production
################################################################################
