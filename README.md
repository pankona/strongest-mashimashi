# strongest-mashimashi

Phrase generator for stronger password/passphrase.  
Visit [http://strongest-mashimashi.appspot.com/](http://strongest-mashimashi.appspot.com/) to see the function.

## This is...

* A phrase generator.
  * A phrase is consists of some adjective and 1 noun.

* For generating strong password/passphrase.
  * Words, that consists a phrase, are choose randomly from English dictionary.
  * But note that words that are too short or too long are excluded not to be the phrase too short or too long.
  * A phrase always ends with noun, it may help us to remember easily than just a random phrase like uuid.

## Use API via curl (or any API client)

For example, using curl, use API as follows.
```bash
$ curl -X GET http://strongest-mashimashi.appspot.com/api/v1/phrase?num=3
```

* phrase length can be specified using query `num`.
  * It is optional. If `num` is not specified, it is assumed to be `3`.
  * For example,
    * If `num` is 1, then a noun will be returned.
    * If `num` is 3, then 2 adjective and 1 noun will be returned.
    * If `num` is 5, then 4 adjective and 1 noun will be returned.
  * Neither `num` cannot be negative, more than equal 6, nor other than integer.

## LICENSE

MIT

## Author

Yosuke Akatsuka (@pankona)
