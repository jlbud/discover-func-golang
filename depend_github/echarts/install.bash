#!/bin/bash

cd /Users/kevin/Software

wget https://bitbucket.org/ariya/phantomjs/downloads/phantomjs-2.1.1-windows.zip
unzip phantomjs-2.1.1-windows.zip
rm -f phantomjs-2.1.1-windows.zip

git clone git@github.com:pesla/highcharts-phantomjs.git