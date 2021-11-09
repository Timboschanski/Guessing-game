FROM golang 

RUN mkdir -p /home/app

COPY . /home/app

WORKDIR /home/app/src

RUN ls
