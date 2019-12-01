# base image
FROM consul:1.6.2

## add the config file
COPY ./config/consul-config.json /consul/config/config.json

# expose ports
EXPOSE 8300 8400 8500 8600

# run consul
ENTRYPOINT ["consul"]
