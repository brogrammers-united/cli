docker build . -t cli:testing
docker run -it --rm --name cli cli:testing 