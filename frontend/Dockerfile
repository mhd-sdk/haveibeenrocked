FROM oven/bun:latest AS base

WORKDIR /app

COPY package*.json ./

# Install dependencies
RUN bun install

COPY . .

RUN bun run build

# Use Nginx as the production image 
FROM nginx

COPY --from=base /app/dist /usr/share/nginx/html

# Just a documentation of wich port the app runs inside the container
EXPOSE 80

# Start Nginx when the container runs
CMD ["nginx", "-g", "daemon off;"]
