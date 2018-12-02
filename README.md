# nsfw
Yahoo's repo using Caffe:cpu as neural network, to recognize Not Safe For work images. Ported to a web interface

## Installation

1. Clone https://github.com/yahoo/open_nsfw
2. Download this repo into the same folder
3. Build the image `docker build -t github.com/stefanolsenn/nsfw .`
4. Start the container `docker run -d -p 8080:8080 github.com/stefanolsenn/nsfw`
5. Create a POST-request against localhost:8080/ with payload:
  `{"path": "URL_TO_IMG"}`
