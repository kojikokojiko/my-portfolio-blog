# Base stage
FROM node:18-alpine AS base

WORKDIR /app

# Copy package.json and package-lock.json first to leverage caching
COPY frontend/package.json frontend/package-lock.json ./
RUN npm install

# Copy the rest of the application code
COPY frontend/ ./

# Development stage
FROM base AS development
CMD ["npm", "start"]

# Production stage
FROM node:18-alpine AS build
WORKDIR /app

# Copy package.json and package-lock.json to install dependencies
COPY frontend/package.json frontend/package-lock.json ./
RUN npm install

# Copy the rest of the application code and build the production files
COPY frontend/ ./
RUN npm run build

# Final stage for serving the production build
FROM node:18-alpine AS production

WORKDIR /app

# Install serve to serve the production build
RUN npm install -g serve

# Copy the build output from the build stage
COPY --from=build /app/build ./build

# Serve the application
CMD ["serve", "-s", "build"]
