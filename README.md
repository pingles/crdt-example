# CRDT Example

Example application for playing with Convergent Replicated Data Types.

## Running

    $ ./crdt-example -name=node1
    $ ./crdt-example -name=node2 -listen=7947 -join=127.0.0.1:7946 -http=:8081
    
    2015/04/14 15:00:28 listening 0.0.0.0:7946
    2015/04/14 15:00:28 member node1: 192.168.138.132:7946
    2015/04/14 15:00:28 http listening :8080
    2015/04/14 15:00:45 [DEBUG] memberlist: Responding to push/pull sync with: [::1]:53318
    2015/04/14 15:00:45 [DEBUG] memberlist: Responding to push/pull sync with: 127.0.0.1:53319
    2015/04/14 15:00:50 [DEBUG] memberlist: Initiating push/pull sync with: 192.168.138.132:7945
    
    $ curl http://localhost:8080
    0
    $ curl http://localhost:8080/inc
    1

    # some time passes :) ...

    $ curl http://localhost:8081
    1