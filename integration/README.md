# Test Script For goval-dictionary
Documentation on testing for developers

## Getting Started
```terminal
$ pip install -r requirements.txt
```

## Run test
Use `127.0.0.1:1325` and `127.0.0.1:1326` to diff the server mode between the latest tag and your working branch.

If you have prepared the two addresses yourself, you can use the following Python script.
```terminal
$ python diff_server_mode.py debian --help
usage: diff_server_mode.py [-h] [--list_path LIST_PATH] [--debug | --no-debug] {cveid,package} {alpine,amazon,debian,oracle,redhat,suse,ubuntu}

positional arguments:
  {cveid,package}       Specify the mode to test.
  {alpine,amazon,debian,oracle,redhat,suse,ubuntu}
                        Specify the OS to be started in server mode when testing.

optional arguments:
  -h, --help            show this help message and exit
  --list_path LIST_PATH
                        A file path containing a line by line list of CVE-IDs or Packages to be diffed in server mode results
  --debug, --no-debug   print debug message
```

[GNUmakefile](../GNUmakefile) has some tasks for testing.  
Please run it in the top directory of the goval-dictionary repository.

**NOTE: Tests for RedHat are commented out by default because fetch takes a long time. Tests for Microsoft are commented out by default because they require API KEY. Please uncomment them if necessary.**

- build-integration: create the goval-dictionary binaries needed for testing
- clean-integration: delete the goval-dictionary process, binary, and docker container used in the test
- fetch-rdb: fetch data for RDB for testing
- fetch-redis: fetch data for Redis for testing
- diff-cveid: Run tests for CVE ID in server mode
- diff-package: Run tests for Package in server mode
- diff-server-rdb: take the result difference of server mode using RDB
- diff-server-redis: take the result difference of server mode using Redis
- diff-server-rdb-redis: take the difference in server mode results between RDB and Redis

## About the CVE ID and Packages list used for testing
Duplicates are removed from the latest fetched data and prepared.  
For example, for sqlite3, you can get it as follows.  
**NOTE: If there are blank lines, the test will fail, so please remove them from the list.**
```terminal
$ sqlite3 oval.sqlite3
SQLite version 3.31.1 2020-01-27 19:55:54
Enter ".help" for usage hints.
# CVE ID
sqlite> .output integration/cveid_redhat.txt
sqlite> SELECT DISTINCT cve_id FROM 'cves' ORDER BY cve_id ASC;

# Packages

```