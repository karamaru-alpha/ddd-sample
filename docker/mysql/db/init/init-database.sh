#!/bin/bash

mysql -u root -ppassword ddd_sample < "/docker-entrypoint-initdb.d/001-ddl.sql"
