FROM node as node

WORKDIR /app

COPY . .

RUN npm install

RUN npm run build --prod

FROM nginx:alpine

EXPOSE 4200

COPY --from=node /app/dist/my-app /usr/share/nginx/html

CMD ["nginx", "-g", "daemon off;"]
