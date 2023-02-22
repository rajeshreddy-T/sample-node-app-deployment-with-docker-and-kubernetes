# node alpine image
From node:16-alpine3.16
# // we specify the image name that we need to our project

WORKDIR /usr/src/app
# // setting the path of the working directory
# // this is the directory where our application will be running
# // this is the directory where we will be copying all the files
# // this is the directory where we will be installing all the dependencies


COPY package*.json ./
# // copying both package.json & package-lock.json

RUN npm install
# // installing all the dependencies

COPY . .
# // copying all the rest of the files inside our project

EXPOSE 3000
# // exposing the port which our application runs on

CMD ["node", "app.js" ]
# CMD [ "ping", "google.com" ]

# // here, you can enter the command which we use to run our application
# // in our case, we are using node app.js
# // this is the command which will be executed when we run our container
# // this is the command which will be executed when we run our image as a container 


