FROM node:latest

# Bad: Running update and install in separate layers
RUN apt-get update
RUN apt-get install -y vim

# Bad: COPY . . before npm install breaks caching
WORKDIR /app
COPY . .

# Bad: Installing dependencies (heavy layer)
RUN npm install

# Bad: Multiple RUN commands
RUN echo "Building app..."
RUN npm run build

EXPOSE 3000
CMD ["npm", "start"]
