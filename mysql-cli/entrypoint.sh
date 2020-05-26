if [ "$*" = init ]; then
    mysql -uroot -proot -hmysql-server < fixtures.sql
elif [ "$*" = "" ]; then
    exit
else
    echo "$*"
    echo "$*" | mysql -uroot -proot -hmysql-server -Dmydb
fi
