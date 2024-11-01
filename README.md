# resub
Tool to create subdomain permutations for active enumeration

## Usage:

Using cutom wordlist:

```
$ cat wordlist.txt 
www
test
test-uat
hello-world

$ go run resub.go FUZZ.example.com -w wordlist.txt 
www.example.com
test.example.com
test-uat.example.com
hello-world.example.com
```

using [n0kovo_subdomains](https://github.com/n0kovo/n0kovo_subdomains) wordlist:

```
go run resub.go FUZZ.example.com -m huge
```

i.e. wordlist path being

```
$HOME/.config/resub/n0kovo/n0kovo_subdomains/n0kovo_subdomains_huge.txt
```
