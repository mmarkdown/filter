# Mmark Filter

mmark-filter is a markdown filter. It takes in a markdown file (possibly on standard input) and
translates the input according to a bunch of filters (implemented as plugins).

## Plugins

Currently these plugins exist:

*  noop, a plugin that does nothing.

*  emph, a plugin that replaces `*emphasis*` with `XXemphasisXX`, added mostly as an example.

*  exec, a plugin the runs a command using the codeblock contents and replaces it with an image.

*  rot13, a plugin that rot13's all text, but leaves all other elements alone.

*  protocol, a plugin that runs [protocol](http://www.luismg.com/protocol/) on the contents of a
   codeblock and replaces the codeblock's content with its output. Comes in handy when writting IETF
   drafts.
