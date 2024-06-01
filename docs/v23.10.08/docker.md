# Docker

If you do not install GO or if you want to use Gocure but your projects it's not a GO project, you can use the Gocure docker image.

### Using ```docker run``` command:

```shell
docker run -d \
-v "$PWD/data:/data" \
-v "/Users:/Users" \
-p "8087:80" \
-p "7087:7087" \
--name gocure \
rodrigoodhin/gocure
```

### Using ```docker-compose``` file:

Create a file called ```docker-compose.yml```.
```yml
version: '3'
services:
  gocure:
    container_name: gocure
    image: rodrigoodhin/gocure
    volumes:
      - "$PWD/data:/data"
      - "/Users:/Users"
    ports:
      - "8087:80"
      - "7087:7087"
    networks:
      - my_network
networks:
  my_network:
```

Then execute the command:
```shell
docker-compose up -d
```

after start your container, we can go to:

- Gocure Website at [http://localhost:8087/](http://localhost:8087/). 
  
  See more info at [Website](/v23.07.24/website) section.
  
![](/_media/scr_website_01.png)

- Gocure REST API at [http://localhost:7087/](http://localhost:7087/). 
  
  See more info at [REST API](/v23.07.24/restAPI) section.
  
![](/_media/scr_api.png)


- Gocure REST API Swagger documentation at [http://localhost:7087/swagger/index.html](http://localhost:7087/swagger/index.html). 
  
  See more info at [REST API](/v23.07.24/restAPI) section.
  
![](/_media/scr_api_swagger.png)
