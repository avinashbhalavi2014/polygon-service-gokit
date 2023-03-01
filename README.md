# polygon-service-gokit

### config file
```
app.yaml
```
### Environment Variable Used
```
ENV: develop
HOST_IP: 127.0.0.1
DB_DRIVER: postgres
DB_CONNECTION_STRING: postgresql://postgres:root@127.0.0.1:5433/postgis_db?sslmode=disable
MIGRATION_URL: file://db/migration
PORT: 8081
SERVICE_NAME: polygon
```

### Backend setup to run strings service

#### BE build
```
go build
```

#### Run Server
```
start polygon-service-gokit
```

### API Details
##### 1) POST API 
```
http://localhost:8081/api/create-polygon
```
##### Request Body Payload
```
 {
       "type": "FeatureCollection",
       "features": [{
           "type": "Feature",
           "geometry": {
               "type": "Polygon",
               "coordinates": [
                   [
                       [0, 0],
                       [0, 1],
                       [1, 0],
                       [1, 1],
                       [0, 0]
                   ]
               ]
           },
           "properties": {
               "name": "polygon1"
           }
       }]
   }
```
##### Response:
```
StatusCode: 201
```

##### 2) GET API 
```
http://localhost:8081/api/get-polygon?name=polygon1
```

##### Response
```
{
    "FeatureCollection": {
        "type": "FeatureCollection",
        "features": [
            {
                "type": "Feature",
                "geometry": {
                    "type": "Polygon",
                    "coordinates": [
                        [
                            [
                                100,
                                0
                            ],
                            [
                                101,
                                0
                            ],
                            [
                                101,
                                1
                            ],
                            [
                                100,
                                1
                            ],
                            [
                                100,
                                0
                            ]
                        ]
                    ]
                },
                "properties": {
                    "name": "polygon1"
                }
            }
        ]
    }
}
```