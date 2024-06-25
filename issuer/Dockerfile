# Note: must call docker build in root directory

# Use a Node.js base image
FROM node:20-slim AS base
RUN corepack enable

# Update the package lists and install Python3 and pip
RUN apt-get update && apt-get install -y python3 python3-pip && ln -s /usr/bin/python3 /usr/bin/python && rm -rf /var/lib/apt/lists/*

FROM base AS build

# Copy root directory
COPY . /workspace
WORKDIR /workspace
RUN ls /workspace
RUN yarn install --frozen-lockfile
RUN yarn build

FROM base AS issuer
# node_modules are symlinked to root directory, so we just copy the whole thing for now
COPY --from=build /workspace /workspace
# COPY --from=build /prod/issuer /prod/issuer
WORKDIR /workspace/apps/issuer
EXPOSE 8080 9090
CMD [ "yarn", "start" ]
