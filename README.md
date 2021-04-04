# quote

A small utility for quoting and escaping shell pipeline inputs. Useful when processing inputs with spaces.

### Usage

```
Usage of dist/quote_linux_amd64/quote:
  -0	output null instead of newline
  -no
    	don't quote
  -single
    	use single quotes
  -space
    	escape spaces
```


### Example

```
% echo "hey" | quote
"hey"
```
