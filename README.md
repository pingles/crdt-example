# CRDT Example

Example application for playing with Convergent Replicated Data Types.

## Running

    $ ./crdt-example -name=node1
    $ ./crdt-example -name=node2 -listen=7947 -join=127.0.0.1:7946
    
    2015/04/13 11:34:02 [DEBUG] memberlist: Initiating push/pull sync with: 127.0.0.1:7946
    2015/04/13 11:34:02 listening 0.0.0.0:7947
    2015/04/13 11:34:02 member node1: 192.168.138.100:7946
    2015/04/13 11:34:02 member node2: 192.168.138.100:7947