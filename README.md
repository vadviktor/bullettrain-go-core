# !!! DOCUMENTATION IS HEAVILY EVOLVING


<img src="http://rawgit.com/caiogondim/bullet-train-oh-my-zsh-theme/master/img/icon.svg" width="100%" />

Bullet Train is a [zsh](http://www.zsh.org/) &
[bash](https://www.gnu.org/software/bash/) shell prompt theme based on
the [Powerline Vim plugin](https://github.com/Lokaltog/vim-powerline).
It aims for simplicity, showing information only when it's relevant.

[IRC](http://webchat.freenode.net?channels=%23bullettrain-sh)
#bullettrain-sh@freenode

Core modules show:
- Time and date
- Current directory
- Exit code of last command
- Execution time of last command
- User and hostname
- Background jobs
- OS icon

External modules can show:
- Git status (https://github.com/bullettrain-sh/bullettrain-go-git)
- Current Python version and/or virtualenv
  (https://github.com/bullettrain-sh/bullettrain-go-python)
- Current Ruby version and/or gemset
  (https://github.com/bullettrain-sh/bullettrain-go-ruby)
- Current Node.js version
  (https://github.com/bullettrain-sh/bullettrain-go-nodejs)
- Current Golang version
  (https://github.com/bullettrain-sh/bullettrain-go-golang)
- Current PHP version
  (https://github.com/bullettrain-sh/bullettrain-go-php)

If you want add some new feature, of fix some bug, open an issue and
lets hack together.

## Requirements

In order to use the theme, you will first need:

* [Nerd fonts](https://nerdfonts.com/)
* Make sure terminal is using 256-colors mode with `export
  TERM="xterm-256color"`
* For [iTerm 2](http://iterm2.com/) users, make sure you go into your
  settings and set both the regular font and the non-ascii font to
  powerline compatible [fonts](https://github.com/powerline/fonts) or
  the prompt separators and special characters will not display
  correctly.

## Compatible terminal emulators

- Linux
- [Tilix](https://gnunn1.github.io/tilix-web/)
- [Terminator](https://gnometerminator.blogspot.ie/p/introduction.html)
- [Konsole](https://konsole.kde.org/) (with some bugs)
- Mac
- [iTerm2](http://iterm2.com/)

## Installing

We have prepare release executables on our release page
https://github.com/bullettrain-sh/bullettrain-go-core/releases.

Of course you are more then welcomed to build your own, customised
version if you feel comfortable with Go.

In your rc files you only need to set the single prompt variable.

(Single quotes are important not to store the evaluated result in the
variable, but to reevaluate on every call.)

### ZSH

.zshrc

`PROMPT='$(bullettrain)'`

Two side prompt feature is planned too.

### BASH

.bashrc

`export PS1='$(bullettrain)'`

## Options

Most of the behaviours can be configured through environment variables,
making you free from the recompiling work.

These are the **core** feature configuration variables and module
configuration information can be found on their respective READMEs.

All envirnment variables must be exported for Go to be able to pick up.

E.g.: `export BULLETTRAIN_CAR_ORDER="time context python ruby"`

### Defining colours and text effects

Anything what https://github.com/mgutz/ansi supports.

`foregroundColor+attributes:backgroundColor+attributes`

Colors

- black
- red
- green
- yellow
- blue
- magenta
- cyan
- white
- [0...255 (256 colors)](https://en.wikipedia.org/wiki/ANSI_escape_code#Colors)

![ansi256](ansi256.png)

Foreground Attributes

- B = Blink
- b = bold
- h = high intensity (bright)
- i = inverse
- s = strikethrough
- u = underline

Background Attributes

- h = high intensity (bright)

### Basic behaviours

| Environment variable               | Description                                                               | Default value                                     |
| :--------------------------------- | :------------------------------------------------------------------------ | :------------------------------------------------ |
| BULLETTRAIN_CARS                   | Control which cars to appear and in what order, using their _callwords_.  | `os time date context dir python status exectime` |
| BULLETTRAIN_CARS_SEPARATE_LINE     | Whether the cars should be on their own line above the prompt.            | false                                             |
| BULLETTRAIN_NO_PAINT               | Whether you wish not to use paint at all, aka black and white mode.       | false                                             |
| BULLETTRAIN_DEBUG                  | Turning debug print mode on to help seeing actual haracter codes.         | false                                             |
| BULLETTRAIN_SEPARATOR_ICON         | Defines the car separator icon.                                           | ``                                              |
| BULLETTRAIN_SEPARATOR_PAINT        | Defines the car separator icon's paint.                                   | calculated on the fly                             |
| BULLETTRAIN_PROMPT_CHAR            | Redefines the end char of the prompt when you are a normal user.          | `$ `                                              |
| BULLETTRAIN_PROMPT_CHAR_PAINT      | Redefines the end char's colour of the prompt when you are a normal user. | green                                             |
| BULLETTRAIN_PROMPT_CHAR_ROOT       | Redefines the end char of the prompt when you are a root user.            | `# `                                              |
| BULLETTRAIN_PROMPT_CHAR_ROOT_PAINT | Redefines the end char's colour of the prompt when you are a root user.   | red                                               |

## Core cars

### Time Car

Showing current time.

**Callword**: `time`

**Options**

| Environment variable                  | Description                                                    | Default value                     |
| :------------------------------------ | :------------------------------------------------------------- | :-------------------------------- |
| BULLETTRAIN_CAR_TIME_SHOW             | Whether the car needs to be shown.                             | false                             |
| BULLETTRAIN_CAR_TIME_SYMBOL_ICON      | Icon displayed on the car.                                     | ``                             |
| BULLETTRAIN_CAR_TIME_SYMBOL_PAINT     | Colour override for the car's symbol.                          | white:black                       |
| BULLETTRAIN_CAR_TIME_PAINT            | Colour override for the car's paint.                           | white:black                       |
| BULLETTRAIN_CAR_TIME_SEPARATOR_PAINT  | Colour override for the car's right hand side separator paint. | Using default painting algorythm. |
| BULLETTRAIN_CAR_TIME_SEPARATOR_SYMBOL | Override the car's right hand side separator symbol.           | Using global symbol.              |

- [ ] ability to display 12H clock

### Date Car

Showing current date. Format: `YYYY-MM-DD`

**Callword**: `date`

**Options**

| Environment variable                  | Description                                                    | Default value                     |
| :------------------------------------ | :------------------------------------------------------------- | :-------------------------------- |
| BULLETTRAIN_CAR_DATE_SHOW             | Whether the car needs to be shown.                             | false                             |
| BULLETTRAIN_CAR_DATE_PAINT            | Colour override for the car's paint.                           | red:black                         |
| BULLETTRAIN_CAR_DATE_SYMBOL_ICON      | Icon displayed on the car.                                     | ``                             |
| BULLETTRAIN_CAR_DATE_SYMBOL_PAINT     | Colour override for the car's symbol.                          | white:black                       |
| BULLETTRAIN_CAR_DATE_SEPARATOR_PAINT  | Colour override for the car's right hand side separator paint. | Using default painting algorythm. |
| BULLETTRAIN_CAR_DATE_SEPARATOR_SYMBOL | Override the car's right hand side separator symbol.           | Using global symbol.              |

**TODO list**

- [ ] make date format configurable

### Context Car

Showing current user and hostname.

**Callword**: `context`

**Options**

| Environment variable                     | Description                                                    | Default value                     |
| :--------------------------------------- | :------------------------------------------------------------- | :-------------------------------- |
| BULLETTRAIN_CAR_CONTEXT_SHOW             | Whether the car needs to be shown.                             | true                              |
| BULLETTRAIN_CAR_CONTEXT_PAINT            | Colour override for the car's paint.                           | black:white                       |
| BULLETTRAIN_CAR_CONTEXT_SEPARATOR_PAINT  | Colour override for the car's right hand side separator paint. | Using default painting algorythm. |
| BULLETTRAIN_CAR_CONTEXT_SEPARATOR_SYMBOL | Override the car's right hand side separator symbol.           | Using global symbol.              |

### Directory Car

Showing current directory.

**Callword**: `dir`

**Options**

| Environment variable                       | Description                                                    | Default value                     |
| :----------------------------------------- | :------------------------------------------------------------- | :-------------------------------- |
| BULLETTRAIN_CAR_DIRECTORY_SHOW             | Whether the car needs to be shown.                             | true                              |
| BULLETTRAIN_CAR_DIRECTORY_PAINT            | Colour override for the car's paint.                           | white:blue                        |
| BULLETTRAIN_CAR_DIRECTORY_SEPARATOR_PAINT  | Colour override for the car's right hand side separator paint. | Using default painting algorythm. |
| BULLETTRAIN_CAR_DIRECTORY_SEPARATOR_SYMBOL | Override the car's right hand side separator symbol.           | Using global symbol.              |

### OS Car

Showing current operating system logo. Mainly purposed as a design
element.

**Callword**: `os`

**Options**

| Environment variable                | Description                                                    | Default value                                      |
| :---------------------------------- | :------------------------------------------------------------- | :------------------------------------------------- |
| BULLETTRAIN_CAR_OS_SHOW             | Whether the car needs to be shown.                             | false                                              |
| BULLETTRAIN_CAR_OS_PAINT            | Colour override for the car's paint.                           | white:cyan                                         |
| BULLETTRAIN_CAR_OS_NAME             | Override the car's text.                                       | Trying to figure out by the given tools of the OS. |
| BULLETTRAIN_CAR_OS_SYMBOL_PAINT     | Colour override for the car's symbol.                          | white:cyan                                         |
| BULLETTRAIN_CAR_OS_SYMBOL_ICON      | Icon displayed on the car.                                     | ``                                              |
| BULLETTRAIN_CAR_OS_SEPARATOR_PAINT  | Colour override for the car's right hand side separator paint. | Using default painting algorythm.                  |
| BULLETTRAIN_CAR_OS_SEPARATOR_SYMBOL | Override the car's right hand side separator symbol.           | Using global symbol.                               |

### Last command exit code Car

Showing last command's exit code.

**Callword**: `status`

**Options**

| Environment variable | Description | Default value |
| :------------------- | :---------- | :------------ |


### Last command execution time Car

Showing last command's total execution time.

**Callword**: `exectime`

**Options**

| Environment variable | Description | Default value |
| :------------------- | :---------- | :------------ |


## Development

### Managing dependencies

We use Glide until some official support comes along.

https://github.com/Masterminds/glide

### Benchmarking

We not only want the prompt to be super sexy but also super snappy.
What'd be the point writting it in Go?! :)

So to bluntly benchmark it's speed, build the executable and then sample
a 10x batch 5 times like this in ZSH:

```
$ go build bullettrain.go
$ repeat 5 (time (repeat 10 ./bullettrain > /dev/null))
( repeat 10; do; ./bullettrain > /dev/null; done; )  0.48s user 0.16s system 107% cpu 0.590 total
( repeat 10; do; ./bullettrain > /dev/null; done; )  0.48s user 0.14s system 107% cpu 0.581 total
( repeat 10; do; ./bullettrain > /dev/null; done; )  0.51s user 0.15s system 107% cpu 0.615 total
( repeat 10; do; ./bullettrain > /dev/null; done; )  0.49s user 0.17s system 107% cpu 0.613 total
( repeat 10; do; ./bullettrain > /dev/null; done; )  0.51s user 0.17s system 107% cpu 0.625 total
```

Be sure to benchmark your code to make sure you are not introducing a
feature that will make the prompt sluggish all of a sudden.


## Credits

This theme is highly inspired by the following themes:
- [Powerline](https://github.com/jeremyFreeAgent/oh-my-zsh-powerline-theme)
- [Agnoster](https://gist.github.com/agnoster/3712874)
