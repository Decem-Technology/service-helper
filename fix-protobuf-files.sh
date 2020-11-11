#!/bin/bash
if ! command -v gsed &> /dev/null
then
    brew install gnu-sed
    export PATH="/usr/local/opt/gnu-sed/libexec/gnubin:$PATH"
fi

echo Find 'ClientConnInterface'
grep -rli 'ClientConnInterface' --exclude=README.md --exclude=fix-protobuf-files.sh *
grep -rli 'ClientConnInterface' --exclude=README.md --exclude=fix-protobuf-files.sh * | xargs -I@ gsed -i "s/ClientConnInterface/ClientConn/g" @
echo '\n'Find 'SupportPackageIsVersion6'
grep -rli 'SupportPackageIsVersion6' --exclude=README.md --exclude=fix-protobuf-files.sh *
grep -rli 'SupportPackageIsVersion6' --exclude=README.md --exclude=fix-protobuf-files.sh * | xargs -I@ gsed -i "s/SupportPackageIsVersion6/SupportPackageIsVersion5/g" @