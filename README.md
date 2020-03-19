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

### Post/Patch File Upload
You need to send the file in `multipart/form-data` as the `file` key

### Responses
The api will **always** respond with the application/json format

### Filenames
The filename is defaulted to the name of the uploaded asset. If you want to generate a filename you can set the `generate_name` key to `true`, the cdn will respond with an object containing a `file_name` value

### Prerequisites

```
Go 1.14+
```

## Deployment

Environment Variables:
  - `CONTENT_ROOT` The folder to save assets to
  - `PORT` The port to run the server on
  - `TOKEN` **Optional** A token which is required to add/replace/delete assets
  - `REVERSE_PROXY` **Optional** If this instance runs behind a reverse proxy (e.g. nginx, HAProxy or apache) and should use forward header

## Built With

* [mux](https://github.com/gorilla/mux) - The http router used

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/DevYukine/cdn/tags). 

## Authors

* **DevYukine** - *Initial work* - [DevYukine](https://github.com/DevYukine)

See also the list of [contributors](https://github.com/DevYukine/cdn/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details