#!/bin/bash
ng build --env=prod &&
go build -o charlie ./backend/ &&
./charlie
