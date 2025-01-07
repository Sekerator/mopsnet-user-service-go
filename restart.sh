#!/bin/bash

APP_NAME="user-service"
APP_DIR="/var/www/user-service"

# Останавливаем приложение
pkill -f $APP_NAME || true

# Переходим в директорию приложения
cd $APP_DIR || exit 1

# Запускаем новую версию приложения
nohup ./$APP_NAME > app.log 2>&1 &
echo "$APP_NAME restarted"
