# Filter

Filter is a markdown filter. It takes in a markdown file (possibly on standard input) and translate
the input according to a bunch of filters (implemented as plugins).

## Plugins

Currently there are three plugins:

* noop, a plugin that does nothing.
* emph, a plugin that replaces `*emphasis*` with `XXemphasisXX`.
* exec, a plugin the runs a command using the codeblock contents and replaces it with an image.
