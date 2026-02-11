FROM python:3.9

# Bad: Separate layers for apt update and install, no cleanup
RUN apt-get update
RUN apt-get install -y netcat-openbsd

WORKDIR /code

# Bad: COPY everything early
COPY . /code

# Bad: No --no-cache-dir, results in larger image
RUN pip install -r requirements.txt

# Bad: Unnecessary environment variables in separate layers
ENV APP_HOME=/code
ENV DEBUG=True

CMD ["python", "app.py"]
