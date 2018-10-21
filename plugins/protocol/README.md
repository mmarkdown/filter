# protocol

The *protocol* plugin will check if a code block has a `protocol` language tag. If found the
the following steps will be performed:

1. `protocol` will be executed with the contents of the code block piped to it's standard input.

1. The code block's contents will then be deleted and *replaced* with the output from `protocol`.

## Error handling

In case of an error, the code block is left as-is.

## See Also

<http://www.luismg.com/protocol/>, and the *exec* plugin.
