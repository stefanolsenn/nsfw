# nsfw
Yahoo's repo using Caffe:cpu as neural network, to recognize Not Safe For work images. Ported to a web interface

## Installation

1. Clone repo
2. Build the image `docker build -t github.com/stefanolsenn/nsfw .`
3. Start the container `docker run -d -p 8080:8080 github.com/stefanolsenn/nsfw`
4. Create a POST-request against localhost:8080/ with payload:
  `{"path": "URL_TO_IMG"}`
  response will be a json object with `{"Score": string}`

## Notes
There is currently no validation of the url nor the http stream's datatype. If the url isn't pointing to a image, the program wil most likely break. Feel free to submit a pull request :)

It was just a little side-project and isn't tweeked to production, the response time is quite heavy (700-1500ms). The pretrained model could be optimized by using https://github.com/TechnikEmpire/NsfwSqueezenet to decrease the response time.
  
### Credits
Copyright 2016, Yahoo Inc.

Copyright (c) 2014-2017 The Regents of the University of California (Regents)
