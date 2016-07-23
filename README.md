# 666

Visually display success or failure of a command

## Install

Pick an msi package [here](https://github.com/mh-cbon/666/releases)!

__deb/rpm__

```sh
curl -L https://raw.githubusercontent.com/mh-cbon/latest/master/install.sh \
| GH=mh-cbon/666 sh -xe
# or
wget -q -O - --no-check-certificate \
https://raw.githubusercontent.com/mh-cbon/latest/master/install.sh \
| GH=mh-cbon/666 sh -xe
```

__go__

```sh
mkdir -p $GOPATH/src/github.com/mh-cbon
cd $GOPATH/src/github.com/mh-cbon
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
