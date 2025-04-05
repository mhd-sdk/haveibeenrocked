FROM oven/bun:latest AS base

WORKDIR /app

COPY package*.json ./

# Install dependencies
RUN bun install

COPY . .

RUN bun run build

# Use Nginx as the production server
FROM nginx

COPY --from=base /app/dist /usr/share/nginx/html

# Expose port 80 for the Nginx server
EXPOSE 80

# Start Nginx when the container runs
CMD ["nginx", "-g", "daemon off;"]
