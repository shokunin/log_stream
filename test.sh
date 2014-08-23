#!/bin/bash

if [ ! -d logstash-1.4.2 ]; then
	wget https://download.elasticsearch.org/logstash/logstash/logstash-1.4.2.tar.gz
	tar -zxf logstash-1.4.2.tar.gz
	rm logstash-1.4.2.tar.gz
fi

touch /tmp/test.log

echo "Things are working"

