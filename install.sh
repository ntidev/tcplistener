#!/bin/bash

cp ./tcplistener /usr/bin/tcplistener
cp ./tcplistener.service /etc/systemd/system/tcplistener.service

systemctl daemon-reload
