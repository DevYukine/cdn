# cdn

serve & manage static hosted assets with ease!

## Getting Started

### Run it

1. Clone the repo `git clone https://github.com/DevYukine/cdn`
2. Run it `go run .`

### Docker
You can get the latest docker image with `docker pull devyukine/cdn`


### Routes
  - `GET` `/` Serve all the assets
  - `POST` `/assets` Upload assets
  - `PATCH` `/assets/{name}` Replace assets
  - `DELETE` `/assets/{name}` Delete assets

### Post/Patch:
You need to send multipart/form-data with a key called `file` & a file as value

### Responses
The api will **always** respond with the application/json format

### Prerequisites

```
Go 1.14+
```

## Deployment

Environment Variables:
  - `CONTENT_ROOT` The folder to save assets to
  - `PORT` The port to run the server on
  - `TOKEN` **Optional** A token which is required to add/replace/delete assets

## Built With

* [mux](https://github.com/gorilla/mux) - The http router used

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/DevYukine/cdn/tags). 

## Authors

* **DevYukine** - *Initial work* - [DevYukine](https://github.com/DevYukine)

See also the list of [contributors](https://github.com/DevYukine/cdn/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details