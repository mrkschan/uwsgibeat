FROM debian:jessie

RUN apt-get update && \
    apt-get install -y build-essential curl libssl-dev openssl

    # Setup Python env
    ENV LANG C.UTF-8
    ENV PYTHONUNBUFFERED 1

RUN apt-get install -y python python-dev && \
    curl -sL 'https://bootstrap.pypa.io/get-pip.py' | python -

RUN pip install uwsgi==2.0.11.2
COPY wsgi.py wsgi.py

EXPOSE 8080
CMD ["uwsgi", "--wsgi-file", "wsgi.py", "--http", ":8080"]