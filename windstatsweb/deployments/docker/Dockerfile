# Stage 0, "build-stage", based on Node.js, to build and compile the frontend
FROM --platform=$BUILDPLATFORM node:12.14.0 as build-stage
ARG TARGETPLATFORM
ARG BUILDPLATFORM
RUN echo "I am running on $BUILDPLATFORM, building for $TARGETPLATFORM"
WORKDIR /app
COPY package*.json /app/
RUN npm install
COPY ./ /app/
ARG configuration=production
RUN npm run build -- --output-path=./dist/out --configuration $configuration

# Stage 1, based on Nginx, to have only the compiled app, ready for production with Nginx
FROM --platform=$TARGETPLATFORM nginx:alpine
#Copy ci-dashboard-dist
COPY --from=build-stage /app/dist/out/ /usr/share/nginx/html
#Copy default nginx configuration
COPY ./deployments/docker/nginx-custom.conf /etc/nginx/conf.d/default.conf

EXPOSE 80