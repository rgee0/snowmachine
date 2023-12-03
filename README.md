# snowmachine

A Go port of the python script by (John Anderson)[https://github.com/sontek/snowmachine] that brings a little festive cheer to your terminal. View it in action here:

![Rainbow ascii tree with snowfall](./images/animatedtree.gif)



Getting Started
---------------
You can make it snow:

```bash
$ snowmachine snow
```
![Example of snow command](./images/snow.png)

or render a tree:

```bash
$ snowmachine tree
```
![Example of tree command](./images/tree.png)

You can also tell it to stack the snow if you prefer.

```bash
$ snowmachine snow -stack
```
![Example of snow stack command](./images/stack.png)

If you don't like the unicode particles you can tell it to use
asterisk or some other character.  If you use cmd.exe for example,
this will be required.

```bash
$ snowmachine snow -stack -particle="*"
```
![Example of snow particle command](./images/particle.png)

You can also change the particle colors if you would like:

```bash
$ snowmachine snow -color=rainbow
```
![Example of rainbow snow command](./images/rainbow.png)

## Full list of available flags

### `snow` command
```sh
$ snowmachine snow -help
Usage of snow:
  -colour string
    	Change the colour of the particles. [red|green|blue|magenta|cyan|yellow] (default "white")
  -particle string
    	Change the particle used.
  -speed int
    	Increase to make it snow faster. (default 14)
  -stack
    	Set snow to pile up.
```

### `tree` command
```sh
$ snowmachine tree -help
Usage of tree:
  -colour string
    	Change the colour of the snow particles. [red|green|blue|magenta|cyan|yellow] (default "green")
  -light-delay int
    	Seconds between light changes (default 1)
  -light-colour string
    	Change the color of the lights. [red|green|blue|magenta|cyan|yellow] (default "rainbow")
  -particle string
    	Change the particle used for the tree. (default "*")
  -snow
    	Whether snow should fall. (default true)
  -snow-colour string
    	Change the colour of the snow particles. [red|green|blue|magenta|cyan|yellow] (default "white")
  -snow-particle string
    	Change the snow particle used.
  -snow-speed int
    	Increase to make it snow faster. (default 14)

```