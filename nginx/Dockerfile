FROM nginx:alpine AS base

# Copy the main nginx configuration file
COPY nginx.conf /etc/nginx/nginx.conf

# Create the sites-enabled directory
RUN mkdir -p /etc/nginx/sites-enabled


FROM base AS dev
COPY sites-available/my_portfolio.dev /etc/nginx/sites-available/my_portfolio

# Create a symbolic link
RUN ln -s /etc/nginx/sites-available/my_portfolio /etc/nginx/sites-enabled/my_portfolio

CMD ["nginx", "-g", "daemon off;"]


FROM base AS prod
COPY sites-available/my_portfolio.prod /etc/nginx/sites-available/my_portfolio

# Create a symbolic link
RUN ln -s /etc/nginx/sites-available/my_portfolio /etc/nginx/sites-enabled/my_portfolio

CMD ["nginx", "-g", "daemon off;"]
