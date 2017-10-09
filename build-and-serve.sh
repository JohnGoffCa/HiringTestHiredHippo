#!/bin/bash
ng build &&
go build -o charlie ./backend/ &&
./charlie
