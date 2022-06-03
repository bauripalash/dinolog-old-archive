# Dinolog Protocol

### Utterly simplified and lightweight blogging protocol

> Dinolog protocol is a simplified blogging protocol which can even be surfed with telnet or netcat


## Specification
### Request:
    
<Me> : I want to read blogs of Mango Man
<MangoManServer> : Sure, Which one?
<Me> : Not sure! Show the list of every entry. `DD~mangoman~ALL`
<MangoManServer> : Okay, Here you go...
<Me> : Too much.. Too much.. Show me the last 2 entries `DD~mangoman~L2`
<MangoManServer> : As you wish...

```
DD<SEPARATOR><UNAME><SEPARATOR>C<COMMAND>
```

Here is the list of available `COMMAND`s :
    * ALL (All entries)
    * L<1-20> (Latest N entries)
    * O<1-20> (Oldest N entries)
    * D<ISO8601 Date> (entries published on DATE)
    * T<MAX> (Entry with TAG )

I know the URL/SLUG/ID of the entry:

`DD<SEPARATOR><UNAME><SEPARATOR>X<SEPARATOR><ID>`

`DD~mangoman~X~I-eat-mango`

### Response Header:

```
D<SEPARATOR><STATUS_NUMBER><SEPARATOR><NUMBER_OF_POSTS_TO_BE_SHOWN><CRLF>
```

### Response Body
UTF-8 Encoded Plain Text

**NOTE**: Entry listing must make entry id/slug visible with clear indication; for easy browsing via any client

* Server must truncate the post body if it exceeds 500 characters, while doing so server must put clear indication.

* entries can have title fields, which can be used to generate slugs/ids. Title is just like any other normal text on entries. If entry have Title field server will send them as the first line of the entry body text. Clients should make them distinguishable

Here is an example listing

```
MangoMan's Blog
================

> my-first-post
I like to play football

> i-played-football-for-first-time
I won playing football

```


* Each `Log` in a Dinolog server should have a `Name` and an `ID Name`. While `Name` should be unique, `ID Name` must be unique.

Example `Log` Pod Structure

```
NAME
UNAME
ENTRIES:
    0:
        SLUG
        TITLE
        SIZE
        TEXT
```
