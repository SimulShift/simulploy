import psycopg2
from psycopg2 import OperationalError

def create_connection(db_name, db_user, db_password, db_host, db_port):
    connection = None
    try:
        connection = psycopg2.connect(
            database=db_name,
            user=db_user,
            password=db_password,
            host=db_host,
            port=db_port,
        )
        print("Connection to PostgreSQL DB successful")
    except OperationalError as e:
        print(f"The error '{e}' occurred")
    return connection

# Replace the arguments with your server's details
connection = create_connection("postgres", "postgres", "SuperPasswOrd1337420!", "146.190.199.222", "5432")

#postgresql://postgres:SuperPasswOrd1337420!@146.190.199.222:5432/simulshift.com