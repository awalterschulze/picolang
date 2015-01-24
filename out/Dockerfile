#  Copyright 2015 Walter Schulze
#
#  Licensed under the Apache License, Version 2.0 (the "License");
#  you may not use this file except in compliance with the License.
#  You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
#  Unless required by applicable law or agreed to in writing, software
#  distributed under the License is distributed on an "AS IS" BASIS,
#  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#  See the License for the specific language governing permissions and
#  limitations under the License.

FROM ubuntu

RUN sed 's/main$/main universe/' -i /etc/apt/sources.list
RUN apt-get update

RUN apt-get install -y python-software-properties git wget mercurial build-essential protobuf-compiler graphviz net-tools

ENV GOVERSION 1.3
ENV GODOWNLOAD http://golang.org/dl/
ENV GOFILENAME go$GOVERSION.linux-amd64.tar.gz
RUN wget $GODOWNLOAD/$GOFILENAME && \
	tar -C / -xzf $GOFILENAME && \
	rm $GOFILENAME

ENV GOROOT /go
ENV GOPATH /gopath
ENV PATH $PATH:$GOPATH/bin:$GOROOT/bin
RUN mkdir $GOPATH

ENV PICOPATH /gopath/src/github.com/awalterschulze/picolang

RUN mkdir -p $PICOPATH
RUN git clone https://github.com/awalterschulze/picolang $PICOPATH && \
	(cd $PICOPATH && git checkout 2a2580c1e982d6982db03f8901a90236a6d9f2f9)

ENV SERVICEPATH /gopath/src/picoservice
RUN mkdir -p $SERVICEPATH
ADD picoservice.go $SERVICEPATH/
RUN (cd $SERVICEPATH && go install ./...)

EXPOSE 8080
USER daemon
CMD ["picoservice"]
