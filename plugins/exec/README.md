# exec

The *exec* plugin will check if a code block has a `exec:CMD` language tag. If found the
the following steps will be performed:

1. `CMD` will be executed with the contents of the code block piped to it's standard input.

1. The ouput from `CMD` (if any) will be used to construct a data URI.

1. The code block will then be deleted and *replaced* with an image containing the data URI.

The image outputted used *must* be a png. The title of the image will be empty: `![](...)`, as there
is no text to put in there.

If the original code block has a block level attribute the figure will be wrapped in a subfigure so
you can reference it from within the document.

## Caption

The caption will be retained and the image will be wrapped in a subfigure.

## Error handling

In case of an error, the code block is left as-is. There is no option handling for `CMD`, so if a
command needs those you need to construct a shell script that can be used as `CMD` instead.
