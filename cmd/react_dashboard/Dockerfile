# Start with a base image containing Node.js runtime
FROM node:20.10.0

# Set the working directory in the builder stage
WORKDIR /app

# Copy package.json file
COPY package*.json ./

# Install dependencies 
RUN npm install

# Copy the keys directory
COPY keys ./keys

# Copy the rest of your app's source code from your host to your image filesystem.
COPY . .

# Expose the listening port of your app
EXPOSE 3000

# Run the app
CMD ["npm", "run", "start"]
