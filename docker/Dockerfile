# Use an official Node.js runtime as a parent image
FROM node:20 as builder

# Set the working directory in the container
WORKDIR /usr/src/app

# Copy package.json and package-lock.json (or yarn.lock if you're using Yarn)
COPY package*.json ./

# Install your project dependencies
RUN npm install

# Copy your application code and any other necessary files
COPY ../../../docker .

# Build the application
RUN npm run build

# Start a new stage from scratch for a cleaner, smaller image
FROM node:20

WORKDIR /usr/src/app

# Copy built assets from the builder stage
COPY --from=builder /usr/src/app/node_modules ./node_modules
COPY --from=builder /usr/src/app/dist ./dist

# Expose the port the app runs on
# 8082 - Rest API
# 9090 - Main GRPC
# 9095 - Waifu GRPC
# Expose the ports your app runs on
EXPOSE 8082 9090 9095

# Command to run the app
CMD ["node", "dist/Runner.js"]
