# protocol

The *protocol* plugin will check if a code block has a `protocol` language tag. If found the the
following steps will be performed:

1. `protocol` will be executed with the contents of the code block given to it on the command line.

1. The code block's contents will then be deleted and *replaced* with the output from `protocol`.

## Error handling

In case of an error, the code block is left as-is.

## Example

This markdown file:

~~~
This is a protocol

``` protocol
Source:16,TTL:8,Reserved:40
```
Figure: This is a protocol.
~~~

Will be transformed with `filter -p protocol < protocol.md | mmark -markdown`, to:

~~~
We describe the following protocols:

```
 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|             Source            |      TTL      |               |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+               +
|                            Reserved                           |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
```
Figure: This is a protocol.
~~~

## See Also

[http://www.luismg.com/protocol/](http://www.luismg.com/protocol/), and the *exec* plugin.
