# [DRAFT]

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
