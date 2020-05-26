while ! (echo "SHOW DATABASES" | mysql -uroot -proot -hmysql-server); do
    echo "Waiting for mysql-server..."
    sleep 1
done

# Start daemon
ls -ahltr /var/app
/var/app/app
