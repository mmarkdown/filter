---
title: 'MMARK-FILTER(1)'
author:
    - Mmark Authors
date: October 2018
---

# NAME

mmark-filter – a Unix filer that changes markdown documents.

# SYNOPSIS

**mmark-filter** [**OPTIONS**] [*FILE...*]

# DESCRIPTION

**Mmark-filter** is a powerful markdown filter. It allows you to rewrite parts of the document and
 output markdown again.

It's powered by plugins that can be loaded with the `-p` flag. Multiple plugins may be chained.
They will be executed in the order specified.

## PLUGINS

The following plugins are available; each of them has their own README.md

* emph 
* exec  
* noop  
* protocol 
* rot13

# OPTIONS

**-l**

:  list all available plugins.

**-p string,string**

:  the plugins to load, separated by commas. Note spaces are now allowed.

**-version**

:  show mmark-filters's version.
