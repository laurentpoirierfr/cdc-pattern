Use Debezium to send Oracle database data to PostgreSQL database through CDC.


# Preparations
1. Sign in https://container-registry.oracle.com/ , then click Database > Enterprise > Tags, select the database version to use.
2. pull oracle database image
```shell
docker login container-registry.oracle.com
```
after successful login, execute the following command
```shell
docker pull container-registry.oracle.com/database/enterprise:12.2.0.1
```
it may take a few minutes.

3. delete all files under the **oracle>ORCLCDB** folder of this project
4. create empty logminer file
```shell
cd oracle
mkdir recovery_area
cd ORCLCDB
touch logminer_tbs.dbf
mkdir ORCLPDB1
cd ORCLPDB1
touch logminer_tbs.dbf
```
5. modify the IP address of KAFKA_CFG_ADVERTISED_LISTENERS=PLAINTEXT://192.168.0.105 in the docker-compose.yml file to the IP of your machine (ifconfig)

# Step
1.0 docker-compose up -d  

1.1 **Please ensure that all containers can start normally and there are no errors before proceeding to the next steps**  
![](https://github.com/jacksparrow414/oracle-debezium-postgres/blob/main/Xnip2023-02-27_19-21-22.jpg)
2.0 configure the Oracle database,Refs [Debezium Oracle Connetor Documentation Start with preparing_the_database](https://debezium.io/documentation/reference/2.1/connectors/oracle.html#_preparing_the_database)  

2.1 [resize Oracle Database redo log](https://logic.edchen.org/how-to-resize-redo-logs-in-oracle/)  

2.1.0 Query the redo log file location  
```code
SQL> select * from v$logfile ;
```

2.1.1 Check Current Redo Logs  
```code
SQL> column group# format 99999;
SQL> column status format a10;
SQL> column mb format 99999;
SQL> select group#, status, bytes/1024/1024 mb from v$log;

GROUP# STATUS	      MB
------ ---------- ------
     1 ACTIVE	       4
     2 ACTIVE	       4
     3 CURRENT	     4
```
2.1.2 Add 3 Groups with New Size (1GB)  
```code
SQL> alter database add logfile group 4 ('/u04/app/oracle/redo/redo004.log') size 1g, group 5 ('/u04/app/oracle/redo/redo005.log') size 1g, group 6 ('/u04/app/oracle/redo/redo006.log') size 1g;

Database altered.

SQL> select group#, status, bytes/1024/1024 mb from v$log;

GROUP# STATUS	      MB
------ ---------- ------
     1 INACTIVE        4
     2 INACTIVE        4
     3 CURRENT	       4
     4 UNUSED	    1024
     5 UNUSED	    1024
     6 UNUSED	    1024

6 rows selected
```
2.1.3 Switch Logfile to New Groups  
```code
SQL> alter system switch logfile;

System altered.

SQL> /

System altered.

SQL> select group#, status, bytes/1024/1024 mb from v$log;

GROUP# STATUS	      MB
------ ---------- ------
     1 INACTIVE        4
     2 INACTIVE        4
     3 ACTIVE	         4
     4 ACTIVE	      1024
     5 CURRENT	    1024
     6 UNUSED	      1024

6 rows selected.
```
2.1.4 Force a CheckPoint  
```code
SQL> alter system checkpoint;

System altered.

SQL> select group#, status, bytes/1024/1024 mb from v$log;

GROUP# STATUS	      MB
------ ---------- ------
     1 INACTIVE        4
     2 INACTIVE        4
     3 INACTIVE        4
     4 INACTIVE     1024
     5 CURRENT	    1024
     6 UNUSED	      1024

6 rows selected.
```
2.1.5 Drop Group 1, 2, 3  
```code
SQL> alter database drop logfile group 1, group 2, group 3;

Database altered.

SQL> select group#, status, bytes/1024/1024 mb from v$log;

GROUP# STATUS	      MB
------ ---------- ------
     4 INACTIVE     1024
     5 CURRENT	    1024
     6 UNUSED	      1024
```
2.1.6 Remove Redo Log Files  
```code
[oracle@test ~]$ rm -i /u04/app/oracle/redo/redo00[1-3].log
rm: remove regular file `/u04/app/oracle/redo/redo001.log'? y
rm: remove regular file `/u04/app/oracle/redo/redo002.log'? y
rm: remove regular file `/u04/app/oracle/redo/redo003.log'? y
```
2.1.7 Add Group 1, 2, 3 with New Size (1GB)  
```code
SQL> alter database add logfile group 1 ('/u04/app/oracle/redo/redo001.log') size 1g, group 2 ('/u04/app/oracle/redo/redo002.log') size 1g, group 3 ('/u04/app/oracle/redo/redo003.log') size 1g;

Database altered.

SQL> select group#, status, bytes/1024/1024 mb from v$log;

GROUP# STATUS	      MB
------ ---------- ------
     1 UNUSED	    1024
     2 UNUSED	    1024
     3 UNUSED	    1024
     4 INACTIVE   1024
     5 CURRENT	  1024
     6 UNUSED	    1024

6 rows selected.
```
2.1.8 Switch Logfile Several Times  
```code
SQL> alter system switch logfile;

System altered.

SQL> /

System altered.

SQL> /

System altered.

SQL> /

System altered.

SQL> /

System altered.

```
2.1.9 Check Status of All Redo Logs  
```code
SQL> select group#, status, bytes/1024/1024 mb from v$log;

GROUP# STATUS	      MB
------ ---------- ------
     1 ACTIVE	    1024
     2 ACTIVE	    1024
     3 ACTIVE	    1024
     4 CURRENT	  1024
     5 ACTIVE	    1024
     6 ACTIVE	    1024

6 rows selected.

SQL> column member format a40;
SQL> select group#, member from v$logfile;

GROUP# MEMBER
------ ----------------------------------------
     1 /u04/app/oracle/redo/redo001.log
     2 /u04/app/oracle/redo/redo002.log
     3 /u04/app/oracle/redo/redo003.log
     4 /u04/app/oracle/redo/redo004.log
     5 /u04/app/oracle/redo/redo005.log
     6 /u04/app/oracle/redo/redo006.log

6 rows selected.
```
2.1.10
```code
SQL> exit;
```
2.2.0 [Creating the connector’s LogMiner user](https://debezium.io/documentation/reference/2.1/connectors/oracle.html#creating-users-for-the-connector)  
2.2.1 create logminer tablespace
```code
[oracle@2b18e16ba29b /]$ sqlplus sys/12345@//localhost:1521/ORCLCDB as sysdba

SQL*Plus: Release 12.2.0.1.0 Production on Mon Feb 27 11:44:28 2023

Copyright (c) 1982, 2016, Oracle.  All rights reserved.

ERROR:
ORA-12514: TNS:listener does not currently know of service requested in connect
descriptor


Enter user-name: sys as sysdba
Enter password:

Connected to:
Oracle Database 12c Enterprise Edition Release 12.2.0.1.0 - 64bit Production
```
```code
SQL> CREATE TABLESPACE logminer_tbs DATAFILE '/opt/oracle/oradata/ORCLCDB/logminer_tbs.dbf' SIZE 25M REUSE AUTOEXTEND ON MAXSIZE UNLIMITED;

Tablespace created.

SQL> exit;
```
```code
[oracle@2b18e16ba29b /]$ sqlplus sys/12345@//localhost:1521/ORCLCDBPDB1 as sysdba

SQL*Plus: Release 12.2.0.1.0 Production on Mon Feb 27 11:46:27 2023

Copyright (c) 1982, 2016, Oracle.  All rights reserved.

ERROR:
ORA-12514: TNS:listener does not currently know of service requested in connect
descriptor


Enter user-name: sys as sysdba
Enter password:

Connected to:
Oracle Database 12c Enterprise Edition Release 12.2.0.1.0 - 64bit Production

SQL> ALTER SESSION SET CONTAINER = ORCLPDB1;

Session altered.

SQL> SHOW CON_NAME

CON_NAME
------------------------------
ORCLPDB1
SQL> CREATE TABLESPACE logminer_tbs DATAFILE '/opt/oracle/oradata/ORCLCDB/ORCLPDB1/logminer_tbs.dbf' SIZE 25M REUSE AUTOEXTEND ON MAXSIZE UNLIMITED;

Tablespace created.

SQL> exit;
```
2.2.2 create user
```code
[oracle@2b18e16ba29b /]$ sqlplus sys/12345@//localhost:1521/ORCLCDB as sysdba

SQL*Plus: Release 12.2.0.1.0 Production on Mon Feb 27 11:49:00 2023

Copyright (c) 1982, 2016, Oracle.  All rights reserved.

ERROR:
ORA-12514: TNS:listener does not currently know of service requested in connect
descriptor


Enter user-name: sys as sysdba
Enter password:

Connected to:
Oracle Database 12c Enterprise Edition Release 12.2.0.1.0 - 64bit Production

SQL> alter session set "_ORACLE_SCRIPT"=true;

Session altered.

SQL> CREATE USER c##dbzuser IDENTIFIED BY dbz DEFAULT TABLESPACE logminer_tbs QUOTA UNLIMITED ON logminer_tbs CONTAINER=ALL;

User created

...
Subsequent SQL omitted. please refer to the Debezium documentation
```
2.2.3 After creating the user, verify whether you can log in, you can use sqlDeveloper or Navicat  

2.2.4 After successful login, execute **init-oracle-contact-table.sql** in C##DBZUSER schema  

2.2.5 enable supplemental logging for captured tables
```sql
ALTER TABLE CONTACT ADD SUPPLEMENTAL LOG DATA (ALL) COLUMNS;
```

2.2.6 create a schema named **postgres** in PostgreSQL Database

3.0 create souce connector in kafka-connect  
```shell
curl -i -X POST -H "Accept:application/json" -H  "Content-Type:application/json" http://localhost:8083/connectors/ -d @oracle-source.json
```

3.1 
```shell
curl -i -X POST -H "Accept:application/json" -H  "Content-Type:application/json" http://localhost:8083/connectors/ -d @postgre-sink.json
```

# Verification
1.0 check if initial data is  synchronized to PostgreSQL  
1.1 insert,update,delete a row in Oracle Database, then check if the data in the postgreSQL database has changed

# Delete Connector
```shell
curl -X DELETE http://localhost:8083/connectors/jdbc-sink-customers-postgress
```
# Troubleshooting
- Data changes not synced to PostgreSQL?
```shell
kafka-console-consumer.sh --from-beginning --bootstrap-server kafka1:9092 --topic server1.C__DBZUSER.CONTACT
```
- ORA-04036: PGA memory used by the instance exceeds PGA_AGGREGATE_LIMIT
```code
SQL> show parameter pga_aggregate_limit

NAME TYPE VALUE
------------------------------------ ----------- ------------------------------
pga_aggregate_limit big integer 1G
SQL>
```
increase pga_aggregate_limit
```code
alter system set pga_aggregate_limit=16G scope=both sid='*';
```
