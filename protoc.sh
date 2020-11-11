#!/bin/bash
if ! command -v gsed &> /dev/null
then
    brew install gnu-sed
    export PATH="/usr/local/opt/gnu-sed/libexec/gnubin:$PATH"
fi

for arg in "$@"
do
    case $arg in
        -f|--file)
            file=$2
            protoc -I. --go_out=plugins=grpc,paths=source_relative:. \
            --micro_out=paths=source_relative:. \
            --proto_path=$GOPATH/src \
            --proto_path=../ \
            $file
            grep -rli 'ClientConnInterface' --exclude=protoc.sh --exclude=README.md --exclude=fix-protobuf-files.sh * | xargs -I@ gsed -i "s/ClientConnInterface/ClientConn/g" @
            grep -rli 'SupportPackageIsVersion6' --exclude=protoc.sh --exclude=README.md --exclude=fix-protobuf-files.sh * | xargs -I@ gsed -i "s/SupportPackageIsVersion6/SupportPackageIsVersion5/g" @
            exit 1        
        ;;
        --fix)
            echo Find 'ClientConnInterface'
            grep -rli 'ClientConnInterface' --exclude=protoc.sh --exclude=README.md --exclude=fix-protobuf-files.sh *
            grep -rli 'ClientConnInterface' --exclude=protoc.sh --exclude=README.md --exclude=fix-protobuf-files.sh * | xargs -I@ gsed -i "s/ClientConnInterface/ClientConn/g" @
            echo '\n'Find 'SupportPackageIsVersion6'
            grep -rli 'SupportPackageIsVersion6' --exclude=protoc.sh --exclude=README.md --exclude=fix-protobuf-files.sh *
            grep -rli 'SupportPackageIsVersion6' --exclude=protoc.sh --exclude=README.md --exclude=fix-protobuf-files.sh * | xargs -I@ gsed -i "s/SupportPackageIsVersion6/SupportPackageIsVersion5/g" @
            exit 1
        ;;
    esac
done

echo usage  [-f generate protobuf file:filepath]
echo '      '[--fix fix protobuf generated file]
