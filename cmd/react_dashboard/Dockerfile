# Start with a base image containing Node.js runtime
FROM node:17.0.0 

# Set the working directory in the builder stage
WORKDIR /app

# Copy package.json file
COPY package*.json ./

# Install dependencies 
RUN npm install

# Create a group and user
#COPY --chown=node:node . .

# Copy the keys directory
COPY keys ./keys

# Copy the rest of your app's source code from your host to your image filesystem.
COPY . .

# Change to the node user
USER node

# Expose the listening port of your app
EXPOSE 3000

# Run the app
CMD ["npm", "run", "start-react"]
