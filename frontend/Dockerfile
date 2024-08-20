# Use an official Node.js image as the base image
FROM node:18 AS development

# Set the working directory in the container
WORKDIR /app

# Install dependencies first (caching layer)
COPY package*.json ./
RUN npm install

# Copy the rest of the application code
COPY . .

# Expose the port that the app will run on
EXPOSE 3000

# Start the React development server
CMD ["npm", "start"]