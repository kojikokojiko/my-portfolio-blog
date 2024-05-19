# Base stage
FROM nginx:alpine AS base

# Copy the main nginx configuration file
COPY nginx.conf /etc/nginx/nginx.conf

# Create the sites-enabled directory and ensure it exists
RUN mkdir -p /etc/nginx/sites-enabled

# Development stage
FROM base AS dev
# Copy the development configuration file
COPY sites-available/my_portfolio.dev /etc/nginx/sites-available/my_portfolio

# Create a symbolic link to the sites-enabled directory
RUN ln -s /etc/nginx/sites-available/my_portfolio /etc/nginx/sites-enabled/my_portfolio

# Run nginx in the foreground
CMD ["nginx", "-g", "daemon off;"]

# Production stage
FROM base AS prod
# Copy the production configuration file
COPY sites-available/my_portfolio.prod /etc/nginx/sites-available/my_portfolio

# Create a symbolic link to the sites-enabled directory
RUN ln -s /etc/nginx/sites-available/my_portfolio /etc/nginx/sites-enabled/my_portfolio

# Run nginx in the foreground
CMD ["nginx", "-g", "daemon off;"]
