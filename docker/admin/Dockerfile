FROM ruby:3.4.10

WORKDIR /app

# Using Node.js v22.x(LTS)
RUN curl -fsSL https://deb.nodesource.com/setup_22.x | bash -

# Add packages
RUN apt-get update && apt-get install -y \
      git \
      nodejs \
      vim

ENV DISABLE_BOOTSNAP 1

# Add yarnpkg for assets:precompile
RUN npm install -g yarn
