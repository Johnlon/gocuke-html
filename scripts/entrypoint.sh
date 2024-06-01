#!/bin/bash
docsify serve ./docs &
./gocureAPI &
apache2ctl -D FOREGROUND
