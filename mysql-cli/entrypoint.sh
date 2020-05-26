if ! (echo "SHOW DATABASES" | mysql -uroot -proot -hmysql-server) > /dev/null; then
    echo "Database not yet initialized."
    exit 1
fi

if [ "$*" = init ]; then
    if mysql -uroot -proot -hmysql-server < fixtures.sql; then
        echo "Initialized mydb.mytable."
    else
        echo "Error initializing mydb.mytable."
        exit 1
    fi
elif [ "$*" = "" ]; then
    exit
else
    echo "$*" | mysql -uroot -proot -hmysql-server -Dmydb
fi
