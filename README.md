# nsfw
Yahoo's repo using Caffe:cpu as neural network, to recognize Not Safe For work images. Ported to a web interface

## Installation

1. Clone repo
2. Build the image `docker build -t github.com/stefanolsenn/nsfw .`
3. Start the container `docker run -d -p 8080:8080 github.com/stefanolsenn/nsfw`
4. Create a POST-request against localhost:8080/ with payload:
  `{"path": "URL_TO_IMG"}`
  reponse will be a json object with `{"score": string}`

## Notes
There is currently no validation of the url nor if it is a picture. The program will most likely fail. Pull requests welcome!
  
### Credits
Copyright 2016, Yahoo Inc.
Copyright (c) 2014-2017 The Regents of the University of California (Regents)
