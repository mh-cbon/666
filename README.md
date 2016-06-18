# 666

Visually display success or failure of a command

# Install

You can grab a pre-built binary file in the [releases page](https://github.com/mh-cbon/666/releases)

```sh
mkdir -p $GOPATH/github.com/mh-cbon
cd $GOPATH/github.com/mh-cbon
git clone https://github.com/mh-cbon/666.git
cd 666
glide install
go install
```

# Usage

```sh
$ t6 echo "ok"
ok
 ✔ Success

$ t6 nopnop "ok"
exec: "nopnop": executable file not found in $PATH
 ✘ Failed
```
