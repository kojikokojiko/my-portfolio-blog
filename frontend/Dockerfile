# Frontend Dockerfile
FROM node:18-alpine AS base

WORKDIR /app

COPY frontend/package.json frontend/package-lock.json ./
RUN npm install

COPY frontend/ ./

# Development stage
FROM base AS development
CMD ["npm", "start"]

# Production stage
FROM base AS production
RUN npm run build
CMD ["npx", "serve", "-s", "build"]
