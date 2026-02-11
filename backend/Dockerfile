FROM openjdk:11

# Bad: No multi-stage build, remains in a large JDK image
WORKDIR /app

# Bad: COPY all source code
COPY . .

# Bad: Running build inside the final image
RUN ./gradlew build

# Bad: Running with standard jar (heavy environment)
EXPOSE 8080
CMD ["java", "-jar", "build/libs/app.jar"]
