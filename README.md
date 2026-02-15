# golang-ovh-agent

The first main goal is to list all the *Block Storage* resources from my OVH account and identify orphans ones. (volumes)

## Prerequesites 

- [x] You to need to get golang ready on your system.
- [x] Create the API token creds : [API creds](https://api.ovh.com/createToken/)

## External services

I'll use the official ovh golang wrapper for OVH API : [go-ovh](https://github.com/ovh/go-ovh)

Installing it using : 
```golang
go get github.com/ovh/go-ovh/ovh
```