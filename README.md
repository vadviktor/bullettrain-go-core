# !!! DOCUMENTATION IS HEAVILY EVOLVING

Shell prompt made from building blocks while staying styleish.

```
 ______ _____ _   _ 
|___  //  ___| | | |
   / / \ `--.| |_| |
  / /   `--. \  _  |
./ /___/\__/ / | | |
\_____/\____/\_| |_/
```                    

# How to use use this prompt:

Single quotes are important not to store the evaluated result in the variable,
but to reevaluate on every call.

`PROMPT='$(prompt)'`

Two side prompt feature is planned too.

This is an example of a two-sided prompt:

```
PROMPT='%F{red}%n%f@%F{blue}%m%f %F{yellow}%1~%f %# '
RPROMPT='[%F{yellow}%?%f]'
```

```
______           _     
| ___ \         | |    
| |_/ / __ _ ___| |__  
| ___ \/ _` / __| '_ \ 
| |_/ / (_| \__ \ | | |
\____/ \__,_|___/_| |_|
```

# How to use use this prompt:

Single quotes are important not to store the evaluated result in the variable,
but to reevaluate on every call.

`export PS1='$(prompt)'`

# Development

We not only want the prompt to be super sexy but also super snappy. What'd be the point writting it in Go?! :)

So to bluntly benchmark it's speed, build the executable and then sample a 10x batch 5 times like this in ZSH:

```
$ go build bullettrain.go
$ repeat 5 (time (repeat 10 ./bullettrain > /dev/null))
( repeat 10; do; ./bullettrain > /dev/null; done; )  0.48s user 0.16s system 107% cpu 0.590 total
( repeat 10; do; ./bullettrain > /dev/null; done; )  0.48s user 0.14s system 107% cpu 0.581 total
( repeat 10; do; ./bullettrain > /dev/null; done; )  0.51s user 0.15s system 107% cpu 0.615 total
( repeat 10; do; ./bullettrain > /dev/null; done; )  0.49s user 0.17s system 107% cpu 0.613 total
( repeat 10; do; ./bullettrain > /dev/null; done; )  0.51s user 0.17s system 107% cpu 0.625 total
```

Be sure to benchmark your code to make sure you are not introducing a feature that will make the prompt sluggish all of a sudden.
