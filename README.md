# pool of cybersec tools using golang 



# MILOK SCRAPER


```
 _      _  _     ____  _  __ _    ____  ____ ____  ____  ____  _____ ____
/ \__/|/ \/ \   /  _ \/ |/ // \  / ___\/   _Y  __\/  _ \/  __\/  __//  __\
| |\/||| || |   | / \||   / | |  |    \|  / |  \/|| / \||  \/||  \  |  \/|
| |  ||| || |_/\| \_/||   \ | |  \___ ||  \_|    /| |-|||  __/|  /_ |    /
\_/  \|\_/\____/\____/\_|\_\\_/  \____/\____|_/\_\\_/ \|\_/   \____\\_/\_\
```

## Description
tool used to recursively downloads the images in a URL

## Features
- saves cookies to memake a real user
- use custom http header that the browser send
- comming feature inspect the metadata of the images and change them

## Usage

```bash
./spider [-rlp] URL
```

### Options

- `-l uint` - The maximum depth level of the recursive download (default: 5)
- `-p string` - The path where the downloaded files will be saved (default: "data")
- `-r boolean` - Recursively downloads the images in a URL

### Examples

```bash
# Basic usage
./spider https://example.com

# Download with custom depth and path
./spider -r -l 3 -p ./downloads https://example.com
```
