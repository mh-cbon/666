# 666

Visually display success or failure of a command

This tool is part of the [go-github-release workflow](https://github.com/mh-cbon/go-github-release)

## Install

Pick an msi package [here](https://github.com/mh-cbon/666/releases)!

__deb/ubuntu/rpm repositories__

```sh
wget -O - https://raw.githubusercontent.com/mh-cbon/latest/master/source.sh \
| GH=mh-cbon/666 sh -xe
# or
curl -L https://raw.githubusercontent.com/mh-cbon/latest/master/source.sh \
| GH=mh-cbon/666 sh -xe
```

__deb/ubuntu/rpm packages__

```sh
curl -L https://raw.githubusercontent.com/mh-cbon/latest/master/install.sh \
| GH=mh-cbon/666 sh -xe
# or
wget -q -O - --no-check-certificate \
https://raw.githubusercontent.com/mh-cbon/latest/master/install.sh \
| GH=mh-cbon/666 sh -xe
```

__chocolatey__

```sh
choco install product666 -y
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
$ 666 echo "ok"
ok
 ✔ Success

$ 666 nopnop "ok"
exec: "nopnop": executable file not found in $PATH
 ✘ Failed
```
