# Hasher

A console program that takes a list of URLs as an argument, sends GET requests to these URLs, and for each response prints a URL plus a hash of the response body. 
It sends requests in parallel. Default amount of parallel requests is 10. Add a flag `-parallel` to manage the amount of parallel requests.

## Usage 

Example with default parallel limitation:
```sh
$ hasher http://example.com/foo hostname.example/bar
http://hostname.example/bar 328c30fae61cd119cd177c061d1ac11f 
http://example.com/foo 84238dfc8092e5d9c0dac8ef93371a07
```

Example with `-parallel` flag:
```sh
$ hasher -parallel 5 adjust.com google.com facebook.com yahoo.com yandex.com twitter.com
http://adjust.com 6fbf18a8a69905ec39cb8b226be50702
http://yandex.com 13c09def248d6c3bc68ce6e55da3a99f
http://facebook.com b54708db27b969ad201c804c686e22fa
...
```