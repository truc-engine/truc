rm -rf ./common/dto/*.*
mkdir -p ./common/dto

find . -path ./common -prune -o -name '*.dto.proto' -exec sh -c 'cp "$0" "./common/dto/dto$(echo "$0" | sed -e "s|/|_|g" -e "s|^.*\/||" -e "s|^\.||")"' {} \;
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative common/dto/*.proto

rm -rf ./common/model/*.*
mkdir -p ./common/model

find . -path ./common -prune -o -name '*.model.proto' -exec sh -c 'cp "$0" "./common/model/model$(echo "$0" | sed -e "s|/|_|g" -e "s|^.*\/||" -e "s|^\.||")"' {} \;
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative common/model/*.proto
